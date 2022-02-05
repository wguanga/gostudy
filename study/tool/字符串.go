package tool

import (
	"fmt"
	"strings"
)

func Sstring() {
	//字符串定义
	var ss1 string = "abcdefg"
	var ss2 = "abcdefg"
	ss3 := "abcdefg"
	//字符串连接
	ss4 := ss1 + ss2                             //方法1
	ss5 := fmt.Sprint("111%c 222%c", ss1, ss2)   //方法2
	ss6 := strings.Join([]string{ss1, ss2}, ",") //方法3
	longss := `
	line 1
	line 2
	line 3
	`
	fmt.Printf("ss1: %v\n", ss1)
	fmt.Printf("ss2: %v\n", ss2)
	fmt.Printf("ss3: %v\n", ss3)
	fmt.Printf("ss4: %v\n", ss4)
	fmt.Printf("ss5: %v\n", ss5)
	fmt.Printf("ss6: %v\n", ss6)
	fmt.Printf("longss: %v\n", longss)

	//字符串切片
	sl1 := "playground"
	ord := 3
	fmt.Printf("sl1[ord]: %c\n", sl1[ord])
	fmt.Printf("%v\n", sl1[0:4]) //前闭后开[,),此处等价于sl1[:4]
	fmt.Printf("len(sl1): %v\n", len(sl1))

	//字符串函数
	fmt.Printf("strings.Split(sl1, \" \"): %v\n", strings.Split(sl1, "g"))           //该函数返回一个数组
	fmt.Printf("strings.Contains(sl1, \"lay\"): %v\n", strings.Contains(sl1, "lay")) //该函数返回bool型
	fmt.Printf("strings.ToLower(sl1): %v\n", strings.ToLower(sl1))
	fmt.Printf("strings.ToUpper(sl1): %v\n", strings.ToUpper(sl1))
	fmt.Printf("strings.HasPrefix(sl1, \"play\"): %v\n", strings.HasPrefix(sl1, "play"))     //判开头
	fmt.Printf("strings.HasSuffix(sl1, \"Ground\"): %v\n", strings.HasSuffix(sl1, "Ground")) //判结尾
	fmt.Printf("strings.Index(sl1, \"oun\"): %v\n", strings.Index(sl1, "oun"))               //oun的第一个字母在原字符串的位置
	fmt.Printf("strings.Index(sl1, \"oo\"): %v\n", strings.Index(sl1, "oo"))
}
