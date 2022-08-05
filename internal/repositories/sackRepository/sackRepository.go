package sackRepository

import (
	"context"

	"github.com/dogukanuhn/delivery-system/domain"
	"github.com/dogukanuhn/delivery-system/internal/repositories/commonRepository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SackRepository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) *SackRepository {
	return &SackRepository{collection: collection}
}

func (r *SackRepository) FindByPoint(ctx context.Context, point int) (map[string]domain.Sack, error) {

	query := []bson.M{{
		"$lookup": bson.M{ // lookup the documents table here
			"from":         "package-sack",
			"localField":   "barcode",
			"foreignField": "sackBarcode",
			"as":           "package",
		}}, {"$match": bson.M{"unloadAt": point}}}

	var sacks []domain.Sack

	err := commonRepository.FindAggregate(ctx, r.collection, query, &sacks)

	if err != nil {
		return nil, err
	}

	sackMap := make(map[string]domain.Sack)

	for _, item := range sacks {
		sackMap[item.Barcode] = item
	}

	return sackMap, nil

}

func (r *SackRepository) BulkWrite(ctx context.Context, sacks map[string]domain.Sack) (*mongo.BulkWriteResult, error) {

	var domain []mongo.WriteModel

	for _, data := range sacks {

		replaceModel := mongo.NewUpdateOneModel().SetFilter(bson.D{{Key: "_id", Value: data.ID}}).SetUpdate(bson.D{
			{Key: "$set", Value: bson.D{{Key: "state", Value: data.State}}}})

		domain = append(domain, replaceModel)
	}

	return commonRepository.BulkWrite(ctx, r.collection, domain)
}
