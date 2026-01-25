package utils

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/zaheershaikh936/gogen/cmd/utils/helper"
	"github.com/zaheershaikh936/gogen/cmd/utils/ui"
)

var Resource = &cobra.Command{
	Use:     "resource [model]",
	Short:   "Generate CRUD resource",
	Long:    "Generates a complete CRUD API resource including controller, service, repository, and routes using the Fiber framework.",
	Example: "gogen resource user\ngogen resource",
	Run: func(cmd *cobra.Command, args []string) {
		ui.PrintLogo()

		var model string
		var output string

		output, _ = cmd.Flags().GetString("output")
		if output == "" {
			output = "./"
		}

		if len(args) > 0 {
			model = args[0]
		} else {
			// Interactive Wizard
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Title("What is the model name?").
						Prompt("? ").
						Placeholder("e.g. user").
						Value(&model).
						Validate(func(str string) error {
							if str == "" {
								return fmt.Errorf("model name is required")
							}
							return nil
						}),
				),
			).WithTheme(huh.ThemeCharm())

			err := form.Run()
			if err != nil {
				fmt.Println("Wizard cancelled.")
				os.Exit(0)
			}
		}

		model = helper.Pluralize(model)

		// Start Spinner
		complete := make(chan bool)
		go ui.ShowSpinner(complete)

		// Create model
		helper.CreateModel(model, output)
		ui.ActionDelay()

		// Routes
		helper.RoutesGenerated(model, output)
		ui.ActionDelay()

		// Controller
		helper.ControllerGenerated(model, output)
		ui.ActionDelay()

		// Service
		helper.ServiceGenerated(model, output)
		ui.ActionDelay()

		// Repository
		helper.RepositoryGenerated(model, output)
		ui.ActionDelay()

		// Stop Spinner
		complete <- true

		// Verification Summary
		header := ui.SuccessStyle.Render("✓ Files generated successfully")

		var lines []string

		// Routes
		routePath := fmt.Sprintf("%s/routes/%s_routes.go", model, helper.ToSnakeCase(model))
		lines = append(lines, fmt.Sprintf("%s %s", ui.InfoStyle.Render("↳"), ui.PathStyle.Render(routePath)))

		// Controller
		controllerPath := fmt.Sprintf("%s/controllers/%s_controller.go", model, helper.ToSnakeCase(model))
		lines = append(lines, fmt.Sprintf("%s %s", ui.InfoStyle.Render("↳"), ui.PathStyle.Render(controllerPath)))

		// Service
		servicePath := fmt.Sprintf("%s/services/%s_service.go", model, helper.ToSnakeCase(model))
		lines = append(lines, fmt.Sprintf("%s %s", ui.InfoStyle.Render("↳"), ui.PathStyle.Render(servicePath)))

		// Repository
		repoPath := fmt.Sprintf("%s/repositories/%s_repository.go", model, helper.ToSnakeCase(model))
		lines = append(lines, fmt.Sprintf("%s %s", ui.InfoStyle.Render("↳"), ui.PathStyle.Render(repoPath)))

		summaryContent := lipgloss.JoinVertical(lipgloss.Left, append([]string{header, ""}, lines...)...)

		fmt.Println(lipgloss.NewStyle().
			Border(lipgloss.DoubleBorder()).
			BorderForeground(ui.SuccessColor).
			Padding(1, 2).
			MarginLeft(2).
			Render(summaryContent))

		fmt.Printf("\n  %s Resource %s is ready to use!\n", ui.SuccessStyle.Render("✨"), ui.FileStyle.Render(model))
	},
	Args: cobra.MaximumNArgs(1),
}
