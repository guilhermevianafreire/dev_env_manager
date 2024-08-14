package component

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/guilhermevianafreire/dev_env_manager/internal/operation"
)

type Destination string

type Action struct {
	name        string
	description string
}

func (a *Action) Name() string {
	return a.name
}

func (a *Action) Description() string {
	return a.description
}

type Software struct {
	code        string
	name        string
	description string
	actions     map[string][]tea.Cmd
}

func (s *Software) Code() string {
	return s.code
}

func (s *Software) Name() string {
	return s.name
}

func (s *Software) Description() string {
	return s.description
}

func (s *Software) Actions() map[string][]tea.Cmd {
	return s.actions
}

// Navigation destinations
const (
	EnvironmentSoftware  Destination = "ENVIRONMENT_SOFTWARE"
	EnvironmentActions               = "ENVIRONMENT_ACTIONS"
	EnvironmentOperation             = "ENVIRONMENT_OPERATION"

	Infrastructure = "INFRASTRUCTURE"
	Configuration  = "CONFIGURATION"
	Application    = "APPLICATION"
	Manual         = "MANUAL"
)

const (
	EnterKey     string = "enter"
	QKey                = "q"
	ControlCKey         = "ctrl+c"
	BackspaceKey        = "backspace"
	HomeKey             = "home"
)

const (
	ActionNameInstall   string = "Install"
	ActionNameUpdate           = "Update"
	ActionNameReinstall        = "Reinstall"
	ActionNameUninstall        = "Uninstall"
	ActionNameValidate         = "Validate"
)

const (
	SoftwareCodeDocker        string = "DOCKER"
	SoftwareNameDocker        string = "Docker"
	SoftwareDescriptionDocker string = "Docker is a platform designed to help developers build, share, and run container applications."

	SoftwareCodeJdk        string = "JDK"
	SoftwareNameJdk        string = "Java Development Kit"
	SoftwareDescriptionJdk string = "Java Development Kit is a software development kit used to develop Java applications."

	SoftwareCodeMaven        string = "MAVEN"
	SoftwareNameMaven        string = "Apache Maven"
	SoftwareDescriptionMaven string = "Apache Maven can manage a project's build, reporting and documentation from a central piece of information."

	SoftwareCodeVsCode        string = "VSCODE"
	SoftwareNameVsCode        string = "Visual Studio Code"
	SoftwareDescriptionVsCode string = "Visual Studio Code is a source code editor developed by Microsoft."

	SoftwareCodeIntelliJCom        string = "INTELLIJ_COMMUNITY"
	SoftwareNameIntelliJCom        string = "IntelliJ Community"
	SoftwareDescriptionIntelliJCom string = "IntelliJ Community is an open-source integrated development environment."

	SoftwareCodeIntelliJUlt        string = "INTELLIJ_ULTIMATE"
	SoftwareNameIntelliJUlt        string = "IntelliJ Ultimate"
	SoftwareDescriptionIntelliJUlt string = "IntelliJ Ultimate is a commercial integrated development environment."

	SoftwareCodeEclipse        string = "ECLIPSE"
	SoftwareNameEclipse        string = "Eclipse"
	SoftwareDescriptionEclipse string = "Eclipse is an integrated development environment used in computer programming."

	SoftwareCodePostman        string = "POSTMAN"
	SoftwareNamePostman        string = "Postman"
	SoftwareDescriptionPostman string = "Postman is a collaboration platform for API development."

	SoftwareCodeGit        string = "GIT"
	SoftwareNameGit        string = "Git"
	SoftwareDescriptionGit string = "Git is a distributed version control system."

	SoftwareCodeDbeaver        string = "DBEAVER"
	SoftwareNameDbeaver        string = "DBeaver"
	SoftwareDescriptionDbeaver string = "DBeaver is a free multi-platform database tool for developers, SQL programmers, database administrators and analysts."
)

func AdditionalFullHelpKeys() []key.Binding {
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

func SoftwareAll() []Software {
	return []Software{
		SoftwareDocker(),
		SoftwareJdk(),
		SoftwareMaven(),
		SoftwareVsCode(),
		SoftwareIntelliJCom(),
		SoftwareIntelliJUlt(),
		SoftwareEclipse(),
		SoftwarePostman(),
		SoftwareGit(),
		SoftwareDbeaver(),
	}
}

func SoftwareDocker() Software {
	return Software{
		code:        SoftwareCodeDocker,
		name:        SoftwareNameDocker,
		description: SoftwareDescriptionDocker,
		actions: map[string][]tea.Cmd{
			ActionNameInstall:   operation.DockerInstall(),
			ActionNameUpdate:    operation.DockerUpdate(),
			ActionNameReinstall: operation.DockerReinstall(),
			ActionNameUninstall: operation.DockerUninstall(),
			ActionNameValidate:  operation.DockerValidate(),
		},
	}
}

func SoftwareJdk() Software {
	return Software{
		code:        SoftwareCodeJdk,
		name:        SoftwareNameJdk,
		description: SoftwareDescriptionJdk,
		actions: map[string][]tea.Cmd{
			ActionNameInstall:   operation.JdkInstall(),
			ActionNameUpdate:    operation.JdkUpdate(),
			ActionNameReinstall: operation.JdkReinstall(),
			ActionNameUninstall: operation.JdkUninstall(),
			ActionNameValidate:  operation.JdkValidate(),
		},
	}
}

func SoftwareMaven() Software {
	return Software{
		code:        SoftwareCodeMaven,
		name:        SoftwareNameMaven,
		description: SoftwareDescriptionMaven,
		actions: map[string][]tea.Cmd{
			ActionNameInstall:   operation.MavenInstall(),
			ActionNameUpdate:    operation.MavenUpdate(),
			ActionNameReinstall: operation.MavenReinstall(),
			ActionNameUninstall: operation.MavenUninstall(),
			ActionNameValidate:  operation.MavenValidate(),
		},
	}
}

func SoftwareVsCode() Software {
	return Software{
		code:        SoftwareCodeVsCode,
		name:        SoftwareNameVsCode,
		description: SoftwareDescriptionVsCode,
		actions: map[string][]tea.Cmd{
			ActionNameInstall:   operation.VsCodeInstall(),
			ActionNameUpdate:    operation.VsCodeUpdate(),
			ActionNameReinstall: operation.VsCodeReinstall(),
			ActionNameUninstall: operation.VsCodeUninstall(),
			ActionNameValidate:  operation.VsCodeValidate(),
		},
	}
}

func SoftwareIntelliJCom() Software {
	return Software{
		code:        SoftwareCodeIntelliJCom,
		name:        SoftwareNameIntelliJCom,
		description: SoftwareDescriptionIntelliJCom,
		actions: map[string][]tea.Cmd{
			ActionNameInstall:   operation.IntelliJComInstall(),
			ActionNameUpdate:    operation.IntelliJComUpdate(),
			ActionNameReinstall: operation.IntelliJComReinstall(),
			ActionNameUninstall: operation.IntelliJComUninstall(),
			ActionNameValidate:  operation.IntelliJComValidate(),
		},
	}
}

func SoftwareIntelliJUlt() Software {
	return Software{
		code:        SoftwareCodeIntelliJUlt,
		name:        SoftwareNameIntelliJUlt,
		description: SoftwareDescriptionIntelliJUlt,
		actions: map[string][]tea.Cmd{
			ActionNameInstall:   operation.IntelliJUltInstall(),
			ActionNameUpdate:    operation.IntelliJUltUpdate(),
			ActionNameReinstall: operation.IntelliJUltReinstall(),
			ActionNameUninstall: operation.IntelliJUltUninstall(),
			ActionNameValidate:  operation.IntelliJUltValidate(),
		},
	}
}

func SoftwareEclipse() Software {
	return Software{
		code:        SoftwareCodeEclipse,
		name:        SoftwareNameEclipse,
		description: SoftwareDescriptionEclipse,
		actions: map[string][]tea.Cmd{
			ActionNameInstall:   operation.EclipseInstall(),
			ActionNameUpdate:    operation.EclipseUpdate(),
			ActionNameReinstall: operation.EclipseReinstall(),
			ActionNameUninstall: operation.EclipseUninstall(),
			ActionNameValidate:  operation.EclipseValidate(),
		},
	}
}

func SoftwarePostman() Software {
	return Software{
		code:        SoftwareCodePostman,
		name:        SoftwareNamePostman,
		description: SoftwareDescriptionPostman,
		actions: map[string][]tea.Cmd{
			ActionNameInstall:   operation.PostmanInstall(),
			ActionNameUpdate:    operation.PostmanUpdate(),
			ActionNameReinstall: operation.PostmanReinstall(),
			ActionNameUninstall: operation.PostmanUninstall(),
			ActionNameValidate:  operation.PostmanValidate(),
		},
	}
}

func SoftwareGit() Software {
	return Software{
		code:        SoftwareCodeGit,
		name:        SoftwareNameGit,
		description: SoftwareDescriptionGit,
		actions: map[string][]tea.Cmd{
			ActionNameInstall:   operation.GitInstall(),
			ActionNameUpdate:    operation.GitUpdate(),
			ActionNameReinstall: operation.GitReinstall(),
			ActionNameUninstall: operation.GitUninstall(),
			ActionNameValidate:  operation.GitValidate(),
		},
	}
}

func SoftwareDbeaver() Software {
	return Software{
		code:        SoftwareCodeDbeaver,
		name:        SoftwareNameDbeaver,
		description: SoftwareDescriptionDbeaver,
		actions: map[string][]tea.Cmd{
			ActionNameInstall:   operation.DbeaverInstall(),
			ActionNameUpdate:    operation.DbeaverUpdate(),
			ActionNameReinstall: operation.DbeaverReinstall(),
			ActionNameUninstall: operation.DbeaverUninstall(),
			ActionNameValidate:  operation.DbeaverValidate(),
		},
	}
}

func ActionAll(software *Software) []Action {
	return []Action{
		ActionInstall(software),
		ActionUpdate(software),
		ActionReinstall(software),
		ActionUninstall(software),
		ActionValidate(software),
	}
}

func ActionInstall(software *Software) Action {
	return Action{
		name:        ActionNameInstall,
		description: GetMenuActionDescription(ActionNameInstall, software),
	}
}

func ActionUpdate(software *Software) Action {
	return Action{
		name:        ActionNameUpdate,
		description: GetMenuActionDescription(ActionNameUpdate, software),
	}
}

func ActionReinstall(software *Software) Action {
	return Action{
		name:        ActionNameReinstall,
		description: GetMenuActionDescription(ActionNameReinstall, software),
	}
}

func ActionUninstall(software *Software) Action {
	return Action{
		name:        ActionNameUninstall,
		description: GetMenuActionDescription(ActionNameUninstall, software),
	}
}

func ActionValidate(software *Software) Action {
	return Action{
		name:        ActionNameValidate,
		description: GetMenuActionDescription(ActionNameValidate, software),
	}
}
