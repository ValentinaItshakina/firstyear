package main

import (
	"fmt"
)

func prints(folders []string) {
	fmt.Println("=== папки ===")
	for i := 0; i < len(folders); i++ {
		fmt.Println(i+1, folders[i])
	}
	fmt.Println("0. Выход")
}

func main() {
	s := []string{"Не удаляй"}
	var choice int
	var choice2 int
	var name string
mainLoop:
	for {
		prints(s)
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			return
		}
		switch {
		case choice == 0:
			break mainLoop
		case choice > 0 && choice <= len(s):
			fmt.Println("[%s]", s[choice-1])
			fmt.Println("Рано пока")
			fmt.Println(`
1. Создать папку
2. Удалить папку
3. Переименовать папку
4. Назад
0. Выход`)
			_, err := fmt.Scan(&choice2)
			if err != nil {
				fmt.Println("Ошибка ввода:", err)
				return
			}
			switch choice2 {
			case 1:
				fmt.Println("Имя:")
				_, err := fmt.Scan(&name)
				if err != nil {
					fmt.Println("Ошибка ввода:", err)
					return
				}
				s = append(s, name)
				fmt.Println("Готово")
			case 2:
				s = append(s[:choice-1], s[choice:]...)
				fmt.Println("Готово")
			case 3:
				fmt.Println("Имя:")
				_, err := fmt.Scan(&name)
				if err != nil {
					fmt.Println("Ошибка ввода:", err)
					return
				}
				s[choice-1] = name
			case 4:
				break
			case 0:
				break mainLoop
			default:
				print("Ты не туда жмал. Попробуй ещё раз")
			}
		default:
			print("Ты не туда жмал. Попробуй ещё раз")
		}
	}
}
