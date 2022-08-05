package distributionService_test

import (
	"testing"

	"github.com/dogukanuhn/delivery-system/domain"
	"github.com/dogukanuhn/delivery-system/domain/types/delivery_status"
	"github.com/dogukanuhn/delivery-system/internal/logger"
	"github.com/dogukanuhn/delivery-system/internal/services/distributionService"
	"github.com/stretchr/testify/assert"
)

///////////////////////////////////
// DISTRIBUTION DELIVERY TEST
///////////////////////////////////

func TestUnloadService_Deliver_ShouldBeNoErrorAndSacksOrPackagesUnloaded(t *testing.T) {

	packages := map[string]domain.Package{
		"P8988000120": {Barcode: "P8988000120", UnloadAt: 2, State: 1},
		"P8988000121": {Barcode: "P8988000121", UnloadAt: 2, State: 1},
		"P8988000124": {Barcode: "P8988000124", UnloadAt: 2, State: 1},
		"P8988000122": {Barcode: "P8988000122", UnloadAt: 2, State: 1},
		"P8988000126": {Barcode: "P8988000126", UnloadAt: 2, State: 1},
	}

	sacks := map[string]domain.Sack{
		"C725800": {Barcode: "C725800", Packages: []domain.Package{
			{Barcode: "P8988000122"}, {Barcode: "P8988000126"},
		}, UnloadAt: 2, State: 1},
	}

	deliveries := []domain.Delivery{
		{Barcode: "P8988000120"},
		{Barcode: "P8988000121"},
		{Barcode: "P8988000125"},
		{Barcode: "C725800"},
	}

	logger := logger.NewInstance()
	distributionService := distributionService.NewHandler(&packages, &sacks, logger)

	distributionService.Deliver(deliveries)

	correctDelivery := 0

	for _, item := range deliveries {
		if item.State == int(delivery_status.Unloaded) {
			correctDelivery++
		}
	}

	correctPackage := 0
	for _, item := range packages {
		if item.State == int(delivery_status.Unloaded) {
			correctPackage++
		}
	}

	assert.Equal(t, correctPackage, 4)
	assert.Equal(t, correctDelivery, 3)
	assert.Equal(t, sacks["C725800"].State, 4)
}

func TestUnloadService_Deliver_WithoutSacks_ShouldBeNoErrorAndOnlyPackagesUnloaded(t *testing.T) {

	packages := map[string]domain.Package{
		"P8988000120": {Barcode: "P8988000120", UnloadAt: 2, State: 1},
		"P8988000121": {Barcode: "P8988000121", UnloadAt: 2, State: 1},
		"P8988000124": {Barcode: "P8988000124", UnloadAt: 2, State: 1},
		"P8988000122": {Barcode: "P8988000122", UnloadAt: 2, State: 1},
		"P8988000126": {Barcode: "P8988000126", UnloadAt: 2, State: 1},
	}

	sacks := map[string]domain.Sack{}

	deliveries := []domain.Delivery{
		{Barcode: "P8988000120"},
		{Barcode: "P8988000121"},
		{Barcode: "P8988000125"},
		{Barcode: "C725800"},
	}

	logger := logger.NewInstance()
	distributionService := distributionService.NewHandler(&packages, &sacks, logger)

	distributionService.Deliver(deliveries)

	correctDelivery := 0

	for _, item := range deliveries {
		if item.State == int(delivery_status.Unloaded) {
			correctDelivery++
		}
	}

	correctPackage := 0
	for _, item := range packages {
		if item.State == int(delivery_status.Unloaded) {
			correctPackage++
		}
	}

	assert.Equal(t, correctPackage, 2)
	assert.Equal(t, correctDelivery, 2)
	assert.Zero(t, len(sacks))

}

func TestUnloadService_Deliver_WithNoData(t *testing.T) {

	packages := map[string]domain.Package{}

	sacks := map[string]domain.Sack{}

	deliveries := []domain.Delivery{}

	logger := logger.NewInstance()
	distributionService := distributionService.NewHandler(&packages, &sacks, logger)

	distributionService.Deliver(deliveries)

	correctDelivery := 0

	for _, item := range deliveries {
		if item.State == int(delivery_status.Unloaded) {
			correctDelivery++
		}
	}

	correctPackage := 0
	for _, item := range packages {
		if item.State == int(delivery_status.Unloaded) {
			correctPackage++
		}
	}

	assert.Zero(t, correctDelivery)
	assert.Zero(t, correctPackage)
	assert.Zero(t, len(sacks))
	assert.Zero(t, len(packages))
	assert.Zero(t, len(deliveries))
}
