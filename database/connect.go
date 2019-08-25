package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go.xstore.local/go-hit-counter/config"
	"go.xstore.local/go-hit-counter/models"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	conf := config.Get()
	db, err := gorm.Open("mysql", conf.DSN)

	if err == nil {
		db.DB().SetMaxIdleConns(conf.MaxIdleConn)
		DB = db
		db.AutoMigrate(&models.HitCounter{})
		return db, err
	}
	return nil, err
}

// Get the count of a particular url
func GetCount(url string) models.HitCounter {
	var hitCounter models.HitCounter
	// Get first matched record
	int0 := 0
	if err := DB.Where("url = ?", url).First(&hitCounter).Error; err != nil {
		return models.HitCounter{
			URL:   url,
			Count: &int0,
		}
	}
	// SELECT * FROM users WHERE name = 'jinzhu' limit 1;
	return hitCounter
}

func AddView(url string) {
	hitCounter := GetCount(url)
	// Make sure the url entry exists
	if *(hitCounter.Count) == 0 {
		DB.Create(&models.HitCounter{
			URL:   url,
			Count: hitCounter.Count,
		})
	}

	// Add 1 to the url count
	oldCount := *hitCounter.Count + 1
	hitCounter.Count = &oldCount
	DB.Model(&hitCounter).Update(&hitCounter)
}
