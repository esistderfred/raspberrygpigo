package raspberrygpigo

// Pin emulates raspberry Pin
type Pin uint8

type Mode uint8
type State uint8
type Pull uint8
type Edge uint8

// Pin mode, a pin can be set in Input or Output, Clock or Pwm mode
const (
	Input Mode = iota
	Output
	Clock
	Pwm
	Spi
)

// State of pin, High / Low
const (
	Low State = iota
	High
)

// Pull Up / Down / Off
const (
	PullOff Pull = iota
	PullDown
	PullUp
)

// Edge events
const (
	NoEdge Edge = iota
	RiseEdge
	FallEdge
	AnyEdge = RiseEdge | FallEdge
)

// Output ...
func (a Pin) Output() {
	i := getInstance()
	i.States[a].Mode = Output
}

// High ...
func (a Pin) High() {
	i := getInstance()
	i.States[a].State = High
}

// Low ...
func (a Pin) Low() {
	i := getInstance()
	i.States[a].State = Low
}

// Toggle ...
func (a Pin) Toggle() {
	i := getInstance()
	switch i.States[a].State {
	case High:
		i.States[a].State = Low
	case Low:
		i.States[a].State = High
	}
}

// Input ...
func (a Pin) Input() {
	i := getInstance()
	i.States[a].Mode = Input
}

// Read ...
func (a Pin) Read() State {
	i := getInstance()
	return i.States[a].State
}

// Mode ...
func (a Pin) Mode(m Mode) {
	i := getInstance()
	i.States[a].Mode = m
}

// Write ...
func (a Pin) Write(s State) {
	i := getInstance()
	i.States[a].State = s
}

// PullUp ...
func (a Pin) PullUp() {
	i := getInstance()
	i.States[a].Pull = PullUp
}

// PullDown ...
func (a Pin) PullDown() {
	i := getInstance()
	i.States[a].Pull = PullDown
}

// PullOff ...
func (a Pin) PullOff() {
	i := getInstance()
	i.States[a].Pull = PullOff
}

// PullOff ...
func (a Pin) Pull(s Pull) {
	i := getInstance()
	i.States[a].Pull = s
}
