package cmd

import "github.com/zaheershaikh936/gogen/cmd/utils"


func init() {
    rootCmd.AddCommand(utils.Resource)
    utils.Resource.Flags().StringP("output", "o", "./", "Output directory")
}
