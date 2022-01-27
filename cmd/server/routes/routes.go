package routes

import (
	"database/sql"

	"github.com/EdigiraldoML/storage_project/cmd/server/handler"
	"github.com/EdigiraldoML/storage_project/internal/section"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()

	r.buildSectionRoutes()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}

func (r *router) buildSectionRoutes() {
	repo := section.NewRepository(r.db)
	service := section.NewService(repo)
	handler := handler.NewSection(service)

	r.rg.GET("/sections", handler.GetAll())
	r.rg.GET("/sections/:id", handler.Get())
	r.rg.POST("/sections", handler.Create())
	r.rg.PATCH("/sections/:id", handler.Update())
	r.rg.DELETE("/sections/:id", handler.Delete())
}
