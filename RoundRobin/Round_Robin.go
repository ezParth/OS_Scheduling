package roundrobin

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Stack struct {
	value []int
}

type Queue struct {
	value []int
}

type Pair struct {
	First  int
	Second int
}

type DSInterface interface {
	Push(val int)
	Pop() int
	Size() int
	Top() int
	Complete() []int
	IsEmpty() bool
}

func (s *Stack) Push(val int) {
	s.value = append(s.value, val)
}

func (s *Stack) Pop() int {
	if len(s.value) == 0 {
		return -1
	}
	val := s.value[len(s.value)-1]
	s.value = s.value[:len(s.value)-1]
	return val
}

func (s *Stack) Size() int {
	return len(s.value)
}

func (s *Stack) Top() int {
	if len(s.value) == 0 {
		return -1
	}
	return s.value[len(s.value)-1]
}

func (s *Stack) Complete() []int {
	return s.value
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (q *Queue) Push(val int) {
	q.value = append(q.value, val)
}

func (q *Queue) Pop() int {
	if len(q.value) == 0 {
		return -1
	}
	val := q.value[0]
	q.value = q.value[1:]
	return val
}

func (q *Queue) Size() int {
	return len(q.value)
}

func (q *Queue) Top() int {
	if len(q.value) == 0 {
		return -1
	}
	return q.value[0]
}

func (q *Queue) Complete() []int {
	return q.value
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func RoundRobin(pair []Pair, n int) {
	// Q = 2
	var Q int = 2
	var readyQueue DSInterface = &Queue{}
	var ans int = 0
	for i := 0; i < n; i++ {
		var remaining_burst_time int = pair[i].Second - min(Q, pair[i].Second)
		ans += 1
		if remaining_burst_time != 0 {
			readyQueue.Push(remaining_burst_time)
		}
	}

	fmt.Println("ready Queue: ", readyQueue.Complete())

	for {
		if readyQueue.IsEmpty() {
			break
		}

		var val int = readyQueue.Top()
		readyQueue.Pop()
		var remaining_burst_time int = val - min(Q, val)
		if remaining_burst_time != 0 {
			readyQueue.Push(remaining_burst_time)
		}
		ans += 1
	}
	fmt.Printf("ans: %d", ans)
}

func RR() {
	// declaring scanner
	var arrival, burst DSInterface
	arrival = &Queue{}
	burst = &Queue{}
	scanner := bufio.NewScanner(os.Stdin)
	var pair []Pair

	// reading first line -> number of processes
	fmt.Printf("Enter n: ")
	scanner.Scan()
	nStr := scanner.Text()
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))

	// scanner.Scan()
	// qStr := scanner.Text()
	// q, _ := strconv.Atoi(strings.TrimSpace(qStr))

	// reding arrival time
	fmt.Printf("Enter arrival time of each process: ")
	scanner.Scan()
	line1 := scanner.Text()
	arr1 := make([]int, 0, n)
	for _, val := range strings.Fields(line1) {
		num, _ := strconv.Atoi(val)
		arr1 = append(arr1, num)
		arrival.Push(num)
	}

	// reding burst time
	fmt.Printf("Enter burst time of each process: ")
	scanner.Scan()
	line2 := scanner.Text()
	arr2 := make([]int, 0, n)
	for _, val := range strings.Fields(line2) {
		num, _ := strconv.Atoi(val)
		arr2 = append(arr2, num)
		burst.Push(num)
	}

	for i := 0; i < n; i++ {
		pair = append(pair, Pair{arr1[i], arr2[i]})
	}

	for i := 0; i < n; i++ {
		fmt.Println(pair[i].First, " ", pair[i].Second)
	}

	sort.Slice(pair, func(i, j int) bool {
		if pair[i].First == pair[j].First {
			return pair[i].Second < pair[j].Second
		}
		return pair[i].First < pair[j].First
	})

	for i := 0; i < n; i++ {
		fmt.Println(pair[i].First, " ", pair[i].Second)
	}

	// fmt.Printf("Round Robin function call -> ")
	fmt.Println("burst: ", burst.Complete())
	fmt.Println("Arrival: ", arrival.Complete())
	// Round_Robin_OS(n, burst, arrival, arr1, arr2)
	RoundRobin(pair, n)

}
