package models

type Config struct {
	Server struct {
		Port	string	`yaml:"port" binding:"required"`
		Secret	string	`yaml:"secret" binding:"required"`
		Metrics	struct {
			Username	string	`yaml:"username"`
			Password	string	`yaml:"password" binding:"required"`
		} `yaml:"metrics"`
	} `yaml:"server"`
	Database DatabaseConfig	`yaml:"database"`
	Mail struct {
		Server		string	`yaml:"server" binding:"required"`
		Address		string	`yaml:"address" binding:"required"`
		Password	string	`yaml:"password" binding:"required"`
		Port		uint	`yaml:"port" binding:"required"`
	}	`yaml:"mail" binding:"required"`
}

type DatabaseConfig struct {
	Mariadb SqlConfig `yaml:"mariadb"`
	Sqlite	string `yaml:"sqlite"`
}

type SqlConfig struct {
	Username	string	`yaml:"username"`
	Password	string	`yaml:"password"`
	Hostname	string	`yaml:"hostname"`
	Port		string	`yaml:"port"`
	Database	string	`yaml:"database"`
}
