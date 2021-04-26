package config

type ApiModel struct {
	Host     string
	Port     string
	Username string
	Password string
}

type ConfigModel struct {
	Api *ApiModel
}

func NewConfigModel() *ConfigModel {
	return &ConfigModel{}
}
