package repository

import (
	"foodOrder/domain/entity"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type RestRepo struct {
	restDb *gorm.DB
}

func NewRestRepo(db *gorm.DB) *RestRepo {
	return &RestRepo{restDb: db}
}

func (r *RestRepo) InitialTable(table *entity.Table) error {
	dbTx := r.restDb.Begin()
	defer dbTx.Rollback()

	if err := dbTx.Create(table).Error; err != nil {
		return err
	}

	return dbTx.Commit().Error
}

func (r *RestRepo) GetALLTable() ([]entity.Table, error) {
	var tables []entity.Table
	if err := r.restDb.Find(&tables).Error; err != nil {
		return nil, err
	}

	return tables, nil
}

func (r *RestRepo) GetTable(tableNo uint8) (*entity.Table, error) {
	var table entity.Table
	if err := r.restDb.Where("table_no = ?", tableNo).First(&table).Error; err != nil {
		return nil, err
	}

	return &table, nil
}

func (r *RestRepo) GiveCustomerTable(tableNo uint8) (ulid.ULID, error) {
	dbTx := r.restDb.Begin()
	defer dbTx.Rollback()

	var table entity.Table
	if err := dbTx.Where("table_no = ?", tableNo).First(&table).Error; err != nil {
		return ulid.ULID{}, err
	}

	table.Status = "occupied"
	if err := dbTx.Save(&table).Error; err != nil {
		return ulid.ULID{}, err
	}

	preference := entity.TablePreference{
		TableNo: tableNo,
	}

	if err := dbTx.Create(&preference).Error; err != nil {
		return ulid.ULID{}, err
	}

	return preference.PreferenceID, dbTx.Commit().Error
}