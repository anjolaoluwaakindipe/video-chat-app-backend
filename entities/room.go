package entities

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Participant describes a single entity in the hashmap
type Participant struct{
	Host bool
	Conn *websocket.Conn
}

// RoomMap in the main hashmap [roomID string] -> ([]Participant)
type RoomMap struct {
	Mutex sync.RWMutex
	Map map[string][]Participant

}

func NewRoomMap() *RoomMap{
	newRoomMap := &RoomMap{}
	newRoomMap.Init()
	return newRoomMap
}

// Init Initializes the RooMap struct
func  (rm *RoomMap) Init(){
	rm.Map = make(map[string][]Participant)
}

// Get will return the array of participants in the romm
func (rm *RoomMap) Get( roomID string) []Participant{
	rm.Mutex.RLock()
	defer rm.Mutex.RUnlock()

	return rm.Map[roomID]
}

// CreateRoom generate a unique room ID and return it -> insert it in the hashmap
func (rm *RoomMap) CreateRoom() string{
	rm.Mutex.RLock()
	defer rm.Mutex.RUnlock()

	rand.Seed(time.Now().Unix())

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b:=make([]rune, 8)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	roomID:= string(b)
	rm.Map[roomID] = []Participant{}

	return roomID
}

// InsertIntoRoom will create a participant and add it in the hashmap
func (rm *RoomMap) InsertIntoRoom(roomID string, host bool, conn *websocket.Conn){
	rm.Mutex.RLock()
	rm.Mutex.RUnlock()

	p:= Participant{Host: host, Conn: conn}
	log.Println("Insertin into Room with Room : ", roomID)
	rm.Map[roomID] =append(rm.Map[roomID], p)
}


// DeleteRoom deletes the room with the roomID
func (rm *RoomMap) DeleteRoom(roomID string){
	rm.Mutex.RLock()
	rm.Mutex.RUnlock()

	delete(rm.Map, roomID)
}