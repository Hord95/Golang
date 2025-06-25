package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Player представляет игрока
type Player struct {
	Location   *Room
	Inventory  []string
	isBackpack bool
	hasKey     bool
}

// Room представляет комнату в игре
type Room struct {
	Name            string
	Description     string
	Items           []string
	ChairItems      []string
	Exits           map[string]*Room
	Commands        map[string]func(*Player, []string) string
	Other           map[string]string
	ExitDescription string
	LockedExists    map[string]bool
}
type Other struct {
	item      string
	direction string
}

var player *Player
var rooms map[string]*Room

func initCommands() {
	// Глобальные команды доступные везде

	goCommand := func(p *Player, args []string) string {
		direction := args[0]
		if exit, ok := p.Location.Exits[direction]; ok {
			if p.Location.LockedExists[direction] == false {
				p.Location = exit
				return exit.ExitDescription
			} else {
				return "дверь закрыта"
			}
		}
		return "нет пути в " + direction
	}
	takeInventory := func(p *Player, args []string) string {
		item := args[0]

		if item == "рюкзак" {
			for i, roomItem := range p.Location.ChairItems {
				if roomItem == item {
					p.Location.ChairItems = append(p.Location.ChairItems[:i], p.Location.ChairItems[i+1:]...)
					p.Inventory = append(p.Inventory, item)
					p.isBackpack = true
					return "Вы надели " + item
				}
			}
		}
		return "Здесь нет такого предмета."
	}

	takeCommand := func(p *Player, args []string) string {

		item := args[0]

		if p.isBackpack {
			for i, roomItem := range p.Location.Items {
				if roomItem == item {
					if item == "ключи" {
						p.hasKey = true
					}
					// Удаляем из комнаты
					p.Location.Items = append(p.Location.Items[:i], p.Location.Items[i+1:]...)
					// Добавляем в инвентарь
					p.Inventory = append(p.Inventory, item)
					return "предмет добавлен в инвентарь: " + item
				}
			}

		} else {
			return "некуда класть"
		}
		return "нет такого"
	}

	useCommand := func(p *Player, args []string) string {
		item := args[0]
		other := args[1]
		hasItem := false
		for _, v := range p.Inventory {
			if item == v {
				hasItem = true
				break
			}
		}
		if hasItem == false {
			return "нет предмета в инвентаре: " + item
		}
		if other == p.Location.Other[item] {
			for _, v := range p.Inventory {
				if v == item {
					if v == "ключи" {
						if p.Location.LockedExists["улица"] == true {
							p.Location.LockedExists["улица"] = false
							return "дверь открыта"
						}
					}
				}
			}

		}

		return "нет к чему применить"
	}
	for _, room := range rooms {
		room.Commands["идти"] = goCommand
		room.Commands["взять"] = takeCommand
		room.Commands["надеть"] = takeInventory
		room.Commands["применить"] = useCommand
	}
}

func handleCommand(input string) string {
	input = strings.TrimSpace(input)
	if input == "" {
		return "Введите команду"
	}

	parts := strings.Split(input, " ")
	command := parts[0]
	args := parts[1:]

	if handler, ok := player.Location.Commands[command]; ok {
		return handler(player, args)
	}

	return "Неизвестная команда"
}
func initGame() {
	// Создаем комнаты

	room := &Room{
		Name:            "комната",
		Description:     "ты в своей комната, можно пройти - коридор",
		Items:           []string{"ключи", "конспекты"},
		ChairItems:      []string{"рюкзак"},
		Exits:           make(map[string]*Room),
		Commands:        make(map[string]func(*Player, []string) string),
		ExitDescription: "ты в своей комнате. можно пройти - коридор",
		LockedExists:    make(map[string]bool),
	}

	hall := &Room{
		Name:            "коридор",
		Description:     "ничего интересного. можно пройти - кухня, комната, улица",
		Items:           []string{},
		Exits:           make(map[string]*Room),
		Commands:        make(map[string]func(*Player, []string) string),
		Other:           make(map[string]string),
		ExitDescription: "ничего интересного. можно пройти - кухня, комната, улица",
		LockedExists:    make(map[string]bool),
	}
	kitchen := &Room{
		Name:            "Кухня",
		Description:     "Ты нахдишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор",
		Items:           []string{"чай"},
		Exits:           make(map[string]*Room),
		Commands:        map[string]func(*Player, []string) string{},
		ExitDescription: "кухня, ничего интересного. можно пройти - коридор",
		LockedExists:    make(map[string]bool),
	}

	street := &Room{
		Name:            "улица",
		Description:     "на улице весна. можно пройти - домой",
		Items:           []string{},
		Exits:           make(map[string]*Room),
		Commands:        map[string]func(*Player, []string) string{},
		Other:           make(map[string]string),
		ExitDescription: "на улице весна. можно пройти - домой",
		LockedExists:    make(map[string]bool),
	}

	// Настраиваем выходы
	room.Exits["коридор"] = hall
	hall.Exits["кухня"] = kitchen
	hall.Exits["комната"] = room
	hall.Exits["улица"] = street
	kitchen.Exits["коридор"] = hall

	hall.LockedExists["улица"] = true

	// Настраиваем команды для комнат
	room.Commands["осмотреться"] = func(p *Player, args []string) string {
		if len(room.Items) == 0 {
			return "Пустая комната. Можно пройти коридор"
		}
		return " на столе: " + strings.Join(room.Items, ", ") + " на стуле" + strings.Join(room.ChairItems, ", ") + "можно пройти в коридор"
	}

	hall.Commands["осмотреться"] = func(p *Player, args []string) string {
		return " На столе : " + strings.Join(hall.Items, ", ") + " можно пройти коридор"
	}

	kitchen.Commands["осмотреться"] = func(p *Player, args []string) string {
		return " ты находишься на кухне, на столе:" + strings.Join(kitchen.Items, ", ") + ", надо собрать рюкзак и идти в универ. можно пройти - коридор"
	}
	hall.Other["ключи"] = "дверь"
	hall.Other["телефон"] = "шкаф"
	street.Other["ключи"] = "дверь"
	rooms = map[string]*Room{
		"комната": room,
		"коридор": hall,
		"кухня":   kitchen,
		"улица":   street,
	}
	// Создаем игрока
	player = &Player{
		Location: kitchen,
		hasKey:   false,
	}

	// Добавляем глобальные команды
	initCommands()

}
func main() {
	initGame()
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		response := handleCommand(input)
		fmt.Println(response)
	}
}
