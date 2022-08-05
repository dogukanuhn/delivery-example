package deliverPointRepository

import (
	"context"

	"github.com/dogukanuhn/delivery-system/domain"
	"github.com/dogukanuhn/delivery-system/internal/repositories/commonRepository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeliveryPointRepository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) *DeliveryPointRepository {
	return &DeliveryPointRepository{collection: collection}
}

func (r *DeliveryPointRepository) FindAll(ctx context.Context) (*[]domain.DeliveryPoint, error) {

	var deliveryPoints []domain.DeliveryPoint
	commonRepository.FindAll(ctx, r.collection, bson.D{{}}, &deliveryPoints)

	return &deliveryPoints, nil
}
