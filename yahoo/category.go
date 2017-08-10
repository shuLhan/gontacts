package yahoo

//
// Category define a contact category.
//
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Meta2
}
