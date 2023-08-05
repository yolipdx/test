package main


func GetKeys[T comparable](dict map[T]any) []T {
	keys := []T{}
	for k := range(dict) {
		keys = append(keys, k)
	}
	return keys
}

func main() {

}