package application

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/guilhermevianafreire/dev_env_manager/internal/component"
	"github.com/guilhermevianafreire/dev_env_manager/internal/page"
	"log"
)

var applicationStyle = lipgloss.NewStyle()

type AppModel struct {
	page     tea.Model
	width    int
	height   int
	quitting bool
	err      error
}

func NewAppModel() *AppModel {
	return &AppModel{
		page: page.NewHomePage(0, 0),
	}
}

func (m AppModel) Init() tea.Cmd {
	return nil
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := applicationStyle.GetFrameSize()
		m.width = msg.Width - h
		m.height = msg.Height - v
		m.page = page.NewHomePage(m.width, m.height)
	case tea.KeyMsg:
		switch msg.String() {
		case component.QKey, component.ControlCKey:
			return m, tea.Quit
		case component.HomeKey:
			m.page = page.NewHomePage(m.width, m.height)
		}
	}

	newPageModel, cmd := m.page.Update(msg)
	m.page = newPageModel
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m AppModel) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := applicationStyle.Render(m.page.View())
	if m.quitting {
		return str + "\n"
	}
	return str
}

func Start() {
	p := tea.NewProgram(NewAppModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
