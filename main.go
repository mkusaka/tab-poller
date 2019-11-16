package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

var logger = log.New(os.Stdout, "http: ", log.LstdFlags)

func server(jobChan chan interface{}) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		jobChan <- r
		http.ServeFile(w, r, "tab_list.txt")
	})
	http.HandleFunc("/term", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hellow, %s", r.URL.Path[1:])
	})
	logger.Println("Server is starting here: http://localhost:30303")
	http.ListenAndServe(":30303", nil)
}

func worker(jobChan <-chan interface{}, file *os.File) {
	log.Println("chan called")
	for job := range jobChan {
		fmt.Printf("chan has reached %v \n", job)
		result, err := exec.Command("osascript", "list_chrome_safari_tabs.applescript").Output()
		if err != nil {
			log.Fatal("command execute faild")
		}
		log.Println(string(result))
		file.Write(result)
	}
	if jobChan == nil {
		log.Println("channel closed")
	}
}

func timer(jobChan chan interface{}) {
	for range time.Tick(5 * time.Second) {
		go func() {
			jobChan <- "execute !"
		}()
	}
}

func main() {
	jobChan := make(chan interface{}, 100)
	dbFile, err := os.Create("tab_list.txt")
	if err != nil {
		log.Fatal("db file cannot open.")
	}
	defer dbFile.Close()
	go func() {
		worker(jobChan, dbFile)
	}()
	go func() {
		timer(jobChan)
	}()
	server(jobChan)
}
