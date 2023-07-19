package algorithm

func Exists(board [][]byte, word string) bool {
	// row
	row := len(board)
	if row == 0 {
		return false
	}
	// column
	column := len(board[0])
	if column == 0 {
		return false
	}
	wordLen := len(word)
	if wordLen == 0 {
		return false
	}

	var dfs func(r, c, index int) bool
	dfs = func(r, c, index int) bool {
		// index equal to wordLen matches finished
		if index == wordLen {
			return true
		}
		// not equal
		if r < 0 || r >= row || c < 0 || c >= column || board[r][c] != word[index] {
			return false
		}
		//set board[r][c] to 0 set it has passed
		temp := board[r][c]
		board[r][c] = '0'
		if dfs(r+1, c, index+1) ||
			dfs(r-1, c, index+1) ||
			dfs(r, c+1, index+1) ||
			dfs(r, c-1, index+1) {
			return true
		}
		board[r][c] = temp
		return false
	}

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
