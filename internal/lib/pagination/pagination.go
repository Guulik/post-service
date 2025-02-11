package pagination

import (
	"errors"
	"posts/internal/constants"
)

func GetLimitAndOffset(page, pageSize *int) (int, int, error) {
	var (
		limit  = -1
		offset int
		err    error
	)

	if page == nil || pageSize == nil {
		limit = -1
		offset = 0
	} else {
		offset = (*page - 1) * *pageSize
		limit = *pageSize
	}

	if page != nil && *page <= 0 {
		err = errors.New(constants.WrongPageError)
		limit = -1
		offset = 0
	}

	if pageSize != nil && *pageSize < 0 {
		err = errors.New(constants.WrongPageSizeError)
		limit = -1
		offset = 0
	}

	return limit, offset, err
}
