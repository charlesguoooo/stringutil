package stringutil

import "fmt"

// Stringer接口定义在fmt包中，该接口包含String()函数。任何类型只要定义了String()函数，进行Print输出时，就可以得到定制输出
// (可以说任何只要实现了String()方法的类型，都表示实现了fmt接口)
//type Person struct {
//	Name string
//	Age int
//}
//
//func (p Person) string() string {
//	return fmt.Sprintf("%v (%v yers)", p.Name, p.Age)
//
//}
//
//func SGER()  {
//	a := Person{"Arthur Dent", 42}
//	z := Person{"Zaphod Beeblebrox", 9001}
//	fmt.Println(a,z)
//}



type IPAddr [4]byte

func (i IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", i[0],i[1],i[2],i[3])
}
func SGER1() {
	hosts := map[string]IPAddr{
		"loopback": {127,0,0,1},
		"googleDNS": {8,8,8,8},
	}

	for name,ip := range hosts{
		fmt.Printf("%v: %v\n",name ,ip)
	}
}