package packageRepository

import (
	"context"

	"github.com/dogukanuhn/delivery-system/domain"
	"github.com/dogukanuhn/delivery-system/internal/repositories/commonRepository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PackageRepository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) *PackageRepository {
	return &PackageRepository{collection: collection}
}

func (r *PackageRepository) FindByPoint(ctx context.Context, point int) (map[string]domain.Package, error) {

	var packages []domain.Package
	err := commonRepository.FindAll(ctx, r.collection, bson.D{{Key: "unloadAt", Value: point}}, &packages)

	if err != nil {
		return nil, err
	}

	packagesMap := make(map[string]domain.Package)

	for _, item := range packages {
		packagesMap[item.Barcode] = item
	}

	return packagesMap, nil
}

func (r *PackageRepository) BulkWrite(ctx context.Context, packages map[string]domain.Package) (*mongo.BulkWriteResult, error) {

	var domain []mongo.WriteModel

	for _, data := range packages {
		replaceModel := mongo.NewUpdateOneModel().SetFilter(bson.D{{Key: "barcode", Value: data.Barcode}}).SetUpdate(bson.D{
			{Key: "$set", Value: bson.D{{Key: "state", Value: data.State}}}})

		domain = append(domain, replaceModel)
	}

	return commonRepository.BulkWrite(ctx, r.collection, domain)
}
