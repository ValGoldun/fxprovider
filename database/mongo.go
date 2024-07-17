package database

import (
	"context"
	"github.com/ValGoldun/bsonregistry"
	"github.com/ValGoldun/fxprovider/appcontext"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
	"strings"
	"time"
)

type MongoConfig struct {
	Address  string
	Username string
	Password string
	Timeout  time.Duration
}

type Database struct {
	*mongo.Client
	hosts []string
}

func (db Database) Source() string {
	return strings.Join(db.hosts, ";")
}

func (db Database) HealthCheck() error {
	return db.Client.Ping(context.Background(), nil)
}

type Mongo interface {
	*mongo.Client
}

func ProvideMongo[M Mongo](appCtx *appcontext.AppContext, lc fx.Lifecycle, cfg MongoConfig) (M, error) {
	var mongoDatabase M

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	client, err := mongo.Connect(
		ctx,
		options.Client().SetAuth(options.Credential{
			Username: cfg.Username,
			Password: cfg.Password,
		}),
		options.Client().ApplyURI(cfg.Address).SetRegistry(bsonregistry.Registry()),
	)
	if err != nil {
		return mongoDatabase, err
	}

	lc.Append(fx.Hook{OnStop: func(ctx context.Context) error {
		return client.Disconnect(ctx)
	}})

	appCtx.WithHealthChecker(Database{
		client,
		options.Client().ApplyURI(cfg.Address).Hosts,
	})

	return client, nil
}
