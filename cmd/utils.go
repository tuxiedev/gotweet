package cmd

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/tuxiedev/gotweet/pkg/app"
)

// RequiredFlag Defines a flag
type RequiredFlag struct {
	value  *string
	name   string
	envVar string
	help   string
}

func buildFlags(pf *pflag.FlagSet, requiredFlags []RequiredFlag) {
	for _, rf := range requiredFlags {
		pf.StringVar(rf.value, rf.name, viper.GetString(rf.envVar), "REQUIRED: "+rf.help)
	}
}

func runApp(outputName string, outputConfig interface{}) {
	app.RunApp(getTwitterConfigs(), outputName, outputConfig)
}
