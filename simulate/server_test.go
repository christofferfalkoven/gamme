package simulate

import (
	"testing"

	"golang.org/x/net/websocket"
)

func TestStartServer(t *testing.T) {
	type args struct {
		board *Board
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StartServer(tt.args.board)
		})
	}
}

func TestReceiveMessages(t *testing.T) {
	type args struct {
		board    *Board
		ws       *websocket.Conn
		startSim *bool
		virus    *Virus
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReceiveMessages(tt.args.board, tt.args.ws, tt.args.startSim, tt.args.virus)
		})
	}
}

func TestSendMessages(t *testing.T) {
	type args struct {
		board   *Board
		ws      *websocket.Conn
		quitSig *bool
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendMessages(tt.args.board, tt.args.ws, tt.args.quitSig)
		})
	}
}
