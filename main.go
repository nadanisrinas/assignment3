package main

import (
	"assignment3/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "assignment3"
	db       *gorm.DB
	err      error
)

func AutoReload() (*models.Warning, error) {
	var warning models.Warning
	payload := models.Warning{
		ID: 2,
	}
	switch {
	case warning.Water < 5:
		payload = models.Warning{
			ID:     payload.ID,
			Status: "aman",
		}
		break
	case warning.Water >= 6 && warning.Water <= 8:
		payload = models.Warning{
			ID:     payload.ID,
			Status: "siaga",
		}
		break
	case warning.Wind < 6:
		payload = models.Warning{
			ID:     payload.ID,
			Status: "aman",
		}
		break
	case warning.Wind >= 7 && warning.Wind <= 15:
		payload = models.Warning{
			ID:     payload.ID,
			Status: "siaga",
		}
		break
	case warning.Wind > 15:
		payload = models.Warning{
			ID:     payload.ID,
			Status: "bahaya",
		}
		break
	default:
		fmt.Printf("%s.\n", "why")
	}

	result := db.Save(&payload)
	if result.Error != nil {
		return &warning, result.Error
	}
	// fmt.Println(&warning, nil)
	return &warning, nil
}
func StartDB() {
	// Sprintf formats according to a format specifier and returns the resulting string.
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	//open koneksi dengan driver
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	errMigrate := db.Debug().AutoMigrate(models.Warning{})
	if errMigrate != nil {
		log.Fatal("error migrate db", errMigrate)
	}
	fmt.Println("db successfully migrated...")
}

func init() {
	StartDB()
}

func getStatus(ctx *gin.Context) {
	var warning models.Warning
	if err := ctx.BindJSON(&warning); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
	}
	result := db.Find(&warning)
	if result.Error != nil {
		log.Fatal("Error get status data", result.Error)
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"data": warning}})

}
func main() {
	for range time.Tick(time.Second * 15) {
		fmt.Println("this function will run every 15 seconds")
		AutoReload()
	}
	warning := models.Warning{Wind: 5, Water: 5}
	result := db.Create(&warning)
	if result.Error != nil {
		log.Fatal("error", result.Error)
	}
	router := gin.Default()
	router.GET("/status", getStatus)

	router.Run("localhost:8080")

	// c := cron.New(cron.WithSeconds())
	// c.AddFunc("*/5 * * * * *", func() { AutoReload() })
	// c.Start()

}
