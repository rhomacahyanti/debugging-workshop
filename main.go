package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"net/http"
	_ "net/http/pprof"
)

const (
	maxDigit int = 6
)

type signalData struct {
	signal []int
	status string
}

func main() {
	c := make(chan signalData)

	go sender(c)
	go receiver(c)

	log.Println(http.ListenAndServe("localhost:6060", nil))
}

func sender(c chan signalData) {
	for {
		var signalData signalData
		signalData.signal = generateSignal()
		c <- signalData
	}
}

func receiver(c chan signalData) {
	idleTimeout := time.Second
	idleTimer := time.NewTimer(idleTimeout)
	defer idleTimer.Stop()

	for {
		idleTimer.Reset(idleTimeout)
		select {
		case signalData := <-c:
			signalData.status = validateSignal(signalData.signal)
			fmt.Println("Receiving Signal ", signalData.signal, " with status ", signalData.status)
		case <-idleTimer.C:
			fmt.Println("Got timeout while receiving the signal")
			return
		}
	}
}

func generateSignal() []int {
	var signal []int

	for i := 0; i < maxDigit; i++ {
		signal = append(signal, randInt(0, 1))
	}

	return signal
}

func validateSignal(signal []int) string {
	// in my understanding, based on the logic explained, we only need to check the last digit of the signal to get the desired outcome
	// less complexity means less code means less bug potential, saving us from future headache of debugging

	switch {
	case signal[maxDigit-1] == 1:
		return "good"
	case signal[maxDigit-1] == 0:
		return "bad"
	default:
		return "bad"
	}
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}
