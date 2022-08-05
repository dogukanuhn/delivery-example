package interfaces

import (
	"context"

	"github.com/dogukanuhn/delivery-system/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type IPackageRepository interface {
	FindByPoint(ctx context.Context, point int) (map[string]domain.Package, error)
	BulkWrite(ctx context.Context, packages map[string]domain.Package) (*mongo.BulkWriteResult, error)
}

type ISackRepository interface {
	FindByPoint(ctx context.Context, point int) (map[string]domain.Sack, error)
	BulkWrite(ctx context.Context, sacks map[string]domain.Sack) (*mongo.BulkWriteResult, error)
}

type IDeliveryPointRepository interface {
	FindAll(ctx context.Context) (*[]domain.DeliveryPoint, error)
}
