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
	for {
		select {
		case signalData := <-c:
			signalData.status = validateSignal(signalData.signal)
			fmt.Println("Receiving Signal ", signalData.signal, " with status ", signalData.status)
		case <-time.After(time.Second * 1):
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
	// if the signal begin with 1 and end with 1, it indicates the signal is good
	// if the signal begin with 0 and end with 0, it inditates the signal is bad
	// if doesn't meet 2 conditions above, evaluate from all node

	switch {
	case signal[0] == 1 && signal[maxDigit-1] == 1:
		return "good"
	case signal[0] == 0 && signal[maxDigit-1] == 0:
		return "bad"
	default:
		return evaluateSignal(signal)
	}
}

func evaluateSignal(signal []int) string {
	var result int
	for i := range signal {
		if i == 0 {
			result = evaluateNode(result, signal[i], signal[i+1])
		} else if i == maxDigit-1 {
			break
		} else {
			result = evaluateNode(result, signal[i], signal[i+1])
		}
	}

	if result == 1 {
		return "good"
	} else {
		return "bad"
	}
}

func evaluateNode(result, node1, node2 int) int {
	switch {
	case node2 > node1:
		return 1
	case node2 < node1:
		return 0
	default:
		return result
	}
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}
