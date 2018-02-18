package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
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

// bubble sort
func sort(data []int) []int {
	arr := data
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func main() {
	sorted := sort(read())
	fmt.Println(sorted)
}
