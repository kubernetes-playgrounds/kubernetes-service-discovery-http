package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	var url string
	if os.Getenv("DEPLOYMENT") == "PRODUCTION" {
		url = "server-svc"
	} else {
		url = "localhost"
	}
	for {
		resp, err := http.Get(fmt.Sprintf("http://%s:8081/Jude", url))
		if err != nil {
			fmt.Println(err.Error())
		} else {
			defer resp.Body.Close()
			fmt.Println("Response status:", resp.Status)
			scanner := bufio.NewScanner(resp.Body)
			for i := 0; scanner.Scan() && i < 5; i++ {
				fmt.Println(scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				fmt.Println(err.Error())
			}

		}
		time.Sleep(5 * time.Second)
	}
}
