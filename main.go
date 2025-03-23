package main

import (
	"fmt"
	"time"
)

// Структура Brand
type Brand struct {
	Name    string
	Country string
}

// Структура Car
type Car struct {
	Brand Brand
	Model string
	Price float64
	Year  int
}

// Метод для перевірки, чи авто старше 10 років
func (c Car) IsOlderThan10Years() bool {
	currentYear := time.Now().Year()
	return currentYear-c.Year > 10
}

// Функція для пошуку найдорожчого автомобіля
func FindMostExpensiveCar(cars []Car) Car {
	var mostExpensiveCar Car
	for _, car := range cars {
		if car.Price > mostExpensiveCar.Price {
			mostExpensiveCar = car
		}
	}
	return mostExpensiveCar
}

func main() {
	// Створення брендів
	brand1 := Brand{Name: "Toyota", Country: "Japan"}
	brand2 := Brand{Name: "BMW", Country: "Germany"}
	brand3 := Brand{Name: "Audi", Country: "Germany"}

	// Створення автомобілів
	car1 := Car{Brand: brand1, Model: "Camry", Price: 25000, Year: 2015}
	car2 := Car{Brand: brand2, Model: "M3", Price: 55000, Year: 2020}
	car3 := Car{Brand: brand1, Model: "Corolla", Price: 18000, Year: 2010}
	car4 := Car{Brand: brand3, Model: "A4", Price: 40000, Year: 2018}
	car5 := Car{Brand: brand3, Model: "Q5", Price: 60000, Year: 2022}

	// Масив автомобілів
	cars := []Car{car1, car2, car3, car4, car5}

	// Пошук найдорожчого автомобіля
	mostExpensiveCar := FindMostExpensiveCar(cars)
	fmt.Println("Найдорожчий автомобіль:", mostExpensiveCar.Model, "-", mostExpensiveCar.Price)

	// Перевірка на вік авто
	for _, car := range cars {
		if car.IsOlderThan10Years() {
			fmt.Println(car.Model, "старше 10 років")
		}
	}

	// Використання map для зберігання автомобілів по країнах
	carsByCountry := make(map[string][]Car)
	for _, car := range cars {
		carsByCountry[car.Brand.Country] = append(carsByCountry[car.Brand.Country], car)
	}

	// Виведення автомобілів по країнах
	fmt.Println("Автомобілі по країнах виробників:")
	for country, cars := range carsByCountry {
		fmt.Println(country)
		for _, car := range cars {
			fmt.Println("  -", car.Model)
		}
	}
}
