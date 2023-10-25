package main

import (
	"database/sql"
	"effective-mobile/internal/controllers"
	"effective-mobile/internal/services"
	"effective-mobile/internal/storage"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres database=emdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	storage := storage.NewPostgresStorage(db)

	ioSvc := services.NewIoSvc()
	personSvc := services.NewPersonSvc(storage)

	ctrl := controllers.NewController(personSvc, personSvc, personSvc, ioSvc)

	router := gin.Default()

	api := router.Group("/api")
	api.POST("/person", ctrl.AddPersonCtrl.AddPerson)
	api.DELETE("/person/:id", ctrl.DeletePersonCtrl.DeletePerson)
	api.PUT("/person", ctrl.UpdatePersonCtrl.UpdatePerson)

	fmt.Println("running...")
	router.Run(":8080")

}
