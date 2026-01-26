package utils

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/zaheershaikh936/gogen/cmd/utils/helper"
	"github.com/zaheershaikh936/gogen/cmd/utils/text"
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
		var framework string

		output, _ = cmd.Flags().GetString("output")
		if output == "" {
			output = "./"
		}

		if len(args) > 0 {
			model = args[0]
		} else {
			// Interactive Wizard for model name
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Title("What is the model name?").
						Prompt("? ").
						Placeholder("e.g. user-order").
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

		// Always prompt for framework
		frameworkForm := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Pick a framework").
					Options(
						huh.NewOption("Fiber", "fiber"),
						huh.NewOption("Gin", "gin"),
					).
					Value(&framework),
			),
		).WithTheme(huh.ThemeCharm())

		err := frameworkForm.Run()
		if err != nil {
			fmt.Println("Framework selection cancelled.")
			os.Exit(0)
		}

		// Standardize model name for the generator
		model = helper.Pluralize(model)

		// Start Spinner
		complete := make(chan bool)
		go ui.ShowSpinner(complete)

		// Create files
		helper.CreateModel(model, output)
		ui.ActionDelay()

		helper.RoutesGenerated(model, output, framework)
		ui.ActionDelay()

		helper.ControllerGenerated(model, output, framework)
		ui.ActionDelay()

		helper.ServiceGenerated(model, output, framework)
		ui.ActionDelay()

		helper.RepositoryGenerated(model, output, framework)
		ui.ActionDelay()

		// Stop Spinner
		complete <- true

		// Verification Summary
		header := ui.SuccessStyle.Render("✓ Files generated successfully")

		var lines []string

		// Map paths for the summary
		snakeName := text.ToSnakeCase(model)

		paths := []string{
			fmt.Sprintf("%s/routes/%s_routes.go", model, snakeName),
			fmt.Sprintf("%s/controllers/%s_controller.go", model, snakeName),
			fmt.Sprintf("%s/services/%s_service.go", model, snakeName),
			fmt.Sprintf("%s/repositories/%s_repository.go", model, snakeName),
		}

		for _, p := range paths {
			lines = append(lines, fmt.Sprintf("%s %s", ui.InfoStyle.Render("↳"), ui.PathStyle.Render(p)))
		}

		summaryContent := lipgloss.JoinVertical(
			lipgloss.Left,
			header,
			lipgloss.JoinVertical(lipgloss.Left, lines...),
		)

		fmt.Println(summaryContent)

		fmt.Printf("\n%s Resource %s is ready to use!\n", ui.SuccessStyle.Render("✨"), ui.FileStyle.Render(model))
	},
	Args: cobra.MaximumNArgs(1),
}
