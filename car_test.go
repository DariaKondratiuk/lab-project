package main

import (
	"testing"
)

func TestIsOlderThan10Years(t *testing.T) {
	tests := []struct {
		name     string
		car      Car
		expected bool
	}{
		{"Новий автомобіль", Car{Year: 2025}, false},
		{"Автомобілю 10 років", Car{Year: 2015}, false},
		{"Авто старше 10 років", Car{Year: 2000}, true},
		{"Авто з майбутнього", Car{Year: 2100}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.car.IsOlderThan10Years()
			if got != tt.expected {
				t.Errorf("IsOlderThan10Years() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFindMostExpensiveCar(t *testing.T) {
	tests := []struct {
		name     string
		cars     []Car
		expected Car
	}{
		{"Найдорожче авто", []Car{{Model: "A", Price: 10000}, {Model: "B", Price: 20000}}, Car{Model: "B", Price: 20000}},
		{"Всі авто однакові за ціною", []Car{{Model: "A", Price: 5000}, {Model: "B", Price: 5000}}, Car{Model: "A", Price: 5000}},
		{"Порожній список", []Car{}, Car{}},
		{"Один автомобіль", []Car{{Model: "C", Price: 3000}}, Car{Model: "C", Price: 3000}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMostExpensiveCar(tt.cars)
			if got != tt.expected {
				t.Errorf("FindMostExpensiveCar() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCarsByCountry(t *testing.T) {
	brand1 := Brand{Name: "Toyota", Country: "Japan"}
	brand2 := Brand{Name: "BMW", Country: "Germany"}

	cars := []Car{
		{Brand: brand1, Model: "Camry", Price: 25000, Year: 2015},
		{Brand: brand2, Model: "M3", Price: 55000, Year: 2020},
	}

	carsByCountry := make(map[string][]Car)
	for _, car := range cars {
		carsByCountry[car.Brand.Country] = append(carsByCountry[car.Brand.Country], car)
	}

	// Тестуємо, чи правильно автомобілі групуються по країнах
	if len(carsByCountry["Japan"]) != 1 || carsByCountry["Japan"][0].Model != "Camry" {
		t.Errorf("Cars by country 'Japan' = %v, want Camry", carsByCountry["Japan"])
	}

	if len(carsByCountry["Germany"]) != 1 || carsByCountry["Germany"][0].Model != "M3" {
		t.Errorf("Cars by country 'Germany' = %v, want M3", carsByCountry["Germany"])
	}
}
