package data

import (
	"challenge-bravo/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Product struct {
	gorm.Model
	Name  string
	Price uint
}

var (
	db     *gorm.DB
	dbHost string
	err    error
)

func StartDB() {

	fmt.Println("• Connecting to DB..")
	if db, err = OpenTestConnection(); err != nil {
		fmt.Println("• DB: Failed to connect to PostgreSQL. %+v", err)
	}
	db.AutoMigrate(&Product{})

	fmt.Println("• DB: init done")
}

func OpenTestConnection() (db *gorm.DB, err error) {
	var c = config.Config.DB
	dbHost = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", c.Host, c.Port, c.User, c.Database, c.Password)

	db, err = gorm.Open("postgres", dbHost)
	db.LogMode(true)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetDBConn() (db *gorm.DB) {
	// Check if database connection is open, otherwise open a new one
	dbPing := db.DB().Ping()
	if dbPing != nil {
		fmt.Println("• DB: Trying to open database again..")
		if db, err = OpenTestConnection(); err != nil {
			fmt.Println("• DB: Failed to connect to MySQL Server. %+v", err)
		}
		return db
	}
	return db
}

func GetProductByID(productID int64) (bool, *Product) {
	var product Product
	if err := db.Where("id = ?", productID).First(&product).Error; err != nil {
		return false, nil
	}
	return true, &product
}

func CreateProduct(name string, price uint) (bool, *Product, error) {
	var product = db.Create(&Product{Name: name, Price: price})
	if err := db.Create(&product).Error; err != nil {
		return false, nil, err
	}
	return true, &product, nil
}
