package main

func main() {

	mapObject := createHashMap()
	mapObject.AddToHashMap("banana", 5)
	mapObject.PrintHashMap()
	mapObject.AddToHashMap("banana", 5)
	mapObject.AddToHashMap("strawbery", 6)
	mapObject.PrintHashMap()
	//// Tworzenie pustej hashmapy
	//hashMap := make(map[string]int)
	//
	//// Dodawanie elementów do hashmapy
	//hashMap["jabłka"] = 10
	//hashMap["banany"] = 5
	//hashMap["pomarańcze"] = 3
	//
	//// Wyświetlanie zawartości hashmapy
	//fmt.Println(hashMap)
	//
	//// Pobieranie wartości z hashmapy
	//iloscJablek := hashMap["jabłka"]
	//fmt.Println("Ilość jabłek:", iloscJablek)
	//
	//// Sprawdzanie czy klucz istnieje w hashmapie
	//_, istnieje := hashMap["banany"]
	//if istnieje {
	//	fmt.Println("Klucz 'banany' istnieje")
	//}
	//
	//// Usuwanie elementu z hashmapy
	//delete(hashMap, "pomarańcze")
	//fmt.Println(hashMap)
}
