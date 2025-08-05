package bankersalgorithm

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

func BankerAlgorithm() {

}

func Bankers() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter n: ")
	scanner.Scan()
	nStr := scanner.Text()
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))
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
		needed = append(needed, maxNeeded[i]-allocation[i])
	}

	BankerAlgorithm()
}

func main() {
	Bankers()
}
