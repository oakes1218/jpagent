package model

type MockInit struct{}

func MockMysqlConn() {
	Conn = &MockInit{}
}

func (m *MockInit) CreateQuote(p *Product) error {
	return nil
}
func (m *MockInit) GetQuote(name string, offset int) ([]Product, error) {
	return []Product{}, nil
}

func (m *MockInit) DeleteQuote(id int64) error {
	return nil
}

func (m *MockInit) UpdateQuote(p *Product) error {
	return nil
}

func (m *MockInit) Close() error {
	return nil
}
