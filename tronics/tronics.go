package tronics

import (
	"context"
	"fmt"
	"log"

	"github.com/berruhanedar/goweb/config"
	"github.com/berruhanedar/goweb/middleware"
	"github.com/berruhanedar/goweb/models"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/go-playground/validator.v9"
)

var (
	e      = echo.New()
	v      = validator.New()
	cfg    config.ConfigDatabase
	client *mongo.Client
	Logger = logrus.New()
)

func init() {
	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		e.Logger.Fatal("Unable to load configuration")
	}
	fmt.Printf("%+v\n", cfg)
}

func connectToMongo() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}
	fmt.Println("Connected to MongoDB!")
}

func serverMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("inside middleware")
		return next(c)
	}
}

func serverMessageDo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("inside middleware Do")
		return next(c)
	}
}

func Start() {

	connectToMongo()

	models.SetMongoCollection(client)

	e.Validator = &models.ProductValidator{Validator: v}

	e.Use(serverMessage)
	e.Use(serverMessageDo)
	e.Use(middleware.LogrusMiddleware(Logger))

	e.GET("/products", models.GetProducts)
	e.GET("/products/:id", models.GetProduct)
	e.POST("/products", models.CreateProduct)
	e.PUT("/products/:id", models.UpdateProduct)
	e.DELETE("/products/:id", models.DeleteProduct)
	e.PATCH("/products/:id", models.PatchProduct)

	e.Logger.Print(fmt.Sprintf("Listening on port %s", cfg.Port))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
