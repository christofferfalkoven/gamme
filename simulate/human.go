package simulate

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

// State of the human health
type State int

// Pos represent the x and y position in the matrix
type Pos struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Human struct
// Contains Pos, state, radius and all the neighbors if they are calculated.
type Human struct {
	Pos         Pos     `json:"pos"`
	State       State   `json:"state"`
	Radius      int     `json:"radius"`
	Age         int     `json:"age"`
	Healthiness float64 `json:"healthiness"`
	Contagious  bool    `json:"contagious"`
	Neighbors   []*Human
	Viruses     []*Virus `json:"virus"`
	mux         sync.Mutex
	Infected    bool	`json:"infected"`
}

// CreateHuman TODO: desc
func CreateHuman(pos Pos, rng int) *Human {
	temp := new(Human)
	temp.Pos = pos
	temp.State = 1
	temp.Radius = rng
	temp.Infected = false
	temp.GenerateAge()
	temp.CalculateHealthiness()
	return temp
}

// GenerateAge generates the age for the human.
func (h *Human) GenerateAge() {
	rand.Seed(time.Now().UTC().UnixNano())
	h.Age = rand.Intn(100) // generate the age between 0 - 100
}

// CalculateHealthiness generates the healthiness for the human
func (h *Human) CalculateHealthiness() {
	rand.Seed(time.Now().UTC().UnixNano())
	if h.Age > 20 {
		h.Healthiness = float64(100 - rand.Intn(h.Age-10))
	} else {
		h.Healthiness = 100.0
	}
}

// SetState TODO: desc
// TODO: Should we have limits when setting SetState?
func (h *Human) SetState(state State) {
	if state >= -2 && state <= 100 {
		h.State = state
	}
}

// GetState TODO: desc
func (h *Human) GetState() State {
	return h.State
}

// SetRadius TODO: desc
func (h *Human) SetRadius(radius int) {
	h.Radius = radius
}

// GetRadius TODO: desc
func (h *Human) GetRadius() int {
	return h.Radius
}

// AddNeighbor TODO: desc
func (h *Human) AddNeighbor(neighbor *Human) {
	h.Neighbors = append(h.Neighbors, neighbor)
}

// SetNeighbors TODO: desc
func (h *Human) SetNeighbors(neighbors []*Human) {
	h.Neighbors = neighbors
}

// RemoveNeighbor TODO: desc
func (h *Human) RemoveNeighbor(neighbor *Human) {
	neighbors := h.GetNeighbors()
	for i := 0; i < len(neighbors); i++ {
		if neighbors[i] == neighbor {
			neighbors = append(neighbors[:i], neighbors[i+1:]...)
			i--
		}
	}
	h.SetNeighbors(neighbors)
}

// GetNeighbors TODO: desc
func (h *Human) GetNeighbors() []*Human {
	return h.Neighbors
}

// CalcNeighbors TODO: desc
func (h *Human) CalcNeighbors(board *Board) {
	xVal := h.Pos.X
	yVal := h.Pos.Y

	for i := int(math.Max(0, float64(xVal-h.Radius))); i <= int(math.Min(float64(board.Size-1), float64(xVal+h.Radius))); i++ {
		for j := int(math.Max(0, float64(yVal-h.Radius))); j <= int(math.Min(float64(board.Size-1), float64(yVal+h.Radius))); j++ {
			if (Pos{i, j} != h.Pos) {
				h.AddNeighbor(board.GetHuman(Pos{i, j}))
			}
		}
	}
}
