package advanced

import (
	"fmt"
	"sync"
	"time"
)

var channelWg = sync.WaitGroup{}

func Channels_Concept(){
	ch := make(chan int)
	channelWg.Add(2)

	go func(){
		i := <- ch // Receiving data from the channel
		fmt.Println(i)
		channelWg.Done()
	}()

	go func(){
		i := 42
		ch <- i  // Sending date for the channel
		i = 27
		channelWg.Done()
	}()

	channelWg.Wait()
}

func Channels_ReceiveAndSendingDataOnly(){
	ch := make(chan int, 50)
	channelWg.Add(2)

	go func(ch <- chan int){ // Receive Only
		for {
			if i, ok := <- ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		channelWg.Done()
	}(ch)

	go func(ch chan <- int){ // Send Only
		ch <- 42
		ch <- 27
		close(ch)
		channelWg.Done()
	}(ch)

	channelWg.Wait()
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{})

const (
	logInfo = "INFO"
	logWarning = "WARNING"
	logError = "ERROR"
)

type logEntry struct {
	time time.Time
	severity string
	message string
}

func Channels_SelectStatements(){
	go logger()
	logCh <- logEntry{time: time.Now(), severity: logInfo, message: "App is starting"}

	logCh <- logEntry{time: time.Now(), severity: logInfo, message: "App is shutting down"}

	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{}
}

func logger(){
	for {
		select {
		case entry := <- logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <- doneCh:
			break
		}
	}
}