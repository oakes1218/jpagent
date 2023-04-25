package main

import (
	"encoding/json"
	"jpagent/model"
	"log"
)

func InsertValidate(data []byte) (*model.Product, error) {
	p := new(model.Product)
	if err := json.Unmarshal(data, &p); err != nil {
		log.Println("InsertValidate func", err)
		return p, err
	}

	request := struct {
		Name         string  `validate:"required"`
		Price        int32   `validate:"numeric,required"`
		Weight       float64 `validate:"numeric,required"`
		Ticket       int32   `validate:"numeric"`
		Freight      int32   `validate:"numeric"`
		Fare         int32   `validate:"numeric"`
		status       int32   `validate:"numeric"`
		ExchangeRate float64 `validate:"numeric,required"`
		Profit       int32   `validate:"numeric,required"`
	}{
		Name:         p.Name,
		Price:        p.Price,
		Weight:       p.Weight,
		Ticket:       p.Ticket,
		Freight:      p.Freight,
		Fare:         p.Fare,
		status:       p.Status,
		ExchangeRate: p.ExchangeRate,
		Profit:       p.Profit,
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

	request := struct {
		Price        int32   `validate:"numeric"`
		Weight       float64 `validate:"numeric"`
		Ticket       int32   `validate:"numeric"`
		Freight      int32   `validate:"numeric"`
		Fare         int32   `validate:"numeric"`
		status       int32   `validate:"numeric"`
		ExchangeRate float64 `validate:"numeric"`
		Profit       int32   `validate:"numeric"`
	}{
		Price:        p.Price,
		Weight:       p.Weight,
		Ticket:       p.Ticket,
		Freight:      p.Freight,
		Fare:         p.Fare,
		status:       p.Status,
		ExchangeRate: p.ExchangeRate,
		Profit:       p.Profit,
	}

	if err := Validate.Struct(request); err != nil {
		return p, err
	}

	return p, nil
}
