package deliveryService

import (
	"context"

	"github.com/dogukanuhn/delivery-system/domain"
	"github.com/dogukanuhn/delivery-system/domain/dto"
	"github.com/dogukanuhn/delivery-system/domain/interfaces"
	"github.com/dogukanuhn/delivery-system/domain/types/delivery_point"
	"github.com/dogukanuhn/delivery-system/internal/logger"
	"github.com/dogukanuhn/delivery-system/internal/services/branchService"
	"github.com/dogukanuhn/delivery-system/internal/services/distributionService"
	"github.com/dogukanuhn/delivery-system/internal/services/transferService"
)

type DeliveryService struct {
	Packages    map[string]domain.Package
	Sacks       map[string]domain.Sack
	PackageRepo interfaces.IPackageRepository
	SackRepo    interfaces.ISackRepository
	logger      logger.ILogger
}

func NewHandler(packageRepo interfaces.IPackageRepository, sackRepo interfaces.ISackRepository, logger logger.ILogger) *DeliveryService {
	return &DeliveryService{PackageRepo: packageRepo, SackRepo: sackRepo, logger: logger}
}

func (h *DeliveryService) Deliver(DeliveryDTO dto.DeliverDTO) {

	for _, point := range DeliveryDTO.Route {

		h.HandleDelivery(&point)
		err := h.UpdateDb()

		if err != nil {
			h.logger.Errorf(err.Error())
		}
	}

}

func (h *DeliveryService) UpdateDb() error {

	_, err := h.PackageRepo.BulkWrite(context.TODO(), h.Packages)

	if len(h.Sacks) > 0 {
		_, err = h.SackRepo.BulkWrite(context.TODO(), h.Sacks)
	}

	return err

}

func (h *DeliveryService) HandleDelivery(delivery *domain.Route) {

	var err error
	h.Packages, err = h.PackageRepo.FindByPoint(context.Background(), delivery.DeliveryPoint)
	h.Sacks, err = h.SackRepo.FindByPoint(context.Background(), delivery.DeliveryPoint)

	if err != nil {
		h.logger.Warning(err.Error())
	}

	var deliveryPointService interfaces.IDeliveryPointService

	switch delivery.DeliveryPoint {
	case delivery_point.BranchPoint:
		deliveryPointService = branchService.NewHandler(&h.Packages, h.logger)
	case delivery_point.DistributionPoint:
		deliveryPointService = distributionService.NewHandler(&h.Packages, &h.Sacks, h.logger)
	case delivery_point.TransferPoint:
		deliveryPointService = transferService.NewHandler(&h.Packages, &h.Sacks, h.logger)
	}

	deliveryPointService.Deliver(delivery.Deliveries)
}
