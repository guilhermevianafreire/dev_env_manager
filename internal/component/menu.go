package component

import "fmt"

type MenuItem struct {
	title       string
	description string
	destination Destination
	software    Software
	action      Action
}

func (i MenuItem) FilterValue() string {
	return i.title
}

func (i MenuItem) Title() string {
	return i.title
}

func (i MenuItem) Description() string {
	return i.description
}

func (i MenuItem) Destination() Destination {
	return i.destination
}

func (i MenuItem) Software() Software {
	return i.software
}

func (i MenuItem) Action() Action {
	return i.action
}

func NewMenuItemDestination(title string, description string, destination Destination) MenuItem {
	return MenuItem{
		title:       title,
		description: description,
		destination: destination,
	}
}

func NewMenuItemSoftware(title string, description string, software Software) MenuItem {
	return MenuItem{
		title:       title,
		description: description,
		software:    software,
	}
}

func NewMenuItemDestinationSoftware(title string, description string, destination Destination, software Software) MenuItem {
	return MenuItem{
		title:       title,
		description: description,
		destination: destination,
		software:    software,
	}
}

func NewMenuItemDestinationSoftwareAction(title string, description string, destination Destination, software Software, action Action) MenuItem {
	return MenuItem{
		title:       title,
		description: description,
		destination: destination,
		software:    software,
		action:      action,
	}
}

func GetMenuActionDescription(action string, software *Software) string {
	return fmt.Sprintf("%s %s application", action, software.Name())
}
