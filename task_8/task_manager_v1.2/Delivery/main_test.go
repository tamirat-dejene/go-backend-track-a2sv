package main

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestCalculate(t *testing.T) {
// 	assert.Equal(t, Calculate(2), 4)

// 	// Table driven tests

// 	var tests = []struct {
// 		input    int
// 		expected int
// 	}{
// 		{2, 4},
// 		{-1, 1},
// 		{0, 2},
// 		{-5, -3},
// 		{99999, 100001},
// 	}

// 	assert2 := assert.New(t)

// 	for _, test := range tests {
// 		assert2.Equal(Calculate(test.input), test.expected, "The two values should be the same.")
// 	}
// }

// func TestStatus(t *testing.T) {
// 	assert := assert.New(t)

// 	status := "Up"
// 	assert.NotEqual(status, "Down", "The status should be up.")
// 	assert.NotNil(status, "The status should be not nil.")
// }

// // Mocking...
// type smsServiceMock struct {
// 	mock.Mock
// }

// func (m *smsServiceMock) SendChargeNotification(value int) error {
// 	fmt.Println("Mocked charge notification function")
// 	fmt.Printf("Value passed in: %d\n", value)

// 	args := m.Called(value)

// 	return args.Error(0)
// }

// func (m *smsServiceMock) DummyFunc() {
// 	fmt.Println("Dummy")
// }

// func TestChargeCustomer(t *testing.T) {
// 	smsService := new(smsServiceMock)

// 	smsService.On("SendChargeNotification", 100).Return(nil)

// 	myService := MyService{smsService}

// 	myService.ChargeCustomer(100)

// 	smsService.AssertExpectations(t)
// }
