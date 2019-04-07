package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/carloslimasis/grpc-go-course/sum/sumpb"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := sumpb.NewSumServiceClient(conn)

	if len(os.Args) < 3 {
		log.Fatalf("The required parameters not passed")
		os.Exit(1)
	}

	firstNumber, err := strconv.ParseFloat(os.Args[1], 32)
	if err != nil {
		log.Fatalf("Error on parse float: %v", err)
	}

	secondNumber, err := strconv.ParseFloat(os.Args[2], 32)
	if err != nil {
		log.Fatalf("Error on parse float: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Sum(ctx, &sumpb.SumRequest{Sum: &sumpb.Sum{FirstNumber: float32(firstNumber), SecondNumber: float32(secondNumber)}})
	if err != nil {
		log.Fatalf("could not sum: %v", err)
	}
	log.Printf("Sum: %s", r)

}
