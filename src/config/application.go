package config

type ApplicationConfigs struct {
	App struct {
		Port string `yaml:"port"`
	} `yaml:"app"`
}
