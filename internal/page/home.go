package page

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/guilhermevianafreire/dev_env_manager/internal/component"
)

var (
	homePageStyle = lipgloss.NewStyle()
	menuStyle     = lipgloss.NewStyle().
			Background(lipgloss.Color("26")).
			Foreground(lipgloss.Color("196")).
			Padding(0, 1)
)

type HomePageModel struct {
	menu         list.Model
	selectedItem list.Item
	width        int
	height       int
}

func NewHomePage(width int, height int) HomePageModel {
	return HomePageModel{
		menu: NewMenu(width, height),
	}
}

func NewMenu(width int, height int) list.Model {
	menuItems := make([]list.Item, 5)

	menuItems[0] = component.NewMenuItem("Environment", "Manage the development environment", component.Environment)
	menuItems[1] = component.NewMenuItem("Infrastructure", "Manage the development environment infrastructure", component.Infrastructure)
	menuItems[2] = component.NewMenuItem("Configuration", "Configure this application", component.Configuration)
	menuItems[3] = component.NewMenuItem("Application", "Manage the development environment applications", component.Application)
	menuItems[4] = component.NewMenuItem("Manual", "Application manual", component.Application)

	delegate := list.NewDefaultDelegate()

	menu := list.New(menuItems, delegate, width, height)

	menu.Title = "Development Environment Manager"
	menu.Styles.Title = menuStyle

	return menu
}

func (m HomePageModel) Init() tea.Cmd {
	return nil
}

func (m HomePageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var commands []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := homePageStyle.GetFrameSize()
		m.width = msg.Width - h
		m.height = msg.Height - v
		m.menu = NewMenu(m.width, m.height)
	case tea.KeyMsg:
		if m.menu.FilterState() == list.Filtering {
			break
		}
		switch msg.String() {
		case "enter":
			i, ok := m.menu.SelectedItem().(component.MenuItem)
			if ok {
				m.selectedItem = i
				// Redirect to the selected page
			}
			return m, nil
		}
	}

	newMenuModel, cmd := m.menu.Update(msg)
	m.menu = newMenuModel
	commands = append(commands, cmd)

	return m, tea.Batch(commands...)
}

func (m HomePageModel) View() string {
	str := homePageStyle.Render(m.menu.View())
	return str
}
