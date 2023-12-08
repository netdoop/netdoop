package store

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/netdoop/netdoop/utils"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

var defaultDB *gorm.DB
var defaultDBOnce sync.Once

var defaultMongoDB *mongo.Database
var defaultMongoDBOnce sync.Once

func initDB() {
	env := utils.GetEnv()
	dsn := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v sslmode=%v",
		env.GetString("db_username"),
		env.GetString("db_password"),
		env.GetString("db_host"),
		env.GetString("db_port"),
		env.GetString("db_db"),
		env.GetString("db_sslmode"),
	)
	var log logger.Interface
	if !utils.DebugMode {
		log = zapgorm2.Logger{
			ZapLogger:                 utils.GetLogger().Named("gorm"),
			LogLevel:                  logger.Warn,
			SlowThreshold:             time.Second,
			SkipCallerLookup:          false,
			IgnoreRecordNotFoundError: true,
			Context:                   nil,
		}
	} else {
		log = zapgorm2.Logger{
			ZapLogger:                 utils.GetLogger().Named("gorm"),
			LogLevel:                  logger.Info,
			SlowThreshold:             time.Second,
			SkipCallerLookup:          false,
			IgnoreRecordNotFoundError: false,
			Context:                   nil,
		}
	}
	db, err := OpenDatabase(dsn, log)
	if err != nil {
		utils.GetLogger().Fatal("db connection error", zap.String("dsn", dsn))
	}
	defaultDB = db
}

func GetDB() *gorm.DB {
	defaultDBOnce.Do(initDB)
	return defaultDB
}

func OpenDatabase(dsn string, logger logger.Interface, opts ...gorm.Option) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), opts...)
	if err != nil {
		return nil, errors.Wrap(err, "open database")
	}

	db.Logger = logger
	return db, nil
}

func initMongoDB() {
	env := utils.GetEnv()
	uri := env.GetString("mongodb_uri")
	database := env.GetString("mongodb_database")

	db, err := OpenMongoDB(uri, database)
	if err != nil {
		utils.GetLogger().Fatal("open mongo db error", zap.Error(err))
	}
	defaultMongoDB = db
}

func GetMongoDatabase() *mongo.Database {
	defaultMongoDBOnce.Do(initMongoDB)
	return defaultMongoDB
}

func OpenMongoDB(uri string, database string) (*mongo.Database, error) {
	clientOptions := options.Client().
		ApplyURI(uri).
		SetMaxPoolSize(10).                  // Set the maximum connection pool size
		SetMinPoolSize(5).                   // Set the minimum connection pool size
		SetMaxConnIdleTime(30 * time.Second) // Set the maximum idle time for a connection

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, errors.Wrap(err, "connect to mongo")
	}
	return client.Database(database), nil
}
