package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Sack struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Barcode  string             `bson:"barcode,omitempty"`
	UnloadAt int                `bson:"unloadAt,omitempty"`
	State    int                `bson:"state,omitempty"`
	Packages []Package          `bson:"package,omitempty"`
}
