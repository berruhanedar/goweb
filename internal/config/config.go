package config

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	AppName  string `env:"APP_NAME" env-default:"tronics"`
	AppEnv   string `env:"APP_ENV" env-default:"development"`
	Port     string `env:"PORT" env-default:"8081"`
	Host     string `env:"HOST" env-default:"localhost"`
	LogLevel string `env:"LOG_LEVEL" env-default:"info"`
}

var Cfg Config
var mongoClient *mongo.Client
var once sync.Once

func LoadConfig() {
	if err := cleanenv.ReadConfig(".env", &Cfg); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	fmt.Printf("Config loaded: %+v\n", Cfg)
}

func GetMongoClient() *mongo.Client {
	once.Do(func() {
		var err error
		mongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			log.Fatalf("Mongo connect error: %v", err)
		}

		if err = mongoClient.Ping(context.TODO(), nil); err != nil {
			log.Fatalf("Mongo ping error: %v", err)
		}
		fmt.Println("MongoDB connected")
	})
	return mongoClient
}
