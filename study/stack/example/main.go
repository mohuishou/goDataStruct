//中缀表达式转后缀表达式

package main

import (
	"fmt"
	"strconv"
	"strings"

	stack "github.com/mohuishou/goDataStruct/study/stack/example/stack"
)

//符号优先级别：
//栈内( :0
//+ - : 1
//* / : 2
//栈外() : 9
var order = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
	"(": 9,
	")": 9,
}

func main() {
	// var e string
	// fmt.Print("请输入一个表达式：")
	// fmt.Scanln(&e)
	// test(e)
	test("(2*(9+6/3-5)+4)")
}

func test(e string) {
	var result string

	s := stack.NewStack()
	expression := strings.Split(e, "")
	for i := range expression {
		if strings.TrimSpace(expression[i]) == "" {
			continue
		}

		//判断是否是数字
		_, err := strconv.Atoi(expression[i])
		if err == nil {
			result = result + expression[i]
			continue
		}

		result = result + compare(s, expression[i])

	}

	//循环结束，检查堆栈内是否还有运算符，有的话依次抛出
	for s.IsEmpty() == false {
		elem, _ := s.Pop()
		result = result + elem.Expression
	}

	fmt.Println(result)
}

func compare(s *stack.Node, e string) (result string) {
	//优先级判断
	lv := order[e]
	//如果堆栈为空或者是优先级较高则进行下面的操作
	if s.Next == nil || lv > s.Next.Data.Lv {

		//如果该运算符为右括号，那么将堆栈的运算符全部弹出，当出现左括号为止
		if e == ")" {
			for s.Next.Data.Expression != "(" {
				elem, err := s.Pop()
				if err != nil {
					panic(err)
				}
				result = result + elem.Expression
			}

			//弹出左括号
			_, _ = s.Pop()
			return result
		}

		//如果该运算符为左括号，入栈之后，左括号的优先级将至最低，为0
		if e == "(" {
			lv = 0
		}

		//运算符入栈
		s.Push(stack.Element{Expression: e, Lv: lv})

	} else if lv <= s.Next.Data.Lv {
		//当优先级小于等于栈顶运算符优先级时，弹出栈顶运算符，并把新运算符入栈
		elem, _ := s.Pop()
		result = result + elem.Expression

		if s.IsEmpty() == true || lv > s.Next.Data.Lv {
			return result
		}

		result = result + compare(s, e)

		// s.Push(stack.Element{Expression: e, Lv: lv})
	}
	return result

}
