package common

import "strconv"

func Filter[T any](slice []T, f func(T) bool) []T {
	var n []T
	for _, e := range slice {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}
func RemoveDuplication[T any](arr []T, f func(T) string) []T {
	set := make(map[string]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		var key = f(v)
		_, ok := set[key]
		if ok {
			continue
		}
		set[key] = struct{}{}
		arr[j] = v
		j++
	}

	return arr[:j]
}

func CalsVolumUnit(originalVolm float64) (volm string, unit string) {
	fVolm := originalVolm
	unit = ""
	if fVolm/1000 > 1 {
		unit = "k"
		fVolm = fVolm / 1000
	}
	if fVolm/1000 > 1 {
		unit = "m"
		fVolm = fVolm / 1000
	}
	if fVolm/1000 > 1 {
		unit = "b"
		fVolm = fVolm / 1000
	}
	return strconv.FormatFloat(fVolm, 'f', 2, 64), unit
}
