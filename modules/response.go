package modules

import (
	"fmt"
	"math"

	"gorm.io/gorm"
)

type Response struct {
	Meta       *Meta       `json:"meta"`
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination"`
}

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Pagination struct {
	Limit      int    `json:"limit"`
	Page       int    `json:"page"`
	Sort       string `json:"sort"`
	SortBy     string `json:"sortBy"`
	TotalRows  int64  `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
}

func (r *Response) SendResponse(message string, code int, data interface{}, pagination *Pagination) *Response {
	res := &Response{
		Meta: &Meta{
			Message: message,
			Code:    code,
		},
		Data:       data,
		Pagination: pagination,
	}

	return res
}

func Paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	if pagination.Limit == 0 {
		pagination.Limit = 10
	}
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.SortBy == "" {
		p.SortBy = "Id"
	}

	if p.Sort == "" {
		p.Sort = fmt.Sprintf("%s asc", p.SortBy)
		return p.Sort
	}
	p.Sort = fmt.Sprintf("%s %s", p.SortBy, p.Sort)
	return p.Sort
}
