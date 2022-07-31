package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "mkstacks",
}

func init() {
	cobra.OnInitialize(func() {})
}
