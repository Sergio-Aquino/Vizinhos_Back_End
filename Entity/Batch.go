package Entity

import "time"

type Batch struct {
	BatchID           int
	Product           Product
	ManufacturingDate time.Time
	Discount          float64
	Quantity          int
}
