package page

// Page godoc
type Page interface {
	IsEmpty() bool
	GetTotalElements() int
	GetTotalItems() int
	GetContent() []interface{}
}

type page struct {
	content       []interface{}
	totalElements int
	totalItems    int
}

// Empty create empty page
func Empty() Page {
	return &page{
		content:       make([]interface{}, 0),
		totalElements: 0,
		totalItems:    0,
	}
}

// From create page from data
func From(content []interface{}, totalItems int) Page {
	return &page{
		content:       content,
		totalElements: len(content),
		totalItems:    totalItems,
	}
}

func (p *page) IsEmpty() bool {
	return p.totalElements == 0
}

func (p *page) GetTotalElements() int {
	return p.totalElements
}

func (p *page) GetTotalItems() int {
	return p.totalItems
}

func (p *page) GetContent() []interface{} {
	return p.content
}
