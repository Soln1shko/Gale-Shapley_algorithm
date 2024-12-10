package main

import (
	"fmt"
)

// Представление предпочтений
type Preferences map[string][]string

// Алгоритм Гэйла-Шепли
func galeShapley(menPrefs Preferences, womenPrefs Preferences) map[string]string {
	// Инициализация
	menFree := make(map[string]bool)
	womenPartner := make(map[string]string)
	menProposals := make(map[string]int)

	// Все мужчины свободны
	for man := range menPrefs {
		menFree[man] = true
		menProposals[man] = 0
	}

	// Пока есть свободный мужчина, который еще не сделал предложения всем женщинам
	for hasFreeMan(menFree) {
		var man string
		// Найти свободного мужчину
		for m := range menFree {
			if menFree[m] {
				man = m
				break
			}
		}

		// Получить предпочтение мужчины
		prefs := menPrefs[man]
		if menProposals[man] >= len(prefs) {
			// Мужчина сделал предложения всем женщинам
			delete(menFree, man)
			continue
		}

		// Выбрать следующую женщину в списке предпочтений
		woman := prefs[menProposals[man]]
		menProposals[man]++

		// Проверить, свободна ли женщина
		if _, exists := womenPartner[woman]; !exists {
			// Женщина свободна, пара создается
			womenPartner[woman] = man
			delete(menFree, man)
		} else {
			// Женщина уже занята, проверить предпочтения
			currentPartner := womenPartner[woman]
			if rank(womenPrefs[woman], man) < rank(womenPrefs[woman], currentPartner) {
				// Женщина предпочитает нового мужчину
				womenPartner[woman] = man
				delete(menFree, man)
				menFree[currentPartner] = true
			}
		}
	}

	return womenPartner
}

// Проверка, есть ли свободный мужчина
func hasFreeMan(menFree map[string]bool) bool {
	for _, free := range menFree {
		if free {
			return true
		}
	}
	return false
}

// Получить ранг участника в списке предпочтений
func rank(prefs []string, person string) int {
	for i, p := range prefs {
		if p == person {
			return i
		}
	}
	return len(prefs) // Если участник не найден
}

func main() {
	// Пример предпочтений мужчин и женщин
	menPrefs := Preferences{
		"A": {"Y", "X", "Z"},
		"B": {"X", "Y", "Z"},
		"C": {"Y", "Z", "X"},
	}

	womenPrefs := Preferences{
		"X": {"B", "A", "C"},
		"Y": {"A", "C", "B"},
		"Z": {"C", "B", "A"},
	}

	// Найти стабильные бракосочетания
	matches := galeShapley(menPrefs, womenPrefs)

	// Вывод результатов
	fmt.Println("Стабильные бракосочетания:")
	for woman, man := range matches {
		fmt.Printf("%s - %s\n", woman, man)
	}
}
