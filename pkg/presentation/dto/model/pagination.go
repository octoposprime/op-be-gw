package presentation

func (p *PaginationInput) Validate() *PaginationInput {
	if p == nil {
		p = &PaginationInput{}
	}
	if p.Limit == nil {
		p.Limit = new(int32)
		*p.Limit = 10
	}
	if p.Limit != nil && *p.Limit < 0 {
		*p.Limit = 10
	}
	if p.Limit != nil && *p.Limit > 100 {
		*p.Limit = 100
	}
	if p.Offset == nil {
		p.Offset = new(int32)
		*p.Offset = 0
	}
	if p.Offset != nil && *p.Offset < 0 {
		*p.Offset = 0
	}
	return p
}
