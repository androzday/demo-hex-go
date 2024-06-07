package repostory

import (
	"context"
	"demo-hex-go/internal/adapter/db"
	"demo-hex-go/internal/core/domain/entity"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type PostgresRepository struct {
	db *db.PostgresRepository
}

type MongoRepository struct {
	db *db.MongoRepository
}

func NewPostgresRepository(repo *db.PostgresRepository) *PostgresRepository {
	return &PostgresRepository{
		db: repo,
	}
}

func NewMongoRepository(repo *db.MongoRepository) *MongoRepository {
	return &MongoRepository{
		db: repo,
	}
}

/*DB Postgres Repo CRUD*/
func (m PostgresRepository) SaveProduct(product entity.Product) error {

	req := m.db.Create(&product)
	if req.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("product not saved: %v", req.Error))
	}
	return nil
}

func (m *PostgresRepository) ReadProducts() ([]*entity.Product, error) {
	var products []*entity.Product
	req := m.db.Find(&products)
	if req.Error != nil {
		return nil, errors.New(fmt.Sprintf("products not found: %v", req.Error))
	}
	return products, nil
}

func (m *PostgresRepository) ReadProduct(id string) (*entity.Product, error) {
	product := &entity.Product{}
	req := m.db.First(&product, "id = ? ", id)
	if req.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("product not found: %v", req.Error))
	}
	return product, nil
}

func (m *PostgresRepository) UpdateProduct(product entity.Product) error {
	req := m.db.Save(product)
	if req.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("product not updated: %v", req.Error))
	}
	return nil
}

func (m *PostgresRepository) DeleteProduct(id string) error {
	product := &entity.Product{}
	req := m.db.Delete(&product, "id = ?", id)
	if req.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("product not deleted: %v", req.Error))
	}
	return nil
}

/**************************************************************************/
/*DB Mongo Repo CRUD*/
func (m MongoRepository) SaveProduct(product entity.Product) error {
	collection := m.db.Database("demo").Collection("product")
	_, req := collection.InsertOne(context.TODO(), product)
	if req != nil {
		return errors.New(fmt.Sprintf("product not saved: %v", req.Error))
	}
	return nil
}

func (m *MongoRepository) ReadProducts() ([]*entity.Product, error) {
	collection := m.db.Database("demo").Collection("product")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("products not found: %v", err.Error))
	}
	defer cursor.Close(context.TODO())

	var products []*entity.Product
	for cursor.Next(context.TODO()) {
		var product *entity.Product
		req := cursor.Decode(&product)
		if req != nil {
			return nil, errors.New(fmt.Sprintf("products not found: %v", req.Error))
		}
		products = append(products, product)
	}

	return products, nil
}
func (m *MongoRepository) ReadProduct(id string) (*entity.Product, error) {
	collection := m.db.Database("demo").Collection("product")

	var product *entity.Product
	filter := bson.M{"id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&product)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("product not found: %v", err))
	}
	return product, nil
}

func (m *MongoRepository) UpdateProduct(product entity.Product) error {
	collection := m.db.Database("demo").Collection("product")
	filter := bson.M{"id": product.Id}
	update := bson.M{"$set": product}

	_, req := collection.UpdateOne(context.Background(), filter, update)
	if req != nil {
		return errors.New(fmt.Sprintf("product not updated: %v", req.Error))
	}
	return nil
}

func (m *MongoRepository) DeleteProduct(id string) error {
	collection := m.db.Database("demo").Collection("product")
	filter := bson.M{"id": id}

	_, req := collection.DeleteOne(context.TODO(), filter)
	if req != nil {
		return errors.New(fmt.Sprintf("product not deleted: %v", req.Error))
	}
	return nil
}
