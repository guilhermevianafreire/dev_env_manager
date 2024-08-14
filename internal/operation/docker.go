package operation

import (
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

func DockerInstall() []tea.Cmd {
	tea.Batch()
	return []tea.Cmd{
		DockerInstallInit,
		DockerInstallPart1,
		DockerInstallPart2,
		DockerInstallPart3,
		DockerInstallPart4,
	}
}

func DockerUpdate() []tea.Cmd {
	return []tea.Cmd{}
}

func DockerReinstall() []tea.Cmd {
	return []tea.Cmd{}
}

func DockerUninstall() []tea.Cmd {
	return []tea.Cmd{}
}

func DockerValidate() []tea.Cmd {
	return []tea.Cmd{}
}

func DockerInstallInit() tea.Msg {
	return ProgressMessage{
		message:         "Initializing",
		detail:          "Initializing Docker Install",
		progress:        .01,
		executing:       true,
		finishedError:   false,
		finishedSuccess: false,
	}
}

func DockerInstallPart1() tea.Msg {
	time.Sleep(3 * time.Second)
	return ProgressMessage{
		message:         "Parte 1",
		detail:          "Docker Install Parte 1",
		progress:        .2,
		executing:       true,
		finishedError:   false,
		finishedSuccess: false,
	}
}

func DockerInstallPart2() tea.Msg {
	time.Sleep(3 * time.Second)
	return ProgressMessage{
		message:         "Parte 2",
		detail:          "Docker Install Parte 2",
		progress:        .5,
		executing:       true,
		finishedError:   false,
		finishedSuccess: false,
	}
}

func DockerInstallPart3() tea.Msg {
	time.Sleep(3 * time.Second)
	return ProgressMessage{
		message:         "Parte 3",
		detail:          "Docker Install Parte 3",
		progress:        .8,
		executing:       true,
		finishedError:   false,
		finishedSuccess: false,
	}
}

func DockerInstallPart4() tea.Msg {
	time.Sleep(3 * time.Second)
	return ProgressMessage{
		message:         "Parte 4",
		detail:          "Docker Install Parte 4",
		progress:        1.0,
		executing:       true,
		finishedError:   false,
		finishedSuccess: true,
	}
}
