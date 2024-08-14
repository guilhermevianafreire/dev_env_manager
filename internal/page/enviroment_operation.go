package page

import (
	"fmt"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/guilhermevianafreire/dev_env_manager/internal/component"
	"github.com/guilhermevianafreire/dev_env_manager/internal/operation"
)

var (
	environmentOperationPageStyle = lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center)
)

type EnvironmentOperationMenuPageModel struct {
	software           component.Software
	action             component.Action
	currentOperation   string
	executing          bool
	finishedSuccess    bool
	finishedError      bool
	progressPercentage float64
	progressBar        progress.Model
	executingSpinner   spinner.Model
	error              string
	width              int
	height             int
}

func NewEnvironmentOperationPage(width int, height int, software component.Software, action component.Action) EnvironmentOperationMenuPageModel {
	return EnvironmentOperationMenuPageModel{
		software:           software,
		action:             action,
		currentOperation:   "Starting...",
		executing:          false,
		finishedError:      false,
		finishedSuccess:    false,
		progressPercentage: .0,
		progressBar:        progress.New(progress.WithWidth(width)),
		width:              width,
		height:             height,
	}
}

func (m EnvironmentOperationMenuPageModel) Init() tea.Cmd {
	var commands = m.software.Actions()[m.action.Name()]
	if len(commands) > 0 {
		return tea.Sequence(commands...)
	}
	return nil
}

func (m EnvironmentOperationMenuPageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var commands []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := environmentOperationPageStyle.GetFrameSize()
		m.width = msg.Width - h
		m.height = msg.Height - v
	case operation.ProgressMessage:
		m.currentOperation = msg.Detail()
		m.progressPercentage = msg.Progress()
		m.finishedSuccess = msg.FinishedSuccess()
		m.finishedError = msg.FinishedError()
		m.executing = msg.Executing()
		m.error = msg.Error()
	case tea.KeyMsg:
		switch msg.String() {
		case component.EnterKey:
			if !m.executing {
				var operations = m.software.Actions()[m.action.Name()]
				if len(operations) > 0 {
					commands = append(commands, operations...)
				}
			} else if m.executing && m.finishedSuccess {
				return NewEnvironmentActionsPage(m.width, m.height, m.software), nil
			} else if m.executing && m.finishedError {
				return NewEnvironmentActionsPage(m.width, m.height, m.software), nil
			}
		case component.BackspaceKey:
			return NewEnvironmentActionsPage(m.width, m.height, m.software), nil
		}
	}

	newProgressModel, cmd := m.progressBar.Update(msg)
	m.progressBar = newProgressModel.(progress.Model)
	commands = append(commands, cmd)

	return m, tea.Sequence(commands...)
}

func (m EnvironmentOperationMenuPageModel) View() string {
	var str = environmentOperationPageStyle.Render("Press Enter to start the operation or Backspace to cancel")
	if m.executing && !m.finishedSuccess && !m.finishedError {
		str = environmentOperationPageStyle.Render(lipgloss.JoinVertical(lipgloss.Center, m.currentOperation, lipgloss.JoinHorizontal(lipgloss.Center, fmt.Sprintf("executing: %v", m.executing), " | ", fmt.Sprintf("finishedSuccess: %v", m.finishedSuccess), " | ", fmt.Sprintf("finishedError: %v", m.finishedError)), m.progressBar.ViewAs(m.progressPercentage)))
	} else if m.executing && m.finishedSuccess {
		str = environmentOperationPageStyle.Render("Operation finished with success press Enter to go back")
	} else if m.executing && m.finishedError {
		str = environmentOperationPageStyle.Render(fmt.Sprintf("Operation finished with error: %s press Enter to go back", m.error))
	}
	return str
}
