package database

// import (
// 	"github.com/pump-shop/backend-repository/models"
// 	"github.com/stretchr/testify/suite"
// 	"gopkg.in/DATA-DOG/go-sqlmock.v1"
// 	"gorm.io/gorm"
// )

// type Suite struct {
// 	suite.Suite
// 	DB   *gorm.DB
// 	mock sqlmock.Sqlmock

// 	// repository Repository
// 	product     *models.Product
// }

// func (s *Suite) SetupSuite() {
// 	var (
// 		db  *sql.DB
// 		err error
// 	)

// 	db, s.mock, err = sqlmock.New()
// 	require.NoError(s.T(), err)

// 	s.DB, err = gorm.Open("postgres", db)
// 	require.NoError(s.T(), err)

// 	s.DB.LogMode(true)

// 	s.repository = CreateRepository(s.DB)
// }