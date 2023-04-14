package config

type WorkSpace struct {
	ShHome     string `mapstructure:"sh_home"`
	WorkingDir string `mapstructure:"working_dir"`
}

type DatabaseConfig struct {
	Host string `mapstructure:"hostname"`
	Port int    `mapstructure:"port"`
	User string `mapstructure:"username"`
	Pass string `mapstructure:"password"`
}

type AppConfig struct {
	Ws WorkSpace      `mapstructure:"ws"`
	Db DatabaseConfig `mapstructure:"db"`
}
