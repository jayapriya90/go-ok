package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	filepath = flag.String("fpath", "input.txt", "input file path")
	persist  = flag.Bool("persist", false, "save the contents")
)

func main() {
	flag.Parse()
	file, err := os.Open(*filepath)
	fileInfo, _ := os.Stat(*filepath)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	if err == nil && fileInfo.IsDir() {
		log.Fatalf("%s is a directory and cannot be opened", *filepath)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Parse the input file. Each line in the file has IP address
	// and port in the format host:port
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var hostPorts []string

	for scanner.Scan() {
		hostPorts = append(hostPorts, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	// for debugging
	fmt.Printf("Input: %v \n", hostPorts)

	// timeout after 5 seconds
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	fmt.Println("HTTP GET is successful for")
	for _, hostPort := range hostPorts {
		resp, err := client.Get("http://" + hostPort)
		if err != nil {
			log.Println(err)
			continue
		}
		if resp.StatusCode == http.StatusOK {
			fmt.Printf("%v \n", hostPort)
		}
	}
}
