package operationtype

import (
	"github.com/jmoiron/sqlx"
	"github.com/leonardogazio/golang-account-crud/db"
	"github.com/leonardogazio/golang-account-crud/models"
)

type repository struct {
	db *sqlx.DB
}

// New ...
func New(db *sqlx.DB) (*repository, error) {
	return &repository{db}, nil
}

// Get ...
func (r *repository) Get(id int) (m []*models.OperationType, err error) {
	err = r.db.Select(&m, "SELECT * FROM operation_types")
	return
}

// GetByID ...
func (r *repository) GetByID(id int) (m models.OperationType, err error) {
	err = r.db.Get(&m, "SELECT * FROM operation_types WHERE id = $1", id)
	return
}

// Insert ...
func (r *repository) Insert(m *models.OperationType) error {
	return db.Persist(r.db, "INSERT INTO operation_types (description) VALUES ($1);", m.Description)
}

// Update ...
func (r *repository) Update(m *models.OperationType) error {
	return db.Persist(r.db, "UPDATE operation_types SET description = $1 WHERE id = $2;", m.Description, m.ID)
}

// Delete ...
func (r *repository) Delete(m *models.OperationType) error {
	return db.Persist(r.db, "DELETE from operation_types WHERE id = $1;", m.ID)
}
