package config

type server struct {
	Name    string `yaml:"name"`
	Port    int    `yaml:"port"`
	Logger  bool   `yaml:"logger"`
	Prefork bool   `yaml:"prefork"`
}

func Server() *server {
	return &Config().Server
}
