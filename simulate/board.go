package simulate

// Board definition of a matrix struct
type Board struct {
	Matrix [][]*Human
	Size   int
}

// AllocateBoard is used to allocate the board
func (board *Board) AllocateBoard(size int) {
	board.Matrix = make([][]*Human, size, size)
	board.Size = size
}

// GetMatrix returns the matrix of pointers to the allocated humans
func (board *Board) GetMatrix() [][]*Human {
	return board.Matrix
}

// SetMatrix sets the board's matrix to the matrix given as an argument
func (board *Board) SetMatrix(matrix [][]*Human) {
	board.Matrix = matrix
}

// SetMatrixRow sets a row in the board's matrix to the row given as an argument
func (board *Board) SetMatrixRow(humans []*Human, row int) {
	board.Matrix[row] = humans
}

// GetHuman returns a pointer to a human in the board's matrix on the position given as argument
func (board *Board) GetHuman(Pos Pos) *Human {
	return board.Matrix[Pos.X][Pos.Y]
}

// SetHuman inserts a human on its correct position in the board's matrix
func (board *Board) SetHuman(human *Human) {
	board.Matrix[human.Pos.X][human.Pos.Y] = human
}

// GenRow creates humans and inserts pointers to them in a slice
func (board *Board) GenRow(row int, signal chan []*Human, rng int) {
	temp := make([]*Human, board.Size)
	for col := 0; col < board.Size; col++ {
		temp[col] = CreateHuman(Pos{row, col}, rng)
	}
	signal <- temp
}

//sequential version of GenRow to be used to test FillBoardSeq
func GenRowSeq(row int, rng int, size int) []*Human {
	temp := make([]*Human, size)
	for col := 0; col < size; col++ {
		temp[col] = CreateHuman(Pos{row, col}, rng)
	}
	return temp
}

// ReceiveSignal TODO: desc
func ReceiveSignal(signal chan []*Human, board *Board, row int) {
	var temp = <-signal
	board.SetMatrixRow(temp, row)
}

// FillBoard TODO: desc and maybe look over if we should directly manipulate humans that are allocated in AllocateBoard()
func (board *Board) FillBoard(rng int) {
	var channels []chan []*Human
	for row := 0; row < board.Size; row++ {
		channels = append(channels, make(chan []*Human))
		go board.GenRow(row, channels[row], rng)
		ReceiveSignal(channels[row], board, row)
	}
}

//sequential version of FillBoard function
//run both versions and see if concurrent is faster
func (board *Board) FillBoardSeq(rng int) {
	for row := 0; row < board.Size; row++ {
		result := GenRowSeq(row, rng, board.Size)
		board.SetMatrixRow(result, row)
	}
}

// CalcAllNeighborsAUX TODO: desc
func (board *Board) CalcAllNeighborsAUX(row int) {
	for col := 0; col < board.Size; col++ {
		board.GetHuman(Pos{row, col}).CalcNeighbors(board)
	}
}

// CalcAllNeighbors TODO: desc
func (board *Board) CalcAllNeighbors() {
	for row := 0; row < board.Size; row++ {
		go board.CalcAllNeighborsAUX(row)
	}
}

// GetHumanStates returns a matrix with the health states of the humans
func (board *Board) GetHumanStates() [][]State {
	ans := make([][]State, board.Size, board.Size)
	for row := 0; row < board.Size; row++ {
		ans[row] = make([]State, board.Size, board.Size)
		for col := 0; col < board.Size; col++ {
			ans[row][col] = board.GetHuman(Pos{row, col}).GetState()
		}
	}
	return ans
}

// GetHumans returns a matrix with the humans
func (board *Board) GetHumans() [][]Human {
	ans := make([][]Human, board.Size, board.Size)
	for row := 0; row < board.Size; row++ {
		ans[row] = make([]Human, board.Size, board.Size)
		for col := 0; col < board.Size; col++ {
			var empty []*Human
			ans[row][col] = *(board.GetHuman(Pos{row, col}))
			ans[row][col].SetNeighbors(empty)
		}
	}
	return ans
}
