package model

import "time"

type (
	// A TODO expresses ...
	TODO struct {
		ID          int
		Subject     string
		Description string
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	// A CreateTODORequest expresses ...
	CreateTODORequest struct {
		Subject     string `json:"subject"`
		Description string `json:"description"`
	}

	// A CreateTODOResponse expresses ...
	CreateTODOResponse struct {
		TODO TODO `json:"todo"`
	}

	// A ReadTODORequest expresses ...
	ReadTODORequest struct {
		ID int `json:"id"`
	}

	// A ReadTODOResponse expresses ...
	ReadTODOResponse struct {
		TODO TODO `json:"todo"`
	}

	// A UpdateTODORequest expresses ...
	UpdateTODORequest struct {
		ID          int    `json:"id"`
		Subject     string `json:"subject"`
		Description string `json:"description"`
	}

	// A UpdateTODOResponse expresses ...
	UpdateTODOResponse struct {
		ID int `json:"id"`
	}

	// A DeleteTODORequest expresses ...
	DeleteTODORequest struct {
		ID int `json:"id"`
	}

	// A DeleteTODOResponse expresses ...
	DeleteTODOResponse struct {
		ID int `json:"id"`
	}
)
