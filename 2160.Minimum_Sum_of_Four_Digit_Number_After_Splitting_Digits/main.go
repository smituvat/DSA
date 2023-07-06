package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Print(minimumSum(2932))
}

func minimumSum(num int) int {
	sN := strconv.Itoa(num)
	sNs := strings.Split(sN, "")
	sort.Strings(sNs)
	n1, _ := strconv.Atoi(sNs[0] + sNs[2])
	n2, _ := strconv.Atoi(sNs[1] + sNs[3])
	return n1 + n2
}
