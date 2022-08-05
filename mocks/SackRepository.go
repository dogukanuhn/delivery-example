package mocks

import (
	"context"

	"github.com/dogukanuhn/delivery-system/domain"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockSackRepo struct {
	mock.Mock
}

func (m *MockSackRepo) FindByPoint(ctx context.Context, point int) (map[string]domain.Sack, error) {

	ret := m.Called(ctx, point)

	var r0 map[string]domain.Sack
	if rf, ok := ret.Get(0).(func(ctx context.Context, point int) map[string]domain.Sack); ok {
		r0 = rf(ctx, point)
	} else {
		r0 = ret.Get(0).(map[string]domain.Sack)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ctx context.Context, point int) error); ok {
		r1 = rf(ctx, point)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *MockSackRepo) BulkWrite(ctx context.Context, Sacks map[string]domain.Sack) (*mongo.BulkWriteResult, error) {

	ret := m.Called(ctx, Sacks)

	var r0 *mongo.BulkWriteResult
	if rf, ok := ret.Get(0).(func(ctx context.Context, Sacks map[string]domain.Sack) *mongo.BulkWriteResult); ok {
		r0 = rf(ctx, Sacks)
	} else {
		r0 = ret.Get(0).(*mongo.BulkWriteResult)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ctx context.Context, Sacks map[string]domain.Sack) error); ok {
		r1 = rf(ctx, Sacks)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
