package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var points int = 20

const pointsPerQuestion = 5

func main() {
	fmt.Println("Привіт, вітаємо у грі!")
	time.Sleep(1 * time.Second)

	for i := 5; i > 1; i-- {
		fmt.Printf("До початку: %v", i)
		time.Sleep(1 * time.Second)
	}

	myPoints := 0
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

	fmt.Printf("Вітаю, ти впорався за %v", timeSpent)
	time.Sleep(10 * time.Second)
}
