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

func divide(arr []int) ([]int, []int) {
	if len(arr) == 0 {
		return nil, nil
	}
	if len(arr) == 1 {
		return arr, nil
	}

	return combine(divide(arr[:len(arr)/2])), combine(divide(arr[len(arr)/2:]))
}

func combine(a []int, b []int) []int {
	if a == nil {
		return nil
	}
	if b == nil {
		return a
	}

	var arr []int
	for i, j := 0, 0; ; {
		if i >= len(a) {
			for j < len(b) {
				arr = append(arr, b[j])
				j++
			}
			break
		}
		if j >= len(b) {
			for i < len(a) {
				arr = append(arr, a[i])
				i++
			}
			break
		}
		if a[i] > b[j] {
			arr = append(arr, b[j])
			j++
		} else {
			arr = append(arr, a[i])
			i++
		}
	}
	return arr
}

// merge sort
func sort(data []int) []int {
	return combine(divide(data))
}

func main() {
	dat := read()
	start := time.Now()
	sorted := sort(dat)
	elapsed := time.Since(start)
	fmt.Println(sorted)
	fmt.Printf("Execution time: %s\n", elapsed)
}
