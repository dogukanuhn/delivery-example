package cfg

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var lock = &sync.Mutex{}

var Database *mongo.Database

func GetDatabase() *mongo.Database {

	if Database == nil {
		lock.Lock()
		defer lock.Unlock()
		if Database == nil {

			client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
			if err != nil {
				log.Fatal(err)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			err = client.Connect(ctx)
			if err != nil {
				log.Fatal(err)
			}

			err = client.Ping(ctx, readpref.Primary())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Database connected!")

			Database = client.Database(os.Getenv("Database"))

		} else {
			fmt.Println("Db instance already created")
		}
	} else {
		fmt.Println("Db instance already created")
	}

	return Database

}
