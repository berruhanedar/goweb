package models

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/go-playground/validator.v9"
)

type Product struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name" validate:"required,min=3"`
	Price float64            `json:"price" bson:"price" validate:"required,gt=0"`
}

type ProductValidator struct {
	Validator *validator.Validate
}

func (p *ProductValidator) Validate(i interface{}) error {
	return p.Validator.Struct(i)
}

var ProductCollection *mongo.Collection

func SetMongoCollection(client *mongo.Client) {
	ProductCollection = client.Database("tronicsdb").Collection("products")
}

// GET /products
func GetProducts(c echo.Context) error {
	cursor, err := ProductCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error fetching products")
	}
	defer cursor.Close(context.Background())

	var products []Product
	if err := cursor.All(context.Background(), &products); err != nil {
		return c.JSON(http.StatusInternalServerError, "Error decoding products")
	}

	return c.JSON(http.StatusOK, products)
}

// GET /products/:id
func GetProduct(c echo.Context) error {
	idParam := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	var product Product
	err = ProductCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&product)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	return c.JSON(http.StatusOK, product)
}

// POST /products
func CreateProduct(c echo.Context) error {
	var newProduct Product

	if err := c.Bind(&newProduct); err != nil {
		return err
	}
	if err := c.Validate(newProduct); err != nil {
		return err
	}

	res, err := ProductCollection.InsertOne(context.Background(), newProduct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error inserting product")
	}

	newProduct.ID = res.InsertedID.(primitive.ObjectID)
	return c.JSON(http.StatusCreated, newProduct)
}

// PUT /products/:id
func UpdateProduct(c echo.Context) error {
	idParam := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	var updatedProduct Product
	if err := c.Bind(&updatedProduct); err != nil {
		return err
	}
	if err := c.Validate(updatedProduct); err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"name":  updatedProduct.Name,
			"price": updatedProduct.Price,
		},
	}

	_, err = ProductCollection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error updating product")
	}

	updatedProduct.ID = objID
	return c.JSON(http.StatusOK, updatedProduct)
}

// DELETE /products/:id
func DeleteProduct(c echo.Context) error {
	idParam := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	res, err := ProductCollection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error deleting product")
	}
	if res.DeletedCount == 0 {
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	return c.JSON(http.StatusOK, "Product deleted successfully")
}

func PatchProduct(c echo.Context) error {
	idParam := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	var updates bson.M = make(bson.M)
	var body map[string]interface{}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid body")
	}

	// Sadece desteklenen alanlar güncellenmeli
	if name, ok := body["name"].(string); ok && name != "" {
		if len(name) < 3 {
			return c.JSON(http.StatusBadRequest, "Name must be at least 3 characters long")
		}
		updates["name"] = name
	}

	if price, ok := body["price"].(float64); ok {
		if price <= 0 {
			return c.JSON(http.StatusBadRequest, "Price must be greater than 0")
		}
		updates["price"] = price
	}

	if len(updates) == 0 {
		return c.JSON(http.StatusBadRequest, "No valid fields provided to update")
	}

	update := bson.M{"$set": updates}
	result, err := ProductCollection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error updating product")
	}
	if result.MatchedCount == 0 {
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Product updated successfully",
		"id":      idParam,
		"updated": updates,
	})
}
