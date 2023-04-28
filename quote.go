package main

import (
	"jpagent/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const ERROR_CODE = 100001
const unit = 1

func CreateQuote(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		log.Println("CreateQuote func", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": ERROR_CODE,
		})

		return
	}

	p, err := InsertValidate(data)
	if err != nil {
		log.Println("CreateQuote func", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": ERROR_CODE,
		})

		return
	}

	if err := model.Conn.CreateQuote(p); err != nil {
		log.Println("CreateQuote func", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": ERROR_CODE,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
	})
}

/**

報價 = 日幣(price) * 營收比(profit)%
成本 = (日幣 * 代購費10%(固定) * 匯率(寫死)) ＋ 運費
運費 = 重量 ＊ 單位價(寫死)
利潤 = 報價 - 成本

**/

func GetQuote(c *gin.Context) {
	name := c.Query("name")
	pOffset := c.Query("offset")
	if pOffset == "" {
		pOffset = "0"
	}

	offset, err := strconv.Atoi(pOffset)
	if err != nil {
		log.Println("GetQuote func", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": ERROR_CODE,
		})
		return
	}

	data, err := model.Conn.GetQuote(name, offset)
	if err != nil {
		log.Println("GetQuote func", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": ERROR_CODE,
		})

		return
	}

	sData := make([]interface{}, 0)
	sComput := make([]interface{}, 0)
	date := make(map[string]interface{})
	reslut := make(map[string]map[string]interface{})
	for _, v := range data {
		comput := make(map[string]interface{})
		res := model.Product{
			ID:           v.ID,
			Name:         v.Name,
			Price:        Multiply(v.Price, 0),
			Weight:       v.Weight,
			Ticket:       v.Ticket,
			Freight:      Multiply(int32(v.Weight), unit),
			Fare:         v.Fare,
			People:       v.People,
			Status:       v.Status,
			ExchangeRate: v.ExchangeRate,
			Profit:       v.Profit,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		}
		sData = append(sData, res)
		//報價 = 日幣 * 營收比%
		quote := Multiply(Multiply(v.Price, (1+Div(v.Profit, 100))), v.ExchangeRate)
		//成本 = (日幣 * 代購費10% * 匯率) ＋ 運費
		cost := Multiply(Multiply(v.Price, 0)+v.Freight, v.ExchangeRate)
		//利潤 = 報價 - 成本
		profit := quote - cost

		comput["id"] = v.ID
		comput["name"] = v.Name
		comput["quote"] = quote
		comput["cost"] = cost
		comput["profit"] = profit
		sComput = append(sComput, comput)
		// log.Printf("品名: %s 報價: %d 成本: %d 利潤: %d 運費: %d", v.Name, quote, cost, profit, Multiply(int32(v.Weight), unit))
	}

	date["row"] = sData
	date["comput"] = sComput
	reslut["data"] = date
	c.JSON(http.StatusOK, reslut)
}

func DeleteQuote(c *gin.Context) {
	pid := c.Param("id")
	id, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		log.Println("DeleteQuote func", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": ERROR_CODE,
		})

		return
	}

	if err := model.Conn.DeleteQuote(id); err != nil {
		log.Println("DeleteQuote func", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": ERROR_CODE,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
	})
}

func UpdateQuote(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		log.Println("UpdateQuote func", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": ERROR_CODE,
		})

		return
	}

	p, err := UpdateValidate(data)
	if err != nil {
		log.Println("UpdateQuote func", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": ERROR_CODE,
		})

		return
	}

	if err := model.Conn.UpdateQuote(p); err != nil {
		log.Println("UpdateQuote func", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": ERROR_CODE,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
	})
}
