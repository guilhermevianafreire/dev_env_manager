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
	selectedItem component.MenuItem
	width        int
	height       int
}

func NewHomePage(width int, height int) HomePageModel {
	return HomePageModel{
		menu:   NewMenu(width, height),
		width:  width,
		height: height,
	}
}

func NewMenu(width int, height int) list.Model {
	menuItems := make([]list.Item, 5)

	menuItems[0] = component.NewMenuItemDestination("Environment", "Manage the development environment", component.EnvironmentSoftware)
	menuItems[1] = component.NewMenuItemDestination("Infrastructure", "Manage the development environment infrastructure", component.Infrastructure)
	menuItems[2] = component.NewMenuItemDestination("Configuration", "Configure this application", component.Configuration)
	menuItems[3] = component.NewMenuItemDestination("Application", "Manage the development environment applications", component.Application)
	menuItems[4] = component.NewMenuItemDestination("Manual", "Application manual", component.Manual)

	delegate := list.NewDefaultDelegate()

	menu := list.New(menuItems, delegate, width, height)

	menu.Title = "Development EnvironmentSoftware Manager"
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
		case component.EnterKey:
			i, ok := m.menu.SelectedItem().(component.MenuItem)
			if ok {
				m.selectedItem = i
				switch m.selectedItem.Destination() {
				case component.EnvironmentSoftware:
					return NewEnvironmentSoftwarePage(m.width, m.height), nil
				}
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
	str := environmentActionsMenuPageStyle.Render(m.menu.View())
	return str
}
