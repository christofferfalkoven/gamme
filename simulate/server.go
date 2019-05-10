package simulate

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"golang.org/x/net/websocket"
)

const address string = "localhost:8080"

// BoardMessageHuman structure for passing a message containing the board towards the client
type BoardMessageHuman struct {
	ID   int       `json:"id"`
	Name string    `json:"name"`
	Data [][]Human `json:"data"`
}


// Message structure for converting a received message containing viruses, and starting / stopping messages.
type Message struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Cont    int    `json:"contagiousness"`
	ContD   int    `json:"contagiousDays"`
	Life    int    `json:"lifespan"`
	Dead    int    `json:"deadliness"`
	Range   int    `json:"range"`
	Message string `json:"Message"`
	PosX    int    `json:"posX"`
	PosY    int    `json:"posY"`
}

var positionX int
var positionY int


// StartServer starts the go-server and saves a dial up connection towards to the client.
func StartServer(board *Board) {
	quitSig := true
	startSimulation := false
	origin := "http://localhost/"
	url := "ws://localhost:8080/ws"
	startVirus := new(Virus)

	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	var cMessage Message
	cMessage.ID = 0
	cMessage.Name = "goserver"
	connectMessage, _ := json.Marshal(cMessage)
	websocket.Message.Send(ws, string(connectMessage))
	go ReceiveMessages(board, ws, &startSimulation, startVirus)
	// While simulation waiting to start
	for true {
		fmt.Println("Not started simulation yet")
		if startSimulation {
			time.Sleep(time.Second)
			quitSig = false
			go SpreadVirus(board.GetHuman(Pos{positionY, positionX}), startVirus, true)
			go SendMessages(board, ws, &quitSig)
			fmt.Println("end of startsim")
			for startSimulation {
				time.Sleep(time.Millisecond * 250)
			}
		}

		quitSig = true
		time.Sleep(time.Second)
	}

}

// ReceiveMessages TODO: desc
func ReceiveMessages(board *Board, ws *websocket.Conn, startSim *bool, virus *Virus) {
	for {
		var msg Message
		time.Sleep(time.Millisecond * 250)
		err := websocket.JSON.Receive(ws, &msg)
		if err != nil {
			log.Fatal(err)
		}
		if msg.ID == 0 {
			virus.Name = msg.Name
			virus.Lifespan = msg.Life
			virus.Contagiousness = msg.Cont
			virus.Range = msg.Range
			virus.Contagiousdays = msg.ContD
			virus.Deadliness = msg.Dead
			positionX = msg.PosX
			positionY = msg.PosY

			fmt.Println("setting virus parameters -- DONE")
			board.FillBoard(virus.Range)
			fmt.Println("filling board")
			time.Sleep(time.Millisecond * 50)
			fmt.Println("Calculating neighbors")
			board.CalcAllNeighbors()
		} else if msg.ID == 3 && msg.Name == "start" {
			*startSim = true
			fmt.Println("STARTING SIMULATION")
		} else if msg.ID == 3 && msg.Name == "stop" {
			*startSim = false
			fmt.Println("STOPPING SIMULATION")
		} else {
			fmt.Println("Got a wrongful Message, cant handle it")
		}
	}
}

// SendMessages TODO: desc
func SendMessages(board *Board, ws *websocket.Conn, quitSig *bool) {
	for {
		if *quitSig {

		} else {
			time.Sleep(time.Millisecond * 250)
			var bMessage BoardMessageHuman
			bMessage.ID = 1
			bMessage.Name = "board"
			bMessage.Data = board.GetHumans()

			// converting the Message to JSON
			bm, _ := json.Marshal(bMessage)
			// Sending the Message to our js server
			websocket.Message.Send(ws, string(bm))
			//fmt.Println("SENT Message")
		}
	}
}
