package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/shashankpn/abc"
)

var regex = []regexp.Regexp{}

func main() {
	// http.HandleFunc("/", handleData)

	matchStrings := []string{"p([a-z]+)ch"}

	for _, s := range matchStrings {
		regex = append(regex, *regexp.MustCompile(s))
	}
	abc.Log("test")
	// fmt.Println("using all cpu cores. ", runtime.NumCPU())
	// runtime.GOMAXPROCS(runtime.NumCPU())

	// fmt.Println("Starting webserver on port 9000")
	// err := http.ListenAndServe(":9000", nil)
	// if err != nil {
	// 	panic(err)
	// }
}

func handleData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := r.URL.Query().Get("data")
	go processLine(data)
	json.NewEncoder(w).Encode("OK")
}

func processLine(data string) {
	for _, ex := range regex {
		if !ex.Match([]byte(data)) {
			continue
		}
		pushToFile(data, "access.log")

		return

	}
	fmt.Println("Discarding the data")
}

func pushToFile(data string, fileName string) {

	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)

	}
	if _, err := f.Write([]byte(data + "\n")); err != nil {
		fmt.Println(err)
	}
	if err := f.Close(); err != nil {
		fmt.Println(err)
	}
}
