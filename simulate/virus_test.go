package simulate

import (
	"reflect"
	"testing"
)

func TestCreateVirus(t *testing.T) {
	type args struct {
		Range          int
		Contagiousness int
		Contagiousdays int
	}
	tests := []struct {
		name string
		args args
		want *Virus
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateVirus(tt.args.Range, tt.args.Contagiousness, tt.args.Contagiousdays); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateVirus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVirus_GetRange(t *testing.T) {
	tests := []struct {
		name string
		v    *Virus
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.GetRange(); got != tt.want {
				t.Errorf("Virus.GetRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVirus_GetContagiousness(t *testing.T) {
	tests := []struct {
		name string
		v    *Virus
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.GetContagiousness(); got != tt.want {
				t.Errorf("Virus.GetContagiousness() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVirus_GetContagiousDays(t *testing.T) {
	tests := []struct {
		name string
		v    *Virus
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.GetContagiousDays(); got != tt.want {
				t.Errorf("Virus.GetContagiousDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHuman_RemoveInfectedNeighbors(t *testing.T) {
	tests := []struct {
		name  string
		human *Human
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.human.RemoveInfectedNeighbors()
		})
	}
}

func TestHuman_Infect(t *testing.T) {
	type args struct {
		target *Human
	}
	tests := []struct {
		name     string
		Infecter *Human
		args     args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.Infecter.Infect(tt.args.target)
		})
	}
}

func TestUpdateHumanState(t *testing.T) {
	type args struct {
		human *Human
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateHumanState(tt.args.human)
		})
	}
}

func TestSpreadVirus(t *testing.T) {
	type args struct {
		human *Human
		virus *Virus
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SpreadVirus(tt.args.human, tt.args.virus)
		})
	}
}

func TestHuman_InfectAndRemove(t *testing.T) {
	type args struct {
		virus *Virus
	}
	tests := []struct {
		name  string
		human *Human
		args  args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.human.InfectAndRemove(tt.args.virus)
		})
	}
}
