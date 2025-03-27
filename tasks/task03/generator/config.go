package generator

type Config struct {
	NumberOfGenerations int `env:"NUMBER_OF_GENERATIONS" envDefault:"1000000" yaml:"NUMBER_OF_GENERATIONS"`
	NumberOfGoroutines  int `env:"NUMBER_OF_GOROUTINES" envDefault:"16" yaml:"NUMBER_OF_GOROUTINES"`
	MinCarBrandLen      int `env:"MIN_CAR_BRAND_LEN" envDefault:"5" yaml:"MIN_CAR_BRAND_LEN"`
	MaxCarBrandLen      int `env:"MAX_CAR_BRAND_LEN" envDefault:"30" yaml:"MAX_CAR_BRAND_LEN"`
	MinCarPrice         int `env:"MIN_CAR_PRICE" envDefault:"5000" yaml:"MIN_CAR_PRICE"`
	MaxCarPrice         int `env:"MAX_CAR_PRICE" envDefault:"10000000" yaml:"MAX_CAR_PRICE"`
	BatchSize           int `env:"BATCH_SIZE" envDefault:"10000" yaml:"BATCH_SIZE"`
}
