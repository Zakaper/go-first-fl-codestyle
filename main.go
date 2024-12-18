package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func attack(charName, charClass string) string {
	if charClass == "Воитель" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randint(3, 5))
	}

	if charClass == "Маг" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randint(5, 10))
	}

	if charClass == "Лекарь" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randint(-3, -1))
	}
	return "неизвестный класс персонажа"
}

// обратите внимание на "if else" и на "else"
func defence(char_name, char_class string) string {
	if char_class == "Воитель" {
		return fmt.Sprintf("%s блокировал %d урона.", char_name, 10+randint(5, 10))
	}
	if char_class == "Маг" {
		return fmt.Sprintf("%s блокировал %d урона.", char_name, 10+randint(-2, 2))
	}
	if char_class == "Лекарь" {
		return fmt.Sprintf("%s блокировал %d урона.", char_name, 10+randint(2, 5))
	}
	return "неизвестный класс персонажа"

}

// обратите внимание на "if else" и на "else"
func special(charName, charClass string) string {
	if charClass == "Воитель" {
		return fmt.Sprintf("%s применил специальное умение `Выносливость %d`", charName, 80+25)
	}
	if charClass == "Маг" {
		return fmt.Sprintf("%s применил специальное умение `Атака %d`", charName, 5+40)
	}
	if charClass == "Лекарь" {
		return fmt.Sprintf("%s применил специальное умение `Защита %d`", charName, 10+30)
	}
	return "неизвестный класс персонажа"

}

// здесь обратите внимание на имена параметров
func start_training(char_name, char_class string) string {
	if char_class == "Воитель" {
		fmt.Printf("%s, ты %s - отличный боец ближнего боя.\n", char_name, char_class)
	}

	if char_class == "Маг" {
		fmt.Printf("%s, ты %s - превосходный укротитель стихий.\n", char_name, char_class)
	}

	if char_class == "Лекарь" {
		fmt.Printf("%s, ты %s - чародей, способный исцелять раны.\n", char_name, char_class)
	}

	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд: attack — чтобы атаковать противника,")
	fmt.Println("defence — чтобы блокировать атаку противника,")
	fmt.Println("special — чтобы использовать свою суперсилу.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var cmd string
	for cmd != "skip" {
		fmt.Print("Введи команду: ")
		fmt.Scanf("%s\n", &cmd)

		if cmd == "attack" {
			fmt.Println(attack(char_name, char_class))
		}

		if cmd == "defence" {
			fmt.Println(defence(char_name, char_class))
		}

		if cmd == "special" {
			fmt.Println(special(char_name, char_class))
		}
	}

	return "тренировка окончена"
}

// обратите внимание на имя функции и имена переменных
func choise_char_class(approve_choice, char_class string) string {
	for approve_choice != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель, Маг, Лекарь: ")
		fmt.Scanf("%s\n", &char_class)
		if char_class == "Воитель" {
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
		} else if char_class == "Маг" {
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
		} else if char_class == "Лекарь" {
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		}
		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanf("%s\n", &approve_choice)
		approve_choice = strings.ToLower(approve_choice)
	}
	return char_class
}

// обратите внимание на имена переменных
func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру назови себя: ")

	var char_name string
	fmt.Scanf("%s\n", &char_name)

	fmt.Printf("Здравствуй, %s\n", char_name)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	char_class := choise_char_class("Воин", " ")

	fmt.Println(start_training(char_name, char_class))
}

func randint(min, max int) int {
	return rand.Intn(max-min) + min
}
