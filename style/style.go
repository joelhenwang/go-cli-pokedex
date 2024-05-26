package style

import "github.com/charmbracelet/lipgloss"

func GetTypeStyling() map[string]lipgloss.Color {
	var typeColors map[string]lipgloss.Color = make(map[string]lipgloss.Color)

	typeColors = map[string]lipgloss.Color{
		"Bug":      lipgloss.Color("#D7FF87"),
		"Electric": lipgloss.Color("#FDFF90"),
		"Fire":     lipgloss.Color("#FF7698"),
		"Flying":   lipgloss.Color("#FF87D7"),
		"Grass":    lipgloss.Color("#75FBAB"),
		"Ground":   lipgloss.Color("#FF875F"),
		"Normal":   lipgloss.Color("#929292"),
		"Poison":   lipgloss.Color("#7D5AFC"),
		"Water":    lipgloss.Color("#00E2C7"),
	}

	return typeColors
}

func GetBaseStyling(re *lipgloss.Renderer) lipgloss.Style {
	return re.NewStyle().
		Padding(0, 1)
}

func GetHeaderStyling(re *lipgloss.Renderer) lipgloss.Style {
	return re.NewStyle().
		Background(lipgloss.Color("240")).
		Foreground(lipgloss.Color("235")).
		Bold(true).
		Padding(0, 1)
}

func GetSelectedStyling(re *lipgloss.Renderer) lipgloss.Style {
	return re.NewStyle().
		Background(lipgloss.Color("235")).
		Foreground(lipgloss.Color("240")).
		Bold(true).
		Padding(0, 1)
}
