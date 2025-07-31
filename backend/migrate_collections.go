package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Endale2/DRPS/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func migrateCollections() {
	// Get the products collection
	productsCollection := config.GetCollection("DRPS", "products")

	// Find all products that have the old collection_id field
	filter := bson.M{"collection_id": bson.M{"$exists": true}}
	cursor, err := productsCollection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal("Error finding products:", err)
	}
	defer cursor.Close(context.Background())

	var count int
	for cursor.Next(context.Background()) {
		var product bson.M
		if err := cursor.Decode(&product); err != nil {
			log.Printf("Error decoding product: %v", err)
			continue
		}

		// Get the old collection_id
		oldCollectionID, exists := product["collection_id"]
		if !exists {
			continue
		}

		// Convert to ObjectID if it's a string
		var collectionOID primitive.ObjectID
		switch v := oldCollectionID.(type) {
		case primitive.ObjectID:
			collectionOID = v
		case string:
			collectionOID, err = primitive.ObjectIDFromHex(v)
			if err != nil {
				log.Printf("Invalid collection_id format: %v", err)
				continue
			}
		default:
			log.Printf("Unexpected collection_id type: %T", oldCollectionID)
			continue
		}

		// Create the new collection_ids array
		newCollectionIDs := []primitive.ObjectID{collectionOID}

		// Update the product
		productID := product["_id"].(primitive.ObjectID)
		update := bson.M{
			"$set": bson.M{
				"collection_ids": newCollectionIDs,
			},
			"$unset": bson.M{
				"collection_id": "",
			},
		}

		_, err := productsCollection.UpdateOne(
			context.Background(),
			bson.M{"_id": productID},
			update,
		)
		if err != nil {
			log.Printf("Error updating product %s: %v", productID.Hex(), err)
			continue
		}

		count++
		fmt.Printf("Migrated product %s: collection_id %s -> collection_ids [%s]\n",
			productID.Hex(), collectionOID.Hex(), collectionOID.Hex())
	}

	fmt.Printf("Migration completed. Migrated %d products.\n", count)
}
