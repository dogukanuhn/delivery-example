package distributionService

import (
	"github.com/dogukanuhn/delivery-system/domain"
	"github.com/dogukanuhn/delivery-system/domain/types/delivery_status"
	"github.com/dogukanuhn/delivery-system/internal/logger"
)

type DistributionService struct {
	Packages *map[string]domain.Package
	Sacks    *map[string]domain.Sack
	logger   logger.ILogger
}

func NewHandler(packages *map[string]domain.Package, sacks *map[string]domain.Sack, logger logger.ILogger) *DistributionService {
	return &DistributionService{Packages: packages, Sacks: sacks, logger: logger}
}

//everything can be unload
func (h *DistributionService) Deliver(deliveryList []domain.Delivery) {

	localList := deliveryList
	const unloaded = int(delivery_status.Unloaded)
	for index, delivery := range localList {

		item, isAvaible := (*h.Packages)[delivery.Barcode]
		if isAvaible {

			item.State = unloaded
			(*h.Packages)[delivery.Barcode] = item
			deliveryList[index].State = unloaded
			continue
		}

		sack, isAvaible := (*h.Sacks)[delivery.Barcode]
		if isAvaible {

			sack.State = unloaded
			deliveryList[index].State = unloaded
			(*h.Sacks)[delivery.Barcode] = sack

			for _, packageItem := range (*h.Sacks)[delivery.Barcode].Packages {
				packageItem.State = unloaded
				(*h.Packages)[packageItem.Barcode] = packageItem
			}
			continue
		}

		h.logger.Warning(delivery.Barcode, " Wrong Delivery Point")
		deliveryList[index].State = int(delivery_status.Loaded)

	}

}
