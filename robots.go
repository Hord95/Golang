package robots

import (
	"fmt"
	"math/rand/v2"
)

type Robot struct {
	Name     string
	Capacity int
	Hp       float32
	Power    float32
}

func newRobot(name string, capacity int, hp float32, power float32) *Robot {
	return &Robot{
		Name:     name,
		Capacity: capacity,
		Hp:       hp,
		Power:    power,
	}
}
func Fight(r1 *Robot, r2 *Robot) string {

	for r1.IsAlive() && r2.IsAlive() {
		if r1.Capacity >= 0 {
			r1.Throw(r1)
		} else {
			r1.Recharge()
		}
		if r2.Capacity >= 0 {
			r2.Throw(r2)
		} else {
			r2.Recharge()
		}

	}

	if !r1.IsAlive() {
		return r2.Name
	} else {
		return r1.Name
	}
}
func (r *Robot) Throw(enemy *Robot) {

	if rand.IntN(2) == 1 {
		enemy.takeDamage(r.Power)
		r.Capacity -= 1

	}

}
func (r *Robot) Recharge() {
	if rand.IntN(2) == 1 {
		r.Capacity = 1
	}
}

func (r *Robot) takeDamage(damage float32) {
	r.Hp -= damage
}

func (r *Robot) IsAlive() bool {
	return r.Hp > 0
}

func main() {
	robot1 := newRobot("Avtobot", 10, 100, 5)
	robot2 := newRobot("Alesha", 10, 100, 5)
	fmt.Printf("Победитель - %s", Fight(robot1, robot2))

}
