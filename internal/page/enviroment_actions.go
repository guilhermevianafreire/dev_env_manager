package page

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/guilhermevianafreire/dev_env_manager/internal/component"
)

var (
	environmentActionsMenuPageStyle = lipgloss.NewStyle()
	environmentActionsMenuStyle     = lipgloss.NewStyle().
					Background(lipgloss.Color("26")).
					Foreground(lipgloss.Color("196")).
					Padding(0, 1)
)

type EnvironmentActionsMenuPageModel struct {
	software     component.Software
	menu         list.Model
	selectedItem component.MenuItem
	width        int
	height       int
}

func NewEnvironmentActionsPage(width int, height int, software component.Software) EnvironmentActionsMenuPageModel {
	return EnvironmentActionsMenuPageModel{
		software: software,
		menu:     NewEnvironmentActionsMenu(width, height, software),
		width:    width,
		height:   height,
	}
}

func NewEnvironmentActionsMenu(width int, height int, software component.Software) list.Model {
	actions := component.ActionAll(&software)

	menuItems := make([]list.Item, len(actions))

	for i, a := range actions {
		menuItems[i] = component.NewMenuItemDestinationSoftwareAction(a.Name(), a.Description(), component.EnvironmentOperation, software, a)
	}

	delegate := list.NewDefaultDelegate()

	menu := list.New(menuItems, delegate, width, height)

	menu.Title = fmt.Sprintf("Environment - %s - Actions", software.Name())
	menu.Styles.Title = environmentActionsMenuStyle

	menu.AdditionalFullHelpKeys = component.AdditionalFullHelpKeys

	return menu
}

func (m EnvironmentActionsMenuPageModel) Init() tea.Cmd {
	return nil
}

func (m EnvironmentActionsMenuPageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var commands []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := environmentActionsMenuPageStyle.GetFrameSize()
		m.width = msg.Width - h
		m.height = msg.Height - v
		m.menu = NewEnvironmentActionsMenu(m.width, m.height, m.software)
	case tea.KeyMsg:
		if m.menu.FilterState() == list.Filtering {
			break
		}
		switch msg.String() {
		case component.EnterKey:
			i, ok := m.menu.SelectedItem().(component.MenuItem)
			if ok {
				m.selectedItem = i
				// Redirect to the selected page
			}
			return NewEnvironmentOperationPage(m.width, m.height, m.software, m.selectedItem.Action()), nil
		case component.BackspaceKey:
			return NewEnvironmentSoftwarePage(m.width, m.height), nil
		}
	}

	newMenuModel, cmd := m.menu.Update(msg)
	m.menu = newMenuModel
	commands = append(commands, cmd)

	return m, tea.Batch(commands...)
}

func (m EnvironmentActionsMenuPageModel) View() string {
	str := environmentActionsMenuPageStyle.Render(m.menu.View())
	return str
}
