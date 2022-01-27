package section

import (
	"errors"

	"github.com/EdigiraldoML/storage_project/internal/domain"
	"github.com/gin-gonic/gin"
)

// Errors
var (
	ErrNotFound = errors.New("section not found")
)

type Service interface {
	GetAll(c *gin.Context) ([]domain.Section, error)
	Get(c *gin.Context, id int) (domain.Section, error)
	Create(c *gin.Context, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id int) (createdSection domain.Section, err error)
	Update(c *gin.Context, id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id int) (uptadedSection domain.Section, err error)
	Delete(c *gin.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// Metodo para obtener todas las sections en base de datos.
func (s *service) GetAll(c *gin.Context) ([]domain.Section, error) {
	return s.repository.GetAll(c)
}

// Metodo para obtener una section por su id.
func (s *service) Get(c *gin.Context, id int) (domain.Section, error) {

	return s.repository.Get(c, id)
}

// Metodo para crear una nueva section.
func (s *service) Create(c *gin.Context, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id int) (createdSection domain.Section, err error) {
	var section domain.Section

	section.SectionNumber = section_number
	section.CurrentTemperature = current_temperature
	section.MinimumTemperature = minimum_temperature
	section.CurrentCapacity = current_capacity
	section.MinimumCapacity = minimum_capacity
	section.MaximumCapacity = maximum_capacity
	section.WarehouseID = warehouse_id
	section.ProductTypeID = product_type_id

	exists := s.repository.Exists(c, section_number)
	if exists {
		err = errors.New("el número de la seccion ya existe")
		return createdSection, err
	}

	id, err := s.repository.Save(c, section)
	if err != nil {
		return createdSection, err
	}

	return s.repository.Get(c, id)
}

// Metodo para actualizar una section.
func (s *service) Update(c *gin.Context, id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id int) (updatedSection domain.Section, err error) {
	var section domain.Section

	section.ID = id
	section.SectionNumber = section_number
	section.CurrentTemperature = current_temperature
	section.MinimumTemperature = minimum_temperature
	section.CurrentCapacity = current_capacity
	section.MinimumCapacity = minimum_capacity
	section.MaximumCapacity = maximum_capacity
	section.WarehouseID = warehouse_id
	section.ProductTypeID = product_type_id

	currentSection, err := s.repository.Get(c, int(id))
	if err != nil {
		return updatedSection, err
	}

	sectionFieldsToSet := selectFields(section, currentSection)

	err = s.repository.Update(c, sectionFieldsToSet)
	if err != nil {
		return updatedSection, err
	}

	updatedSection = sectionFieldsToSet
	return updatedSection, nil
}

// Devuelve la section actualidada para ser guardada en base de datos.
func selectFields(section domain.Section, currentSection domain.Section) (sectionFieldsToSet domain.Section) {
	sectionFieldsToSet = currentSection
	if section.SectionNumber != 0 {
		sectionFieldsToSet.SectionNumber = section.SectionNumber
	}
	if section.CurrentTemperature != 0 {
		sectionFieldsToSet.CurrentTemperature = section.CurrentTemperature
	}
	if section.MinimumTemperature != 0 {
		sectionFieldsToSet.MinimumTemperature = section.MinimumTemperature
	}
	if section.CurrentCapacity != 0 {
		sectionFieldsToSet.CurrentCapacity = section.CurrentCapacity
	}
	if section.MinimumCapacity != 0 {
		sectionFieldsToSet.MinimumCapacity = section.MinimumCapacity
	}
	if section.MaximumCapacity != 0 {
		sectionFieldsToSet.MaximumCapacity = section.MaximumCapacity
	}
	if section.WarehouseID != 0 {
		sectionFieldsToSet.WarehouseID = section.WarehouseID
	}
	if section.ProductTypeID != 0 {
		sectionFieldsToSet.ProductTypeID = section.ProductTypeID
	}

	return sectionFieldsToSet
}

// Metodo para eliminar una section.
func (s *service) Delete(c *gin.Context, id int) error {
	return s.repository.Delete(c, id)
}
