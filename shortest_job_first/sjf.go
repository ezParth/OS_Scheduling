package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	First   int
	Second  int
	Process string
}

type Grannt struct {
	Process string
	Time    int
}

func min(a int, b int) int {
	if a > b {
		return b
	}

	return a
}

func min_element(arr []int) int {
	var ans int = arr[0]
	for _, val := range arr {
		ans = min(ans, val)
	}

	return ans
}

type Process struct {
	Arrival_Time int
	Burst_Time   int
	Process_Id   string
}

type ProcessMetaData struct {
	UnitProcess  Process
	Accessed     bool
	AccessIndex  int
	AccessedTime int
}

func create_process_id(idx int) string {
	// c := strconv.Itoa(idx + '0')
	var s string = "P" + strconv.Itoa(idx)
	return s
}

func findProcess(metaData *[]ProcessMetaData, n, min_arrival_time int, AccessedTime *int, idx *int) (int, int) {
	var temp []int
	for i := 0; i < n; i++ {
		if !(*metaData)[i].Accessed && (*metaData)[i].UnitProcess.Arrival_Time <= min_arrival_time {
			temp = append(temp, i)
		}
	}

	for _, ll := range temp {
		fmt.Println((*idx), ": ", ll)
	}

	if len(temp) == 0 {
		(*AccessedTime)--
		return (min_arrival_time + 1), 0
	}

	var val int = temp[0]
	for _, v := range temp {
		if (*metaData)[v].UnitProcess.Burst_Time < (*metaData)[val].UnitProcess.Burst_Time {
			val = v
		} else if (*metaData)[v].UnitProcess.Burst_Time == (*metaData)[val].UnitProcess.Burst_Time {
			if (*metaData)[v].UnitProcess.Arrival_Time < (*metaData)[val].UnitProcess.Arrival_Time {
				val = v
			}
		}
	}

	(*metaData)[val].Accessed = true
	(*metaData)[val].AccessIndex = (*idx)
	(*idx) = (*idx) + 1
	(*metaData)[val].AccessedTime = min_arrival_time
	return ((*metaData)[val].UnitProcess.Burst_Time + min_arrival_time), 1
}

func shortest_job_first(pair []Pair, n int, min_arrival_time int, processes []Process) {
	var AccessedProcesses = 0
	var metaData []ProcessMetaData
	var GranntChart []Grannt
	var idx int = 1
	for _, p := range processes {
		metaData = append(metaData, ProcessMetaData{
			UnitProcess: p,
			Accessed:    false,
			AccessIndex: 0,
		})
	}

	var AccessedTime int = 0
	for AccessedProcesses < n {
		// if AccessedProcesses == n {
		// break
		// }
		var inc int = 0
		min_arrival_time, inc = findProcess(&metaData, n, min_arrival_time, &AccessedTime, &idx)
		AccessedTime++
		AccessedProcesses += inc
	}

	// for i, value := range GranntChart {
	// 	fmt.Println(i+1, ". ", value)
	// }

	fmt.Println("Answer:")
	for _, value := range metaData {
		fmt.Println(value.AccessIndex, " -> ", value.UnitProcess.Process_Id, " -> ", value.Accessed, " -> ", value.AccessedTime)
		GranntChart = append(GranntChart, Grannt{Process: value.UnitProcess.Process_Id, Time: value.AccessedTime})
	}

	sort.Slice(GranntChart, func(i, j int) bool {
		return GranntChart[i].Time > GranntChart[i].Time
	})

	for _, value := range GranntChart {
		fmt.Printf("(%s, %d)", value.Process, value.Time)
	}
}

func SJF() {
	var pair []Pair
	var arrival []int
	var burst []int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter n: ")
	scanner.Scan()
	nStr := scanner.Text()
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))

	fmt.Printf("Enter Arrival time of each process: ")
	scanner.Scan()
	arrival_time := scanner.Text()
	var min_arrival_time int = 1e9
	for _, val := range strings.Fields(arrival_time) {
		num, _ := strconv.Atoi(val)
		min_arrival_time = min(min_arrival_time, num)
		arrival = append(arrival, num)
	}

	fmt.Printf("Enter Burst time of each process: ")
	scanner.Scan()
	burst_time := scanner.Text()
	for _, val := range strings.Fields(burst_time) {
		num, _ := strconv.Atoi(val)
		burst = append(burst, num)
	}
	// fmt.Printf("Ended input")

	var process []Process

	for i := 0; i < n; i++ {
		pair = append(pair, Pair{arrival[i], burst[i], "p" + strconv.Itoa(i)})
		process = append(process, Process{Arrival_Time: arrival[i], Burst_Time: burst[i], Process_Id: create_process_id(i)})
	}

	// sort.Slice(pair, func(i, j int) bool {
	// 	if pair[i].Second == pair[j].Second {
	// 		return pair[i].First < pair[j].First
	// 	}

	// 	return pair[i].Second < pair[j].Second
	// })

	shortest_job_first(pair, n, min_arrival_time, process)

	// for i := 0; i < n; i++ {
	// 	fmt.Println(pair[i].First, " ", pair[i].Second, " ", pair[i].Process)
	// }

	// for i := 0; i < n; i++ {
	// 	fmt.Printf("%s ", pair[i].Process)
	// }
}

func main() {
	SJF()
}
