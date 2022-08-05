package cfg

import (
	"context"
	"fmt"
	"log"

	"github.com/dogukanuhn/delivery-system/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/exp/slices"
)

// if database and collections are not initialized already, import datas and create db and collections
func MockDatabase() {

	var mockDataList = map[string][]interface{}{
		"delivery-points": {
			bson.D{{"name", "Branch"}, {"id", 1}},
			bson.D{{"name", "Distribution Centre"}, {"id", 2}},
			bson.D{{"name", "Transfer Centre"}, {"id", 3}},
		},
		"sacks": {
			bson.D{{"barcode", "C725799"}, {"state", 1}, {"unloadAt", 2}},
			bson.D{{"barcode", "C725800"}, {"state", 1}, {Key: "unloadAt", Value: 3}},
		},
		"packages": {
			bson.D{{"barcode", "P7988000121"}, {"unloadAt", 1}, {"state", 1}, {"desi", "5"}},
			bson.D{{"barcode", "P7988000122"}, {"unloadAt", 1}, {"state", 1}, {"desi", "5"}},
			bson.D{{"barcode", "P7988000123"}, {"unloadAt", 1}, {"state", 1}, {"desi", "9"}},
			bson.D{{"barcode", "P8988000120"}, {"unloadAt", 2}, {"state", 1}, {"desi", "33"}},
			bson.D{{"barcode", "P8988000121"}, {"unloadAt", 2}, {"state", 1}, {"desi", "17"}},
			bson.D{{"barcode", "P8988000122"}, {"unloadAt", 2}, {"state", 1}, {"desi", "26"}},
			bson.D{{"barcode", "P8988000124"}, {"unloadAt", 2}, {"state", 1}, {"desi", "35"}},
			bson.D{{"barcode", "P8988000125"}, {"unloadAt", 2}, {"state", 1}, {"desi", "1"}},
			bson.D{{"barcode", "P8988000123"}, {"unloadAt", 2}, {"state", 1}, {"desi", "200"}},
			bson.D{{"barcode", "P8988000126"}, {"unloadAt", 2}, {"state", 1}, {"desi", "50"}},
			bson.D{{"barcode", "P9988000126"}, {"unloadAt", 3}, {"state", 1}, {"desi", "15"}},
			bson.D{{"barcode", "P9988000127"}, {"unloadAt", 3}, {"state", 1}, {"desi", "16"}},
			bson.D{{"barcode", "P9988000128"}, {"unloadAt", 3}, {"state", 1}, {"desi", "55"}},
			bson.D{{"barcode", "P9988000129"}, {"unloadAt", 3}, {"state", 1}, {"desi", "28"}},
			bson.D{{"barcode", "P9988000130"}, {"unloadAt", 3}, {"state", 1}, {"desi", "17"}},
		},
		"package-sack": {
			bson.D{{"barcode", "P8988000122"}, {"sackBarcode", "C725799"}},
			bson.D{{"barcode", "P8988000126"}, {Key: "sackBarcode", Value: "C725799"}},
			bson.D{{"barcode", "P9988000128"}, {"sackBarcode", "C725800"}},
			bson.D{{"barcode", "P9988000129"}, {"sackBarcode", "C725800"}},
		},
	}

	result, err := GetDatabase().ListCollectionNames(
		context.TODO(),
		bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	query := []bson.M{{
		"$lookup": bson.M{ // lookup the documents table here
			"from":         "package-sack",
			"localField":   "barcode",
			"foreignField": "sackBarcode",
			"as":           "package",
		}}}

	showLoadedCursor, err := GetDatabase().Collection("sacks").Aggregate(context.TODO(), query)
	if err != nil {
		panic(err)
	}

	type PodcastEpisode struct {
		ID       primitive.ObjectID `bson:"_id,omitempty"`
		Package  []domain.Package   `bson:"package,omitempty"`
		Barcode  string             `bson:"barcode,omitempty"`
		UnloadAt int                `bson:"unloadAt,omitempty"`
	}

	var showsLoaded []PodcastEpisode

	if err = showLoadedCursor.All(context.TODO(), &showsLoaded); err != nil {
		panic(err)
	}

	if len(result) > 0 {
		log.Println("Database already migrated")
		return
	}

	for key, element := range mockDataList {

		if !slices.Contains(result, key) {
			log.Println(fmt.Sprintf("Collection is imported for '%s' ", key))
			_, err := GetDatabase().Collection(key).InsertMany(context.TODO(), element)

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	log.Println("Migration completed")

}
