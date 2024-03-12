package resources

import "time"

type Resource struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	JSONSchema  any       `json:"json_schema"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
