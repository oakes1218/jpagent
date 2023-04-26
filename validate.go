package main

import (
	"encoding/json"
	"errors"
	"jpagent/model"
	"log"
)

const ExchangeRate = 0.23

func InsertValidate(data []byte) (*model.Product, error) {
	p := new(model.Product)
	//先寫固定
	p.ExchangeRate = ExchangeRate
	if err := json.Unmarshal(data, &p); err != nil {
		log.Println("InsertValidate func", err)
		return p, err
	}

	if p.Status == 1 {
		if p.Remark == "" {
			return p, errors.New("備註不能為空")
		}
	}

	request := struct {
		Name    string  `validate:"required"`
		Price   int32   `validate:"numeric,required"`
		Weight  float64 `validate:"numeric,required"`
		Ticket  int32   `validate:"numeric"`
		Freight int32   `validate:"numeric"`
		Fare    int32   `validate:"numeric"`
		People  int32   `validate:"numeric"`
		Status  int32   `validate:"numeric,min=0,max=1"`
		Profit  int32   `validate:"numeric,required"`
		Note    string  `validate:"numeric,min=0,max=255"`
		Remark  string  `validate:"numeric,min=0,max=255"`
	}{
		Name:    p.Name,
		Price:   p.Price,
		Weight:  p.Weight,
		Ticket:  p.Ticket,
		Freight: p.Freight,
		Fare:    p.Fare,
		People:  p.People,
		Status:  p.Status,
		Profit:  p.Profit,
		Note:    p.Note,
		Remark:  p.Remark,
	}

	if err := Validate.Struct(request); err != nil {
		return p, err
	}

	return p, nil
}

func UpdateValidate(data []byte) (*model.Product, error) {
	p := new(model.Product)
	if err := json.Unmarshal(data, &p); err != nil {
		log.Println("InsertValidate func", err)
		return p, err
	}

	if p.Status == 1 {
		if p.Remark == "" {
			return p, errors.New("說明欄不能為空")
		}
	}

	request := struct {
		Price        int32   `validate:"numeric"`
		Weight       float64 `validate:"numeric"`
		Ticket       int32   `validate:"numeric"`
		Freight      int32   `validate:"numeric"`
		Fare         int32   `validate:"numeric"`
		People       int32   `validate:"numeric"`
		Status       int32   `validate:"numeric,min=0,max=1"`
		ExchangeRate float64 `validate:"numeric"`
		Profit       int32   `validate:"numeric"`
		Note         string  `validate:"numeric,min=0,max=255"`
		Remark       string  `validate:"numeric,min=0,max=255"`
	}{
		Price:        p.Price,
		Weight:       p.Weight,
		Ticket:       p.Ticket,
		Freight:      p.Freight,
		Fare:         p.Fare,
		People:       p.People,
		Status:       p.Status,
		ExchangeRate: p.ExchangeRate,
		Profit:       p.Profit,
		Note:         p.Note,
		Remark:       p.Remark,
	}

	if err := Validate.Struct(request); err != nil {
		return p, err
	}

	return p, nil
}
