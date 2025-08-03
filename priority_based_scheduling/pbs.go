package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var n int
var MinimumArrivalTime int = math.MaxInt32
var arrival_time []int
var burst_time []int
var priority []int
var priorityobject []PriorityObject
var po []PO
var GranntObj []Grantt

type PriorityObject struct {
	Arrival  int
	Burst    int
	Priority int
	Process  string
}

type PO struct {
	PriorityObj PriorityObject
	Accessed    bool
}

type Grantt struct {
	priorityobject PriorityObject
	ScheduleTime   int
}

func PrintGranttChart() {
	for i := range n {
		fmt.Printf("{%s, %d}", GranntObj[i].priorityobject.Process, GranntObj[i].ScheduleTime)
		if i < n-1 {
			fmt.Print(", ")
		} else {
			fmt.Print("\n")
		}
	}
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func ScheduleProcess() (int, int) {
	// var tempPriorityObject []PriorityObject
	var ans int = -1
	var index int = -1
	for i := range n {
		if !po[i].Accessed && po[i].PriorityObj.Arrival < MinimumArrivalTime && po[i].PriorityObj.Priority > ans {
			ans = po[i].PriorityObj.Priority
			index = i
		}
	}

	if index == -1 {
		return MinimumArrivalTime + 1, 0
	}

	GranntObj = append(GranntObj, Grantt{po[index].PriorityObj, MinimumArrivalTime})
	po[index].Accessed = true

	return MinimumArrivalTime + po[index].PriorityObj.Burst, 1
}

func PBS_Implementation() {
	AccessedProcesses := 0
	var temp int
	for AccessedProcesses < n {
		temp = 0
		MinimumArrivalTime, temp = ScheduleProcess()
		AccessedProcesses += temp
	}

	PrintGranttChart()
}

func PBS() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter n: ")
	scanner.Scan()
	nStr := scanner.Text()
	n, _ = strconv.Atoi(strings.TrimSpace(nStr))
	fmt.Printf("Enter Arrival Time for each %d processes: ", n)
	scanner.Scan()
	at := scanner.Text()
	for _, val := range strings.Fields(at) {
		v, _ := strconv.Atoi(strings.TrimSpace(val))
		MinimumArrivalTime = Min(MinimumArrivalTime, v)
		arrival_time = append(arrival_time, v)
	}

	fmt.Printf("Enter Burst Time for %d processes: ", n)
	scanner.Scan()
	bt := scanner.Text()
	for _, val := range strings.Fields(bt) {
		v, _ := strconv.Atoi(strings.TrimSpace(val))
		burst_time = append(burst_time, v)
	}

	fmt.Printf("Enter Priority for %d processes: ", n)
	scanner.Scan()
	pt := scanner.Text()
	for _, val := range strings.Fields(pt) {
		v, _ := strconv.Atoi(strings.TrimSpace(val))
		priority = append(priority, v)
	}

	for i := range n {
		var prs string = "P" + strconv.Itoa(i+1)
		var temp_priorityobject *PriorityObject = &PriorityObject{arrival_time[i], burst_time[i], priority[i], prs}
		priorityobject = append(priorityobject, *temp_priorityobject)
		po = append(po, PO{*temp_priorityobject, false})
	}

	PBS_Implementation()
}

func main() {
	PBS()
}
