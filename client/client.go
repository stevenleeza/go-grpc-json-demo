package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"stevenleeza/grpc-json-demo/homeaffairspb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	iterations := [4]int{1, 100, 1000, 10000}

	for _, numRequests := range iterations {
		run(numRequests)
	}
}

func run(numRequests int) {

	fmt.Printf("Requests: %v\n", numRequests)

	cc, err := grpc.Dial("server:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer cc.Close()

	c := homeaffairspb.NewHomeAffairsClient(cc)

	cReq := &homeaffairspb.CitizenRequest{
		IdNumber: "12345",
	}

	c.GetCitizen(context.Background(), cReq) // Send a request to initialise gRPC connection

	// gRPC Requests
	// -------------
	start := time.Now()

	var fastestRPCtime time.Duration

	for i := 0; i < numRequests; i++ {
		requestStart := time.Now()

		_, err := c.GetCitizen(context.Background(), cReq)
		if err != nil {
			log.Printf("gRPC server responded with error: %v\n", err)
		}

		requestElapsed := time.Since(requestStart)

		if i == 0 {
			fastestRPCtime = requestElapsed
		} else if requestElapsed.Seconds() != 0 && requestElapsed.Seconds() < fastestRPCtime.Seconds() {
			fastestRPCtime = requestElapsed
		}
	}

	fmt.Printf("gRPC server requests completed in %v\n", time.Since(start))
	fmt.Printf("gRPC server fastest response: %v\n", fastestRPCtime)

	// JSON requests
	// -------------
	start = time.Now()

	var fastestJSONtime time.Duration

	for i := 0; i < numRequests; i++ {
		requestStart := time.Now()

		payload := map[string]string{"idNumber": "12345"}

		payloadJSON, _ := json.Marshal(payload)

		_, err := http.Post("http://server:50051/getCitizen", "application/json", bytes.NewBuffer(payloadJSON))
		if err != nil {
			log.Printf("JSON server responded with error: %v\n", err)
		}

		requestElapsed := time.Since(requestStart)

		if i == 0 {
			fastestJSONtime = requestElapsed
		} else if requestElapsed.Seconds() != 0 && requestElapsed.Seconds() < fastestJSONtime.Seconds() {
			fastestJSONtime = requestElapsed
		}
	}

	fmt.Printf("JSON server requests completed in %v\n", time.Since(start))
	fmt.Printf("JSON server fastest response: %v\n\n", fastestJSONtime)
}
