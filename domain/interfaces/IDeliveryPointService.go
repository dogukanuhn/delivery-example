package interfaces

import "github.com/dogukanuhn/delivery-system/domain"

type IDeliveryPointService interface {
	Deliver(deliveryList []domain.Delivery)
}
