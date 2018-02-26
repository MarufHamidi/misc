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

func lomuto(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)-1]
	i, j := -1, 0
	for j < len(arr)-1 {
		if arr[j] < pivot {
			i++
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		j++
	}

	i++
	if i <= j {
		arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]
	}

	if arr[0] == pivot {
		arr2 := lomuto(arr[i+1:])
		arr = []int{pivot}
		for i := 0; i < len(arr2); i++ {
			arr = append(arr, arr2[i])
		}
	} else if arr[len(arr)-1] == pivot {
		arr = append(lomuto(arr[0:i]), pivot)
	} else {
		arr2 := lomuto(arr[:i])
		arr3 := lomuto(arr[i+1:])
		var arr []int
		for i := 0; i < len(arr2); i++ {
			arr = append(arr, arr2[i])
		}
		arr = append(arr, pivot)
		for i := 0; i < len(arr3); i++ {
			arr = append(arr, arr3[i])
		}
	}

	return arr
}

// quick sort
func sort(data []int) []int {
	return lomuto(data)
}

func main() {
	dat := read()
	start := time.Now()
	sorted := sort(dat)
	elapsed := time.Since(start)
	fmt.Println(sorted)
	fmt.Printf("Execution time: %s\n", elapsed)
}
