package main

import (
	"fmt"
	"time"
)

type order struct{
	id int
	name string
	amount float64
	status string
	createdAt time.Time
}

func (ord *order) changeStatus(status string) {
	ord.status = status
}

func changeByRef(o *order){
	o.status = "approved"
}


func findPendingOrders(orders []order) []order {
	var pendingOrders []order 
	for _, order := range orders {
		switch order.status{
			case "pending":
				pendingOrders = append(pendingOrders, order)
		}
	}
	return pendingOrders
}

func main(){
	var orders []order
	
	orders = append(orders, order{
		id: 1,
		name: "Jake Peralta",
		amount: 100.00,
		status: "pending",
		createdAt: time.Now(),
	})
	sampleOrder := order{
		id: 2,
		name: "Charles Boyle",
		amount: 200.00,
		status: "pending",
		createdAt: time.Now(),
	}

	sampleOrder1 := order{
		id: 3,
		name: "Raymond Holt",
		amount: 400.00,
		status: "pending",
		createdAt: time.Now(),
	}

	
	fmt.Println(sampleOrder)
	
	changeByRef(&sampleOrder)
	
	orders = append(orders, sampleOrder)
	orders = append(orders, sampleOrder1)

	fmt.Println(sampleOrder)

	fmt.Println(orders)

	pendingOrders := findPendingOrders(orders)
	
	
	fmt.Println("Pending", len(pendingOrders),"/",len(orders),"orders", pendingOrders)
	orders[2].changeStatus("approved")
	pendingOrders = findPendingOrders(orders)

	fmt.Println("Pending", len(pendingOrders),"/",len(orders),"orders", pendingOrders)
}