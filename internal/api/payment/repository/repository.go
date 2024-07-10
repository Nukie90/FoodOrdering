package repository

import (
	"foodOrder/domain/entity"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type PaymentRepo struct {
	Paymentdb *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) *PaymentRepo {
	return &PaymentRepo{Paymentdb: db}
}


func (r *PaymentRepo) GetOrderByTableNo(tableNo uint) ([]entity.Order, error) {
	var orders []entity.Order
	err := r.Paymentdb.Where("table_no = ?", tableNo).Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *PaymentRepo) GetFoodByID(foodId uint) entity.Food {
	dbTx := r.Paymentdb.Begin()
	defer dbTx.Rollback()

	var food entity.Food
	if err := dbTx.Where("food_id = ?", foodId).First(&food).Error; err != nil {
		return entity.Food{}
	}

	return food
}

func (r *PaymentRepo) GetPreferenceID(tableNo uint) (ulid.ULID, error) {
	dbTx := r.Paymentdb.Begin()
	defer dbTx.Rollback()

	var preference entity.TablePreference
	if err := dbTx.Where("table_no = ?", tableNo).First(&preference).Error; err != nil {
		return ulid.ULID{}, err
	}

	return preference.PreferenceID, nil
}

func (r *PaymentRepo) GetTableNo(preferenceID string) uint8 {
	dbTx := r.Paymentdb.Begin()
	defer dbTx.Rollback()

	preferenceIDULID, err := ulid.Parse(preferenceID)
	if err != nil {
		return 0
	}


	var preference entity.TablePreference
	if err := dbTx.Where("preference_id = ?", preferenceIDULID).First(&preference).Error; err != nil {
		return 0
	}

	return preference.TableNo
}

func (r *PaymentRepo) CreateBill(bill *entity.Payment) (ulid.ULID, error) {
	dbTx := r.Paymentdb.Begin()
	defer dbTx.Rollback()

	err := r.Paymentdb.Create(bill).Error
	if err != nil {
		return ulid.ULID{}, err
	}

	tableNo := r.GetTableNo(bill.PreferenceID.String())
	
	var orders []entity.Order
	err = r.Paymentdb.Where("table_no = ?", tableNo).Find(&orders).Error
	if err != nil {
		return ulid.ULID{}, err
	}

	for _, v := range orders {
		if v.Status == "done" || v.Status == "end" {
			v.Status = "paid"
		}else{
			v.Status = "cancel"
		}
		err = r.Paymentdb.Save(&v).Error
		if err != nil {
			return ulid.ULID{}, err
		}
	}

	var preferenceTable entity.TablePreference
	err = r.Paymentdb.Where("table_no = ?", tableNo).Find(&preferenceTable).Error
	if err != nil {
		return ulid.ULID{}, err
	}
	preferenceTable.Status = "paid"
	err = r.Paymentdb.Save(&preferenceTable).Error
	if err != nil {
		return ulid.ULID{}, err
	}
	
	var table entity.Table
	err = r.Paymentdb.Where("table_no = ?", tableNo).Find(&table).Error
	if err != nil {
		return ulid.ULID{}, err
	}
	table.Status = "available"
	if err := dbTx.Save(&table).Error; err != nil {
		return ulid.ULID{}, err
	}

	return bill.ReceiptID, dbTx.Commit().Error
}