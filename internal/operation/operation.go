package operation

type ProgressMessage struct {
	message         string
	detail          string
	error           string
	progress        float64
	executing       bool
	finishedSuccess bool
	finishedError   bool
}

func (p *ProgressMessage) Message() string {
	return p.message
}

func (p *ProgressMessage) Detail() string {
	return p.detail
}

func (p *ProgressMessage) Error() string {
	return p.error
}

func (p *ProgressMessage) Progress() float64 {
	return p.progress
}

func (p *ProgressMessage) Executing() bool {
	return p.executing
}

func (p *ProgressMessage) FinishedSuccess() bool {
	return p.finishedSuccess
}

func (p *ProgressMessage) FinishedError() bool {
	return p.finishedError
}
