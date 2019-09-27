package matrix_ops

func nullify(m matrix) bool {
	if len(m) == 0 || len(m[0]) == 0 {
		return false
	}

	// we will use first row and first column for storage, so first
	// flag whether or not first row or first column has any zeroes
	nullifyFirstColumn := false
	for i := range m {
		if m[i][0] == 0 {
			nullifyFirstColumn = true
		}
	}
	nullifyFirstRow := false
	for j := range m[0] {
		if m[0][j] == 0 {
			nullifyFirstRow = true
		}
	}

	// // flag rows and columns to nullify
	for i := 1; i < len(m); i++ {
		for j := 1; j < len(m[0]); j++ {
			if m[i][j] == 0 {
				m[i][0] = 0
				m[0][j] = 0
			}
		}
	}

	// nullify flagged rows
	for i := 1; i < len(m); i++ {
		if m[i][0] == 0 {
			for j := 1; j < len(m[0]); j++ {
				m[i][j] = 0
			}
		}
	}

	// // nullify flagged columns
	for j := 1; j < len(m[0]); j++ {
		if m[0][j] == 0 {
			for i := 1; i < len(m); i++ {
				m[i][j] = 0
			}
		}
	}

	if nullifyFirstColumn {
		for i := range m {
			m[i][0] = 0
		}
	}

	if nullifyFirstRow {
		for j := range m[0] {
			m[0][j] = 0
		}
	}

	return true
}
