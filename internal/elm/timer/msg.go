package timer

import tea "github.com/charmbracelet/bubbletea"

func WorkCompleted() tea.Msg {
	return WorkCompletedMsg{}
}

type WorkCompletedMsg struct{}
