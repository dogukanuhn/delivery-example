package transferService

import (
	"github.com/dogukanuhn/delivery-system/domain"
	"github.com/dogukanuhn/delivery-system/domain/types/delivery_status"
	"github.com/dogukanuhn/delivery-system/internal/logger"
)

type TransferService struct {
	Packages *map[string]domain.Package
	Sacks    *map[string]domain.Sack
	logger   logger.ILogger
}

func NewHandler(packages *map[string]domain.Package, sacks *map[string]domain.Sack, logger logger.ILogger) *TransferService {
	return &TransferService{Packages: packages, Sacks: sacks, logger: logger}
}

func Contains(deliveryList []domain.Package, barcode string) (domain.Package, bool) {
	for _, value := range deliveryList {
		if value.Barcode == barcode {
			return value, true
		}
	}
	return domain.Package{}, false
}

func (h *TransferService) Deliver(deliveryList []domain.Delivery) {
	if len(*h.Sacks) > 0 {

		localList := deliveryList
		var state int

		for _, sack := range *h.Sacks {

			for deliveryIndex, deliveryItem := range localList {

				_, isContain := Contains(sack.Packages, deliveryItem.Barcode)

				if isContain {
					state = int(delivery_status.Unloaded)
				} else {
					h.logger.Warning(deliveryItem.Barcode, " Wrong Delivery Point")
					state = int(delivery_status.Loaded)
				}

				deliveryList[deliveryIndex].State = state
				packageItem := (*h.Packages)[deliveryItem.Barcode]
				packageItem.State = state
				(*h.Packages)[deliveryItem.Barcode] = packageItem

			}

			sack.State = int(delivery_status.Unloaded)
			(*h.Sacks)[sack.Barcode] = sack
		}

	}
}
