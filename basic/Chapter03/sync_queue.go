// main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
	"math/rand"
	"time"
)

// constants
const (
	messagePassStart = iota
	messageTicketStart
	messagePassEnd
	messageTicketEnd
)

// SyncQueue class
type SyncQueue struct {
	waitPass    int
	waitTicket  int
	playPass    bool
	playTicket  bool
	queuePass   chan int
	queueTicket chan int
	message     chan int
}

// New method initialises SyncQueue
func (SyncQueue *SyncQueue) New() {

	SyncQueue.message = make(chan int)
	SyncQueue.queuePass = make(chan int)
	SyncQueue.queueTicket = make(chan int)

	go func() {
		var message int
		for {
			select {
			case message = <-SyncQueue.message:
				switch message {
				case messagePassStart:
					SyncQueue.waitPass++
				case messagePassEnd:
					SyncQueue.playPass = false
				case messageTicketStart:
					SyncQueue.waitTicket++
				case messageTicketEnd:
					SyncQueue.playTicket = false
				}
				if SyncQueue.waitPass > 0 && SyncQueue.waitTicket > 0 && !SyncQueue.playPass && !SyncQueue.playTicket {
					SyncQueue.playPass = true
					SyncQueue.playTicket = true
					SyncQueue.waitTicket--
					SyncQueue.waitPass--
					SyncQueue.queuePass <- 1
					SyncQueue.queueTicket <- 1
				}
			}
		}
	}()
}

// StartTicketIssue starts the ticket issue
func (SyncQueue *SyncQueue) StartTicketIssue() {
	SyncQueue.message <- messageTicketStart
	<-SyncQueue.queueTicket
}

// EndTicketIssue ends the ticket issue
func (SyncQueue *SyncQueue) EndTicketIssue() {
	SyncQueue.message <- messageTicketEnd
}

// StartPass ends the Pass SyncQueue
func (SyncQueue *SyncQueue) StartPass() {
	SyncQueue.message <- messagePassStart
	<-SyncQueue.queuePass
}

// EndPass ends the Pass SyncQueue
func (SyncQueue *SyncQueue) EndPass() {
	SyncQueue.message <- messagePassEnd
}

// ticketIssue starts and ends the ticket issue
func ticketIssue(SyncQueue *SyncQueue) {
	for {
		// Sleep up to 10 seconds.
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		SyncQueue.StartTicketIssue()
		fmt.Println("Ticket Issue starts")
		// Sleep up to 2 seconds.
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println("Ticket Issue ends")
		SyncQueue.EndTicketIssue()
	}
}

// passenger method starts and ends the pass SyncQueue
func passenger(SyncQueue *SyncQueue) {
	for {
		// Sleep up to 10 seconds.
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		SyncQueue.StartPass()
		fmt.Println("  Passenger starts")
		// Sleep up to 2 seconds.
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println(" Passenger ends")

		SyncQueue.EndPass()
	}
}

// main method
func SyncQueueMain() {
	var SyncQueue *SyncQueue = &SyncQueue{}
	SyncQueue.New()
	fmt.Println(SyncQueue)
	var i int
	for i = 0; i < 10; i++ {
		go passenger(SyncQueue)
	}
	var j int
	for j = 0; j < 5; j++ {
		go ticketIssue(SyncQueue)
	}
	select {}
}
