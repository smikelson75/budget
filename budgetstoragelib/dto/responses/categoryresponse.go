package responses

type CategoryResponse struct {
	Categories []CategoryInternal `json:"categories"`
}

type CategoryInternal struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
