package pagination

import (
	"errors"
	"posts/internal/constants"
)

func GetLimitAndOffset(page, pageSize *int) (int, int, error) {
	var (
		limit  int
		offset int
		err    error
	)

	if page != nil && *page <= 0 {
		err = errors.New(constants.WrongPageError)
	}

	if pageSize != nil && *pageSize < 0 {
		err = errors.New(constants.WrongPageSizeError)
	}

	if page == nil || pageSize == nil {
		limit = -1
		offset = 0
	} else {
		offset = (*page - 1) * *pageSize
		limit = *pageSize
	}
	return limit, offset, err
}
