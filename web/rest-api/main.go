package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Data struct {
	Data string `json:"data"`
}

var PORT int
var LENGTH int
var RAND = rand.NewSource(time.Now().UnixNano())

const (
	defaultPort   = 8081
	defaultLength = 512
	portHelp      = "port number to listen - integer"
	lengthHelp    = "data(character) length - integer"
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func generateRandomString(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, RAND.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = RAND.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func getJSONData(length int) string {
	dt := &Data{
		Data: generateRandomString(length),
	}
	data, _ := json.Marshal(dt)
	return string(data)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, getJSONData(LENGTH))
}

func indexWithLength(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	length, err := strconv.Atoi(vars["len"])
	if err != nil {
		fmt.Fprintf(w, getJSONData(LENGTH))
	} else {
		fmt.Fprintf(w, getJSONData(length))
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", index)
	myRouter.HandleFunc("/{len}", indexWithLength)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), myRouter)
}

func init() {
	flag.IntVar(&PORT, "port", defaultPort, portHelp)
	flag.IntVar(&PORT, "p", defaultPort, portHelp+" (shorthand)")
	flag.IntVar(&LENGTH, "length", defaultLength, lengthHelp)
	flag.IntVar(&LENGTH, "l", defaultLength, lengthHelp+" (shorthand)")
	flag.IntVar(&LENGTH, "n", defaultLength, lengthHelp+" (shorthand)")
	flag.Parse()
	fmt.Printf("Simulated data length %d\n", LENGTH)
	fmt.Printf("Server will be listening on port %d\n", PORT)
}

func main() {
	handleRequests()
}
