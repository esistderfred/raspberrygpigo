package raspberrygpigo

type GoPin interface {
	Output()
	High()
	Low()
	Toggle()

	Input()
	Read() State

	Mode(Mode)
	Write(State)

	PullUp()
	PullDown()
	PullOff()

	Pull(Pull)
}
