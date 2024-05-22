package main

import "github.com/AndreeJait/exam-platform-service/common/constant"

func main() {
	cmdRoot.Flags().StringP(constant.FlagEnvMode, constant.FlagShortEnvMode,
		string(constant.Development), "to define app mode")

	err := cmdRoot.Execute()
	if err != nil {
		panic(err)
	}
}
