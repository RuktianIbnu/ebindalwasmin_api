package response

// Body ...
type Body struct {
	OK         bool        `json:"ok"`
	Code       int         `json:"code,omitempty"`
	Msg        string      `json:"msg,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Pagination *pagination `json:"pagination,omitempty"`
}

type pagination struct {
	CurrentPage interface{} `json:"current_page,omitempty"`
	// NextPage     uint        `json:"next_page,omitempty"`
	// PreviousPage uint        `json:"previous_page,omitempty"`
	// TotalPages   uint        `json:"total_pages,omitempty"`
	PerPage      interface{} `json:"per_page,omitempty"`
	TotalEntries interface{} `json:"total_entries,omitempty"`
}

// Format ...
func Format(code int, err error, data ...interface{}) (b *Body) {
	var (
		ok  bool = true
		msg string
		d   interface{}

		pg = pagination{}
	)

	if err != nil {
		msg = err.Error()
		ok = false
	}

	if len(data) >= 1 {
		d = data[0]
	}

	if len(data) > 1 {
		pg.TotalEntries = data[1]
		pg.CurrentPage = data[2]
		pg.PerPage = data[3]
	}

	b = &Body{
		OK:         ok,
		Code:       code,
		Data:       d,
		Msg:        msg,
		Pagination: &pg,
	}

	return
}
