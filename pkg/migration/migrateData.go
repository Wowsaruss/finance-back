package migration

import (
	"context"
	"fmt"
	"log"

	"github.com/Wowsaruss/financial-back-go/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// You will be using this Trainer type later in the program
// type Trainer struct {
//     Name string
//     Age  int
//     City string
// }

func migration() {
	fmt.Println("HIT!")
	cfg := config.NewConfig()

	clientOptions := options.Client().ApplyURI(cfg.MongoDBURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}
