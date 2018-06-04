package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

type measureTime func(int) int

func trackFunc(fn measureTime, number int, name string) int {
	timeStart := time.Now()

	defer timeTrack(timeStart, name)

	return fn(number)
}

func main() {

	fibStore = make(map[int]int)

	if len(os.Args) < 2 {
		log.Println("You must pass a value to calculate its fib number")
		os.Exit(1)
	}

	arg := os.Args[1]

	n, _ := strconv.Atoi(arg)

	trackFunc(fibLoop, n, "loop")
	trackFunc(fibRec, n, "recursive")
	trackFunc(fibMemo, n, "memoised")

	log.Println("---------------------First test over---------------------")

	trackFunc(fibLoop, n, "loop")
	trackFunc(fibRec, n, "recursive")
	trackFunc(fibMemo, n, "memoised")
}

func fibLoop(i int) int {
	v1 := 0
	v2 := 1

	for l := 0; l < i; l++ {
		// Uncomment to see all iterations, impacts performance
		// log.Println(v1)
		tmp := v2
		v2 = v1 + v2
		v1 = tmp
	}

	return v2
}

func fibRec(i int) int {
	if i <= 1 {
		return 1
	}

	return fibRec(i-1) + fibRec(i-2)
}

var fibStore map[int]int

func fibMemo(i int) int {
	if _, ok := fibStore[i]; ok {
		return fibStore[i]
	}

	if i <= 1 {
		return 1
	}

	fibStore[i] = fibMemo(i-1) + fibMemo(i-2)
	return fibStore[i]
}

func timeTrack(start time.Time, name string) {
	t := time.Since(start)
	log.Printf("%s took %s to perform", name, t)
}
