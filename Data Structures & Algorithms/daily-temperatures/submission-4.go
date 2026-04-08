func dailyTemperatures(temperatures []int) []int {
	var st = make([]int, len(temperatures))
	for i := 0; i < len(st); i++ {
		st[i] = 0
	}

	var sortedArr = []Obj{}

	for i := 0; i < len(temperatures); i++ {
		sortedArr = ReArrange(sortedArr, Obj{
			Value: temperatures[i],
			Index: i,
		}, st)
	}

	return st
}

type Obj struct {
	Value int
	Index int
}

func ReArrange(arr []Obj, value Obj, st []int) []Obj {
	if len(arr) == 0 {
		return append(arr, value)
	}
	position := FindPosition(arr, value.Value)

	if position == -1 {
		var j = len(arr) - 1
		for j >= 0 && arr[j].Value <= value.Value {
			if arr[j].Value == value.Value {
				j--
				continue
			}
			if st[arr[j].Index] == 0 {
				st[arr[j].Index] = value.Index - arr[j].Index
			}
			j--
		}
		return append([]Obj{value}, arr...)
	} else if position == len(arr) {
		var j = len(arr) - 1
		for j >= 0 && arr[j].Value <= value.Value {
			if arr[j].Value == value.Value {
				j--
				continue
			}
			if st[arr[j].Index] == 0 {
				st[arr[j].Index] = value.Index - arr[j].Index
			}			
			j--
		}
		return append(arr, value)
	}

	var j = position - 1
	for j >= 0 && arr[j].Value <= value.Value {
		if arr[j].Value == value.Value {
			j--
			continue
		}
		if st[arr[j].Index] == 0 {
				st[arr[j].Index] = value.Index - arr[j].Index
			}
		j--
	}

	return append(arr[:position], append([]Obj{value}, arr[position:]...)...)
}

func FindPosition(arr []Obj, value int) int {
	for i := 0; i < len(arr); i++ {
		if value < arr[i].Value {
			return i
		}
	}

	return len(arr)
}