package page

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/guilhermevianafreire/dev_env_manager/internal/component"
)

var (
	environmentMenuPageStyle = lipgloss.NewStyle()
	environmentMenuStyle     = lipgloss.NewStyle().
					Background(lipgloss.Color("26")).
					Foreground(lipgloss.Color("196")).
					Padding(0, 1)
)

type EnvironmentMenuPageModel struct {
	menu         list.Model
	selectedItem list.Item
	width        int
	height       int
}

func NewEnvironmentPage(width int, height int) EnvironmentMenuPageModel {
	return EnvironmentMenuPageModel{
		menu:   NewEnvironmentMenu(width, height),
		width:  width,
		height: height,
	}
}

func NewEnvironmentMenu(width int, height int) list.Model {
	menuItems := make([]list.Item, 5)

	menuItems[0] = component.NewMenuItem("Install", "Install applications", component.Environment)
	menuItems[1] = component.NewMenuItem("Update", "Update applications", component.Infrastructure)
	menuItems[2] = component.NewMenuItem("Reinstall", "Reinstall applications", component.Configuration)
	menuItems[3] = component.NewMenuItem("Uninstall", "Uninstall applications", component.Application)
	menuItems[4] = component.NewMenuItem("Validate", "Validate applications", component.Application)

	delegate := list.NewDefaultDelegate()

	menu := list.New(menuItems, delegate, width, height)

	menu.Title = "Environment"
	menu.Styles.Title = environmentMenuStyle

	menu.AdditionalFullHelpKeys = additionalFullHelpKeys

	return menu
}

func additionalFullHelpKeys() []key.Binding {
	keys := make([]key.Binding, 2)

	keys[0] = key.NewBinding(
		key.WithKeys("backspace"),
		key.WithHelp("backspace", "go back"),
	)

	keys[1] = key.NewBinding(
		key.WithKeys("shift+backspace"),
		key.WithHelp("shift+backspace", "go back to the main menu"),
	)

	return keys
}

func (m EnvironmentMenuPageModel) Init() tea.Cmd {
	return nil
}

func (m EnvironmentMenuPageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var commands []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := environmentMenuPageStyle.GetFrameSize()
		m.width = msg.Width - h
		m.height = msg.Height - v
		m.menu = NewEnvironmentMenu(m.width, m.height)
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
		case "backspace":
			return NewHomePage(m.width, m.height), nil
		case "shift+backspace":
			return NewHomePage(m.width, m.height), nil
		}
	}

	newMenuModel, cmd := m.menu.Update(msg)
	m.menu = newMenuModel
	commands = append(commands, cmd)

	return m, tea.Batch(commands...)
}

func (m EnvironmentMenuPageModel) View() string {
	var debug string
	//debug = fmt.Sprintf("\nresolution (environment_menu): %d x %d", m.width, m.height)
	str := environmentMenuPageStyle.Render(m.menu.View(), debug)
	return str
}
