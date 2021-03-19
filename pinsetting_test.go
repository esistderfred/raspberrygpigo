package raspberrygpigo

import (
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"
)

func TestPinserver(t *testing.T) {
	i := getInstance()
	handler := i.ServeHTTP

	p10 := GetPin(10)
	p10.Output()
	p10.High()
	p1 := GetPin(1)
	p1.Input()
	s := p1.Read()
	log.Print("p1 state ", s)
	p1.High()
	s = p1.Read()
	log.Print("p1 state ", s)
	req := httptest.NewRequest("GET", "http://localhost:3414", nil)
	w := httptest.NewRecorder()
	handler(w, req)
	resp := w.Result()
	x := &emulatorboard{}
	json.NewDecoder(resp.Body).Decode(x)
	t.Log("instance:", x)
	t.Fail()
}
