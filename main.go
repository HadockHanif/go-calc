package main
import (
	"fmt"
	"strconv"
)
var arr [1002]int
var idx int
var postfix [1002]int
var infix [1002]int
var preInfix [1002]int
var brakcetOpen int
var brakcetClose int
var plus int
var minus int
var divide int
var mul int
var mark [1002]int
func Push(a int){
	if idx == 1000 {
		return
	}
	idx++
	arr[idx] = a
}
func Pop(){
	if idx == 0 {
		return 
	}
	idx--
}

func checkPostfix(LL int){
	//fmt.Printf("%d\n",idx)
	for i := 0; i<LL; i++ {
		
		if postfix[i] == plus{
			fmt.Printf("+")
		}else if postfix[i] == divide{
			fmt.Printf("/")
		}else if postfix[i] == minus{
			fmt.Printf("-")
		}else if postfix[i] == mul{
			fmt.Printf("*")
		}else{
			fmt.Printf("%d",postfix[i])
		}
	}
	fmt.Println()
}

func infixToPostFix (LL int) {
	idx = 0
	var index int
	index = 0
	//fmt.Println(idx)	
	for i := 0; i<LL; i++ {
		if infix[i] > brakcetOpen{
			postfix[index] = infix[i];
			index++
		} else{
			if infix[i] == brakcetOpen {
				Push(brakcetOpen)
			} else if infix[i] == brakcetClose {

				for j := idx; j>0 && arr[idx] != brakcetOpen; j-- {
					postfix[index] = arr[idx]
					index++
					idx--
				}
				idx--
			} else if infix[i] == plus {
				Push(plus)
			} else if infix[i] == minus {
				Push(minus)
			} else if infix[i] == mul {
				Push(mul)
			} else if infix[i] == divide {
				Push(divide)
			}		
		}
		//fmt.Printf("yes %d %d\n",i+1,idx)
	}
	//fmt.Printf("%d\n",idx)
	//checkPostfix(index)
	result(index)
}

func result(LL int){
	var ret int
	idx = 0
	ret = 0
	var a int
	var b int
	//fmt.Printf("Yes\n")
	for i := 0; i<LL;i++{
		if postfix[i] > brakcetOpen {
			Push(postfix[i])
		} else if postfix[i] == plus {
			a = arr[idx]
			idx--
			b = arr[idx]
			idx--
			Push(a + b)	
		} else if postfix[i] == minus {
			a = arr[idx]
			idx--
			b = arr[idx]
			idx--
			Push(b-a)	
		} else if postfix[i] == mul {
			a = arr[idx]
			idx--
			b = arr[idx]
			idx--
			Push(a * b)	
		} else if postfix[i] == divide {
			a = arr[idx]
			idx--
			b = arr[idx]
			idx--
			Push(b / a)	
		}
	}
	fmt.Printf("ans: %d\n",arr[1])
}

func checkInfix(LL int){
	for i := 0; i<LL; i++ {
		if infix[i] == plus {
			fmt.Printf("+")
		} else if infix[i] == divide{
			fmt.Printf("/")
		} else if infix[i] == minus{
			fmt.Printf("-")
		} else if infix[i] == mul{
			fmt.Printf("*")
		} else if infix[i] == brakcetOpen{
			fmt.Printf("(")
		} else if infix[i] == brakcetClose{
			fmt.Printf(")")
		} else {
			fmt.Printf("%d",infix[i])
		}
	}
	fmt.Println()
}

func Differentiate(s string)string{
	ret := ""
	var index int
	index = 0
	LL := len(s)
	for i := 0; i<LL;i++ {
		if s[i] == ' ' || s[i] == '(' || s[i] == ')' || s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
			if len(ret) != 0 {

				ii, err := strconv.Atoi(ret)
				if err != nil {
					return "";
				}
				//fmt.Println("",ii)
				if infix[index-1] == minus {
					infix[index-1] = plus
					ii *= -1
				}
				infix[index] = ii
				index++
			}
			if s[i] == '+' 	{
				infix[index] = plus
				index++
			}else if s[i] == '-'{
				infix[index] = minus
				index++
			}else if s[i] == '*'{				
				infix[index] = mul
				index++
			}else if s[i] == '/'{
				infix[index] = divide
				index++
			}else if s[i] == '('{
				infix[index] = brakcetOpen
				index++
			}else if s[i] == ')'{
				infix[index] = brakcetClose
				index++
			}

			ret = ""
		}else{
			ret += string(s[i])
		}
	}
//	checkInfix(index)
	call(index)
	//fmt.Println("yes")
	//processInfix(index)
// 	infixToPostFix(index)			
	return ret;
}

func initialize(){
	brakcetOpen = -10001
	brakcetClose = -10002
	plus = -10003
	minus = -10004
	divide = -10005
	mul = -10006
	arr[0]	= -10000
}
func call(LL int){
	var dic bool
	dic = false
	index := 0
	//fmt.Println(index)
	for i := 0; i < LL; {
		if infix[i] == mul && infix[i-1] > brakcetOpen && infix[i+1] > brakcetOpen {
			preInfix[index-1] = infix[i-1] * infix[i+1]
			//index++
			dic = true
			for j := i+2; j<LL;j++{
				preInfix[index] = infix[j]
				index++
			}
			i = LL
			
		}else if infix[i] == divide && infix[i-1] > brakcetOpen && infix[i+1] > brakcetOpen {
			preInfix[index-1] = infix[i-1] / infix[i+1]
			//index++
			dic = true
			for j := i+2; j<LL;j++{
				preInfix[index] = infix[j]
				index++
			}
			i = LL
		}else{
			preInfix[index] = infix[i]
			index++
			i++
		}
	}
	if dic == true{
		copyToInfix(index)
		call(index)
	}else{
		processInfix(index)
	}
}
func processInfix(LL int){
	var index int
	index = 0
	//checkInfix(LL)
	for i := 0; i<LL; {
		if infix[i] < 0 {
			if infix[i] == mul && infix[i-1]>brakcetOpen && infix[i+1]>brakcetOpen && mark[i-1] == 0 {
				preInfix[index-1] = brakcetOpen
				preInfix[index] = infix[i-1]
				index++
				preInfix[index] = infix[i]
				index++
				preInfix[index] = infix[i+1]
				index++;
				preInfix[index] = brakcetClose
				index++	
				mark[i-1] = 1
				mark[i+1] = 1
				i += 2
			}else if infix[i] == divide && infix[i-1]>brakcetOpen && infix[i+1]>brakcetOpen && mark[i-1] == 0 {
				preInfix[index-1] = brakcetOpen
				preInfix[index] = infix[i-1]
				index++
				preInfix[index] = divide
				index++
				preInfix[index] = infix[i+1]
				index++;
				mark[i-1] = 1
				mark[i+1] = 1
				preInfix[index] = brakcetClose
				index++	
				i += 2
			}else if infix[i] == minus && infix[i-1]>brakcetOpen && infix[i+1]>brakcetOpen && mark[i-1] == 0 {
				preInfix[index-1] = brakcetOpen
				preInfix[index] = infix[i-1]
				index++
				preInfix[index] = infix[i]
				index++
				preInfix[index] = infix[i+1]
				index++;
				mark[i-1] = 1
				mark[i+1] = 1
				preInfix[index] = brakcetClose
				index++	
				i += 2
			}else{
				preInfix[index] = infix[i]
				index++
				i++
			}	
		}else {
			preInfix[index] = infix[i]
			index++
			i++
		}
	}
	copyToInfix(index)
	infixToPostFix(LL)
}

func copyToInfix(LL int){
	for i := 0; i<LL; i++ {
		infix[i] = preInfix[i]
	}
	//checkInfix(LL)
	//infixToPostFix(LL)
}
func StackCheck(){
	var ops string
	var val int
	for i := 1; i<=100; i++ {
		fmt.Scanf("%s",&ops)
		if ops == "POP"{
			Pop()	
		}else if ops == "PUSH"{
			fmt.Scanf("%d",&val)
			Push(val)
		}else{
			fmt.Println(arr[idx])
		}		
	}
}

func main() {
	idx = 0
	initialize()
//	fmt.Println(minus)
	var s string
	var tmp string
	tmp = "("
	fmt.Scanf("%s",&s)
	tmp += s
	tmp += ")"
//	fmt.Printf("%s\n",tmp)
	Differentiate(tmp)
	
	
}
