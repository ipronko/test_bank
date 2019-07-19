package config

import (
	"fmt"
	"github.com/FZambia/viper-lite"
	"github.com/spf13/pflag"
	"log"
	"strings"
)

const (
	DebugMode Flag = "debug"
	ServeAddr Flag = "serve-addr"
)

func New() *viper.Viper {
	initPFlags()
	cfg := viper.New()
	cfg.AutomaticEnv()
	err := cfg.BindPFlags(pflag.CommandLine)
	if err != nil {
		log.Fatal(err)
	}
	cfg.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	return cfg
}

type Flag string

func (f Flag) String() string {
	return string(f)
}

func (f Flag) EnvStyle() string {
	upper := strings.ToUpper(string(f))
	return strings.NewReplacer("-", "_").Replace(upper)
}

func initPFlags() {
	pflag.Bool(DebugMode.String(),
		false,
		fmt.Sprintf(" enable debug mode\n env: %s\n", DebugMode.EnvStyle()))

	pflag.String(ServeAddr.String(),
		":8080",
		fmt.Sprintf(" serve addr\n env: %s\n", ServeAddr.EnvStyle()))

	pflag.Parse()
}
