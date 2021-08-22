package config

type jwt struct {
	Secret string `yaml:"secret"`
}

func Jwt() *jwt {
	return &Config().Jwt
}
