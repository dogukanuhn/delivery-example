package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Package struct {
	ID       primitive.ObjectID `bson:"_id"`
	Barcode  string             `bson:"barcode,omitempty"`
	Desi     string             `bson:"desi,omitempty"`
	UnloadAt int                `bson:"unloadAt,omitempty"`
	State    int                `bson:"state,omitempty"`
}
