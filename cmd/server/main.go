package main

import (
	"github.com/EdigiraldoML/storage_project/cmd/server/routes"
	"github.com/EdigiraldoML/storage_project/internal/db"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()

	router := routes.NewRouter(r, db.StorageDB)
	router.MapRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}
}
