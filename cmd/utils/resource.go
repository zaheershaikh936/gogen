package utils

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zaheershaikh936/gogen/cmd/utils/helper"
)

var Resource = &cobra.Command{
    Use: "resource",
    Short: "Generate CRUD resource",
    Long: "Generates a complete CRUD API resource including controller, service, repository, and routes using the Fiber framework.",
    Example: "gogen resource [model] --output [output]",
    Run: func(cmd *cobra.Command, args []string) {
        model := helper.Pluralize(args[0])
        output, _ := cmd.Flags().GetString("output")
        if output == "" { output = "./" }

        fmt.Println("🚀 Generating CRUD resource...")
        fmt.Println(model, ":Model")
        fmt.Println()

        fmt.Println("Files to be generated:")
        // Create model
        helper.CreateModel(model, output)

        // Routes
        fmt.Printf("  ✓ %s/routes/%s_routes.go\n", model, helper.ToSnakeCase(model))
        helper.RoutesGenerated(model, output)

        // Controller
        fmt.Printf("  ✓ %s/controllers/%s_controller.go\n", model, helper.ToSnakeCase(model))
        helper.ControllerGenerated(model, output)

        // Service
        fmt.Printf("  ✓ %s/services/%s_service.go\n", model, helper.ToSnakeCase(model))
        helper.ServiceGenerated(model, output)

        // Repository
        fmt.Printf("  ✓ %s/repositories/%s_repository.go\n", model, helper.ToSnakeCase(model))
        helper.RepositoryGenerated(model, output)
    },
    Args: cobra.ExactArgs(1),
}