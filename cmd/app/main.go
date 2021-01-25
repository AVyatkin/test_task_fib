package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	fibCache = []int64{0, 1}
)

func main() {
	println("Start server")

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":5555", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path[1:], "/")

	if len(parts) < 3 {
		fmt.Fprintf(w, "Parameters error: require minimum two parameters, /fib/10/20")
		return
	}

	if parts[0] != "fib" {
		fmt.Fprintf(w, "Parameters error: method not allowed")
		return
	}

	from, _ := strconv.ParseInt(parts[1], 10, 64)
	to, _ := strconv.ParseInt(parts[2], 10, 64)

	if to < from {
		fmt.Fprintf(w, "Parameters error: to must not be less then from")
		return
	}
	if to < 1 || from < 1 {
		fmt.Fprintf(w, "Parameters error: to and from must be more then 0")
		return
	}
	if to > 93 || from > 93 {
		fmt.Fprintf(w, "Parameters error: to and from must be less then 94")
		return
	}

	calcFibCache(to)
	fmt.Println(fibCache)

	data := fibCache[(from - 1):to]

	jsData, _ := json.Marshal(data)
	fmt.Fprintf(w, string(jsData))
}

func calcFibCache(n int64) []int64 {
	if n < 1 {
		return []int64{}
	}
	if n <= int64(len(fibCache)) {
		return fibCache[0:n]
	}

	length := int64(len(fibCache))
	a := fibCache[length-2]
	b := fibCache[length-1]
	c := int64(0)
	for length < n {
		length++

		c = a + b
		a = b
		b = c
		fibCache = append(fibCache, c)
	}

	return fibCache
}
