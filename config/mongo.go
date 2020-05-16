package config

import (
	"fmt"
	"github.com/geeks/miniproject/logger"
)

// mongo is the config for the db vars
type mongo struct {
	host         string
	port         string
	rootPassword string
	user         string
	userPassword string
	database     string
}

// load loads the config for the mongodb
func (mongoConfig *mongo) load() {
	logger.Log.Info("Reading mongo config...")
	mongoConfig.host = "localhost"
	mongoConfig.port = "27017"
	mongoConfig.rootPassword = ""
	mongoConfig.user = ""
	mongoConfig.userPassword = ""
	mongoConfig.database = "miniproject"
}

// Host returns the mongo host
func (mongoConfig *mongo) Host() string {
	return mongoConfig.host
}

// Port returns the mongo Port
func (mongoConfig *mongo) Port() string {
	return mongoConfig.port
}

// RootPassword returns the mongo RootPassword
func (mongoConfig *mongo) RootPassword() string {
	return mongoConfig.rootPassword
}

// User returns the mongo user
func (mongoConfig *mongo) User() string {
	return mongoConfig.user
}

// UserPassword returns the mongo userPassword
func (mongoConfig *mongo) UserPassword() string {
	return mongoConfig.userPassword
}

// Database returns the mongo Database
func (mongoConfig *mongo) Database() string {
	return mongoConfig.database
}

func (mongoConfig *mongo) ConnectionString() string {
	return fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", mongoConfig.user, mongoConfig.userPassword, mongoConfig.host, mongoConfig.port, mongoConfig.database)
}
