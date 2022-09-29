package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	current := l
	if current == nil {
		return current
	} else {
		for current != nil {
			index := current.Next
			for index != nil {
				if current.Data > index.Data {
					temp := current.Data
					current.Data = index.Data
					index.Data = temp
				}
				index = index.Next
			}
			current = current.Next
		}
	}
	return l
}
