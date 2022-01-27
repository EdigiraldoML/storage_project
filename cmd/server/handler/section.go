package handler

import (
	"net/http"
	"strconv"

	"github.com/EdigiraldoML/storage_project/internal/section"
	"github.com/EdigiraldoML/storage_project/pkg/web"
	"github.com/gin-gonic/gin"
)

type requestSectionCreate struct {
	SectionNumber      int `json:"section_number" binding:"required"`
	CurrentTemperature int `json:"current_temperature" binding:"required"`
	MinimumTemperature int `json:"minimum_temperature" binding:"required"`
	CurrentCapacity    int `json:"current_capacity" binding:"required"`
	MinimumCapacity    int `json:"minimum_capacity" binding:"required"`
	MaximumCapacity    int `json:"maximum_capacity" binding:"required"`
	WarehouseID        int `json:"warehouse_id" binding:"required"`
	ProductTypeID      int `json:"product_type_id" binding:"required"`
}

type requestSectionUpdate struct {
	ID                 int `json:"id"`
	SectionNumber      int `json:"section_number"`
	CurrentTemperature int `json:"current_temperature"`
	MinimumTemperature int `json:"minimum_temperature"`
	CurrentCapacity    int `json:"current_capacity"`
	MinimumCapacity    int `json:"minimum_capacity"`
	MaximumCapacity    int `json:"maximum_capacity"`
	WarehouseID        int `json:"warehouse_id"`
	ProductTypeID      int `json:"product_type_id"`
}

type Section struct {
	sectionService section.Service
}

func NewSection(s section.Service) *Section {
	return &Section{
		sectionService: s,
	}
}

// GetAll godoc
// @Summary List all sections in database.
// @Tags Section GetAll
// @Description List all sections that are recorded in database.
// @Description respond with status code 200 on success.
// @Description respond with status code 400 on failure.
// @Accept json
// @Produce json
// @Success 200 {object} web.SuccessResponse
// @Failure 400 {object} web.ErrorResponse
// @Router /api/v1/sections/ [get]
func (s *Section) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := s.sectionService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		web.Success(c, http.StatusOK, users)
	}
}

// Get godoc
// @Summary List section given the id.
// @Tags Section GetById
// @Description List section given the id as a param in url.
// @Description respond with status code 200 on success.
// @Description respond with status code 400 if param in url is not an integer.
// @Description respond with status code 403 if requested section was not found.
// @Accept json
// @Produce json
// @Param id path int true "section id"
// @Success 200 {object} web.SuccessResponse
// @Failure 400 {object} web.ErrorResponse
// @Failure 403 {object} web.ErrorResponse
// @Router /api/v1/sections/{id} [get]
func (s *Section) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		user, err := s.sectionService.Get(c, int(id))
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}

		web.Success(c, http.StatusOK, user)
	}
}

// Create godoc
// @Summary Create a new section.
// @Tags Section Create
// @Description Create a new section given params in body as a json object.
// @Description respond with status code 201 on success.
// @Description respond with status code 422 if the sent object is not as expected.
// @Description respond with status code 409 if the section number of the new object already exists.
// @Description respond with status code 500 if the server fails for other reason.
// @Accept json
// @Produce json
// @Param section body requestSectionCreate true "section to create"
// @Success 201 {object} web.SuccessResponse
// @Failure 422 {object} web.ErrorResponse
// @Failure 409 {object} web.ErrorResponse
// @Failure 500 {object} web.ErrorResponse
// @Router /api/v1/sections/ [post]
func (s *Section) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var request requestSectionCreate

		err := c.ShouldBindJSON(&request)
		if err != nil {
			web.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		createdSection, err := s.sectionService.Create(
			c,
			request.SectionNumber,
			request.CurrentTemperature,
			request.MinimumTemperature,
			request.CurrentCapacity,
			request.MinimumCapacity,
			request.MaximumCapacity,
			request.WarehouseID,
			request.ProductTypeID,
		)
		if err != nil {
			if err.Error() == "el número de la seccion ya existe" {
				web.Error(c, http.StatusConflict, err.Error())
				return
			}

			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}

		web.Success(c, http.StatusCreated, createdSection)
	}
}

// Update godoc
// @Summary Update to an existing section.
// @Tags Section Update
// @Description Update to an existing section given the id as a path param.
// @Description respond with status code 200 on success.
// @Description respond with status code 400 if the id param is not an integer.
// @Description respond with status code 404 if a section with the given id do not exists.
// @Accept json
// @Produce json
// @Param section body requestSectionUpdate true "section to update"
// @Param id path int true "section id"
// @Success 200 {object} web.SuccessResponse
// @Failure 400 {object} web.ErrorResponse
// @Failure 404 {object} web.ErrorResponse
// @Router /api/v1/sections/{id} [patch]
func (s *Section) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		var request requestSectionUpdate

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		err = c.ShouldBindJSON(&request)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		request.ID = int(id)

		updatedSection, err := s.sectionService.Update(
			c,
			request.ID,
			request.SectionNumber,
			request.CurrentTemperature,
			request.MinimumTemperature,
			request.CurrentCapacity,
			request.MinimumCapacity,
			request.MaximumCapacity,
			request.WarehouseID,
			request.ProductTypeID,
		)
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}

		web.Success(c, http.StatusOK, updatedSection)
	}
}

// Delete godoc
// @Summary Delete an existing section.
// @Tags Section Delete
// @Description Delete an existing section given the id as an path param.
// @Description respond with status code 204 on success.
// @Description respond with status code 400 if the id param is not an integer.
// @Description respond with status code 404 if a section with the given id do not exists.
// @Accept json
// @Produce json
// @Param id path int true "section id"
// @Success 204 {object} web.SuccessResponse
// @Failure 400 {object} web.ErrorResponse
// @Failure 404 {object} web.ErrorResponse
// @Router /api/v1/sections/{id} [delete]
func (s *Section) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		err = s.sectionService.Delete(c, int(id))
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}

		succesMsg := "eliminado satisfactoriamente"

		web.Success(c, http.StatusNoContent, succesMsg)
	}
}
