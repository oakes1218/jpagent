package main

import "github.com/shopspring/decimal"

func Multiply(a int32, b float64) int32 {
	if b == 0 {
		b = 1.1 //固定10%
	}

	ad := decimal.NewFromFloat(float64(a))
	bd := decimal.NewFromFloat(b)
	res := ad.Mul(bd).IntPart()
	return int32(res)
}

func Div(a, b int32) float64 {
	ad := decimal.NewFromFloat(float64(a))
	bd := decimal.NewFromFloat(float64(b))
	res, _ := ad.Div(bd).Float64()

	return res
}
