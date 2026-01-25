package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	// Colors
	PrimaryColor   = lipgloss.Color("#5f5fff")
	SecondaryColor = lipgloss.Color("#AE91FF")
	SuccessColor   = lipgloss.Color("#04B575")
	ErrorColor     = lipgloss.Color("#FF4C4C")
	DimmedColor    = lipgloss.Color("#777777")

	// Styles
	HeaderStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true).
			Padding(0, 1)

	SubtitleStyle = lipgloss.NewStyle().
			Foreground(DimmedColor).
			Italic(true).
			Padding(0, 1)

	BoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(PrimaryColor).
			Padding(1, 4).
			Margin(1, 2)

	SuccessStyle = lipgloss.NewStyle().
			Foreground(SuccessColor).
			Bold(true)

	InfoStyle = lipgloss.NewStyle().
			Foreground(SecondaryColor)

	PathStyle = lipgloss.NewStyle().
			Foreground(DimmedColor)

	FileStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true)

	// Components
	Logo         = "G O G E N"
	LogoSubtitle = "The professional Go/Fiber generator"
)

func PrintLogo() {
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		HeaderStyle.Render(Logo),
		SubtitleStyle.Render(LogoSubtitle),
	)
	fmt.Println(BoxStyle.Render(content))
}

func GetFormTheme() *lipgloss.Style {
	s := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(PrimaryColor).
		PaddingLeft(2)
	return &s
}
