package main

import (
	"fmt"
	"io"
	"io/ioutil"
	mylog "log"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		logger(<-ch)
	}
	err := logger(fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds()))
	if err != nil {
		os.Exit(1)
	}
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nBytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nBytes, url)
}

func logger(message string) error {
	file, err := os.OpenFile(`./fetch.log`, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic("can not open fetch.log:" + err.Error())
	}
	defer file.Close()

	mylog.SetOutput(io.MultiWriter(file, os.Stdout))
	mylog.SetFlags(mylog.Ldate | mylog.Ltime)
	mylog.Println("| ", message)
	return nil
}
