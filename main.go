package main

import (
	"assignment3/models"
	"fmt"
	"log"
	"math/rand"
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

func CreateWarning(warningPayload models.Warning) (*models.Warning, error) {
	var warning models.Warning

	payload := models.Warning{
		Water: warningPayload.Water,
		Wind:  warningPayload.Wind,
	}
	switch {
	case warningPayload.Water < 5:
		fmt.Println("findWarning0")

		payload = models.Warning{
			ID:     1,
			Water:  warningPayload.Water,
			Wind:   warningPayload.Wind,
			Status: "aman",
		}
		break
	case warningPayload.Water >= 6 && warningPayload.Water <= 8:
		fmt.Println("findWarning3")

		payload = models.Warning{
			ID:     1,
			Water:  warningPayload.Water,
			Wind:   warningPayload.Wind,
			Status: "siaga",
		}
		break
	case warningPayload.Wind < 6:
		fmt.Println("findWarning4")

		payload = models.Warning{
			ID:     1,
			Water:  warningPayload.Water,
			Wind:   warningPayload.Wind,
			Status: "aman",
		}
		break
	case warningPayload.Wind >= 7 && warningPayload.Wind <= 15:
		fmt.Println("findWarning5")

		payload = models.Warning{
			ID:     1,
			Status: "siaga",
			Water:  warningPayload.Water,
			Wind:   warningPayload.Wind,
		}
		break
	case warningPayload.Wind > 15:
		fmt.Println("findWarning6")
		payload = models.Warning{
			ID:     1,
			Status: "bahaya",
			Water:  warningPayload.Water,
			Wind:   warningPayload.Wind,
		}
		break
	default:
		fmt.Printf("%s.\n", "why")
	}

	resultCreate := db.Create(&payload)
	if resultCreate.Error != nil {
		log.Fatal("resultCreate.Error", resultCreate.Error)
		return &warning, resultCreate.Error
	}

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

// func GetBookings(c *gin.Context) {
// 	var bookings []models.Booking

//		if err := models.DB.Find(&bookings).Error; err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//			return
//		}
//		c.JSON(http.StatusOK, bookings)
//	}
func getStatus(ctx *gin.Context) {
	var warning = []models.Warning{}
	err := db.Find(&warning).Error
	if err != nil {
		fmt.Println("a", err.Error())
	}
	// fmt.Println("get status result", &result)
	// because of get no need to bind to json
	// if err := ctx.ShouldBindJSON(&result); err != nil {
	// 	fmt.Println("err", err)
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest,
	// 		gin.H{
	// 			"status":  "VALIDATEERR-1",
	// 			"message": "Invalid inputs. Please check your inputs"})
	// 	return
	// }

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error find data", "message": err.Error()})
		return
	}

	// if result.Error != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error result", "message": err.Error()})
	// 	return
	// }
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"data": warning}})
	return

}

func AutoReload(warning models.Warning) {
	payload := models.Warning{
		ID:    1,
		Water: warning.Water,
		Wind:  warning.Wind,
	}
	switch {
	case warning.Water < 5:
		fmt.Println("findWarning0")

		payload = models.Warning{
			ID:     1,
			Water:  warning.Water,
			Wind:   warning.Wind,
			Status: "aman",
		}
		break
	case warning.Water >= 6 && warning.Water <= 8:
		fmt.Println("findWarning3")

		payload = models.Warning{
			ID:     1,
			Water:  warning.Water,
			Wind:   warning.Wind,
			Status: "siaga",
		}
		break
	case warning.Wind < 6:
		fmt.Println("findWarning4")

		payload = models.Warning{
			ID:     1,
			Water:  warning.Water,
			Wind:   warning.Wind,
			Status: "aman",
		}
		break
	case warning.Wind >= 7 && warning.Wind <= 15:
		fmt.Println("findWarning5")

		payload = models.Warning{
			ID:     1,
			Status: "siaga",
			Water:  warning.Water,
			Wind:   warning.Wind,
		}
		break
	case warning.Wind > 15:
		fmt.Println("findWarning6")
		payload = models.Warning{
			ID:     1,
			Status: "bahaya",
			Water:  warning.Water,
			Wind:   warning.Wind,
		}
		break
	default:
		fmt.Printf("%s.\n", "why")
	}

	resultUpdate := db.Save(&payload)
	if resultUpdate.Error != nil {
		log.Fatal("resultUpdate.Error", resultUpdate.Error)
	}
}

func AutoReloadTick(warning models.Warning) {
	for range time.Tick(time.Second * 15) {
		fmt.Println("this function will run every 15 seconds")
		AutoReload(warning)

	}
}

func randU32(min, max uint32) uint32 {
	var a = rand.Uint32()
	a %= (max - min)
	a += min
	return a
}

func main() {
	var min uint32 = 1
	var max uint32 = 100
	randWater := randU32(min, max)
	randWind := randU32(min, max)

	warning := models.Warning{Wind: randWind, Water: randWater}

	resultCreateWarning, err := CreateWarning(warning)
	if err != nil {
		log.Fatal("resultCreateWarning", resultCreateWarning)
	}
	go AutoReloadTick(warning)

	// fmt.Println("warning1", warning.ID)
	router := gin.Default()
	router.GET("/status", getStatus)

	router.Run("localhost:8080")

	// c := cron.New(cron.WithSeconds())
	// c.AddFunc("*/5 * * * * *", func() { AutoReload() })
	// c.Start()

}
