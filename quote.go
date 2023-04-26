package main

import (
	"jpagent/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

const ERROR_CODE = 100001

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

	if err := model.CreateQuote(p); err != nil {
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

報價 = 日幣 * 營收比%
成本 = (日幣 * 代購費10% * 匯率) ＋ 運費
運費 = 重量 ＊ 單位價
利潤 = 報價 - 成本

**/

func Multiply(a int32) int32 {
	ad := decimal.NewFromFloat(float64(a))
	bd := decimal.NewFromFloat(1.1) //固定10%
	res := ad.Mul(bd).IntPart()
	return int32(res)
}

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

	data, err := model.GetQuote(name, offset)
	if err != nil {
		log.Println("GetQuote func", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": ERROR_CODE,
		})

		return
	}

	sData := make([]interface{}, 0)
	reslut := make(map[string][]interface{})
	// cost := make()
	for _, v := range data {
		res := model.Product{
			ID:           v.ID,
			Name:         v.Name,
			Price:        Multiply(v.Price),
			Weight:       v.Weight,
			Ticket:       v.Ticket,
			Freight:      v.Freight,
			Fare:         v.Fare,
			People:       v.People,
			Status:       v.Status,
			ExchangeRate: v.ExchangeRate,
			Profit:       v.Profit,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		}
		sData = append(sData, res)
	}

	reslut["data"] = sData
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

	if err := model.DeleteQuote(id); err != nil {
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

	if err := model.UpdateQuote(p); err != nil {
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
