package tools



type Config struct {
	OpenAIKey string `mapstructure:"OpenAIKey"`
	Username string `mapstructire:"Username"`
}


var defaultConfig = Config{
	OpenAIKey: "",
	Username: "",
}