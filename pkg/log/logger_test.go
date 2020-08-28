package log

import (
	"testing"
)

func TestInitLog(t *testing.T) {
	Init()
	Logger.Debug("21312")
}

func TestInitConfig(t *testing.T) {
	// config = &Config{
	// 	Name:      "渣渣辉",
	// 	Path:      "./runtime/log",
	// 	Level:     "",
	// 	Formatter: "json",
	// 	Stdout:    true,
	// }
	//
	// cfg.Viper().Set("logger", config)
	// if err := cfg.Viper().WriteConfig(); err != nil {
	// 	t.Fatal(err)
	// }
	//
	// fmt.Println(cfg.Viper().GetStringMap("logger"))
}
