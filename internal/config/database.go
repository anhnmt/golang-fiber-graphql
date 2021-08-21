package config

type database struct {
	Driver  string `yaml:"driver"`
	Url     string `yaml:"url"`
	Migrate bool   `yaml:"migrate"`
	Logger  bool   `yaml:"logger"`
}

func Database() *database {
	return &Config().Database
}
