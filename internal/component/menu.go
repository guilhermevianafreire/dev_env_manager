package component

type MenuItem struct {
	title       string
	description string
	destination Destination
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

func NewMenuItem(title string, description string, destination Destination) MenuItem {
	return MenuItem{
		title:       title,
		description: description,
		destination: destination,
	}
}
