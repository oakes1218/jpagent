package model

var Conn MysqlConnInterface

type MysqlConnInterface interface {
	CreateQuote(p *Product) error
	GetQuote(name string, offset int) ([]Product, error)
	DeleteQuote(id int64) error
	UpdateQuote(p *Product) error
	Close() error
}
