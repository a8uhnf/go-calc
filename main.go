package main
import (
	"fmt"
	"strconv"
)
var arr [1002]int
var idx int
var postfix [1002]int
var infix [1002]int
var brakcetOpen int
var brakcetClose int
var plus int
var minus int
var divide int
var mul int
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
	fmt.Printf("%d\n",idx)
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
		if infix[i] >= 0{
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
		if postfix[i]>=0{
			Push(postfix[i])
		} else if postfix[i] == plus{
			a = arr[idx]
			idx--
			b = arr[idx]
			idx--
			Push(a + b)	
		} else if postfix[i] == minus{
			a = arr[idx]
			idx--
			b = arr[idx]
			idx--
			Push(a - b)	
		} else if postfix[i] == mul{
			a = arr[idx]
			idx--
			b = arr[idx]
			idx--
			Push(a * b)	
		} else if postfix[i] == divide{
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
//	fmt.Printf(ret);
//	fmt.Println();
	checkInfix(index)
	infixToPostFix(index)			
	return ret;
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
	var s string
	brakcetOpen = -1
	brakcetClose = -2
	plus = -3
	minus = -4
	divide = -5
	mul = -6
	arr[0]	= -10
	fmt.Scanf("%s",&s)
	Differentiate(s)
	
	
}