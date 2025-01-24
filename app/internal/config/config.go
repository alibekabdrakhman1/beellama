package config

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Database struct {
		Postgres struct {
			Host     string `mapstructure:"host"`
			Port     int    `mapstructure:"port"`
			User     string `mapstructure:"user"`
			Password string `mapstructure:"password"`
			DBName   string `mapstructure:"dbname"`
			SSLMode  string `mapstructure:"sslmode"`
		} `mapstructure:"postgres"`
		Mongo struct {
			URI        string `mapstructure:"uri"`
			Database   string `mapstructure:"database"`
			Collection string `mapstructure:"collection"`
		} `mapstructure:"mongo"`
	} `mapstructure:"database"`

	Redis struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
	} `mapstructure:"redis"`

	Logging struct {
		LogDir string `mapstructure:"log_dir"`
	} `mapstructure:"logging"`

	Ollama struct {
		URL string `mapstructure:"url"`
	} `mapstructure:"ollama"`
}
