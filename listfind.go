package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	for iter := l.Head; iter != nil; iter = iter.Next {
		if comp(iter.Data, ref) {
			return &iter.Data
		}
	}
	return nil
}
