package requests

type CategoryRequest struct {
	User       string        `json:"user"`
	Patch      bool          `json:"patch,omitempty"`
	Categories []CategoryDto `json:"categories"`
}

type CategoryDto struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Remove bool `json:"remove,omitempty"`
}
