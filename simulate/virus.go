package simulate

import (
	"math"
	"math/rand"
	"time"
)

// Virus TODO: desc
type Virus struct {
	Name           string `json:"name"`
	Contagiousness int    `json:"contagiousness"` //probability (in %) that a human is Infected
	Contagiousdays int    `json:"contagiousdays"` //days that human is contagious
	Lifespan       int    `json:"lifespan"`       //days that the virus lives - lifespan-contagious days = infected days(days when infected but not contagious)
	Deadliness     int    `json:"deadliness"`     //probability (in %) that a human dies (if not dead = immune)
	Range          int    `json:"range"`          //how far the virus can infect
}

// CreateVirus TODO: desc
func CreateVirus(Range int, Contagiousness int, Contagiousdays int) *Virus {
	temp := new(Virus)
	temp.Range = Range
	temp.Contagiousness = Contagiousness
	temp.Contagiousdays = Contagiousdays
	return temp
}

// GetRange TODO: desc
func (v *Virus) GetRange() int {
	return v.Range
}

// GetContagiousness TODO: desc
func (v *Virus) GetContagiousness() int {
	return v.Contagiousness
}

// GetContagiousDays TODO: desc
func (v *Virus) GetContagiousDays() int {
	return v.Contagiousdays
}

// RemoveInfectedNeighbors TODO: desc
func (human *Human) RemoveInfectedNeighbors() {
	for i := 0; i < len(human.Neighbors); i++ {
		if human.Neighbors[i].State != 1 {
			human.RemoveNeighbor(human.Neighbors[i])
		}
	}
}

// Infect TODO: desc
func (Infecter *Human) Infect(target *Human) {
	for i := range Infecter.Viruses {
		go SpreadVirus(target, Infecter.Viruses[i], false)
	}
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// UpdateHumanState TODO: desc
func (human *Human) UpdateHumanState() {

	var amountOfStates = 1 + human.Viruses[0].Lifespan
	var contagiousTime = amountOfStates - human.Viruses[0].Contagiousdays
	var virus = human.Viruses[0]
	var DecrementValue = math.Ceil((human.Healthiness / float64(virus.Lifespan)))
	var d = (float64(virus.Deadliness) / 100)
	var pl = math.Pow(float64(virus.Lifespan), -1)
	var probabilityToDie = math.Pow(d, pl)
	//var startHealth = human.Healthiness

	for i := 2; i <= amountOfStates; i++ {
		if human.State == 0 || human.State == -2 {
			human.Contagious = false
			return
		}

		if i >= contagiousTime {
			human.Contagious = true
		}

		human.SetState(State(i))

		rand.Seed(time.Now().UTC().UnixNano())
		var nr = rand.Float64()
		//fmt.Println("probability: ", probabilityToDie)
		//fmt.Println("random nr : ", nr)

		if nr <= probabilityToDie {
			//fmt.Println("loop nr: ", i)
			//fmt.Println("Human healthiness before decrease: ", human.Healthiness)
			human.Healthiness = human.Healthiness - DecrementValue
			//fmt.Println("Human healthiness after decrease: ", human.Healthiness)

		}
		if human.Healthiness <= 0 {
			human.SetState(-1)
			human.Contagious = false
			return
		}

		time.Sleep(time.Second)
	}
	human.SetState(0)
	human.Contagious = false
	//fmt.Println("human done updating!")
}

// TryToInfect tries to infect a human with the contagiousness as argument.
func (human *Human) TryToInfect(virus *Virus) bool {
	rand.Seed(time.Now().UTC().UnixNano())
	var InfectSeed = rand.Intn(100)
	if InfectSeed <= virus.GetContagiousness() {
		return true
	}
	return false
}

// SpreadVirus TODO: desc
func SpreadVirus(human *Human, virus *Virus, force bool) {
	human.mux.Lock()
	defer human.mux.Unlock()
	if !human.Infected {
		if force {
			human.Viruses = append(human.Viruses, virus)
			human.Infected = true
			go human.UpdateHumanState()
			go human.infectNeighbors(virus)

		} else {
			if human.TryToInfect(virus) {
				human.Viruses = append(human.Viruses, virus)
				human.Infected = true
				//start a goroutine that updates the state of the human
				go human.UpdateHumanState()

				//make this loop not infect all neighbors
				go human.infectNeighbors(virus)
			} else {
				human.SetState(-2)
				human.Contagious = false
			}
		}
	}
}

func (human *Human) infectNeighbors(virus *Virus) {
	var stopped = false
	for len(human.Neighbors) > 0 {
		human.RemoveInfectedNeighbors()
		for human.Contagious {
			stopped = true
			//fmt.Println("contagious")
			human.InfectAndRemove(virus)
			time.Sleep(time.Millisecond * 800)
		}
		if stopped == true {
			return
		}
		time.Sleep(time.Millisecond * 800)
	}
}

// InfectAndRemove TODO: desc
func (human *Human) InfectAndRemove(virus *Virus) {
	var length = len(human.Neighbors)
	if length > 1 {
		var index = rand.Intn(length - 1)
		go human.Infect(human.Neighbors[index])
		go human.RemoveNeighbor(human.Neighbors[index])
	} else if length == 1 {
		go human.Infect(human.Neighbors[0])
		go human.RemoveNeighbor(human.Neighbors[0])
		//fmt.Println("last neighbor being removed!")
	} else {
		//human.Contagious = false
		//fmt.Println("no neighbors left in list")
	}
}
