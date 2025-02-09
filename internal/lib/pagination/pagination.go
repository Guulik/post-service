package pagination

func GetOffsetAndLimit(page, pageSize *int) (int, int) {
	var (
		limit  int
		offset int
	)

	if page != nil && *page <= 0 {
		page = nil
	}

	if pageSize != nil && *pageSize < 0 {
		pageSize = nil
	}

	if page == nil || pageSize == nil {
		limit = -1
		offset = 0
	} else {
		offset = (*page - 1) * *pageSize
		limit = *pageSize
	}
	return limit, offset
}
