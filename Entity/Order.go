package Entity

import "time"

type Order struct {
	OrderID    int
	User       User
	Batch      Batch
	Price      float64
	Quantity   int
	Date       time.Time
	Status     string
	LastUpdate time.Time
}
