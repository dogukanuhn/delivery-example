package commonRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindAggregate[T any](ctx context.Context, collection *mongo.Collection, query []bson.M, data *[]T) error {

	showLoadedCursor, err := collection.Aggregate(ctx, query)
	if err != nil {
		return err
	}

	err = showLoadedCursor.All(ctx, data)

	if err != nil {
		return err
	}
	return nil
}

func FindAll[T any](ctx context.Context, collection *mongo.Collection, filter primitive.D, data *[]T) error {

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return err
	}

	err = cur.All(ctx, data)

	if err != nil {
		return err
	}
	return nil
}

func BulkWrite(ctx context.Context, collection *mongo.Collection, domain []mongo.WriteModel) (*mongo.BulkWriteResult, error) {

	opts := options.BulkWrite().SetOrdered(true)

	return collection.BulkWrite(context.TODO(), domain, opts)

}
