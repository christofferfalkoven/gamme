package simulate

import (
	"reflect"
	"testing"
)

//testing functions/methods in board.go except allocateBoard(), CalcAllNeighbours(), FillBoard() and PrintBoard()

var TestBoard1 = new(Board)
var TestBoard2 = new(Board)
var TestBoard3 = new(Board)
var TestBoard4 = new(Board)
var TestBoard5 = new(Board)
var TestHuman1 = new(Human)
var TestHuman2 = new(Human)
var TestHuman3 = new(Human)
var TestHuman4 = new(Human)

//set up a test board to use in the test functions
func init() {
	TestBoard1.AllocateBoard(2)
	//set fields in humans
	for i := 0; i < 2; i++ {
		TestBoard1.matrix[i] = make([]*Human, 2, 2)
		for j := 0; j < 2; j++ {
			TestBoard1.matrix[i][j] = new(Human)
			TestBoard1.matrix[i][j].Pos =Pos{i, j}
			TestBoard1.matrix[i][j].State = 1
			TestBoard1.matrix[i][j].Radius = 1
		}
	}
}

func TestBoard_GetMatrix(t *testing.T) {
	type fields struct {
		matrix [][]*Human
		size   int
	}
	var ans = make([][]*Human, 2, 2)
	for i := 0; i < 2; i++ {
		ans[i] = make([]*Human, 2, 2)
		for j := 0; j < 2; j++ {
			ans[i][j] = new(Human)
			ans[i][j].Pos =Pos{i, j}
			ans[i][j].State = 1
			ans[i][j].Radius = 1
		}
	}
	tests := []struct {
		name  string
		fields fields
		want  [][]*Human
	}{
		{"GetMatrix:", fields{TestBoard1.matrix, TestBoard1.size}, ans},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				matrix:	tt.fields.matrix,
				size:	tt.fields.size,
			}
			if got := b.GetMatrix(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Board.GetMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_SetMatrix(t *testing.T) {
	type fields struct {
		matrix [][]*Human
		size   int
	}
	//allocate TestBoard2, matrix is zeroed
	TestBoard2.AllocateBoard(2)
	//create a new matrix to be added to TestBoard2
	var NewMatrix = make([][]*Human, 2, 2)
	for i := 0; i < 2; i++ {
		NewMatrix[i] = make([]*Human, 2, 2)
		for j := 0; j < 2; j++ {
			NewMatrix[i][j] = new(Human)
			NewMatrix[i][j].Pos =Pos{i, j}
			NewMatrix[i][j].State = 1
			NewMatrix[i][j].Radius = 1
		}
	}
	type args struct {
		matrix [][]*Human
	}
	tests := []struct {
		name  string
		fields fields
		args  args
		want   [][]*Human
	}{
		{"SetMatrix:", fields{TestBoard2.matrix, TestBoard2.size}, args{NewMatrix}, NewMatrix  },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				matrix:	tt.fields.matrix,
				size: 	tt.fields.size,
			}
			b.SetMatrix(tt.args.matrix)
			if got := b.matrix; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Board.SetMatrix() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestBoard_SetMatrixRow(t *testing.T) {
	type fields struct {
		matrix [][]*Human
		size   int
	}
	//allocate TestBoard2, matrix is zeroed
	TestBoard2.AllocateBoard(2)

	//create a new matrix row (slice) to be added to TestBoard2
	var newMatrixrow = make([]*Human, 2, 2)
	newMatrixrow[0] = TestHuman1
	newMatrixrow[1] = TestHuman2

	type args struct {
		humans []*Human
		row 	int
	}
	tests := []struct {
		name  string
		fields fields
		args  args
		want  []*Human
	}{
		{"SetMatrixRow:", fields{TestBoard2.matrix, TestBoard2.size}, args{newMatrixrow, 1}, newMatrixrow },
	}
	for _, tt := range tests {
		b := &Board{
			matrix:	tt.fields.matrix,
			size: 	tt.fields.size,
		}
		b.SetMatrixRow(tt.args.humans, tt.args.row)
		if got := b.matrix[tt.args.row]; !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Board.SetMatrixRow() = %v, want %v", got, tt.want)
		}
	}
}

func TestBoard_GetHuman(t *testing.T) {
	type fields struct {
		matrix [][]*Human
		size   int
	}
	//allocate TestBoard2, matrix is zeroed
	TestBoard2.AllocateBoard(2)

	//create matrix of humans to be added to TestBoard2
	HumanMatrix := [][]*Human{{TestHuman1, TestHuman2}, {TestHuman3, TestHuman4}}

	//set the position of the humans
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			HumanMatrix[i][j].Pos =Pos{i, j}
		}
	}

	//add the matrix to TestBoard2
	TestBoard2.matrix = HumanMatrix

	type args struct {
		Pos Pos
	}
	tests := []struct {
		name  string
		fields fields
		args  args
		want  *Human
	}{
		{"GetHuman upper left:", fields{TestBoard2.matrix, TestBoard2.size}, args{Pos{0, 0}}, TestHuman1 },
		{"GetHuman upper right:", fields{TestBoard2.matrix, TestBoard2.size}, args{Pos{0,1}}, TestHuman2 },
		{"GetHuman lower left:", fields{TestBoard2.matrix, TestBoard2.size}, args{Pos{1, 0}}, TestHuman3 },
		{"GetHuman lower right:", fields{TestBoard2.matrix, TestBoard2.size}, args{Pos{1, 1}}, TestHuman4 },

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				matrix:	tt.fields.matrix,
				size:	tt.fields.size,
			}
			if got := b.GetHuman(tt.args.Pos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Board.GetHuman() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_SetHuman(t *testing.T) {
	type fields struct {
		matrix [][]*Human
		size   int
	}
	//set the positions of the test humans
	TestHuman1.Pos = Pos{0,0}
	TestHuman4.Pos = Pos{1,1}

	type args struct {
		human *Human
	}
	tests := []struct {
		name  string
		fields fields
		args  args
		want *Human
	}{
		{"SetHuman upper left:", fields{TestBoard1.matrix, TestBoard1.size}, args{TestHuman1}, TestHuman1 },
		{"SetHuman lower right:", fields{TestBoard1.matrix, TestBoard1.size}, args{TestHuman4}, TestHuman4 },
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				matrix:	tt.fields.matrix,
				size:	tt.fields.size,
			}
			b.SetHuman(tt.args.human)
			if got := b.matrix[i][i] ; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Board.SetHuman() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenRow(t *testing.T) {
	type fields struct {
		matrix	[][]*Human
		size	int
	}
	//make a channel that can be sent as an argument to function
	//these channels need to be buffered, otherwise there will be a deadlock executing this test!
	//deadlock because at the end of GenRow func we try to write to the Receiver channel, but the write will block because there is no reader
	//reader is only created in line 269 and this line will never be reached
	//alternative solution: put the write in a separate routine, so that "main" can continure executing, reaching line 269
	Receiver1 := make(chan []*Human)
	Receiver2 := make(chan []*Human)

	TestBoard3.AllocateBoard(2)

	//make a slice of human's positions that we expect to get
	AnswerSlice1 := []Pos{{0, 0}, {0, 1}}
	AnswerSlice2 := []Pos{{1, 0}, {1, 1}}

	type args struct {
		row    int
		signal chan []*Human
		rng    int
	}
	tests := []struct {
		name string
		fields fields
		args args
		want []Pos
	}{
		{"GenRow1:",fields{TestBoard3.matrix, TestBoard3.size }, args{0, Receiver1, 1}, AnswerSlice1 },
		{"GenRow1:",fields{TestBoard3.matrix, TestBoard3.size }, args{1, Receiver2, 1}, AnswerSlice2 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go GenRow(tt.args.row, tt.args.signal, tt.args.rng)
			ReceivedSlice := <- tt.args.signal
			GotSlice := []Pos{ReceivedSlice[0].Pos, ReceivedSlice[1].Pos}
			if got := GotSlice; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReceiveSignal(t *testing.T) {
	type fields struct {
		matrix [][]*Human
		size int
	}
	//allocate for TestBoard4
	TestBoard4.AllocateBoard(2)

	//create slices of human pointers that can be sent on a channel
	TestSlice1 := []*Human{TestHuman1, TestHuman2}
	TestSlice2 := []*Human{TestHuman3, TestHuman4}

	AggregateSlice := [][]*Human{TestSlice1, TestSlice2}

	//make a channel that can be sent as an argument to function to ReceiveSignal()
	ReadFrom := make(chan []*Human)

	type args struct {
		signal chan []*Human
		board  *Board
		row    int
	}
	tests := []struct {
		name string
		fields fields
		args args
		want []*Human
	}{
		{"ReceiveSignal to row 1:", fields{TestBoard4.matrix, TestBoard4.size}, args{ReadFrom, TestBoard4, 0 }, TestSlice1},
		{"ReceiveSignal to row 2:", fields{TestBoard4.matrix, TestBoard4.size}, args{ReadFrom, TestBoard4, 1 }, TestSlice2},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//send slices of human pointers on the channel
			 go func() {
				 tt.args.signal <- AggregateSlice[i]
			 }()
			ReceiveSignal(tt.args.signal, tt.args.board, tt.args.row)
			if got := TestBoard4.matrix[i]; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReceiveSignal() = %v, want %v", got, tt.want)
			}
		})
	}
}

//test sequential version of FillBoard first
func TestFillBoardSeq(t *testing.T) {
	type fields struct {
		matrix	[][]*Human
		size	int
	}
	TestBoard5.AllocateBoard(2)
	//create a matrix of human's positions that we expect to get
	AnswerSlice1 := []Pos{{0, 0}, {0, 1}}
	AnswerSlice2 := []Pos{{1, 0}, {1, 1}}
	AnswerMatrix := [][]Pos{AnswerSlice1, AnswerSlice2}

	type args struct {
		rng    int
	}
	tests := []struct {
		name string
		fields fields
		args args
		want [][]Pos
	}{
		{"FillBoardSeq:",fields{TestBoard5.matrix, TestBoard5.size }, args{1}, AnswerMatrix },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				matrix: tt.fields.matrix,
				size: 	tt.fields.size,
			}
			b.FillBoardSeq(tt.args.rng)
			GotMatrix := make([][]Pos, 2, 2)
			GotMatrix[0] = []Pos{TestBoard5.matrix[0][0].Pos, TestBoard5.matrix[0][1].Pos}
			GotMatrix[1] = []Pos{TestBoard5.matrix[1][0].Pos, TestBoard5.matrix[1][1].Pos}
			if got := GotMatrix; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FillBoardSeq() = %v, want %v", got, tt.want)
			}
		})
	}
}

//test of FillBoard will be to benchmark it and compare it to the sequential version FillBoardSeq
func TestBoard_FillBoard(t *testing.T) {
	type args struct {
		rng int
	}
	tests := []struct {
		name  string
		board *Board
		args  args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.board.FillBoard(tt.args.rng)
		})
	}
}

//TODO: Fix test - this test is failing
func TestBoard_CalcAllNeighborsAUX(t *testing.T) {
	type fields struct {
		matrix	[][]*Human
		size	int
	}
	//create a slice of slices of human pointers that represents the neighbour lists that we want
	ListOfNeighbourLists1 := [][]*Human{{TestHuman2, TestHuman3, TestHuman4}, {TestHuman2, TestHuman3, TestHuman4}}
	ListOfNeighbourLists2 := [][]*Human{{TestHuman1, TestHuman2, TestHuman4}, {TestHuman1, TestHuman2, TestHuman3}}

	type args struct {
		row int
	}
	tests := []struct {
		name  string
		fields fields
		args  args
		want [][]*Human
	}{
		{"CalcAllNeighboursAUX row 1:", fields{TestBoard2.matrix, TestBoard2.size}, args{0}, ListOfNeighbourLists1},
		{"CalcAllNeighboursAUX row 2:", fields{TestBoard2.matrix, TestBoard2.size}, args{1}, ListOfNeighbourLists2},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				matrix: tt.fields.matrix,
				size:  	tt.fields.size,
			}
			b.CalcAllNeighborsAUX(tt.args.row)
			gotList := make([][]*Human, 2, 2)
				for j := 0; j < TestBoard2.size; j++{
					gotList[j] = make([]*Human, 3, 3)
					gotList[j] = TestBoard2.matrix[i][j].Neighbors
				}
			if got := gotList ; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Board.CalcAllNeighboursAUX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_GetHumanStates(t *testing.T) {
	type fields struct {
		matrix	[][]*Human
		size	int
	}
	//set states of humans to 0 in TestBoard2
	TestBoard2.AllocateBoard(2)
	for i := 0; i < 2; i++ {
		TestBoard2.matrix[i] = make([]*Human, 2, 2)
		for j := 0; j < 2; j++ {
			TestBoard2.matrix[i][j] = new(Human)
			TestBoard2.matrix[i][j].State = 0
		}
	}

	//create matrices of states
	//because we set pos (0, 0) and (1, 1) to hold a TestHuman(uninitialized) before, the states on these positions == 0
	AnswerMatrix1 := [][]State{{0, 1}, {1, 0}}
	AnswerMatrix2 := [][]State{{0, 0}, {0, 0}}

	tests := []struct {
		name  string
		fields fields
		want  [][]State
	}{
		{"GetHumanStates TestBoard1:", fields{TestBoard1.matrix, TestBoard1.size}, AnswerMatrix1  },
		{"GetHumanStates TestBoard2:", fields{TestBoard2.matrix, TestBoard2.size}, AnswerMatrix2  },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				matrix:	tt.fields.matrix,
				size:	tt.fields.size,
			}
			if got := b.GetHumanStates(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Board.GetHumanStates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_GetHumans(t *testing.T) {
	type fields struct {
		matrix [][]*Human
		size	int
	}
	//create the matrix that we expect to get as return value
	var Human1 = *(TestBoard1.matrix[0][0])
	var Human2 = *(TestBoard1.matrix[0][1])
	var Human3 = *(TestBoard1.matrix[1][0])
	var Human4 = *(TestBoard1.matrix[1][1])

	AnswerMatrix := [][]Human{{Human1, Human2}, {Human3, Human4}}

	tests := []struct {
		name  string
		fields fields
		want  [][]Human
	}{
		{"GetHumans:", fields{TestBoard1.matrix, TestBoard1.size}, AnswerMatrix },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				matrix:	tt.fields.matrix,
				size:	tt.fields.size,
			}
			if got := b.GetHumans(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Board.GetHumans() = %v, want %v", got, tt.want)
			}
		})
	}
}
