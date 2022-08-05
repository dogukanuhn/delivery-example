package branchService_test

import (
	"testing"

	"github.com/dogukanuhn/delivery-system/domain"
	"github.com/dogukanuhn/delivery-system/domain/types/delivery_status"
	"github.com/dogukanuhn/delivery-system/internal/logger"
	"github.com/dogukanuhn/delivery-system/internal/services/branchService"
	"github.com/stretchr/testify/assert"
)

///////////////////////////////////
// BRANCH DELIVERY TEST
///////////////////////////////////
func TestUnloadService_Deliver_WithoutPackages_ShouldBeNoErrorAndNoUnloaded(t *testing.T) {

	packages := map[string]domain.Package{}

	sacks := map[string]domain.Sack{
		"C725800": {Barcode: "C725800", Packages: nil, UnloadAt: 1, State: 1},
	}

	deliveries := []domain.Delivery{
		{Barcode: "P8988000123"},
		{Barcode: "P8988000124"},
		{Barcode: "P8988000125"},
		{Barcode: "C725799"},
	}

	logger := logger.NewInstance()
	branchService := branchService.NewHandler(&packages, logger)

	branchService.Deliver(deliveries)

	correct := 0

	for _, item := range deliveries {
		if item.State == int(delivery_status.Unloaded) {
			correct++
		}
	}

	assert.Zero(t, correct)
	assert.Equal(t, sacks["C725800"].State, 1)
}

func TestUnloadService_Deliver_PackagesUnloadedAndSacksNotUnloaded(t *testing.T) {

	packages := map[string]domain.Package{
		"P9988000126": {Barcode: "P9988000126", UnloadAt: 1, State: 1},
		"P9988000127": {Barcode: "P9988000127", UnloadAt: 1, State: 1},
	}

	sacks := map[string]domain.Sack{
		"C725800": {Barcode: "C725800", Packages: nil, UnloadAt: 1, State: 1},
	}

	deliveries := []domain.Delivery{
		{Barcode: "P9988000126"},
		{Barcode: "P9988000127"},
		{Barcode: "P8988000125"},
		{Barcode: "C725799"},
	}

	logger := logger.NewInstance()
	branchService := branchService.NewHandler(&packages, logger)

	branchService.Deliver(deliveries)

	correct := 0

	for _, item := range deliveries {
		if item.State == int(delivery_status.Unloaded) {
			correct++
		}
	}

	assert.Equal(t, correct, 2)
	assert.Equal(t, sacks["C725800"].State, 1)
}

func TestUnloadService_Deliver_WithNoData(t *testing.T) {

	packages := map[string]domain.Package{}

	sacks := map[string]domain.Sack{}

	deliveries := []domain.Delivery{}

	logger := logger.NewInstance()
	branchService := branchService.NewHandler(&packages, logger)

	branchService.Deliver(deliveries)

	correct := 0

	for _, item := range deliveries {
		if item.State == int(delivery_status.Unloaded) {
			correct++
		}
	}

	assert.Zero(t, correct)
	assert.Zero(t, len(sacks))
	assert.Zero(t, len(packages))
	assert.Zero(t, len(deliveries))
}

///////////////////////////////////
// BRANCH DELIVERY TEST END
///////////////////////////////////
