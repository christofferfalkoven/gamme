package simulate

import (
	"reflect"
	"testing"
	"time"
)

var testHuman1 = new(Human)
var testHuman2 = new(Human)
var testHuman3 = new(Human)
var testHuman4 = new(Human)
var testHuman5 = new(Human)

var neighbor1 = new(Human)
var neighbor2 = new(Human)
var neighbor3 = new(Human)

var neighborList1 = []*Human{neighbor1, neighbor2, neighbor1}
var neighborList2 = []*Human{neighbor2, neighbor2, neighbor3}
var neighborList3 = []*Human{neighbor3, neighbor1, neighbor2}

func TestHuman_CreateHuman(t *testing.T) {
	type args struct {
		Pos Pos
	}

	//middle
	testHuman1.Pos = Pos{1, 1}
	testHuman1.State = 0
	testHuman1.Radius = 1

	//corner case
	testHuman2.Pos = Pos{0, 2}
	testHuman2.State = 0
	testHuman2.Radius = 1

	//side case
	testHuman3.Pos = Pos{1, 2}
	testHuman3.State = 0
	testHuman3.Radius = 1

	//corner case
	testHuman4.Pos = Pos{2, 0}
	testHuman4.State = 0
	testHuman4.Radius = 1

	//corner case
	testHuman5.Pos = Pos{2, 2}
	testHuman5.State = 0
	testHuman5.Radius = 1

	tests := []struct {
		name string
		args args
		want *Human
	}{
		{"CreateHuman1:", args{Pos{1, 1}}, testHuman1},
		{"CreateHuman2:", args{Pos{0, 2}}, testHuman2},
		{"CreateHuman3:", args{Pos{1, 2}}, testHuman3},
		{"CreateHuman4:", args{Pos{2, 0}}, testHuman4},
		{"CreateHuman5:", args{Pos{2, 2}}, testHuman5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateHuman(tt.args.Pos, 1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateHuman() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHuman_setState(t *testing.T) {
	type fields struct {
		Pos       Pos
		State     State
		Radius    int
		Neighbors []*Human
	}
	type args struct {
		State State
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   State
	}{
		{"SetState 42:", fields{testHuman1.Pos, testHuman1.State, testHuman1.Radius, testHuman1.Neighbors}, args{42}, 0},
		{"SetState 0:", fields{testHuman2.Pos, testHuman2.State, testHuman2.Radius, testHuman2.Neighbors}, args{0}, 0},
		{"SetState 1:", fields{testHuman3.Pos, testHuman3.State, testHuman3.Radius, testHuman3.Neighbors}, args{1}, 1},
		{"SetState 6:", fields{testHuman4.Pos, testHuman4.State, testHuman4.Radius, testHuman4.Neighbors}, args{6}, 6},
		{"SetState -42:", fields{testHuman5.Pos, testHuman5.State, testHuman5.Radius, testHuman5.Neighbors}, args{-42}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Human{
				Pos:       tt.fields.Pos,
				State:     tt.fields.State,
				Radius:    tt.fields.Radius,
				Neighbors: tt.fields.Neighbors,
			}
			h.SetState(tt.args.State)
			if h.State != tt.want {
				t.Errorf("got SetState(%v) = %v, wanted: %v", tt.args.State, h.State, tt.want)
			}
		})
	}
}

func TestHuman_GetState(t *testing.T) {
	type fields struct {
		Pos       Pos
		State     State
		Radius    int
		Neighbors []*Human
	}
	tests := []struct {
		name   string
		fields fields
		want   State
	}{
		//TODO: add more testHumans to actually test GetState correctly
		{"GetState:", fields{testHuman1.Pos, 1, testHuman1.Radius, testHuman1.Neighbors}, 1},
		{"GetState:", fields{testHuman1.Pos, 2, testHuman1.Radius, testHuman1.Neighbors}, 2},
		{"GetState:", fields{testHuman1.Pos, 5, testHuman1.Radius, testHuman1.Neighbors}, 5},
		{"GetState:", fields{testHuman1.Pos, 4, testHuman1.Radius, testHuman1.Neighbors}, 4},
		{"GetState:", fields{testHuman1.Pos, 0, testHuman1.Radius, testHuman1.Neighbors}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Human{
				Pos:       tt.fields.Pos,
				State:     tt.fields.State,
				Radius:    tt.fields.Radius,
				Neighbors: tt.fields.Neighbors,
			}
			if got := h.GetState(); got != tt.want {
				t.Errorf("got GetState() = %v, wanted: %v", got, tt.want)
			}
		})
	}
}

func TestHuman_setRadius(t *testing.T) {
	type fields struct {
		Pos       Pos
		State     State
		Radius    int
		Neighbors []*Human
	}
	type args struct {
		Radius int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		//TODO: Fix testcases, what should the limits of Radius be?
		{"SetRadius 2:", fields{testHuman1.Pos, testHuman1.State, testHuman1.Radius, testHuman1.Neighbors}, args{2}, 2},
		{"SetRadius 0:", fields{testHuman2.Pos, testHuman2.State, testHuman2.Radius, testHuman2.Neighbors}, args{0}, 0},
		{"SetRadius -2:", fields{testHuman3.Pos, testHuman3.State, testHuman3.Radius, testHuman3.Neighbors}, args{-2}, -2},
		{"SetRadius SIZE+1:", fields{testHuman4.Pos, testHuman4.State, testHuman4.Radius, testHuman4.Neighbors}, args{SIZE + 1}, 4},
		{"SetRadius -SIZE-1:", fields{testHuman5.Pos, testHuman5.State, testHuman5.Radius, testHuman5.Neighbors}, args{-SIZE - 1}, -4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Human{
				Pos:       tt.fields.Pos,
				State:     tt.fields.State,
				Radius:    tt.fields.Radius,
				Neighbors: tt.fields.Neighbors,
			}
			h.SetRadius(tt.args.Radius)
			if h.Radius != tt.want {
				t.Errorf("got SetRadius(%v) = %v, wanted: %v", tt.args.Radius, h.Radius, tt.want)
			}
		})
	}
}

func TestHuman_GetRadius(t *testing.T) {
	type fields struct {
		Pos       Pos
		State     State
		Radius    int
		Neighbors []*Human
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add more test cases.
		// TODO: create more testHumans to test GetRadius correctly
		{"GetRadius:", fields{testHuman1.Pos, testHuman1.State, testHuman1.Radius, testHuman1.Neighbors}, 1},
		{"GetRadius:", fields{testHuman2.Pos, testHuman2.State, 5, testHuman2.Neighbors}, 5},
		{"GetRadius:", fields{testHuman3.Pos, testHuman3.State, 3, testHuman3.Neighbors}, 3},
		{"GetRadius:", fields{testHuman4.Pos, testHuman4.State, 8, testHuman4.Neighbors}, 8},
		{"GetRadius:", fields{testHuman5.Pos, testHuman5.State, 2, testHuman5.Neighbors}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Human{
				Pos:       tt.fields.Pos,
				State:     tt.fields.State,
				Radius:    tt.fields.Radius,
				Neighbors: tt.fields.Neighbors,
			}
			if got := h.GetRadius(); got != tt.want {
				t.Errorf("got GetRadius() = %v, wanted: %v", got, tt.want)
			}
		})
	}
}

func TestHuman_addNeighbor(t *testing.T) {
	type fields struct {
		Pos       Pos
		State     State
		Radius    int
		Neighbors []*Human
	}
	type args struct {
		neighbor *Human
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Human
	}{
		// TODO: Add test cases.
		{"addNeighbor1 neighbor:", fields{testHuman1.Pos, testHuman1.State, testHuman1.Radius, testHuman1.Neighbors}, args{neighbor1}, neighbor1},
		{"addNeighbor2 neighbor:", fields{testHuman2.Pos, testHuman2.State, testHuman2.Radius, neighborList1}, args{neighbor3}, neighbor3},
		{"addNeighbor3 neighbor:", fields{testHuman2.Pos, testHuman2.State, testHuman2.Radius, neighborList2}, args{neighbor1}, neighbor1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Human{
				Pos:       tt.fields.Pos,
				State:     tt.fields.State,
				Radius:    tt.fields.Radius,
				Neighbors: tt.fields.Neighbors,
			}
			h.AddNeighbor(tt.args.neighbor)

			var ans = func(Neighbors []*Human) bool {
				for i := range Neighbors {
					if Neighbors[i] == neighbor1 {
						return true
					}
				}
				return false
			}(h.Neighbors)

			if !ans {
				t.Errorf("AddNeighbor(%v), was not added to Neighbors: %v", neighbor1, tt.fields.Neighbors)
			}

		})
	}
}

func TestHuman_GetNeighbors(t *testing.T) {
	type fields struct {
		Pos       Pos
		State     State
		Radius    int
		Neighbors []*Human
	}
	var hum = new(Human)
	var ans1 = append(neighborList1, hum)
	var ans2 = append(neighborList2, hum)
	var ans3 = append(neighborList3, hum)

	tests := []struct {
		name   string
		fields fields
		want   []*Human
	}{
		// TODO: Add test cases.
		{"GetNeighbors:", fields{testHuman1.Pos, testHuman1.State, testHuman1.Radius, neighborList1}, ans1},
		{"GetNeighbors:", fields{testHuman2.Pos, testHuman2.State, testHuman2.Radius, neighborList2}, ans2},
		{"GetNeighbors:", fields{testHuman3.Pos, testHuman3.State, testHuman3.Radius, neighborList3}, ans3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Human{
				Pos:       tt.fields.Pos,
				State:     tt.fields.State,
				Radius:    tt.fields.Radius,
				Neighbors: tt.fields.Neighbors,
			}
			h.AddNeighbor(hum)
			if got := h.GetNeighbors(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Human.GetNeighbors() = %v, want %v", got, tt.want)
			}
		})
	}
}

//TODO: Write more testcases

func TestHuman_calcNeighbors(t *testing.T) {
	type fields struct {
		Pos       Pos
		State     State
		Radius    int
		Neighbors []*Human
	}
	type args struct {
		board *Board
	}

	var ans []*Human
	board := new(Board)
	board.AllocateBoard(SIZE)
	board.FillBoard(1)

	time.Sleep(50 * time.Millisecond)

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			ans = append(ans, board.GetHuman(Pos{i, j}))
		}
	}
	ans = append(ans[:4], ans[5:]...)

	var ans2= []*Human{board.GetHuman(Pos{0, 1}), board.GetHuman(Pos{1, 1}), board.GetHuman(Pos{1, 2})}
	var ans3= []*Human{board.GetHuman(Pos{0, 1}), board.GetHuman(Pos{0, 2}), board.GetHuman(Pos{1, 1}), board.GetHuman(Pos{2, 1}), board.GetHuman(Pos{2, 2})}
	var ans4= []*Human{board.GetHuman(Pos{1, 0}), board.GetHuman(Pos{1, 1}), board.GetHuman(Pos{2, 1})}
	var ans5 = []*Human{board.GetHuman(Pos{1, 1}), board.GetHuman(Pos{1, 2}), board.GetHuman(Pos{2, 1})}

		var listOfAns = [][]*Human{ans, ans2, ans3, ans4, ans5}

		tests := []struct{
		name   string
		fields fields
		args   args
		want   []*Human
	}{
		// TODO: Add test cases.
	{"calcNeighbors1:", fields{testHuman1.Pos, testHuman1.State, 1, []*Human{}}, args{board}, ans},
	{"calcNeighbors2:", fields{testHuman2.Pos, testHuman2.State, 1, []*Human{}}, args{board}, ans2},
	{"calcNeighbors3:", fields{testHuman3.Pos, testHuman3.State, 1, []*Human{}}, args{board}, ans3},
	{"calcNeighbors4:", fields{testHuman4.Pos, testHuman4.State, 1, []*Human{}}, args{board}, ans4},
	{"calcNeighbors5:", fields{testHuman5.Pos, testHuman5.State, 1, []*Human{}}, args{board}, ans5},
	}
		for i, tt := range tests{
		t.Run(tt.name, func (t *testing.T){
		h := &Human{
		Pos:       tt.fields.Pos,
		State:     tt.fields.State,
		Radius:    tt.fields.Radius,
		Neighbors: tt.fields.Neighbors,
	}
		h.CalcNeighbors(tt.args.board)
		if got := h.GetNeighbors(); !reflect.DeepEqual(got, tt.want){
		t.Errorf("CalcNeighbors(),%v was not equal to Neighbors: %v", h.Neighbors, listOfAns[i])
	}
	})
	}
	}

	func TestHuman_RemoveNeighbour(t * testing.T) {
		type fields struct {
			Pos       Pos
			State     State
			Radius    int
			Neighbors []*Human
		}

		type args struct {
			neighbor *Human
		}

		var ans1= neighborList1[1:2]
		var ans2= neighborList2[:len(neighborList2)-1]
		var ans3= make([]*Human, len(neighborList3))
		copy(ans3, neighborList3)
		ans3 = append(ans3[:1], ans3[2:]...)

		var listOfAns= [][]*Human{ans1, ans2, ans3}

		tests := []struct {
			name   string
			fields fields
			args   *Human
			want   []*Human
		}{
			// TODO: Add test cases.
			{"removeNeighbors1:", fields{testHuman1.Pos, testHuman1.State, 1, neighborList1}, neighbor1, ans1},
			{"removeNeighbors2:", fields{testHuman2.Pos, testHuman2.State, 1, neighborList2}, neighbor3, ans2},
			{"removeNeighbors3:", fields{testHuman3.Pos, testHuman3.State, 1, neighborList3}, neighbor1, ans3},
		}
		for i, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				h := &Human{
					Pos:       tt.fields.Pos,
					State:     tt.fields.State,
					Radius:    tt.fields.Radius,
					Neighbors: tt.fields.Neighbors,
				}
				h.RemoveNeighbor(tt.args)
				if got := h.GetNeighbors(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Removed list,%v was not equal to expected list: %v", h.Neighbors, listOfAns[i])
				}
			})
		}
	}

	func TestHuman_SetNeighbours(t * testing.T) {
		type fields struct {
			Pos       Pos
			State     State
			Radius    int
			Neighbors []*Human
		}
		type args struct {
			Neighbors []*Human
		}

		var ans1 = neighborList1
		var ans2 = neighborList2
		var ans3 = neighborList3

		var listOfAns = [][]*Human{ans1, ans2, ans3}

		tests := []struct {
			name   string
			fields fields
			args   args
			want   []*Human
		}{
			// TODO: Add test cases.
			{"SetNeighbors1:", fields{testHuman1.Pos, testHuman1.State, 1, []*Human{}}, args{neighborList1}, ans1},
			{"SetNeighbors2:", fields{testHuman2.Pos, testHuman2.State, 1, []*Human{}}, args{neighborList2}, ans2},
			{"SetNeighbors3:", fields{testHuman3.Pos, testHuman3.State, 1, []*Human{}}, args{neighborList3}, ans3},
		}
		for i, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				h := &Human{
					Pos:       tt.fields.Pos,
					State:     tt.fields.State,
					Radius:    tt.fields.Radius,
					Neighbors: tt.fields.Neighbors,
				}
				h.SetNeighbors(tt.args.Neighbors)
				if got := h.GetNeighbors(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Set list,%v was not equal to expected list: %v", h.Neighbors, listOfAns[i])
				}
			})
		}
	}