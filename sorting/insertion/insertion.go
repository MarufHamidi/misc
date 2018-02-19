package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

var inputFile = "../input"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read() []int {
	dat, err := ioutil.ReadFile(inputFile)
	check(err)
	datStr := strings.TrimSpace(string(dat))
	datStrArr := strings.Split(datStr, " ")
	var arr []int
	for i := 0; i < len(datStrArr); i++ {
		n, _ := strconv.Atoi(datStrArr[i])
		arr = append(arr, n)
	}
	return arr
}

// insertion sort
func sort(data []int) []int {
	arr := data
	for i := 1; i < len(arr); i++ {
		n := arr[i]
		j := i - 1
		for j >= 0 && n < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = n
	}
	return arr
}

func main() {
	dat := read()
	start := time.Now()
	sorted := sort(dat)
	elapsed := time.Since(start)
	fmt.Println(sorted)
	fmt.Printf("Execution time: %s\n", elapsed)
}
