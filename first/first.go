//go module
/* go mod init 项目名 -----用来链接模块
   go mod tidy ------------检测链接
   go mod download --------下载依赖模块
   go mod vender -----------将依赖模块导入项目*/
//go install 项目名------生成可执行文件

//go build 文件名--------生成可执行文件
//pkgm 快速生成主函数
package main

import (
	"first/morestring"
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func main() {
	//fp----生成输出函数
	fmt.Println("hello,world")
	//ff----生成格式化输出函数
	name := "abc"
	fmt.Printf("name=%v", name)
	//name.print
	fmt.Printf("name: %v\n", name)
	//for
	count := 10
	for i := 0; i < count; i++ {
		fmt.Printf("i= %v\n", i)
	}

	s := morestring.ReverseRunes("test")
	fmt.Printf("s: %v\n", s) //逆置test

	//字符串比较
	fmt.Println(cmp.Diff("hello,world", "hello,go"))

}
