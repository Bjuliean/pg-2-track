### Протестировать падение производителности при исползовании pgbouncer в разных режимах: statement, transaction, session:
##### statement:
![statement1](../../images/task05/13.png)
![statement2](../../images/task05/14.png)
##### transaction:
![transaction1](../../images/task05/15.png)
![transaction2](../../images/task05/16.png)
##### session:
![session1](../../images/task05/17.png)
![session2](../../images/task05/18.png)
#### Итог:
#### statement - наиболее агрессивный режим, вернет соединение в пул, как только будет обработан первый запрос
#### transaction - удерживает соединение на момент выполнения транзакции, наблюдается небольшой упадок в производительности, поскольку в тесте транзакции короткие
#### session - наименее агрессивный режим, соединение удерживается для всей сессии клиента, значительный упадок в производительности, в сравнении со statement

### Предварительная работа по лекции:
![lecture](../../images/task05/1.png)
![lecture](../../images/task05/2.png)
![lecture](../../images/task05/3.png)
![lecture](../../images/task05/4.png)
![lecture](../../images/task05/5.png)
![lecture](../../images/task05/6.png)
![lecture](../../images/task05/7.png)
![lecture](../../images/task05/8.png)
![lecture](../../images/task05/9.png)
![lecture](../../images/task05/10.png)
![lecture](../../images/task05/11.png)
![lecture](../../images/task05/12.png)