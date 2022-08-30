package main

type Student struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float64 `json:"score"`
}

type StudentSlice []Student

func (s StudentSlice) Len() int {
	return len(s)
}
func (s StudentSlice) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}
func (s StudentSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
