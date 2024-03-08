package maps

import (
	"fmt"
)

//type StringIntMap map[string]int
//
//func createHashMap() StringIntMap {
//
//	hashMap := make(map[string]int)
//	fmt.Println("New HashMap created!")
//	return hashMap
//}
//
//

type StringIntMap map[string]int

func CreateStringIntMap() StringIntMap {

	hashMap := make(map[string]int)

	//alternative way to declare and initialize map
	//hashMap2 := map[string]int{} //-> composite literal
	//return hashMap2

	return hashMap

}

func (m StringIntMap) AddToHashMap(key string, value int) bool {

	_, exists := m[key]

	if exists {
		fmt.Printf("key %s already exist in map!", key)
		return false
	} else {
		m[key] = value
		fmt.Printf("key %s added to map!\n", key)
		return true
	}

}

func (m StringIntMap) RemoveFromHashMap(key string) bool {

	_, exists := m[key]

	if exists {
		delete(m, key)
		fmt.Printf("key %s deleted from map!", key)
		return true
	} else {
		fmt.Printf("key %s do not exists in map!", key)
		return false
	}
}

func (m StringIntMap) PrintMap() {

	fmt.Println("Hash Map: ")
	for key, value := range m {
		fmt.Printf("key: [%s], value: [%d]\n", key, value)

	}

}
