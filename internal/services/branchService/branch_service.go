package branchService

import (
	"github.com/dogukanuhn/delivery-system/domain"
	"github.com/dogukanuhn/delivery-system/domain/types/delivery_status"
	"github.com/dogukanuhn/delivery-system/internal/logger"
)

type BranchService struct {
	Packages *map[string]domain.Package
	logger   logger.ILogger
}

func NewHandler(packages *map[string]domain.Package, logger logger.ILogger) *BranchService {
	return &BranchService{Packages: packages, logger: logger}
}

//Only packages can be unloaded at Branch, no need sack validation
func (h *BranchService) Deliver(deliveryList []domain.Delivery) {

	localList := deliveryList
	for index, delivery := range localList {

		item, isAvaible := (*h.Packages)[delivery.Barcode]

		if isAvaible {
			const state = int(delivery_status.Unloaded)

			item.State = state
			(*h.Packages)[delivery.Barcode] = item
			deliveryList[index].State = state
			continue
		}

		h.logger.Warning(delivery.Barcode, " Wrong Delivery Point")
		deliveryList[index].State = int(delivery_status.Loaded)

	}
}
