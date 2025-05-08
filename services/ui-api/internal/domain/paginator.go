package domain

type Paginator struct {
	Page     int
	PageSize int
}

func (p *Paginator) GetPage() int {
	if p.Page > 0 {
		return p.Page
	}
	return 1
}

func (p *Paginator) GetLimit() int {
	if p.PageSize > 0 {
		return p.PageSize
	}
	return 10
}

func (p *Paginator) GetOffset() int64 {
	if p.Page > 1 {
		return int64(p.PageSize * (p.Page - 1))
	}
	return 0
}
