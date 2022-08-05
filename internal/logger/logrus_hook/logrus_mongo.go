package logrus_hook

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogrusMongoHook struct {
	c *mongo.Collection
}

type M bson.M

func NewLogrusMongoHook(collection *mongo.Collection) (*LogrusMongoHook, error) {
	return &LogrusMongoHook{c: collection}, nil
}

func (h *LogrusMongoHook) Fire(entry *logrus.Entry) error {

	data := make(logrus.Fields)

	data["Level"] = entry.Level.String()
	data["Time"] = entry.Time
	data["Message"] = entry.Message

	for k, v := range entry.Data {
		if errData, isError := v.(error); logrus.ErrorKey == k && v != nil && isError {
			data[k] = errData.Error()
		} else {
			data[k] = v
		}
	}

	_, err := h.c.InsertOne(context.TODO(), M(data))

	if err != nil {
		return fmt.Errorf("Failed to send log entry to mongodb: %v", err)
	}

	return nil
}

func (h *LogrusMongoHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
