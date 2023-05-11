package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
)

type animal struct{
	food, loco, noice, string
}

func (a animal) Eat(){
	fmt.Println(a.food)
}
func (a animal) Move(){
	fmt.Println(a.loco)
}
func (a animal) Speak(){
	fmt.Println(a.noice)
}

func main(){
	var ani animal
	for{
		fmt.Printf("> please enter request : name and info : ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		req := scanner.Text()
		name := strings.Split(req," ")[0]
		info := strings.Split(req," ")[1]

		switch name{
		case "cow":
			ani = animal{"grass","walk","moo"}
		case "bird":
			ani = animal{"worms","fly","peep"}
		case "snake":
			ani = animal{"mice","slither","hsss"}
		default:
			fmt.Printf("%s is not in cow,bird, snake",name)
			return
		}

		switch info{
		case "eat":
			ani.Eat()
		case "move":
			ani.Move()
		case "speak":
			ani.Speak()
		default:
			fmt.Printf("invalid choice")
			return
		}


	}
}