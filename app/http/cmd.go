package main

import (
	"github.com/AndreeJait/template-service-go/common/constant"
	"github.com/spf13/cobra"
)

var cmdRoot = &cobra.Command{
	Use:   "",
	Short: "run http service",
	Long:  "this command use to run http server",
	Run: func(cmd *cobra.Command, args []string) {
		appModeStr, err := cmd.Flags().GetString(constant.FlagEnvMode)
		if err != nil {
			appModeStr = string(constant.Development)
		}
		appMode := constant.GetAppMode(appModeStr)
		initServer(appMode)
	},
}
