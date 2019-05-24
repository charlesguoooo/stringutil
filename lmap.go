package stringutil

import (
	"fmt"
)

/*
定义 map
声明变量，默认 map 是 nil
var map_variable map[key_data_type]value_data_type

使用 make 函数
map_variable := make(map[key_data_type]value_data_type)
*/
type Lmap struct {
	Lat, Long float64
}

//var m map[string]Lmap



func LMAP()  {
	m := make(map[string]Lmap) // 定义了一个map,默认值为nil
	m["Bell Labs"] = Lmap{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"],m)
}