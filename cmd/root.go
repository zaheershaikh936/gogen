package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gogen",
	Short: "A professional CRUD generator for Go Fiber applications",
	Long: `Gogen is a powerful CLI tool designed to accelerate Go Fiber web development
by generating idiomatic CRUD resources including controllers, services, 
repositories, and routes using clean architecture principles.`,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}