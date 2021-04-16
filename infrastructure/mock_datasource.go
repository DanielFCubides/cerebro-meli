package infrastructure

import (
	"github.com/jinzhu/gorm"
	mocket "github.com/selvatico/go-mocket"
)

type MockConnection struct {
	db *gorm.DB
}

func NewMockConnection() Connection {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true
	db, _ := gorm.Open(mocket.DriverName, "connection_mock")

	return &MockConnection{db: db}
}

func (c *MockConnection) GetDatabase() *gorm.DB {
	return c.db
}
