package main

import (
	"bufio"
	"fmt"
	"net/http"
	"time"
)

func main() {
	for {
		resp, err := http.Get("http://server-svc:8081/Jude")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		fmt.Println("Response status:", resp.Status)
		scanner := bufio.NewScanner(resp.Body)
		for i := 0; scanner.Scan() && i < 5; i++ {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		time.Sleep(5 * time.Second)
	}
}
