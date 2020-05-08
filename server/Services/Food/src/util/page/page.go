package page

// Page godoc
type Page interface {
	IsEmpty() bool
	GetTotalElements() int
	GetTotalPages() int
	GetContent() []interface{}
}

type page struct {
	content       []interface{}
	totalElements int
	totalPages    int
}

// Empty create empty page
func Empty() Page {
	return &page{
		content:       make([]interface{}, 0),
		totalElements: 0,
		totalPages:    0,
	}
}

// From create page from data
func From(content []interface{}, totalPages int) Page {
	return &page{
		content:       content,
		totalElements: len(content),
		totalPages:    totalPages,
	}
}

func (p *page) IsEmpty() bool {
	return p.totalElements == 0
}

func (p *page) GetTotalElements() int {
	return p.totalElements
}

func (p *page) GetTotalPages() int {
	return p.totalPages
}

func (p *page) GetContent() []interface{} {
	return p.content
}
