package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Hello NeuronOS !!!")
	fmt.Println("\nWait Ctrl+C for exit...")

	// Create a channel to receive OS signals
	c := make(chan os.Signal, 1)
	// Notify the channel when SIGINT is received
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		// Block until a signal is received
		fmt.Println("\nCtrl+C pressed. Exit...")
		os.Exit(0)
	}()

	commander := NewCommander()

	server := &http.Server{
		Addr:    ":8080",
		Handler: handleRequests(commander),
	}

	log.Fatal(server.ListenAndServe())
}

func handleRequests(cmdr Commander) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/execute", handleCommand(cmdr))
	return mux
}

func handleCommand(cmdr Commander) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cr CommandResponse
		unsupportedMethod := "Unsupported method: %v  Use: GET or POST\n"
		unsupportedCommand := "Unsupported command: %v  Use: ping or sysinfo\n"
		command := ""
		host := ""

		// Parse request and execute command
		fmt.Printf("--- Method: %v\n", r.Method)
		for name, headers := range r.Header {
			for _, h := range headers {
				fmt.Printf("%v: %v\n", name, h) // Output on screen
				fmt.Fprintf(w, "%v: %v\n", name, h)
			}
		}
		if r.Method != "POST" && r.Method != "GET" {
			fmt.Printf(unsupportedMethod, r.Method)
			fmt.Fprintf(w, unsupportedMethod, r.Method)
			return
		}
		if r.Method == "POST" {
			r.ParseForm() // Input as Form data
			command = r.Form.Get("command")
			host = r.Form.Get("host")
		} else {
			queryValues := r.URL.Query()
			command = queryValues.Get("command")
			host = queryValues.Get("host")
		}

		if command != "ping" && command != "sysinfo" {
			fmt.Printf(unsupportedCommand, command)
			fmt.Fprintf(w, unsupportedCommand, command)
			return
		}

		if command == "ping" {
			pr, err := cmdr.Ping(host)
			cr.Data = pr
			cr.Success = pr.Successful
			if err != nil {
				cr.Error = err.Error()
			}
		} else {
			si, err := cmdr.GetSystemInfo()
			cr.Data = si
			cr.Success = true
			if err != nil {
				cr.Error = err.Error()
				cr.Success = false
			}
		}

		// Send result via http
		crb, _ := json.Marshal(cr)
		crs := string(crb)
		fmt.Printf("%v\n", crs)
		fmt.Fprintf(w, "%v\n", crs)
	}
}
