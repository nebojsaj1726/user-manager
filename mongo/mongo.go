package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database interface {
	Collection(string) Collection
	Client() Client
}

type Collection interface {
	InsertOne(context.Context, interface{}) (interface{}, error)
	DeleteOne(context.Context, interface{}) (int64, error)
	Find(context.Context, interface{}, ...*options.FindOptions) (Cursor, error)
	UpdateOne(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)
}

type SingleResult interface {
	Decode(interface{}) error
}

type Cursor interface {
	Close(context.Context) error
	Next(context.Context) bool
	Decode(interface{}) error
	All(context.Context, interface{}) error
}

type Client interface {
	Database(string) Database
	Disconnect(context.Context) error
	StartSession() (mongo.Session, error)
	UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error
	Ping(context.Context) error
}

type mongoClient struct {
	cl *mongo.Client
}
type mongoDatabase struct {
	db *mongo.Database
}
type mongoCollection struct {
	coll *mongo.Collection
}

type mongoSingleResult struct {
	sr *mongo.SingleResult
}

type mongoCursor struct {
	mc *mongo.Cursor
}

type mongoSession struct {
	mongo.Session
}

func NewClient(connection string) (Client, error) {

	time.Local = time.UTC
	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connection))

	return &mongoClient{cl: c}, err

}

func (mc *mongoClient) Ping(ctx context.Context) error {
	return mc.cl.Ping(ctx, readpref.Primary())
}

func (mc *mongoClient) Database(dbName string) Database {
	db := mc.cl.Database(dbName)
	return &mongoDatabase{db: db}
}

func (mc *mongoClient) UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error {
	return mc.cl.UseSession(ctx, fn)
}

func (mc *mongoClient) StartSession() (mongo.Session, error) {
	session, err := mc.cl.StartSession()
	return &mongoSession{session}, err
}

func (mc *mongoClient) Disconnect(ctx context.Context) error {
	return mc.cl.Disconnect(ctx)
}

func (md *mongoDatabase) Collection(colName string) Collection {
	collection := md.db.Collection(colName)
	return &mongoCollection{coll: collection}
}

func (md *mongoDatabase) Client() Client {
	client := md.db.Client()
	return &mongoClient{cl: client}
}

func (mc *mongoCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return mc.coll.UpdateOne(ctx, filter, update, opts[:]...)
}

func (mc *mongoCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	id, err := mc.coll.InsertOne(ctx, document)
	return id.InsertedID, err
}

func (mc *mongoCollection) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
	count, err := mc.coll.DeleteOne(ctx, filter)
	return count.DeletedCount, err
}

func (mc *mongoCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (Cursor, error) {
	findResult, err := mc.coll.Find(ctx, filter, opts...)
	return &mongoCursor{mc: findResult}, err
}

func (sr *mongoSingleResult) Decode(v interface{}) error {
	return sr.sr.Decode(v)
}

func (mr *mongoCursor) Close(ctx context.Context) error {
	return mr.mc.Close(ctx)
}

func (mr *mongoCursor) Next(ctx context.Context) bool {
	return mr.mc.Next(ctx)
}

func (mr *mongoCursor) Decode(v interface{}) error {
	return mr.mc.Decode(v)
}

func (mr *mongoCursor) All(ctx context.Context, result interface{}) error {
	return mr.mc.All(ctx, result)
}
