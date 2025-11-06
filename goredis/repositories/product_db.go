package repositories

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type productRepositoryDB struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewProductRepositoryDB(db *gorm.DB, redisClient *redis.Client) ProductRepository {
	db.AutoMigrate(&product{})
	mockData(db)
	return productRepositoryDB{db: db, redisClient: redisClient}
}

// Method ใน productRepositoryDB
func (r productRepositoryDB) GetProducts() (products []product, err error) {
	err = r.db.Order("quantity desc").Limit(30).Find(&products).Error
	return products, err
}
