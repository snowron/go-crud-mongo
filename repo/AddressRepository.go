package repo

import (
	"addressproject/model"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AddressRepository struct {
	Client *mongo.Client
}

func (r AddressRepository) SearchAddressWithText(keywords string) []model.Address {

	collection := r.Client.Database("address").Collection("addressCollection")
	filter := bson.M{"$text": bson.M{"$search": keywords}}
	var result []model.Address
	results, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error: Database problem", err)
	}

	for results.Next(context.TODO()) {
		var iterate model.Address
		err := results.Decode(&iterate)
		if err != nil {

		}
		result = append(result, iterate)
		fmt.Println(iterate)
	}
	return result

}
func (r AddressRepository) FindAddressWithCoordinates(location model.Location) model.Address {
	collection := r.Client.Database("address").Collection("addressCollection")
	filter := bson.M{"Location.Latitude": location.Latitute, "Location.Longitude": location.Longitude}
	var result model.Address

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Printf("Error: Database Problem", err)
	}

	fmt.Println(result)
	return result
}

func (r AddressRepository) AddAddress(address model.Address) string {
	collection := r.Client.Database("address").Collection("addressCollection")
	asd, err := collection.InsertOne(context.TODO(), address)

	if err != nil {
		log.Printf("Error: Cannot insert to collection", err)
	}

	return asd.InsertedID.(primitive.ObjectID).Hex()
}

func (r AddressRepository) UpdateAddress(Id string, address model.Address) {
	collection := r.Client.Database("address").Collection("addressCollection")
	fmt.Println(address.Id)
	convertedId, _ := primitive.ObjectIDFromHex(Id)
	query := bson.D{{Key: "_id", Value: convertedId}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "City", Value: address.City},
				{Key: "Country", Value: address.Country},
				{Key: "FullText", Value: address.FullText},
				{Key: "Location", Value: bson.M{"Longitude": address.Location.Longitude, "Latitude": address.Location.Latitute}},
			},
		}}
	_, err := collection.UpdateOne(context.TODO(), query, update)

	if err != nil {
		log.Printf("Error: Cannot update address", err)
	}
}

func (r AddressRepository) DeleteAddress(Id string) {
	collection := r.Client.Database("address").Collection("addressCollection")
	fmt.Println(Id)
	objectID, err2 := primitive.ObjectIDFromHex(Id)
	if err2 != nil {
		log.Printf("Error: Cant Convert to object ID", err2)
	}
	query := bson.M{"_id": objectID}
	_, err := collection.DeleteOne(context.TODO(), query)

	if err != nil {
		log.Printf("Error: Cannot delete address", err)
	}
}
