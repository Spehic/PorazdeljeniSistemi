package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"slices"
	"strconv"
	"time"

	"github.com/DistributedClocks/GoVector/govec"
)

var start chan bool
var stopHeartbeat bool
var N int
var id int

var Logger *govec.GoLog
var opts govec.GoLogOptions

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func receive(addr *net.UDPAddr) string {

	conn, err := net.ListenUDP("udp", addr)
	checkError(err)

	fmt.Println(conn)
	defer conn.Close()

	deadline := time.Now().Add(5 * time.Second)
	conn.SetDeadline(deadline)

	buffer := make([]byte, 1024)
	var msg []byte

	conn.Read(buffer)
	Logger.UnpackReceive("Prejeto sporocilo ", buffer, &msg, opts)

	if len(msg) == 0 {
		return ""
	}

	return string(rune(msg[0]))
}

func send(addr *net.UDPAddr, msg int) {
	// Odpremo povezavo
	conn, err := net.DialUDP("udp", nil, addr)
	checkError(err)
	defer conn.Close()
	// Pripravimo sporočilo

	Logger.LogLocalEvent("Priprava sporocila", opts)
	sMsg := strconv.Itoa(msg)
	sMsgVC := Logger.PrepareSend("Poslano sporocilo ", []byte(sMsg), opts)
	conn.Write(sMsgVC)
	fmt.Println("Proces", id, "poslal sporočilo", sMsg, "procesu na naslovu", addr)
	//fmt.Println("endsend", id)
}

func getRandomNumbers(numOfProcesses, spread int) []int {
	arr := make([]int, 0)
	for i := 0; i < spread; i++ {
		guess := rand.Intn(numOfProcesses-1) + 1
		if !slices.Contains(arr, guess) {
			arr = append(arr, guess)
		} else {
			i--
		}
	}
	return arr
}

func mainProcess(port, numOfProcesses, numOfMessages, spread int) {
	for i := 0; i < numOfMessages; i++ {
		arr := getRandomNumbers(numOfProcesses, spread)

		for _, pid := range arr {
			curPort := port + pid
			addr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("localhost:%d", curPort))
			send(addr, i)
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func normalProcess(port, numOfProcesses, spread int) {
	addr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("localhost:%d", (port+id)))
	timeout := time.After(5 * time.Second)

	localMap := make(map[string]int)
	for {
		select {
		case <-timeout:
			fmt.Println("Timeout reached, stopping the loop. Process: ", id)
			return
		default:
			fmt.Println(addr)
			msg := receive(addr)

			if len(msg) == 0 {
				continue
			}
			fmt.Println("returnd", msg)
			if _, ok := localMap[msg]; !ok {
				arr := getRandomNumbers(numOfProcesses, spread)

				for _, pid := range arr {
					addr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("localhost:%d", (port+pid)))
					send(addr, pid)
				}
			}
		}
	}

}

func main() {
	// Preberi argumente
	portPtr := flag.Int("p", 9000, "# start port")
	processId := flag.Int("id", 0, "# process id")
	numOfProcesses := flag.Int("n", 2, "total number of processes")
	numOfMessages := flag.Int("m", 5, "# number of messages")
	spread := flag.Int("k", 2, "# number of spread")
	flag.Parse()

	// dnevnik z vektorsko uro
	Logger = govec.InitGoVector("Process-"+strconv.Itoa(id), "Log-Process-"+strconv.Itoa(id), govec.GetDefaultConfig())
	opts = govec.GetDefaultLogOptions()

	id = *processId
	if *processId == 0 {
		mainProcess(*portPtr, *numOfProcesses, *numOfMessages, *spread)
		fmt.Println("Glavni proces")
	} else {
		normalProcess(*portPtr, *numOfProcesses, *spread)
	}
}
