package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var n int
var arrival_time []int
var burst_time []int
var priority []int

func PBS() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter n:")
	scanner.Scan()
	nStr := scanner.Text()
	n, _ = strconv.Atoi(strings.TrimSpace(nStr))
	fmt.Printf("Enter Arrival Time for each %d processes", n)
	scanner.Scan()
	at := scanner.Text()
	for _, val := range strings.Fields(at) {
		v, _ := strconv.Atoi(strings.TrimSpace(val))
		arrival_time = append(arrival_time, v)
	}

	fmt.Printf("Enter Burst Time for %d processes", n)
	scanner.Scan()
	bt := scanner.Text()
	for _, val := range strings.Fields(bt) {
		v, _ := strconv.Atoi(strings.TrimSpace(val))
		burst_time = append(burst_time, v)
	}

	fmt.Printf("Enter Priority for %d processes", n)
	scanner.Scan()
	pt := scanner.Text()
	for _, val := range strings.Fields(pt) {
		v, _ := strconv.Atoi(strings.TrimSpace(val))
		priority = append(priority, v)
	}
}

func main() {
	PBS()
}
