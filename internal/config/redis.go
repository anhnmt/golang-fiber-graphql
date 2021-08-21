package config

type redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

func Redis() *redis {
	return &Config().Redis
}
