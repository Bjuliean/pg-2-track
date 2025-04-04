gcloud beta compute --project=celtic-house-266612 instances create postgres4 --zone=us-central1-a --machine-type=e2-medium --subnet=default --network-tier=PREMIUM --maintenance-policy=MIGRATE --service-account=933982307116-compute@developer.gserviceaccount.com --scopes=https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring.write,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/trace.append --image-family=ubuntu-2404-lts-amd64 --image-project=ubuntu-os-cloud --boot-disk-size=100GB --boot-disk-type=pd-ssd --boot-disk-device-name=postgres4 --no-shielded-secure-boot --shielded-vtpm --shielded-integrity-monitoring --reservation-affinity=any

gcloud compute ssh postgres4

sudo apt update && sudo DEBIAN_FRONTEND=noninteractive apt upgrade -y && sudo sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list' && wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add - && sudo apt-get update && sudo DEBIAN_FRONTEND=noninteractive apt -y install postgresql-17 unzip atop iotop

sudo su postgres

-- посмотрим pg_bench в несколько потоков
cd ~ && wget https://storage.googleapis.com/thaibus/thai_small.tar.gz && tar -xf thai_small.tar.gz && psql < thai.sql

cat > ~/workload.sql << EOL
\set r random(1, 5000000)
SELECT id, fkRide, fio, contact, fkSeat FROM book.tickets WHERE id = :r;
EOL
------------------------------------
-- for 17 PG
/usr/lib/postgresql/17/bin/pgbench -c 8 -j 4 -T 10 -f ~/workload.sql -n -U postgres thai
--11 k

-- в 100 юзеров
/usr/lib/postgresql/17/bin/pgbench -c 100 -j 4 -T 10 -f ~/workload.sql -n -U postgres thai
-- часть не смогли получить коннект


-- развернем pg_bouncer и проверим на той же машине
-- https://www.pgbouncer.org/
sudo DEBIAN_FRONTEND=noninteractive apt install -y pgbouncer

sudo systemctl status pgbouncer

sudo systemctl stop pgbouncer

cat > temp.cfg << EOF
[databases]
thai = host=0.0.0.0 port=5432 dbname=thai
[pgbouncer]
logfile = /var/log/postgresql/pgbouncer.log
pidfile = /var/run/postgresql/pgbouncer.pid
listen_addr = *
listen_port = 6432
auth_type = scram-sha-256
auth_file = /etc/pgbouncer/userlist.txt
admin_users = admindb
EOF
cat temp.cfg | sudo tee -a /etc/pgbouncer/pgbouncer.ini

cat > temp2.cfg << EOF
"admindb" "admin123#"
"postgres" "admin123#"
EOF
cat temp2.cfg | sudo tee -a /etc/pgbouncer/userlist.txt

sudo systemctl start pgbouncer


-- зададим пароль юзеру postgres
sudo -u postgres psql -c "ALTER USER postgres WITH PASSWORD 'admin123#';";
--sudo -u postgres psql -c "create user admindb with password 'admin123#';";

-- скрам-ша-256 в некоторых версиях не работает для юзерлиста
-- !!! обратите внимание, если встречаются спецсимволы скрипты могут криво отрабатывать с такими паролями в скрам-ша-256!!!
sudo -u postgres psql -c "select usename,passwd from pg_shadow;"

sudo -u postgres psql -c "select sha256('pass');"

echo "localhost:5432:thai:postgres:admin123#" | sudo tee -a /var/lib/postgresql/.pgpass && sudo chmod 600 /var/lib/postgresql/.pgpass
sudo chown postgres:postgres /var/lib/postgresql/.pgpass

sudo su postgres
psql -h localhost -U postgres
psql -h localhost -U postgres -d thai

SHOW password_encryption;

\! nano /etc/pgbouncer/pgbouncer.ini
\! nano /etc/pgbouncer/userlist.txt

-- трейс логов
\! tail /var/log/postgresql/postgresql-17-main.log
\! tail /var/log/postgresql/pgbouncer.log


psql -p 6432 -h 127.0.0.1 -d thai -U postgres

-- зайти в админку pgbouncer
psql -p 6432 -h 127.0.0.1 -d pgbouncer -U admindb

show clients;

-- Просмотр статистики в баунсере
show servers;
SHOW STATS_TOTALS;
show pools;

-- Поставить на паузу коннекты:
-- пауза на прием новых коннектов, старые будут выполняться
pause thai;

-- Возобновить коннект:
resume thai;


-- сравним скорости на разных сетевых интерфейсах
-- psql -p 6432 -h 127.0.0.1 -d thai -U postgres
/usr/lib/postgresql/17/bin/pgbench -c 8 -j 4 -T 10 -f ~/workload.sql -n -U postgres thai
-- 10700

-- это был линукс сокет
-- а теперь TCP
/usr/lib/postgresql/17/bin/pgbench -c 8 -j 4 -T 10 -f ~/workload.sql -n -U postgres -p 5432 -h localhost thai
-- 7900


-- admin123#
/usr/lib/postgresql/17/bin/pgbench -c 8 -j 4 -T 10 -f ~/workload.sql -n -U postgres -p 6432 -h 127.0.0.1 thai
-6400

-- 100+ юзеров
/usr/lib/postgresql/17/bin/pgbench -c 120 -j 4 -T 10 -f ~/workload.sql -n -U postgres -p 5432 -h localhost thai
/usr/lib/postgresql/17/bin/pgbench -c 120 -j 4 -T 10 -f ~/workload.sql -n -U postgres -p 6432 -h 127.0.0.1 thai
--5200

nano /etc/postgresql/17/main/pg_hba.conf






-- pool_mode
max_client_conn = 2000
# How many server connections to allow per user/database pair
default_pool_size = 80


gcloud compute instances delete postgres4

-- есть мнение, что временная таблица сначала создается на диске (с) Сергей
