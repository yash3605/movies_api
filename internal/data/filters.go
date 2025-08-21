package data

import (
	"github.com/greenlight/internal/validator"
	"strings"
)

type Filters struct {
	Page			int
	PageSize		int
	Sort			string
	SortSafeList	[]string
}

func ValidateFilters(v *validator.Validator, f Filters){
	v.Check(f.Page > 0, "page", "Must be Greater than zero")
	v.Check(f.Page <= 10_000_000, "page", "Must be a Maximum of 10 million")
	v.Check(f.PageSize > 0, "page_size", "Must be greater than zero")
	v.Check(f.PageSize <= 100, "page_size", "Must be a maximum of 100")

	v.Check(validator.PermittedValue(f.Sort, f.SortSafeList...), "sort", "Invalid Sort Value")
}

func (f Filters) sortColumn() string {
	for _, safeValue := range f.SortSafeList {
		if f.Sort == safeValue {
			return strings.TrimPrefix(f.Sort, "-")
		}
	}

	panic("unsafe sort parameter: " + f.Sort)
}

func (f Filters) sortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}

	return "ASC"
}
