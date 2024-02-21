package maps

import "fmt"

type StringIntMap map[string]int

func createHashMap() StringIntMap {

	hashMap := make(map[string]int)
	fmt.Println("New HashMap created!")
	return hashMap
}

func (m StringIntMap) AddToHashMap(key string, value int) bool {

	_, exists := m[key]
	if exists {
		fmt.Printf("Object: %v already exist in HashMap!\n", key)
		return false
	} else {
		m[key] = value
		fmt.Printf("Object: %v added to HashMap!\n", key)
		return true
	}
}

func (m StringIntMap) RemoveFromHashMap(key string) bool {

	_, exists := m[key]
	if exists {
		fmt.Printf("Object: %v not exists in HashMap!\n", key)
		return false
	} else {
		delete(m, key)
		fmt.Printf("Object: %v removed from HashMap!\n", key)
		return true
	}
}

func (m StringIntMap) PrintHashMap() {
	fmt.Println("HashMap: ")

	for key, value := range m {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}

}
