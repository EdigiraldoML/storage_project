package section

import (
	"context"
	"database/sql"
	"log"

	"github.com/EdigiraldoML/storage_project/internal/db"
	"github.com/EdigiraldoML/storage_project/internal/domain"
)

// Repository encapsulates the storage of a section.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Section, error)
	Get(ctx context.Context, id int) (domain.Section, error)
	Exists(ctx context.Context, cid int) bool
	Save(ctx context.Context, s domain.Section) (int, error)
	Update(ctx context.Context, s domain.Section) error
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Section, error) {
	var sections []domain.Section
	db := db.StorageDB

	getQuery := "SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM sections"
	rows, err := db.Query(getQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var section domain.Section
		if err := rows.Scan(&section.ID, &section.SectionNumber, &section.CurrentTemperature, &section.MinimumTemperature, &section.CurrentCapacity, &section.MinimumCapacity, &section.MaximumCapacity, &section.WarehouseID, &section.ProductTypeID); err != nil {
			log.Fatal(err)
			return nil, err
		}

		sections = append(sections, section)
	}

	return sections, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Section, error) {
	var section domain.Section
	db := db.StorageDB
	getQuery := "SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM sections WHERE id = ?"
	rows, err := db.Query(getQuery, id)
	if err != nil {
		return section, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&section.ID, &section.SectionNumber, &section.CurrentTemperature, &section.MinimumTemperature, &section.CurrentCapacity, &section.MinimumCapacity, &section.MaximumCapacity, &section.WarehouseID, &section.ProductTypeID); err != nil {
			return section, err
		}
	}
	return section, nil
}

func (r *repository) Exists(ctx context.Context, cid int) bool {
	return false
}
func (r *repository) Save(ctx context.Context, section domain.Section) (int, error) {
	db := db.StorageDB
	insertStatement := "INSERT INTO sections(section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id) VALUES( ?, ?, ?, ?, ?, ?, ?, ? )"
	stmt, err := db.Prepare(insertStatement)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var result sql.Result
	result, err = stmt.Exec(section.SectionNumber, section.CurrentTemperature, section.MinimumTemperature, section.CurrentCapacity, section.MinimumCapacity, section.MaximumCapacity, section.WarehouseID, section.ProductTypeID)
	if err != nil {
		return 0, err
	}

	insertedId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(insertedId), nil
}

func (r *repository) Update(ctx context.Context, s domain.Section) error {
	updateQuery := "UPDATE sections SET section_number=?, current_temperature=?, minimum_temperature=?, current_capacity=?, minimum_capacity=?, maximum_capacity=?, warehouse_id=?, product_type_id=? WHERE id=?;"
	stmt, err := r.db.Prepare(updateQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, &s.SectionNumber, &s.CurrentTemperature, &s.MinimumTemperature, &s.CurrentCapacity, &s.MinimumCapacity, &s.MaximumCapacity, &s.WarehouseID, &s.ProductTypeID, &s.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
func (r *repository) Delete(ctx context.Context, id int) error {
	return nil
}
