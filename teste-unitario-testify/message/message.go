package message

import (
	"fmt"
)


type MessageService interface {
    SendChargeNotification(int) bool
}

type MessageServiceImp struct {}

type MyService struct {
    MessageService MessageService
}

func (sms MessageServiceImp) SendChargeNotification(value int) bool {
    fmt.Println("Sending Charge Notification")
    return true
}

func (a MyService) ChargeCustomer(value int) error {
    a.MessageService.SendChargeNotification(value)
    fmt.Printf("Charging Customer For the value of %d\n", value)
    return nil
}
