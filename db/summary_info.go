package db

import (
	"context"
	 "fmt"
	 

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/hendry19901990/stori-tech-challenge/mappers"
)


// Function to save SummaryInfo data in MongoDB
func SaveSummaryInfo(summary mappers.SummaryInfo) error {
	// Set up MongoDB connection
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	// Access the database and collection
	collection := client.Database("db_txs").Collection("summary_info")

	// Insert the SummaryInfo data into MongoDB
	_, err = collection.InsertOne(context.Background(), summary)
	if err != nil {
		return err
	}

	fmt.Println("SummaryInfo data saved in MongoDB!")
	return nil
}
