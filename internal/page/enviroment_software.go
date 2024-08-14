package page

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/guilhermevianafreire/dev_env_manager/internal/component"
)

var (
	environmentSoftwareMenuPageStyle = lipgloss.NewStyle()
	environmentSoftwareMenuStyle     = lipgloss.NewStyle().
						Background(lipgloss.Color("26")).
						Foreground(lipgloss.Color("196")).
						Padding(0, 1)
)

type EnvironmentSoftwareMenuPageModel struct {
	menu         list.Model
	selectedItem component.MenuItem
	width        int
	height       int
}

func NewEnvironmentSoftwarePage(width int, height int) EnvironmentSoftwareMenuPageModel {
	return EnvironmentSoftwareMenuPageModel{
		menu:   NewEnvironmentSoftwareMenu(width, height),
		width:  width,
		height: height,
	}
}

func NewEnvironmentSoftwareMenu(width int, height int) list.Model {
	softwareArray := component.SoftwareAll()
	menuItems := make([]list.Item, len(softwareArray))

	for i, s := range softwareArray {
		menuItems[i] = component.NewMenuItemDestinationSoftware(s.Name(), s.Description(), component.EnvironmentActions, s)
	}

	delegate := list.NewDefaultDelegate()

	menu := list.New(menuItems, delegate, width, height)

	menu.Title = "Environment - Software"
	menu.Styles.Title = environmentSoftwareMenuStyle

	menu.AdditionalFullHelpKeys = component.AdditionalFullHelpKeys

	return menu
}

func (m EnvironmentSoftwareMenuPageModel) Init() tea.Cmd {
	return nil
}

func (m EnvironmentSoftwareMenuPageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var commands []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := environmentSoftwareMenuPageStyle.GetFrameSize()
		m.width = msg.Width - h
		m.height = msg.Height - v
		m.menu = NewEnvironmentSoftwareMenu(m.width, m.height)
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
				case component.EnvironmentActions:
					return NewEnvironmentActionsPage(m.width, m.height, m.selectedItem.Software()), nil
				}
			}
			return m, nil
		case component.BackspaceKey:
			return NewHomePage(m.width, m.height), nil
		}
	}

	newMenuModel, cmd := m.menu.Update(msg)
	m.menu = newMenuModel
	commands = append(commands, cmd)

	return m, tea.Batch(commands...)
}

func (m EnvironmentSoftwareMenuPageModel) View() string {
	str := environmentSoftwareMenuPageStyle.Render(m.menu.View())
	return str
}
