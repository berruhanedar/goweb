package service

import (
	"context"

	"github.com/berruhanedar/goweb/internal/config"
	"github.com/berruhanedar/goweb/internal/errors"
	"github.com/berruhanedar/goweb/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = config.GetMongoClient().Database("tronicsdb").Collection("products")

func parseObjectID(id string) (primitive.ObjectID, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, errors.NewBadRequest("Invalid product ID format")
	}
	return objID, nil
}

func GetAllProducts() ([]model.Product, error) {
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, errors.NewInternal("Failed to fetch products from database")
	}
	var products []model.Product
	if err := cursor.All(context.TODO(), &products); err != nil {
		return nil, errors.NewInternal("Failed to decode products")
	}
	return products, nil
}

func GetProductByID(id string) (*model.Product, error) {
	objID, err := parseObjectID(id)
	if err != nil {
		return nil, err
	}
	var product model.Product
	if err := collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&product); err != nil {
		return nil, errors.NewNotFound("Product not found")
	}
	return &product, nil
}

func CreateProduct(product model.Product) (*model.Product, error) {
	res, err := collection.InsertOne(context.TODO(), product)
	if err != nil {
		return nil, errors.NewInternal("Failed to create product")
	}
	product.ID = res.InsertedID.(primitive.ObjectID)
	return &product, nil
}

func UpdateProduct(id string, updateData map[string]interface{}) error {
	objID, err := parseObjectID(id)
	if err != nil {
		return err
	}
	if len(updateData) == 0 {
		return errors.NewBadRequest("No fields provided for update")
	}
	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, bson.M{"$set": updateData})
	if err != nil {
		return errors.NewInternal("Failed to update product")
	}
	return nil
}

func DeleteProduct(id string) error {
	objID, err := parseObjectID(id)
	if err != nil {
		return err
	}
	res, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return errors.NewInternal("Failed to delete product")
	}
	if res.DeletedCount == 0 {
		return errors.NewNotFound("Product not found")
	}
	return nil
}

func PatchProduct(id string, fields map[string]interface{}) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.NewBadRequest("Invalid product ID")
	}
	if len(fields) == 0 {
		return errors.NewBadRequest("No fields to patch")
	}

	update := bson.M{"$set": fields}
	result, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if err != nil {
		return errors.NewInternal("Failed to patch product")
	}

	if result.MatchedCount == 0 {
		return errors.NewNotFound("Product not found")
	}

	return nil
}
