package main

import (
	"fmt"
	"os"
	"time"
	"net"
	"strings"
	"net/http"
	"net/url"
	"io"
	"strconv"
	"github.com/mackerelio/go-osstat/memory"
)

func main() {
	fmt.Println("Start data collector.")
	for {
		collect_memory()
		time.Sleep(10 * time.Second)
	}

}

func collect_memory() {
	timestamp := time.Now()

	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	var ipv4 net.IP
	for _, addr := range addrs {
		if ipv4 = addr.To4(); ipv4 != nil {
			fmt.Print("")
		}   
	}

	systemMemory, err := memory.Get()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print("Sending data：")
	fmt.Printf("Timestamp=%s, ip=%s, free=%d",timestamp.String(), ipv4, systemMemory.Free)

	params := url.Values{}
	params.Add("timestamp", timestamp.String())
	params.Add("ip", ipv4.String())
	params.Add("free", strconv.FormatUint(uint64(systemMemory.Free), 10))
	body := strings.NewReader(params.Encode())
	server_addr := "http://" + os.Getenv("serverAddr") + ":5000/memory/record"
	req, err := http.NewRequest("POST", server_addr, body)
	if err != nil {
		fmt.Print(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
        fmt.Println(err)
    }
	fmt.Println("\n-> " + string(bodyBytes))

}
