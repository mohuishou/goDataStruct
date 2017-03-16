//中缀表达式转后缀表达式

package main

import (
	"fmt"

	"github.com/mohuishou/goDataStruct/study/stack/expression"
)

func main() {
	var e string
	fmt.Print("请输入一个表达式：")
	fmt.Scanln(&e)
	result := expression.Infix2Postfix(e)
	fmt.Println("中缀表达式为：", result)
}
