package raspberrygpigo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type pinsetting struct {
	State State
	Mode  Mode
	Pull  Pull
}

func (p *pinsetting) String() string {
	sw := &bytes.Buffer{}
	mode_str := ""
	switch p.Mode {
	case Input:
		mode_str = "Input"
	case Output:
		mode_str = "Output"
	}

	state_string := ""
	switch p.State {
	case High:
		state_string = "High"
	case Low:
		state_string = "Low"
	}

	fmt.Fprintf(sw, "%s %s", mode_str, state_string)

	return sw.String()
}

type emulatorboard struct {
	States [27]pinsetting `json:"states"`
}

func (e *emulatorboard) String() string {
	sw := &bytes.Buffer{}

	for i, pin := range e.States {
		fmt.Fprintln(sw, i+1, &pin)
	}

	return sw.String()
}

var (
	instance *emulatorboard
	lock     = &sync.Mutex{}
)

func getInstance() *emulatorboard {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = &emulatorboard{}
	}

	return instance
}

// ServeHTTP implements default handler
func (e *emulatorboard) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(e)
	}
}
