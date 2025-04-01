package mock

import (
	"gin/models"

	"github.com/stretchr/testify/mock"
)

// DatabaseInterface ใช้เป็น Interface สำหรับ Mock
type DatabaseInterface interface {
	CreateStaff(staff *models.Staff) error
}

// MockDatabase จำลอง Database แทน Gorm
type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) CreateStaff(staff *models.Staff) error {
	args := m.Called(staff)
	return args.Error(0)
}
