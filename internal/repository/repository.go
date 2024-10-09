package repository

import (
	"chess/pkg/config"
	"chess/pkg/log"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	TX_KEY = "tx_key"
)

type Repository struct {
	db *gorm.DB
	//rdb    *redis.Client
	logger *log.Logger
}

func NewRepository(logger *log.Logger, db *gorm.DB) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}

func NewDb() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.ConfigInstance.GetString("data.mysql.user")))
	if err != nil {
		panic(err)
	}
	return db
}

func GetTxDB(ctx *gin.Context) (*gorm.DB, error) {
	txKey, _ := ctx.Get(TX_KEY)
	tx, isExist := ctx.Get(TX_KEY + "_" + txKey.(string))
	if !isExist {
		return nil, errors.New("tx key not found")
	}
	return tx.(*gorm.DB), nil
}
