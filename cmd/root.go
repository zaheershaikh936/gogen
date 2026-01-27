package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version = "v1.1.0"

var rootCmd = &cobra.Command{
	Use:   "gogen",
	Short: "A professional CLI tool to generate CRUD resources for Go Fiber and Gin",
	Long: `Gogen - Go CRUD Generator CLI A command-line tool for generating CRUD resource scaffolding for Go Fiber and Gin web frameworks. gogen automates the creation of controllers, services, repositories, and routes.`,
	// Adding the version field
	Version: Version, // this enables the --version flag
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}