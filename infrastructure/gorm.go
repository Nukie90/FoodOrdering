package infrastructure

import (
	"fmt"
	"foodOrder/domain/entity"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

/*
database:
  driver: postgres
  host: localhost
  port: 5432
  username: postgres
  password: postgres
  database: food_ordering
*/

type GormConfig struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func NewGormConfig(v *viper.Viper) *GormConfig {
	return &GormConfig{
		Driver:   v.GetString("database.driver"),
		Host:     v.GetString("database.host"),
		Port:     v.GetInt("database.port"),
		User:     v.GetString("database.username"),
		Password: v.GetString("database.password"),
		DBName:   v.GetString("database.database"),
	}
}

func (gc *GormConfig) Connection() (*gorm.DB, error) {
	if gc.Driver == "postgres" {
		return gc.PostgresConnection()
	} else if gc.Driver == "mysql" {
		return gc.MySQLConnection()
	}

	return nil, fmt.Errorf("unsupported database driver: %s", gc.Driver)
}

func (gc *GormConfig) PostgresConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", gc.Host, gc.User, gc.Password, gc.DBName, gc.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (gc *GormConfig) MySQLConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", gc.User, gc.Password, gc.Host, gc.Port, gc.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (gc *GormConfig) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Food{})
}