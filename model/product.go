package model

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

var ProductM *gorm.DB

const ProductTable = "product"

type Product struct {
	ID           int64     `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"id,omitempty"`
	Name         string    `gorm:"unique_index:name_p" json:"name,omitempty"`
	Price        int32     `json:"price,omitempty" gorm:"column:price"`
	Weight       float64   `json:"weight,omitempty" gorm:"type:decimal(10,2)"`
	Ticket       int32     `json:"ticket,omitempty" gorm:"column:ticket"`
	Freight      int32     `json:"freight,omitempty" gorm:"column:freight"`
	Fare         int32     `json:"fare,omitempty" gorm:"column:fare"`
	People       int32     `json:"people,omitempty" gorm:"column:people"`
	ExchangeRate float64   `json:"exchange_rate,omitempty" gorm:"type:decimal(10,2)"`
	Profit       int32     `gorm:"unique_index:name_p" json:"profit,omitempty"`
	Status       int32     `gorm:"type:int(1);" json:"status,omitempty"`
	CreatedAt    time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt    time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

func CreateQuote(p *Product) error {
	res := ProductM.Table(ProductTable).Create(&p)
	return res.Error
}

func GetQuote(name string, offset int) ([]Product, error) {
	var data []Product
	//先寫死分頁數量
	limit := 10
	offset = offset * 10
	res := ProductM.Table(ProductTable).
		Where("name like ?", "%"+name+"%").
		Offset(offset).
		Limit(limit).
		Scan(&data)

	return data, res.Error
}

func DeleteQuote(id int64) error {
	res := ProductM.Table(ProductTable).
		Where("id = ?", id).
		Delete(&Product{})

	return res.Error
}

func UpdateQuote(p *Product) error {
	updateDate := make(map[string]interface{})
	if p.ID == 0 {
		return errors.New("ID 有誤")
	}

	if p.Price != 0 {
		updateDate["price"] = p.Price
	}
	if p.Weight != 0 {
		updateDate["weight"] = p.Weight
	}
	if p.Ticket != 0 {
		updateDate["Ticket"] = p.Ticket
	}
	if p.Freight != 0 {
		updateDate["freight"] = p.Freight
	}
	if p.Fare != 0 {
		updateDate["fare"] = p.Fare
	}
	if p.Status != 0 {
		updateDate["status"] = p.Status
	}
	if p.Profit != 0 {
		updateDate["profit"] = p.Profit
	}

	res := ProductM.Model(&Product{}).Where("id = ?", p.ID).Update(updateDate)

	return res.Error
}
