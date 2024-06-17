package infrastructure

import (
	"database/sql"
	"fmt"
	"foodOrder/domain/entity"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

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
	} else if gc.Driver == "sqlite3" {
		return gc.SQLiteConnection()
	}

	return nil, fmt.Errorf("unsupported database driver: %s", gc.Driver)
}

func (gc *GormConfig) PostgresConnection() (*gorm.DB, error) {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", gc.Host, gc.User, gc.Password, gc.DBName, gc.Port)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    
    if err != nil {
        log.Printf("Error connecting to database: %v", err)
        
        // Connect to the default 'postgres' database to create the target database
        adminDSN := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%d sslmode=disable TimeZone=Asia/Shanghai", gc.Host, gc.User, gc.Password, gc.Port)
        adminDb, err := sql.Open("postgres", adminDSN)
        if err != nil {
            return nil, fmt.Errorf("failed to connect to the admin database: %w", err)
        }
        defer adminDb.Close()

        // Create the new database
        _, err = adminDb.Exec("CREATE DATABASE " + gc.DBName)
        if err != nil {
            return nil, fmt.Errorf("failed to create database: %w", err)
        }
        log.Printf("Database %s created successfully.", gc.DBName)

        // Reconnect to the newly created database using GORM
        db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err != nil {
            return nil, fmt.Errorf("failed to connect to the new database: %w", err)
        }

		gc.AutoMigrate(db)
		MockData(db)
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

func (gc *GormConfig) SQLiteConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s.db", "internal/db/" + gc.DBName)
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (gc *GormConfig) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Food{})
	db.AutoMigrate(&entity.Order{})
	db.AutoMigrate(&entity.TablePreference{})
	db.AutoMigrate(&entity.Cart{})
	db.AutoMigrate(&entity.Table{})
}

func MockData(db *gorm.DB) {
	CreateStaff(db)
	CreateAdmin(db)
	InitialTable(db)
}

func CreateStaff(db *gorm.DB) {
	db.Create(&entity.User{
		Username: "staff",
		Password: "123",
		Type:     "staff",
	})
}

func CreateAdmin(db *gorm.DB) {
	db.Create(&entity.User{
		Username: "cooker",
		Password: "123",
		Type:     "cooker",
	})
}

func InitialTable(db *gorm.DB) {
	for i := 1; i <= 10; i++ {
		db.Create(&entity.Table{
			TableNo: uint8(i),
			Status:  "available",
		})
	}
}