package main

type Datapoint struct {
	value     string
	timestamp int
}

type TimeMap struct {
	data map[string][]Datapoint
}

func Constructor2() TimeMap {
	return TimeMap{
		data: make(map[string][]Datapoint),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.data[key] = append(this.data[key], Datapoint{
		value:     value,
		timestamp: timestamp,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	items, ok := this.data[key]
	if !ok {
		return ""
	}
	left, right := 0, len(items)
	var mid int
	for left < right {
		mid = left + (right-left)/2
		if timestamp < items[mid].timestamp {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if left > 0 {
		return items[left-1].value
	}
	return ""
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
