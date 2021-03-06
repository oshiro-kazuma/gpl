package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
		fmt.Print(<-ch + "\n")
	}
}

func fetch(url string, ch chan<- string) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	re := make(chan string)

	go func() {
		start := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			re <- fmt.Sprint(err)
			return
		}

		nbytes, err := io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		if err != nil {
			re <- fmt.Sprintf("while reading %s: %v", url, err)
			return
		}

		secs := time.Since(start).Seconds()
		re <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
	}()
	for {
		select {
		case msg := <-re:
			ch <- msg
		case <-time.After(10 * time.Second):
			ch <- "timeout"
		}
	}
}
