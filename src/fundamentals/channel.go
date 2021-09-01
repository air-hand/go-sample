package fundamentals

import (
	"math/rand"
	"time"
)

func randomNumber(maxNum int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(maxNum)
	return value
}

func calculateValue(intChan chan int, maxNum int) {
	number := randomNumber(maxNum)
	intChan <- number
}

func ChannelSample(maxNum int) int {
	intChan := make(chan int)
	defer close(intChan) // close resource when this func ends

	go calculateValue(intChan, maxNum)
	num := <-intChan // listen the channel
	return num
}
