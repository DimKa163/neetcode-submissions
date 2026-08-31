func isValidSudoku(board [][]byte) bool {
	rows := 9
	cols := 9
	colMaps := make(map[int]map[byte]struct{})
	rowMaps := make(map[int]map[byte]struct{})
	boxMap := make(map[int]map[byte]struct{})
	for r := range rows {
		row := board[r]
		for c := range cols {
			if row[c] == '.' {
				continue
			}
			if !checkMap(colMaps, c, row[c]) {
				return false
			}
			if !checkMap(rowMaps, r, row[c]) {
				return false
			}
			boxID := (r / 3) * 3 + (c / 3)
			if !checkMap(boxMap, boxID, row[c]) {
				return false
			}
		}
	}
	return true
}

func checkMap(set map[int]map[byte]struct{}, idx int,  value byte) bool {
	m, ok := set[idx]
	if !ok {
		set[idx] = make(map[byte]struct{})
		set[idx][value] = struct{}{}
		return true
	}
	_, ok = m[value]
	if ok {
		return false
	}
	m[value] = struct{}{}
	return true
}
