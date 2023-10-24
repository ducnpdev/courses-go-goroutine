package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

func init() {
}

func main() {
	wg := sync.WaitGroup{}
	var receiverOrderCh = make(chan Order)
	var outOrderCh = make(chan Order)
	go receiverOrder(&wg, receiverOrderCh)
	//
	go checkOrder(receiverOrderCh, outOrderCh)
	// time.Sleep(time.Second)
	wg.Add(1)
	go func() {
		order := <-outOrderCh
		fmt.Printf("received: %v\n", order)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("stop")
	// fmt.Printf("output %v", listOrder)
}

func checkOrder(receiverOrderCh chan Order, outOrderCh chan Order) {
	tmpOrder := <-receiverOrderCh
	// process
	if tmpOrder.Quantity <= 0 {
		fmt.Printf("outData %v", outOrderCh)
	} else {
		outOrderCh <- tmpOrder
		fmt.Println("else value")
	}
}

func receiverOrder(wg *sync.WaitGroup, outOrder chan Order) {
	defer wg.Done()
	now := time.Now()
	for _, itemOrder := range rawOrders {
		var order Order
		if err := json.Unmarshal([]byte(itemOrder), &order); err != nil {
			fmt.Println(err)
			continue
		}
		// khong biet trước số lượng
		listOrder = append(listOrder, order)
		fmt.Println("push channel")
		outOrder <- order
	}
	fmt.Println(time.Since(now).Microseconds())
}

var listOrder = []Order{}

type Order struct {
	ProductCode int
	Quantity    float64
	Status      orderStatus
}

func (o Order) String() string {
	return fmt.Sprintf("Product code: %v, Quantity: %v, Status: %v", o.ProductCode, o.Quantity, orderStatusToText(o.Status))
}

func orderStatusToText(o orderStatus) string {
	switch o {
	case none:
		return "none"
	case new:
		return "new"
	case received:
		return "received"
	case reserved:
		return "reserved"
	case filled:
		return "filled"
	default:
		return "unknown status"
	}
}

type orderStatus int

const (
	none orderStatus = iota
	new
	received
	reserved
	filled
)

var rawOrders = []string{
	`{"productCode": 1111, "quantity": 5, "status": 1}`,
	`{"productCode": 2222, "quantity": 42.3, "status": 2}`,
	`{"productCode": 3333, "quantity": 19, "status": 3}`,
	`{"productCode": 4444, "quantity": 8, "status": 4}`,
}
