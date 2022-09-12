package util

import (
	"net/url"
	"strconv"
	"strings"
)

// QueryParams: struct hold every query params request.
type QueryParams struct {
	From      int
	Page      int
	Perpage   int
	OrderBy   string
	Ascending bool
	Author    string
	Query     string
}

// GetParamsValue: function to get every request query params value.
func GetParamsValue(params url.Values) *QueryParams {
	queryParams := &QueryParams{
		Page:    1,
		Perpage: 10,
		OrderBy: "-created_at",
	}

	page := params.Get("page")
	perpage := params.Get("perpage")
	orderby := params.Get("orderby")
	author := params.Get("author")
	query := params.Get("query")

	if page != "" {
		queryParams.Page, _ = strconv.Atoi(page)
	}

	if perpage != "" {
		queryParams.Perpage, _ = strconv.Atoi(perpage)
	}

	if orderby != "" {
		queryParams.OrderBy = orderby
	}

	if author != "" {
		queryParams.Author = author
	}

	if query != "" {
		queryParams.Query = query
	}

	queryParams.getFromValue()
	queryParams.getOrderValue()
	return queryParams
}

// getFromValue: QueryParams method to get starter row for current page.
func (qp *QueryParams) getFromValue() {
	qp.From = (qp.Perpage * (qp.Page - 1))
}

// getOrderValue: QueryParams method to get the given field is sorted wether ascending or descending.
func (qp *QueryParams) getOrderValue() {
	descending := strings.HasPrefix(qp.OrderBy, "-")

	if descending {
		qp.Ascending = false
		qp.OrderBy = qp.OrderBy[1:]
	} else {
		qp.Ascending = true
	}
}
