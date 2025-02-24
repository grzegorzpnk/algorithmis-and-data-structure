package binarySearch

type timedValue struct {
	timestamp int
	value     string
}

type TimeMap struct {
	//maps that contains [key]Leaf
	m map[string][]timedValue
}

func Constructor() TimeMap {
	return TimeMap{m: make(map[string][]timedValue)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], timedValue{value: value, timestamp: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	arr := this.m[key]

	if len(arr) == 0 {
		return ""
	}

	l, r := 0, len(arr)
	var minVal timedValue

	for l < r {
		mid := (l + r) / 2

		if arr[mid].timestamp == timestamp {
			return arr[mid].value
		} else if arr[mid].timestamp < timestamp {
			l = mid + 1
			minVal = arr[mid]
			continue
		} else {
			r = mid
		}
	}
	return minVal.value

}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
