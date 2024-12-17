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

var conn *net.UDPConn

var Logger *govec.GoLog
var opts govec.GoLogOptions

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func receive() string {
	fmt.Println(id, "Receiving", conn)

	buffer := make([]byte, 1024)
	var msg []byte
	var strMsg string

	res, err := conn.Read(buffer)
	checkError(err)
	fmt.Println("conn.Read", res)

	Logger.UnpackReceive("Prejeto sporocilo", buffer, &msg, opts)

	if len(msg) == 0 {
		strMsg = ""
	} else {
		strMsg = string(rune(msg[0]))
	}

	fmt.Println("Prejeto sporocilo " + strMsg)
	return strMsg
}

func send(addr *net.UDPAddr, msg int) {
	fmt.Println(id, "Sending", msg)

	// Odpremo povezavo
	sendConn, err := net.DialUDP("udp", nil, addr)
	checkError(err)
	defer sendConn.Close()
	// Pripravimo sporočilo

	Logger.LogLocalEvent("Priprava sporocila", opts)
	sMsg := strconv.Itoa(msg)
	sMsgVC := Logger.PrepareSend("Poslano sporocilo "+sMsg, []byte(sMsg), opts)
	sendConn.Write(sMsgVC)
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
		time.Sleep(500 * time.Millisecond)
		arr := getRandomNumbers(numOfProcesses, spread)

		for _, pid := range arr {
			curPort := port + pid
			addr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("localhost:%d", curPort))
			send(addr, i)
		}
	}
}

func normalProcess(port, numOfProcesses, spread int) {
	addr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("localhost:%d", (port+id)))
	timeout := time.After(5 * time.Second)

	udpConn, err := net.ListenUDP("udp", addr)
	conn = udpConn
	fmt.Println(id, "Conn created", conn)
	checkError(err)
	deadline := time.Now().Add(5 * time.Second)
	conn.SetDeadline(deadline)
	defer conn.Close()

	localMap := make(map[string]bool)
	for {
		select {
		case <-timeout:
			fmt.Println("Timeout reached, stopping the loop. Process: ", id)
			return
		default:
			fmt.Println(addr)
			msg := receive()

			if len(msg) == 0 {
				continue
			}
			fmt.Println("returnd", msg)

			_, ok := localMap[msg]
			if !ok {
				arr := getRandomNumbers(numOfProcesses, spread)

				for _, pid := range arr {
					addr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("localhost:%d", (port+pid)))
					intMsg, err := strconv.Atoi(msg)
					checkError(err)
					send(addr, intMsg)
				}
				localMap[msg] = true
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
	id = *processId
	Logger = govec.InitGoVector("Process-"+strconv.Itoa(id), "Log-Process-"+strconv.Itoa(id), govec.GetDefaultConfig())
	opts = govec.GetDefaultLogOptions()

	if *processId == 0 {
		mainProcess(*portPtr, *numOfProcesses, *numOfMessages, *spread)
		fmt.Println("Glavni proces")
	} else {
		normalProcess(*portPtr, *numOfProcesses, *spread)
	}
}
