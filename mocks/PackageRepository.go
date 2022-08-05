package mocks

import (
	"context"

	"github.com/dogukanuhn/delivery-system/domain"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockPackageRepo struct {
	mock.Mock
}

func (m *MockPackageRepo) FindByPoint(ctx context.Context, point int) (map[string]domain.Package, error) {

	ret := m.Called(ctx, point)

	var r0 map[string]domain.Package
	if rf, ok := ret.Get(0).(func(ctx context.Context, point int) map[string]domain.Package); ok {
		r0 = rf(ctx, point)
	} else {
		r0 = ret.Get(0).(map[string]domain.Package)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ctx context.Context, point int) error); ok {
		r1 = rf(ctx, point)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (m *MockPackageRepo) BulkWrite(ctx context.Context, packages map[string]domain.Package) (*mongo.BulkWriteResult, error) {

	ret := m.Called(ctx, packages)

	var r0 *mongo.BulkWriteResult
	if rf, ok := ret.Get(0).(func(ctx context.Context, packages map[string]domain.Package) *mongo.BulkWriteResult); ok {
		r0 = rf(ctx, packages)
	} else {
		r0 = ret.Get(0).(*mongo.BulkWriteResult)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ctx context.Context, packages map[string]domain.Package) error); ok {
		r1 = rf(ctx, packages)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
