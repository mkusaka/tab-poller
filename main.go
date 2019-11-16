package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"text/template"
	"time"
)

var logger = log.New(os.Stdout, "http: ", log.LstdFlags)

func server(jobChan chan interface{}) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		jobChan <- r
		http.ServeFile(w, r, "tab_list.txt")
	})
	http.HandleFunc("/activate", func(w http.ResponseWriter, r *http.Request) {
		targetURL := r.URL.Query()["url"][0]
		err := activate(targetURL)
		if err != nil {
			log.Fatalf("activate failed: %s", err)
		}
		fmt.Fprintf(w, "activate: %s", targetURL)
	})
	http.HandleFunc("/term", func(w http.ResponseWriter, r *http.Request) {
		// TODO: add endpoint for terminate this app
		fmt.Fprintf(w, "term, %s", r.URL.Path[1:])
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
		jobChan <- "execute !"
	}
}

var activateTemplate = `
set _tab_was_found to false
      tell application "Google Chrome"
                set _window_index to 1

      repeat with _window in windows
        try
          set _tab_count to (count of tabs in _window)
          set _tab_index to 1
          repeat with _tab in tabs of _window
                    if (url of _tab as string is "{{.}}") then
                set _tab_was_found to true
      activate
      tell _window
        set active tab index to  _tab_index
        set index to 1
      end tell
      exit repeat

      end if


            set _tab_index to _tab_index + 1
          end repeat
        end try
        set _window_index to _window_index + 1
      end repeat

      end tell

              if (_tab_was_found) then
        -- Bring window to front
        tell application "System Events" to tell process "Google Chrome"
          perform action "AXRaise" of window 1
          -- account for instances when the window doesn't switch fast enough
          delay 0.5
          perform action "AXRaise" of window 1
          -- Prevent other running browsers from potentially activating
          return
        end tell
      end if

`

func activate(url string) error {
	// TODO: if url not exit in managed tab list, returns error.
	parsedString, err := processString(activateTemplate, url)
	if err != nil {
		return err
	}
	log.Println(parsedString)

	result, err := exec.Command("osascript", "-e", parsedString).Output()
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}

// cf. https://dev.to/kirklewis/go-text-template-processing-181d
func processString(str string, vars interface{}) (string, error) {
	tmpl, err := template.New("tmpl").Parse(str)

	if err != nil {
		return "", err
	}
	proecssed, err := process(tmpl, vars)
	if err != nil {
		return "", err
	}
	return proecssed, nil
}

func process(t *template.Template, vars interface{}) (string, error) {
	var tmplBytes bytes.Buffer

	err := t.Execute(&tmplBytes, vars)
	if err != nil {
		return "", err
	}
	return tmplBytes.String(), nil
}

func main() {
	jobChan := make(chan interface{}, 100)
	dbFile, err := os.Create("tab_list.txt")
	if err != nil {
		log.Fatal("db file cannot open.")
	}
	defer dbFile.Close()
	defer os.Remove("tab_list.txt")
	go func() {
		worker(jobChan, dbFile)
	}()
	go func() {
		timer(jobChan)
	}()
	server(jobChan)
}
