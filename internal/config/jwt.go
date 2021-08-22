package config

type jwt struct {
	SecretKey    string `yaml:"secretKey"`
	AccessToken  token  `yaml:"accessToken"`
	RefreshToken token  `yaml:"refreshToken"`
}

type token struct {
	ExpiredAt int `yaml:"expiredAt"`
}

func Jwt() *jwt {
	return &Config().Jwt
}
