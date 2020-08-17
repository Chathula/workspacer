package app

// WorkspacerSchema struct
type WorkspacerSchema struct {
	Version int    `json:"version"`
	Home    string `json:"home,omitempty"`
	Git     Git    `json:"git"`
	Envs    []Envs `json:"envs,omitempty"`
}

// Repositories struct
type Repositories struct {
	URL    string `json:"url"`
	Name   string `json:"name,omitempty"`
	Delete bool   `json:"delete,omitempty"`
}

// Git struct
type Git struct {
	Path         string         `json:"path"`
	Repositories []Repositories `json:"repositories"`
}

// Envs struct
type Envs struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Delete bool   `json:"delete,omitempty"`
}
