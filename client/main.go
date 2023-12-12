package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/yokaracho/demo-grpc/invoicer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var (
	addr = flag.String("addr", "localhost:8089", "the address to connect to")
	//name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := invoicer.NewInvoicerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	request := &invoicer.CreateRequest{
		Amount: &invoicer.Amount{
			Amount:   1000,  // Установите нужное значение
			Currency: "USD", // Установите нужное значение
		},
		From: "Вова",   // Установите нужное значение
		To:   "Андрей", // Установите нужное значение
	}

	r, err := client.Create(ctx, request)
	if err != nil {
		log.Fatalf("ошибка вызова Create: %v", err)
	}
	log.Printf("Result: %s", r)
	fmt.Printf("PDF: %v\n", r.Pdf)
	fmt.Printf("Docx: %v\n", r.Docx)
}
