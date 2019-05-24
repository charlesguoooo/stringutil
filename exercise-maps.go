package stringutil

import (
	"golang.org/x/tour/wc"
	"strings"
)

// strings.Fields(s) 按照空格来分割字符串，最终返回[]string的切片
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	// 创建一个key为string，value为int的map
	wordCountMap := make(map[string]int)

	// 取出words切片中的元素值
	// 好聪明啊
	for _,word := range words{
		//fmt.Println(word)
		// 如果v没有在map中，m[v]的值为初始值 0，当下一次再遇到同样的v的时候，m[v]的值就+1了
		wordCountMap[word]++
		//wordCountMap[word] = wordCountMap[word]+1
	}
	return wordCountMap
}

func EM() {
	wc.Test(WordCount)
}

