package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"stevenleeza/grpc-json-demo/homeaffairspb"
	"sync"

	"google.golang.org/grpc"
)

type homeAffairsServer struct {
	homeaffairspb.UnimplementedHomeAffairsServer
}

type Citizen struct {
	FirstName, Surname string
	Age                int
}

type CitizenRequest struct {
	IdNumber string
}

var citizens = map[string]Citizen{
	"12345": {
		FirstName: "John",
		Surname:   "Smith",
		Age:       18,
	},
}

func (hs *homeAffairsServer) GetCitizen(ctx context.Context, cReq *homeaffairspb.CitizenRequest) (*homeaffairspb.CitizenResponse, error) {
	idNumber := cReq.GetIdNumber()
	citizen := citizens[idNumber]

	cRes := &homeaffairspb.CitizenResponse{
		FirstName: citizen.FirstName,
		Surname:   citizen.Surname,
		Age:       int32(citizen.Age),
	}

	return cRes, nil
}

func startRpcServer() {
	addr := ":50050"

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	cServer := &homeAffairsServer{}

	s := grpc.NewServer()
	homeaffairspb.RegisterHomeAffairsServer(s, cServer)

	log.Printf("gRPC server listening at %s\n", addr)

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}

func getCitizen(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var cr CitizenRequest
	err := decoder.Decode(&cr)
	if err != nil {
		panic(err)
	}

	citizen := citizens[cr.IdNumber]

	citizenJSON, _ := json.Marshal(citizen)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	fmt.Fprintf(w, string(citizenJSON))
}

func startJsonServer() {
	const restJsonServerAddr = ":50051"
	http.HandleFunc("/getCitizen", getCitizen)
	log.Printf("JSON server listening at %s\n", restJsonServerAddr)
	http.ListenAndServe(restJsonServerAddr, nil)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go startRpcServer()
	go startJsonServer()

	wg.Wait()
}
