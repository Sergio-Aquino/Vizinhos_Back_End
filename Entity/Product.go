package Entity

type Product struct {
	ProductID          int
	Store              StoreOrAddress
	Category           Category
	DaysToExp          int
	SellPrice          float64
	ManufacturingPrice float64
	Size               float64
}
