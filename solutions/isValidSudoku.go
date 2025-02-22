package solutions

func IsValidSudoku(board [][]byte) bool {
	row_set := make(map[byte]bool, 9)
	col_set := make(map[byte]bool, 9)
	box_set := make(map[byte]bool, 9)

	section_set := [][]byte{}
	first_sec := []byte{}
	sec_sec := []byte{}
	third_sec := []byte{}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			row_num := board[i][j]
			if row_num != '.' {
				// for row
				if _, found := row_set[row_num]; found {
					return false
				} else {
					row_set[row_num] = true
				}

				// for box
				switch {
				case j < 3:
					first_sec = append(first_sec, row_num)
				case j < 6:
					sec_sec = append(sec_sec, row_num)
				default:
					third_sec = append(third_sec, row_num)
				}
			}

			// for column
			col_num := board[j][i]
			if col_num != '.' {
				if _, found := col_set[col_num]; found {
					return false
				} else {
					col_set[col_num] = true
				}
			}
		}

		// works on each 3th row
		if (i+1)%3 == 0 {
			section_set = append(section_set, first_sec, sec_sec, third_sec)
			for i := 0; i < len(section_set); i++ {
				for j := 0; j < len(section_set[i]); j++ {
					section_num := section_set[i][j]
					if _, found := box_set[section_num]; found {
						return false
					} else {
						box_set[section_num] = true
					}
				}
				box_set = make(map[byte]bool, 9)

			}
			first_sec = []byte{}
			sec_sec = []byte{}
			third_sec = []byte{}
		}

		row_set = make(map[byte]bool, 9)
		col_set = make(map[byte]bool, 9)
	}

	return true
}
