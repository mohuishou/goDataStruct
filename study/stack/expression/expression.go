package expression

import (
	"strconv"
	"strings"
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

//Infix2Postfix 中缀表达式转后缀表达式
//返回格式化成功的表达式，以空格间隔
func Infix2Postfix(e string) (result string) {

	s := NewStack()
	expression := strings.Split(e, "")
	//标记位，标记两个数字之间是否有运算符间隔，没有运算符间隔则为一个连续的整数
	tag := 0
	for i := range expression {
		if strings.TrimSpace(expression[i]) == "" {
			continue
		}

		//判断是否是数字
		_, err := strconv.Atoi(expression[i])
		if err == nil {
			if tag == 0 {
				result = result + " "
			}
			result = result + expression[i]
			tag++
			continue
		}
		tag = 0

		//优先级比较
		result = result + compare(s, expression[i])

	}

	//循环结束，检查堆栈内是否还有运算符，有的话依次抛出
	for s.IsEmpty() == false {
		elem, _ := s.Pop()
		result = result + " " + elem.Expression
	}
	return result
}

//compare 优先级比较
func compare(s *Node, e string) (result string) {
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
		s.Push(Element{Expression: e, Lv: lv})

	} else if lv <= s.Next.Data.Lv {
		//当优先级小于等于栈顶运算符优先级时，弹出栈顶运算符
		elem, _ := s.Pop()
		result = result + elem.Expression

		//如果此时栈空或者优先级较高，将运算符入栈并且返回结果
		if s.IsEmpty() == true || lv > s.Next.Data.Lv {
			s.Push(Element{Expression: e, Lv: lv})
			return result
		}

		//优先级不比栈顶优先级高时，需要继续比较
		result = result + compare(s, e)

	}
	return result
}
