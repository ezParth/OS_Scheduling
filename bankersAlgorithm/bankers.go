package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var maxNeeded []int
var allocation []int
var needed []int
var process []Process
var available int
var n int
var ans []string

type Process struct {
	maxNeeded  int
	allocated  int
	needed     int
	done       bool
	processStr string
}

func BankerAlgorithm() bool {
	for i := 0; i < n; i++ {
		if !CheckDeadLock() {
			return false
		}
	}

	for i := range ans {
		print(i, " ")
	}
	return true
}

func CheckDeadLock() bool {
	for i := 0; i < n; i++ {
		if !process[i].done && process[i].needed <= available {
			process[i].done = true
			available += process[i].allocated
			ans = append(ans, process[i].processStr)
			return true
		}
	}
	return false
}

func Bankers() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter n: ")
	scanner.Scan()
	nStr := scanner.Text()
	n, _ = strconv.Atoi(strings.TrimSpace(nStr))
	fmt.Printf("Enter max required by %d processes: ", n)
	scanner.Scan()
	maxi := scanner.Text()
	for _, val := range strings.Fields(maxi) {
		v, _ := strconv.Atoi(strings.TrimSpace(val))
		maxNeeded = append(maxNeeded, v)
	}

	fmt.Printf("Enter allocation for %d processes: ", n)
	scanner.Scan()
	allocateString := scanner.Text()
	for _, val := range strings.Fields(allocateString) {
		v, _ := strconv.Atoi(strings.TrimSpace(val))
		allocation = append(allocation, v)
	}

	for i := range n {
		p := strconv.Itoa(i + 1)
		needed = append(needed, maxNeeded[i]-allocation[i])
		process = append(process, Process{maxNeeded: maxNeeded[i], allocated: allocation[i], needed: needed[i], done: false, processStr: "P" + p})
	}
	DeadLock := BankerAlgorithm()
	if !DeadLock {
		fmt.Println("DeadLock is Present!")
	}
}

func main() {
	Bankers()
}
