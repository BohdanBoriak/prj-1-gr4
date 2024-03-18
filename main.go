package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"prj-1-gr4/domain"
	"sort"
	"strconv"
	"time"
)

var id uint64 = 1

const (
	pointsPerQuestion = 20
	totalPoints       = 20
)

func main() {
	fmt.Println("Привіт, вітаємо у грі!")
	time.Sleep(1 * time.Second)

	var users []domain.User
	users = append(users, domain.User{Id: 1, Name: "Vasyl", Time: 5 * time.Second})
	users = append(users, domain.User{Id: 2, Name: "Mykola", Time: 3 * time.Second})
	users = append(users, domain.User{Id: 3, Name: "Artem", Time: 10 * time.Second})

	sortAndSave(users)
	// for {
	// 	menu()
	// 	punct := ""
	// 	fmt.Scan(&punct)

	// 	switch punct {
	// 	case "1":
	// 		u := play()
	// 		users = append(users, u)
	// 	case "2":
	// 		for _, user := range users {
	// 			fmt.Printf("Name: %s, Time: %v\n", user.Name, user.Time)
	// 		}
	// 	case "3":
	// 		return
	// 	default:
	// 		fmt.Println("Зробіть правильний вибір")
	// 	}
	// }
}

func menu() {
	fmt.Println("1. Почати гру")
	fmt.Println("2. Переглянути рейтинг")
	fmt.Println("3. Вийти")
}

func play() domain.User {
	for i := 5; i > 1; i-- {
		fmt.Printf("До початку: %v\n", i)
		time.Sleep(1 * time.Second)
	}

	myPoints := 0
	points := totalPoints
	start := time.Now()
	for points > 0 {
		x, y := rand.Intn(100), rand.Intn(100)
		res := x + y
		fmt.Printf("%v + %v = ", x, y)

		ans := ""
		fmt.Scan(&ans)

		ansInt, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Println("Спробуй ще!")
		} else {
			if res == ansInt {
				myPoints += pointsPerQuestion
				points -= pointsPerQuestion
				fmt.Println("Правильно! Ура!")
				fmt.Printf("У тебе балів: %v, залишилось зібрати: %v\n", myPoints, points)
			} else {
				fmt.Println("Палюча поразка! Вчи математику!")
			}
		}
	}
	end := time.Now()
	timeSpent := end.Sub(start)

	fmt.Printf("Вітаю, ти впорався за %v\n", timeSpent)

	fmt.Println("Введіть своє ім'я:")

	name := ""
	fmt.Scan(&name)

	user := domain.User{
		Id:   id,
		Name: name,
		Time: timeSpent,
	}
	id++

	return user
}

func sortAndSave(users []domain.User) {
	sort.SliceStable(users, func(i, j int) bool {
		return users[i].Time < users[j].Time
	})

	file, err := os.OpenFile("users.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
	}(file)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(users)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
}
