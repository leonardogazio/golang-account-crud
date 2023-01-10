package models

import "time"

// OperationType ...
type OperationType struct {
	ID          int64      `db:"id" json:"id"`
	Description string     `db:"description" json:"description" binding:"required"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

// SwaggerCreateTemplate ...
type SwaggerCreateTemplate struct {
	Description string `json:"description"`
}

// SwaggerUpdateTemplate ...
type SwaggerUpdateTemplate struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
}
