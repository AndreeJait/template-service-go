package main

import (
	"github.com/AndreeJait/template-service-go/common/constant"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "root",
}

var migrateRoot = &cobra.Command{
	Use:   "migrate",
	Short: "to execute migrations files",
}

var migrateUp = &cobra.Command{
	Use:     "up",
	Aliases: []string{"u"},
	Short:   "to update table with newest file migrate",
	Run: func(cmd *cobra.Command, args []string) {
		appModeStr, err := cmd.Flags().GetString(constant.FlagEnvMode)
		if err != nil {
			appModeStr = string(constant.Development)
		}
		appMode := constant.GetAppMode(appModeStr)
		StartMigrate(constant.MigrateTypeUp, appMode)
	},
}

var migrateDown = &cobra.Command{
	Use:     "down",
	Aliases: []string{"d"},
	Short:   "to drop last update table,start with latest file migrate",
	Run: func(cmd *cobra.Command, args []string) {
		appModeStr, err := cmd.Flags().GetString(constant.FlagEnvMode)
		if err != nil {
			appModeStr = string(constant.Development)
		}
		appMode := constant.GetAppMode(appModeStr)
		StartMigrate(constant.MigrateTypeDown, appMode)
	},
}

var migrateFresh = &cobra.Command{
	Use:     "fresh",
	Aliases: []string{"f"},
	Short:   "to drop last update table,start with latest file migrate",
	Run: func(cmd *cobra.Command, args []string) {
		appModeStr, err := cmd.Flags().GetString(constant.FlagEnvMode)
		if err != nil {
			appModeStr = string(constant.Development)
		}
		appMode := constant.GetAppMode(appModeStr)
		StartMigrate(constant.MigrateTypeFresh, appMode)
	},
}
