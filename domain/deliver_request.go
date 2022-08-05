package domain

type Route struct {
	DeliveryPoint int
	Deliveries    []Delivery
}
type Delivery struct {
	Barcode string `json:"barcode"`
	State   int    `json:"state" swaggerignore:"true"`
}
