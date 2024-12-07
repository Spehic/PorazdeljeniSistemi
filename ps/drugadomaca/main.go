/*
Primer uporabe paketa socialNetwork
*/
package main

import (
	"fmt"
	"drugadomaca/socialNetwork"
	"time"
	"math"
	"sync"
	"strings"
	"unicode"
	"flag"
)


var producer socialNetwork.Q

var curGoroutines int
var maxGoroutines int

var invIndex map[string][]uint64 = make(map[string][]uint64)
var lock sync.Mutex
var killWorkerChan = make(chan int, 100)
var killControllerChan = make(chan int)

func validCharacter(r rune) rune {
	if unicode.IsLetter(r) || unicode.IsNumber(r) {
		return unicode.ToLower(r)
	}

	return ' '
}

//pocisti niz v skladu z mojimi zeljami v funkciji validCharacter
func cleanString(s string) string {
	return strings.Map( validCharacter, s );	
}

// worker goroutine works until it reads from exitChan
func worker() {
	for {
		select {
			case <-killWorkerChan:
				return
			default:
				process()
		}
	}
}

// this function processes the requests
func process() {
	task := <- producer.TaskChan
	

	clean := cleanString(task.Data)
	words := strings.Split( clean , " " )
	
	//fmt.Println( words, task.Id )
	
	for _, word := range words {
			if len( word ) < 4 {
				continue;
			}
			
			//fmt.Println( word, task.Id )
			lock.Lock()
			invIndex[ word ] = append( invIndex[word], task.Id)
			lock.Unlock()
	}

	//fmt.Println( words )
	return
}

func controller() {
	for {
		select {
			case <-killControllerChan:
				cleanUp()
				return
			default:
				adjust()	
		}
	}
	return
}

// destroys all workers 
func cleanUp() {
	// wait until queue is empty
	for !producer.QueueEmpty() { 
	}

	// then kill all remaining workers
	for i:= 0; i < curGoroutines; i++ {
		killWorkerChan <- 0
	}
}

func adjust() {
	expected := expectedGoroutines()
	diff := curGoroutines - expected
	fmt.Println( "current:", curGoroutines, "expected:" ,expected, "--------------------------------------------", "diff:", diff)
	if ( diff <= 0 ) {
		for i:= 0; i < -diff; i++ {
			curGoroutines += 1
			go worker()
		}
	} else {
		for i:= 0; i < diff; i++ {
			curGoroutines -= 1
			killWorkerChan <- 0
		}
	}

	time.Sleep(time.Microsecond * 10000)
}

func expectedGoroutines() int {
	curBuf := len( producer.TaskChan )
	full := float64( curBuf ) / float64(10000.0)
	
	//fmt.Println("Full", full)

	if full > 0.5 {
		return maxGoroutines
	}


	return int(math.Max( full * float64( maxGoroutines ), 1.0))
}

func main() {
	// Definiramo nov generator
	// Inicializiramo generator. Parameter določa zakasnitev med zahtevki
	
	goNums := flag.Int("g", 8, "num of goroutines")
	difficulty := flag.Int("d", 5000, "difficulty")
	flag.Parse()

	producer.New(*difficulty)
	maxGoroutines = *goNums
	start := time.Now()

	// Zazenemo prvega delavca
	curGoroutines = 1
	go worker()

	go controller()

	// Zaženemo generator
	go producer.Run()
	// Počakamo 5 sekund
	time.Sleep(time.Second * 5)
	// Ustavimo generator
	producer.Stop()
	// Počakamo, da se vrsta sprazni
	for !producer.QueueEmpty() {
	}
	
	//fmt.Println("GenStopped, killing controller")
	//fmt.Println(invIndex["hard"])
	killControllerChan <- 0

	elapsed := time.Since(start)
	// Izpišemo število generiranih zahtevkov na sekundo
	fmt.Printf("Processing rate: %f MReqs/s\n", float64(producer.N)/float64(elapsed.Seconds())/1000000.0)
	// Izpišemo povprečno dolžino vrste v čakalnici
	fmt.Printf("Average queue length: %.2f %%\n", producer.GetAverageQueueLength())
	// Izpišemo največjo dolžino vrste v čakalnici
	fmt.Printf("Max queue length %.2f %%\n", producer.GetMaxQueueLength())

	return
}

