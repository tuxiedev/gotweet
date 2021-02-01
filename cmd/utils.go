package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// RequiredFlag Defines a flag
type RequiredFlag struct {
	value *string
	name  string
	help  string
}

func buildFlagsAndMarkThemRequired(pf *pflag.FlagSet, requiredFlags []RequiredFlag) {
	for _, rf := range requiredFlags {
		pf.StringVar(rf.value, rf.name, "", "REQUIRED: "+rf.help)
		cobra.MarkFlagRequired(pf, rf.name)
	}
}

