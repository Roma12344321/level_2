package main

import (
	"fmt"
	"os"
)

func task1() {
	client := &RealNTPClient{}
	currentTime, err := client.GetTime()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Ошибка получения времени:", err)
		os.Exit(1)
	}
	fmt.Println("Точное время:", currentTime)
}
