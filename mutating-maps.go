package stringutil

import "fmt"


// 插入或修改元素 m[key] = value
// 获取元素 a = m[key]
// 删元素 delete(m,key)
// 检查key是否存在 elem,ok = m[key]
// key在m中，ok为true，否则ok为false

//当读取的key不存在时，结果是元素类型的0值
func MM()  {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:",m["Answer"])

	m["Answer1"] = 48
	fmt.Println("The value:",m["Answer1"])

	delete(m,"Answer")
	fmt.Println("The value:",m["Answer"])

	// val1, isPresent = map1[key1]
	// isPresent 返回一个bool值，如果key1存在于map1，val1就是key1对应的value值，并且isPresent为true
	// 如果key1不存在，var1为空，并且isPresent会返回false
	s,ok := m["Answer1"]
	fmt.Println("The value:",s,"Present?",ok)

}
