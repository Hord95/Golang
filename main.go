package main

import "fmt"

var hero gameCase
var room gameCase

type gameCase struct {
	step    int
	command string
	answer  string
}

func main() {
	initGame()
	fmt.Println(hero.answer)

}
func initGame() {
	hero.step = 1
	hero.command = "осмотреться"
	hero.answer = "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор"

}
func handleCommand(commands gameCase) {
	//var set,command,answer =room

}
