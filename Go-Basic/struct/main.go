package main

import (
	"encoding/json"
	"fmt"
)

type Player struct {
	Username string `json:"Username"`
	Lavel    uint   `json:"lavel"`
	Status   string `json:"status"`
	Class    string `json:"class"`
}

func (p *Player) LavelUp() {
	p.Lavel++
}

func main() {
	Player1 := Player{
		Username: "Rio",
		Lavel:    60,
		Status:   "active",
		Class:    "mage"}

	jsondata, _ := json.MarshalIndent(&Player1, "", "\t")
	fmt.Println(string(jsondata))

	Player1.LavelUp()
	jsondata, _ = json.MarshalIndent(&Player1, "", "\t")
	fmt.Println(string(jsondata))

}
