// configurator/models.go
package configurator

type BotConfig struct {
	Name          string `mapstructure:"name"`
	DefaultStatus string `mapstructure:"default_status"`
}

type GlobalConfig struct {
	Bot BotConfig `mapstructure:"bot"`
}