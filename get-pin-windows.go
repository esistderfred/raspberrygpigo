package raspberrygpigo

import (
	"log"
	"net/http"
)

func GetPin(p uint8) GoPin {
	return Pin(p)
}

func Open() error {
	return nil
}

func Close() error {
	return nil
}

func init() {
	go func() {
		instance := getInstance()
		log.Println("Listen to PI-Port 3414")
		http.ListenAndServe(":3414", instance)
	}()
}
