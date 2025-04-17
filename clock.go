package main

import (
	"fmt"
	"os"
	"time"
)

type Timer interface {
	Start()
}

type CountdownTimer struct {
	duration time.Duration
	name     string
}

func (t *CountdownTimer) Start() {
	fmt.Printf("Старт таймера на %v\n", t.name, t.duration)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	endtime := time.Now().Add(t.duration)

	for {
		select {
		case <-ticker.C:
			remaining := time.Until(endtime).Round(time.Second)

			if remaining <= 0 {
				fmt.Printf("Час вийшов\n", t.name)
				return
			}
			fmt.Printf("[%s] Залишок часу: %v\n", t.name, remaining)
		}
	}

}

type Alarm struct {
	targetTime time.Time
	name       string
}

func (a *Alarm) Start() {
	fmt.Printf("Будильник встановлено на %v\n", a.name, a.targetTime.Format("2006-01-02 15:04:05"))

	for {
		now := time.Now()
		if now.After(a.targetTime) {
			fmt.Printf("Підйом%v\n", a.name)
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func ShowManu() {
	fmt.Println("Оберіть опцію")
	fmt.Println("1. Таймер зворотнього відліку")
	fmt.Println("2. Будильник")
	fmt.Println("3. Вийти")
}

func main() {
	for {
		ShowManu()

		var choice int
		fmt.Print("Ващ вибір (1-3)")
		_, err := fmt.Scanln(&choice)

		if err != nil {
			fmt.Println("Помилка вводу! Введіть число.")
			var clear string
			fmt.Scanln(&clear)
			continue
		}

		switch choice {
		case 1:

			fmt.Println("Введіть назву таймера")
			var name string
			fmt.Scanln(&name)

			fmt.Println("Тривалість в секундах")
			var seconds int
			fmt.Scanln(&seconds)

			timer := &CountdownTimer{
				duration: time.Duration(seconds) * time.Second,
				name:     name,
			}
			timer.Start()

		case 2:
			fmt.Println("Введіть назву будильника")
			var name string
			fmt.Scanln(&name)

			fmt.Println("Введіть час (години хвилини)")
			var hours, minutes int
			fmt.Scanln(&hours, &minutes)

			if err != nil {
				fmt.Println("Некоректний формат часу.")
				continue
			}

			now := time.Now()
			alarmTime := time.Date(
				now.Year(), now.Month(), now.Day(),
				hours, minutes, 0, 0, now.Location())

			if alarmTime.Before(now) {
				alarmTime = alarmTime.Add(24 * time.Hour)
			}

			alarm := &Alarm{
				targetTime: alarmTime,
				name:       name,
			}
			alarm.Start()

		case 3:
			fmt.Println("До побачення!")
			os.Exit(0)

		default:
			fmt.Println("Невірний вибір, спробуйте ще раз")
		}
	}
}
