package section

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/EdigiraldoML/storage_project/internal/domain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	txdb.Register("txdb", "mysql", "meli_sprint_user:Meli_Sprint#123@/storage")
}

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("txdb", uuid.New().String())

	if err == nil {
		return db, db.Ping()
	}

	defer db.Close()

	return db, err
}

func TestGetOk(t *testing.T) {
	id := 1

	db, err := InitDb()
	assert.NoError(t, err)

	repo := NewRepository(db)
	ctx := context.Background()

	expectedSection := domain.Section{
		ID:                 1,
		SectionNumber:      1,
		CurrentTemperature: 24,
		MinimumTemperature: 20,
		CurrentCapacity:    50,
		MinimumCapacity:    30,
		MaximumCapacity:    60,
		WarehouseID:        1,
		ProductTypeID:      1}

	section, err := repo.Get(ctx, id)

	assert.NoError(t, err)
	assert.Equal(t, expectedSection, section)
}

func TestSaveOk(t *testing.T) {
	db, err := InitDb()
	assert.NoError(t, err)

	repo := NewRepository(db)

	idNewUser := 12
	ctx := context.Background()
	newSection := domain.Section{
		ID:                 idNewUser,
		SectionNumber:      idNewUser,
		CurrentTemperature: 24,
		MinimumTemperature: 20,
		CurrentCapacity:    50,
		MinimumCapacity:    30,
		MaximumCapacity:    60,
		WarehouseID:        idNewUser,
		ProductTypeID:      idNewUser,
	}

	idCreated, err := repo.Save(ctx, newSection)
	assert.NoError(t, err)
	assert.Equal(t, idNewUser, idCreated)
}

func TestGetAllOk(t *testing.T) {
	db, err := InitDb()
	assert.NoError(t, err)

	repo := NewRepository(db)
	ctx := context.Background()

	expectedfirstSection := domain.Section{
		ID:                 1,
		SectionNumber:      1,
		CurrentTemperature: 24,
		MinimumTemperature: 20,
		CurrentCapacity:    50,
		MinimumCapacity:    30,
		MaximumCapacity:    60,
		WarehouseID:        1,
		ProductTypeID:      1}

	sections, err := repo.GetAll(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedfirstSection, sections[0])
	assert.Len(t, sections, 12)
}

func TestUpdateOk(t *testing.T) {
	db, err := InitDb()
	assert.NoError(t, err)

	repo := NewRepository(db)
	ctx := context.TODO()

	expectedUpdateThirdSection := domain.Section{
		ID:                 3,
		SectionNumber:      3,
		CurrentTemperature: 25,
		MinimumTemperature: 16,
		CurrentCapacity:    54,
		MinimumCapacity:    35,
		MaximumCapacity:    66,
		WarehouseID:        45,
		ProductTypeID:      35}

	err = repo.Update(ctx, expectedUpdateThirdSection)
	assert.NoError(t, err)

	updatedSection, err := repo.Get(ctx, 3)
	assert.NoError(t, err)
	assert.Equal(t, expectedUpdateThirdSection, updatedSection)
}
