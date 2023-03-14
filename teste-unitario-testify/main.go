package main

import (
	"fmt"

	"github.com/kustavo/tutorial-go/teste-unitario-testify/message"
)

func main() {
	// A "Production" Example
    fmt.Println("Tutorial teste unitário com mock usando testify")

    messageService := message.MessageServiceImp{}
    myService := message.MyService{MessageService: messageService}
    myService.ChargeCustomer(100)

}