package raspberrygpigo

import (
	rpio "github.com/stianeikeland/go-rpio/v4"
)

type Pin uint8

type GoPin interface {
	Output()
	High()
	Low()
	Toggle()

	Input()
	Read() rpio.State
}
