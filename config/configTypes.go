package config

import "github.com/qdriven/qfluent-go/pkg/orm"

// DBConfig is the configurations for connecting database
type DBConfig struct {
	Driver string // db driver name: sqlite, mysql, postgres
	DSN    string // db connection string
}

// HTTPConfig is the configurations for HTTP server
type HTTPConfig struct {
	Addr string // listen address: ":8080"
	//Https       bool   `json:"https"`    // enable https?
	//TLSCertPath string `json:"tls_cert_path"` // path to tls cert file
	//TLSKeyPath  string `json:"tls_key_path"`  // path to tls key file
}

// BaseConfig includes common config for services
type BaseConfig struct {
	DB       DBConfig   // database config
	HTTP     HTTPConfig // http listen config
	LogLevel string     // log level
}

type Server struct {
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email   Email   `mapstructure:"email" json:"email" yaml:"email"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// auto
	AutoCode Autocode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	// gorm
	Mysql  orm.Mysql       `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Mssql  orm.Mssql       `mapstructure:"mssql" json:"mssql" yaml:"mssql"`
	Pgsql  orm.Pgsql       `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Oracle orm.Oracle      `mapstructure:"oracle" json:"oracle" yaml:"oracle"`
	Sqlite orm.Sqlite      `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	// oss
	Local      Local      `mapstructure:"local" json:"local" yaml:"local"`
	Qiniu      Qiniu      `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	AliyunOSS  AliyunOSS  `mapstructure:"aliyun-oss" json:"aliyun-oss" yaml:"aliyun-oss"`
	HuaWeiObs  HuaWeiObs  `mapstructure:"hua-wei-obs" json:"hua-wei-obs" yaml:"hua-wei-obs"`
	TencentCOS TencentCOS `mapstructure:"tencent-cos" json:"tencent-cos" yaml:"tencent-cos"`
	AwsS3      AwsS3      `mapstructure:"aws-s3" json:"aws-s3" yaml:"aws-s3"`

	Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`
	Timer Timer `mapstructure:"timer" json:"timer" yaml:"timer"`

	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}
