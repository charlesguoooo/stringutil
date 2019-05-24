package stringutil

import "fmt"

// map和struct相似，不过map需要key
type Lmap1 struct {
	Lat,Long float64
}

// 如果定期类型只是一个类型名，我们可以在文法中忽略
var m = map[string]Lmap1{
	"Bell Labs": Lmap1{40.68433,-74.39967},
	"Goole": {37.42202,-122.085008},
}

var s = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
}

var s1 = []string{"h","j","k","h","d"}
func ML()  {
	for _,value := range s1{
		fmt.Println(s[value]) //  如果v没有在map中，s[v]的值为初始值 0，到第二次再遇到h的时候，s[v]的值就已经变成1了
		s[value]=s[value]+1
	}
}