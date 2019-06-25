package jw

type Job struct {
	ID   int64
	Name string
	Args map[string]interface{}
	Res  map[string]interface{}
	Func func(map[string]interface{}) map[string]interface{}

}
