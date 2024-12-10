package main

import (
	"reflect"
	"testing"
)

// Тестовый пример 1: Стандартный случай
func TestGaleShapley_Standard(t *testing.T) {
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

	expected := map[string]string{
		"X": "B",
		"Y": "A",
		"Z": "C",
	}

	result := galeShapley(menPrefs, womenPrefs)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получили %v", expected, result)
	}
}

// Тестовый пример 2: Все участники имеют одинаковые предпочтения
func TestGaleShapley_UniformPreferences(t *testing.T) {
	menPrefs := Preferences{
		"A": {"X", "Y", "Z"},
		"B": {"X", "Y", "Z"},
		"C": {"X", "Y", "Z"},
	}

	womenPrefs := Preferences{
		"X": {"A", "B", "C"},
		"Y": {"A", "B", "C"},
		"Z": {"A", "B", "C"},
	}

	expected := map[string]string{
		"X": "A",
		"Y": "B",
		"Z": "C",
	}

	result := galeShapley(menPrefs, womenPrefs)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получили %v", expected, result)
	}
}

// Тестовый пример 3: Один мужчина, одна женщина
func TestGaleShapley_SinglePair(t *testing.T) {
	menPrefs := Preferences{
		"A": {"X"},
	}

	womenPrefs := Preferences{
		"X": {"A"},
	}

	expected := map[string]string{
		"X": "A",
	}

	result := galeShapley(menPrefs, womenPrefs)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получили %v", expected, result)
	}
}

// Тестовый пример 4: Мужчины и женщины без предпочтений
func TestGaleShapley_NoPreferences(t *testing.T) {
	menPrefs := Preferences{
		"A": {},
		"B": {},
		"C": {},
	}

	womenPrefs := Preferences{
		"X": {},
		"Y": {},
		"Z": {},
	}

	expected := map[string]string{}

	result := galeShapley(menPrefs, womenPrefs)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получили %v", expected, result)
	}
}

// Тестовый пример 5: Несбалансированные предпочтения (мужчин больше, чем женщин)
func TestGaleShapley_UnbalancedPreferences(t *testing.T) {
	menPrefs := Preferences{
		"A": {"X"},
		"B": {"X"},
		"C": {"X"},
	}

	womenPrefs := Preferences{
		"X": {"A", "B", "C"},
	}

	expected := map[string]string{
		"X": "A",
	}

	result := galeShapley(menPrefs, womenPrefs)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получили %v", expected, result)
	}
}
