package message

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type messageServiceMock struct {
    mock.Mock
}

func (m *messageServiceMock) SendChargeNotification(value int) bool {
    args := m.Called(value)
    return args.Bool(0)
}

// TestChargeCustomer is where the magic happens
// here we create our MessageService mock
func TestChargeCustomer(t *testing.T) {
    messageService := new(messageServiceMock)
    messageService.On("SendChargeNotification", 100).Return(true)

    // next we want to define the service we wish to test
    myService := MyService{messageService}
    myService.ChargeCustomer(100)

    // at the end, we verify that our myService.ChargeCustomer
    // method called our mocked SendChargeNotification method
    messageService.AssertExpectations(t)
}