package component

type Destination string

const (
	Home           Destination = "HOME"
	Environment                = "ENVIRONMENT"
	Infrastructure             = "INFRASTRUCTURE"
	Configuration              = "CONFIGURATION"
	Application                = "APPLICATION"
	Manual                     = "MANUAL"
)
