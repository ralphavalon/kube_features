package data

import (
	"fmt"
	"kube_features/api/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
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

var firstTime bool

func StartDB() {

	fmt.Println("• Connecting to DB..")
	if db, err = OpenTestConnection(); err != nil {
		fmt.Println("• DB: Failed to connect to PostgreSQL. %+v", err)
	}
	db.AutoMigrate(&Product{})

	fmt.Println("• DB: init done")
}

func OpenTestConnection() (db *gorm.DB, err error) {
	var c = config.Config
	dbHost = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", c.DBHost, c.DBPort, c.DBUser, c.DBName, c.DBPassword)

	db, err = gorm.Open("postgres", dbHost)
	// db, err = gorm.Open("sqlite3", "/tmp/gorm.db")
	db.LogMode(true)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Database is accessible.")
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
	if !firstTime {
		StartDB()
	}
	var product = Product{Name: name, Price: price}
	if err := db.Create(&product).Error; err != nil {
		fmt.Println(err)
		return false, nil, err
	}
	firstTime = true
	return true, &product, nil
}
