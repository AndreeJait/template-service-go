package main

import "github.com/AndreeJait/template-service-go/common/constant"

func main() {

	migrateUp.Flags().StringP(constant.FlagEnvMode, constant.FlagShortEnvMode,
		string(constant.Development), "to define app mode")
	migrateRoot.AddCommand(migrateUp)
	migrateDown.Flags().StringP(constant.FlagEnvMode, constant.FlagShortEnvMode,
		string(constant.Development), "to define app mode")
	migrateRoot.AddCommand(migrateDown)
	migrateFresh.Flags().StringP(constant.FlagEnvMode, constant.FlagShortEnvMode,
		string(constant.Development), "to define app mode")
	migrateRoot.AddCommand(migrateFresh)
	rootCmd.AddCommand(migrateRoot)

	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
