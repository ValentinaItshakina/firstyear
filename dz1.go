package main

import (
	"errors"
	"fmt"
)

func prints(folders []string) {
	fmt.Println("=== папки ===")
	for i := 0; i < len(folders); i++ {
		fmt.Println(i+1, folders[i])
	}
	fmt.Println("0. Выход")
}

type Task struct {
	title string
	done  bool
}

func (t *Task) Complete() error {
	if t.done {
		return errors.New("уже выполнена")
	}
	t.done = true
	return nil
}

func main() {
	s := []string{"Не удаляй"}
	tasks := map[string][]Task{"Не удаляй": []Task{}}
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
			fmt.Println("=== Задачи ===")
			for i := 0; i < len(tasks[s[choice-1]]); i++ {
				fmt.Println(i+1, tasks[s[choice-1]][i].title)
			}
			fmt.Println(` Опции:
1. Создать папку
2. Удалить папку
3. Переименовать папку
4. Создать задачу в текущей папке
5. показать статус задач
6. Изменить статус задачи
7. Назад
0. Выход`)
			_, err := fmt.Scan(&choice2)
			if err != nil {
				fmt.Println("Ошибка ввода:", err)
				return
			}
			switch choice2 {
			case 1:
				fmt.Println("Имя (в одно слово):")
				_, err := fmt.Scan(&name)
				if err != nil {
					fmt.Println("Ошибка ввода:", err)
					return
				}
				s = append(s, name)
				tasks[name] = []Task{}
				fmt.Println("Готово")
			case 2:
				s = append(s[:choice-1], s[choice:]...)
				delete(tasks, s[choice-1])
				fmt.Println("Готово")
			case 3:
				fmt.Println("Имя (в одно слово):")
				_, err := fmt.Scan(&name)
				if err != nil {
					fmt.Println("Ошибка ввода:", err)
					return
				}
				cop := tasks[s[choice-1]]
				delete(tasks, s[choice-1])
				s[choice-1] = name
				tasks[name] = cop
			case 4:
				fmt.Println("Имя (в одно слово):")
				_, err := fmt.Scan(&name)
				if err != nil {
					fmt.Println("Ошибка ввода:", err)
					return
				}
				tasks[s[choice-1]] = append(tasks[s[choice-1]], Task{title: name})
			case 5:
				fmt.Println("=== Задачи ===")
				for i := 0; i < len(tasks[s[choice-1]]); i++ {
					fmt.Println(i+1, tasks[s[choice-1]][i].title, "-", tasks[s[choice-1]][i].done)
				}
			case 6:
				fmt.Println("Выберите задачу")
				_, err := fmt.Scan(&choice2)
				if err != nil {
					fmt.Println("Ошибка ввода:", err)
					return
				}
				err = tasks[s[choice-1]][choice2-1].Complete()
				if err != nil {
					fmt.Println(err)
				}
			case 7:
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
