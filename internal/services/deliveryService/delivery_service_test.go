package deliveryService_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dogukanuhn/delivery-system/domain"
	"github.com/dogukanuhn/delivery-system/domain/dto"
	"github.com/dogukanuhn/delivery-system/domain/types/delivery_status"
	"github.com/dogukanuhn/delivery-system/internal/logger"
	"github.com/dogukanuhn/delivery-system/internal/services/deliveryService"
	"github.com/dogukanuhn/delivery-system/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestHandleDeliveryBranch(t *testing.T) {

	mockPackageRepo := new(mocks.MockPackageRepo)
	mockSackRepo := new(mocks.MockSackRepo)

	packages := map[string]domain.Package{
		"P7988000121": {Barcode: "P7988000121", UnloadAt: 1, State: 1},
		"P7988000122": {Barcode: "P7988000122", UnloadAt: 1, State: 1},
		"P7988000123": {Barcode: "P7988000123", UnloadAt: 1, State: 1},
	}

	sacks := map[string]domain.Sack{}

	mockPackageRepo.On("FindByPoint", context.Background(), 1).Return(packages, nil)
	mockSackRepo.On("FindByPoint", context.Background(), 1).Return(sacks, nil)

	mockPackageRepo.On("BulkWrite", context.Background(), packages).Return(&mongo.BulkWriteResult{}, nil)
	mockSackRepo.On("BulkWrite", context.Background(), sacks).Return(&mongo.BulkWriteResult{}, nil)

	log := logger.NewInstance()

	service := deliveryService.NewHandler(mockPackageRepo, mockSackRepo, log)

	delivery := domain.Route{
		DeliveryPoint: 1,
		Deliveries: []domain.Delivery{
			{Barcode: "P7988000121"},
			{Barcode: "P7988000122"},
			{Barcode: "P7988000123"},
			{Barcode: "P8988000121"},
			{Barcode: "C725799"},
		},
	}

	deliveryRoute := dto.DeliverDTO{
		Route: []domain.Route{delivery},
	}

	service.Deliver(deliveryRoute)

	deliveredCount := 0
	notDeliveredCount := 0
	for _, item := range delivery.Deliveries {
		if item.State == int(delivery_status.Unloaded) {
			deliveredCount++
		} else {
			notDeliveredCount++
		}
	}

	assert.Equal(t, 3, deliveredCount)
	assert.Equal(t, 2, notDeliveredCount)
}

func TestHandleDeliveryDistribution(t *testing.T) {

	mockPackageRepo := new(mocks.MockPackageRepo)
	mockSackRepo := new(mocks.MockSackRepo)

	packages := map[string]domain.Package{
		"P8988000120": {Barcode: "P8988000120", UnloadAt: 2, State: 1},
		"P8988000121": {Barcode: "P8988000121", UnloadAt: 2, State: 1},
		"P8988000122": {Barcode: "P8988000122", UnloadAt: 2, State: 1},
		"P8988000124": {Barcode: "P8988000124", UnloadAt: 2, State: 1},
		"P8988000125": {Barcode: "P8988000125", UnloadAt: 2, State: 1},
		"P8988000123": {Barcode: "P8988000123", UnloadAt: 2, State: 1},
		"P8988000126": {Barcode: "P8988000126", UnloadAt: 2, State: 1},
	}

	packagesInSacks := []domain.Package{
		{Barcode: "P8988000122", UnloadAt: 2, State: 1},
		{Barcode: "P8988000126", UnloadAt: 2, State: 1},
	}

	sacks := map[string]domain.Sack{
		"C725799": {Barcode: "C725799", Packages: packagesInSacks, UnloadAt: 2, State: 1},
	}

	mockPackageRepo.On("FindByPoint", context.Background(), 2).Return(packages, nil)
	mockSackRepo.On("FindByPoint", context.Background(), 2).Return(sacks, nil)

	mockPackageRepo.On("BulkWrite", context.Background(), packages).Return(&mongo.BulkWriteResult{}, nil)
	mockSackRepo.On("BulkWrite", context.Background(), sacks).Return(&mongo.BulkWriteResult{}, nil)

	log := logger.NewInstance()

	service := deliveryService.NewHandler(mockPackageRepo, mockSackRepo, log)

	delivery := domain.Route{
		DeliveryPoint: 2,
		Deliveries: []domain.Delivery{
			{Barcode: "P8988000123"},
			{Barcode: "P8988000124"},
			{Barcode: "P8988000125"},
			{Barcode: "C725799"},
		},
	}

	deliveryRoute := dto.DeliverDTO{
		Route: []domain.Route{delivery},
	}

	service.Deliver(deliveryRoute)

	deliveredCount := 0
	notDeliveredCount := 0

	for _, item := range delivery.Deliveries {
		if item.State == int(delivery_status.Unloaded) {
			deliveredCount++
		} else {
			notDeliveredCount++
		}
	}

	assert.Equal(t, 4, deliveredCount)
	assert.Equal(t, sacks["C725799"].State, 4)
}

func TestHandleDeliveryTransfer(t *testing.T) {

	mockPackageRepo := new(mocks.MockPackageRepo)
	mockSackRepo := new(mocks.MockSackRepo)

	packages := map[string]domain.Package{
		"P9988000126": {Barcode: "P9988000126", UnloadAt: 3, State: 1},
		"P9988000127": {Barcode: "P9988000127", UnloadAt: 3, State: 1},
		"P9988000128": {Barcode: "P9988000128", UnloadAt: 3, State: 1},
		"P9988000129": {Barcode: "P9988000129", UnloadAt: 3, State: 1},
		"P9988000130": {Barcode: "P9988000130", UnloadAt: 3, State: 1},
	}

	packagesInSacks := []domain.Package{
		{Barcode: "P9988000128", UnloadAt: 3, State: 1},
		{Barcode: "P9988000129", UnloadAt: 3, State: 1},
	}

	sacks := map[string]domain.Sack{
		"C725800": {Barcode: "C725800", Packages: packagesInSacks, UnloadAt: 3, State: 1},
	}

	mockPackageRepo.On("FindByPoint", context.Background(), 3).Return(packages, nil)
	mockSackRepo.On("FindByPoint", context.Background(), 3).Return(sacks, nil)

	mockPackageRepo.On("BulkWrite", context.Background(), packages).Return(&mongo.BulkWriteResult{}, nil)
	mockSackRepo.On("BulkWrite", context.Background(), sacks).Return(&mongo.BulkWriteResult{}, nil)

	log := logger.NewInstance()

	service := deliveryService.NewHandler(mockPackageRepo, mockSackRepo, log)

	delivery := domain.Route{
		DeliveryPoint: 3,
		Deliveries: []domain.Delivery{
			{Barcode: "P9988000126"},
			{Barcode: "P9988000127"},
			{Barcode: "P9988000128"},
			{Barcode: "P9988000129"},
			{Barcode: "P9988000130"},
			{Barcode: "P9988000131"},
		},
	}

	deliveryRoute := dto.DeliverDTO{
		Route: []domain.Route{delivery},
	}

	service.Deliver(deliveryRoute)

	deliveredCount := 0
	notDeliveredCount := 0

	for _, item := range delivery.Deliveries {
		if item.State == int(delivery_status.Unloaded) {
			deliveredCount++
		} else {
			notDeliveredCount++
		}
	}
	assert.Equal(t, 2, deliveredCount)
	assert.Equal(t, sacks["C725800"].State, 4)
}

func TestDeliveryService_UpdateDb(t *testing.T) {

	mockPackageRepo := new(mocks.MockPackageRepo)
	mockSackRepo := new(mocks.MockSackRepo)

	packages := map[string]domain.Package{
		"P9988000126": {Barcode: "P9988000126", UnloadAt: 3, State: 4},
		"P9988000127": {Barcode: "P9988000127", UnloadAt: 3, State: 4},
		"P9988000128": {Barcode: "P9988000128", UnloadAt: 3, State: 4},
		"P9988000129": {Barcode: "P9988000129", UnloadAt: 3, State: 4},
		"P9988000130": {Barcode: "P9988000130", UnloadAt: 3, State: 4},
	}

	packagesInSacks := []domain.Package{
		{Barcode: "P9988000128", UnloadAt: 3, State: 4},
		{Barcode: "P9988000129", UnloadAt: 3, State: 4},
	}

	sacks := map[string]domain.Sack{
		"C725800": {Barcode: "C725800", Packages: packagesInSacks, UnloadAt: 3, State: 4},
	}

	mockPackageRepo.On("FindByPoint", context.Background(), 3).Return(packages, nil)
	mockSackRepo.On("FindByPoint", context.Background(), 3).Return(sacks, nil)

	mockPackageRepo.On("BulkWrite", context.Background(), packages).Return(&mongo.BulkWriteResult{}, nil)
	mockSackRepo.On("BulkWrite", context.Background(), sacks).Return(&mongo.BulkWriteResult{}, nil)

	log := logger.NewInstance()

	service := deliveryService.NewHandler(mockPackageRepo, mockSackRepo, log)
	service.Packages = packages
	service.Sacks = sacks

	err := service.UpdateDb()

	assert.NoError(t, err)
}

func TestDeliveryService_UpdateDb_Error(t *testing.T) {

	mockPackageRepo := new(mocks.MockPackageRepo)
	mockSackRepo := new(mocks.MockSackRepo)

	packages := map[string]domain.Package{
		"P9988000126": {Barcode: "P9988000126", UnloadAt: 3, State: 4},
		"P9988000127": {Barcode: "P9988000127", UnloadAt: 3, State: 4},
		"P9988000128": {Barcode: "P9988000128", UnloadAt: 3, State: 4},
		"P9988000129": {Barcode: "P9988000129", UnloadAt: 3, State: 4},
		"P9988000130": {Barcode: "P9988000130", UnloadAt: 3, State: 4},
	}

	packagesInSacks := []domain.Package{
		{Barcode: "P9988000128", UnloadAt: 3, State: 4},
		{Barcode: "P9988000129", UnloadAt: 3, State: 4},
	}

	sacks := map[string]domain.Sack{
		"C725800": {Barcode: "C725800", Packages: packagesInSacks, UnloadAt: 3, State: 4},
	}

	mockPackageRepo.On("FindByPoint", context.Background(), 3).Return(packages, nil)
	mockSackRepo.On("FindByPoint", context.Background(), 3).Return(sacks, nil)

	mockPackageRepo.On("BulkWrite", context.Background(), packages).Return(&mongo.BulkWriteResult{}, errors.New("fail"))
	mockSackRepo.On("BulkWrite", context.Background(), sacks).Return(&mongo.BulkWriteResult{}, errors.New("fail"))

	log := logger.NewInstance()

	service := deliveryService.NewHandler(mockPackageRepo, mockSackRepo, log)
	service.Packages = packages
	service.Sacks = sacks

	err := service.UpdateDb()

	assert.Error(t, err)
}
