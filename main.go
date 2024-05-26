package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/joelhenwang/go-cli-pokedex/style"
	"github.com/joelhenwang/go-cli-pokedex/utils"
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.table = m.table.Width(msg.Width)
		m.table = m.table.Height(msg.Height)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
		}
	}
	return m, cmd
}

func (m model) View() string {
	return "\n" + m.table.String() + "\n"
}

func main() {
	re := lipgloss.NewRenderer(os.Stdout)
	baseStyle := style.GetBaseStyling(re)
	headerStyle := style.GetHeaderStyling(re)
	//selectedStyle := style.GetSelectedStyling(re)

	pokemon_list, err := utils.LoadPokemonCsv("pokemon_gen_1_to_8.csv")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	headers := []string{"#", "Name", "Gen", "Type 1", "Type 2", "HP", "Attack", "Defense", "Sp.Atk", "Sp.Def", "Speed", "Total"}

	rows := make([][]string, len(pokemon_list))
	for i, p := range pokemon_list {
		rows[i] = []string{p.Id, p.Name, p.Gen, p.Type1, p.Type2, p.Hp, p.Attack, p.Defense, p.SpAttack, p.SpDefense, p.Speed, p.TotalPoints}
	}

	t := table.New().
		Headers(headers...).
		Rows(rows...).Border(lipgloss.NormalBorder()).
		BorderStyle(re.NewStyle().Foreground(lipgloss.Color("241"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == 0 {
				return headerStyle
			}

			switch col {
			case 3, 4:
				typeColor := style.GetTypeStyling()[rows[row-1][col]]
				return baseStyle.Copy().Foreground(typeColor)
			}

			return baseStyle.Copy().Foreground(lipgloss.Color("241"))
		}).Border(lipgloss.ThickBorder())

	m := model{table: t}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		log.Fatal("Error running program:", err)
		os.Exit(1)
	}

}

type model struct {
	table *table.Table
}
