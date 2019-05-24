package main

import "fmt"

func main() {
    // Структура 
	type Person struct {
        Name  string
        Likes []string // срез, что любим
    }
	
	// срез указателей на людей
    var people []*Person
	
	// заполняем срез указателей начальными значениями 
	people = append(people, &Person{"Roman1", []string{"Хлеб", "Мясо", "Кофе"}})
	people = append(people, &Person{"Roman2", []string{"Булка", "Мясо", "Мороженое"}})
	people = append(people, &Person{"Roman3", []string{"Хлеб", "Пельмени", "Кофе"}})
	people = append(people, &Person{"Roman4", []string{"Вода", "Мясо", "Кофе"}})
	people = append(people, &Person{"Roman5", []string{"Хлеб", "Яблоко", "Кофе"}})
	people = append(people, &Person{"Roman6", []string{"Хлеб", "Мясо", "Банан"}})
	people = append(people, &Person{"Roman7", []string{"Мир", "Мясо", "Кофе"}})

	// карта Строка - указатель
	likes := make(map[string][]*Person)
	
	// по срезу людей
	for _, p := range people {
		// по срезу любимых
		for _, l := range p.Likes {
			// создаем элемент карты или дописываем в него указатель на людя
			likes[l] = append(likes[l], p)
		}
	}
	
	// напечатаем ФИО из указателей
	for key, val := range likes {
		fmt.Printf("Key %v: Val:", key)
		
		for _, p := range val {
			fmt.Printf("[%v] ", p.Name)
		}
		
		fmt.Printf("\n")
	}
	
	fmt.Println(len(likes["bacon"]), "people like bacon.")
}
