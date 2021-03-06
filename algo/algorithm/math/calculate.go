package math

import (
	"strings"
)

/**
https://leetcode-cn.com/problems/basic-calculator-ii/
227. 基本计算器 II
*/
func calculate(s string) int {
	preChar := '+'
	s = strings.Replace(s," ","",-1)
	stack := make([]int, 0)
	num := 0
	for i, char := range s {
		isDigst :=  char >='0' && char <= '9'
		if isDigst {
			num = num*10+ int(char-'0')
		}
		if !isDigst || i == len(s)-1{
			switch preChar {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			num = 0
			preChar= char
		}
	}
	res := 0
	for _, i := range stack {
		res += i
	}
	return res
}

/**
https://leetcode-cn.com/problems/basic-calculator/
224. 基本计算器 给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
//(1-2-(4+5-2))-(6+8)
 */
func calculate3(s string) int {
	s = strings.Replace(s," ","",-1)
	num := 0
	stack := make([]int, 0)
	preFlag := '+'
	minusCount := 0 // 括号前减号的数量
	opStack := make([]int, 0) // 存储（前的符号，-为1，+为0
	for i, char := range  s {

		isDigit :=  char >= '0' && char <= '9'
		if isDigit {
			num = 10*num+int(char-'0')
		}
		if !isDigit || i==len(s)-1{
			pos := 1
			if minusCount %2 == 1 {
				pos = -1
			}
			if char == '(' {
				if preFlag == '+' {
					if num != 0 {
						stack = append(stack, pos*num)
					}
					opStack = append(opStack,0)
				} else if preFlag == '-' {
					if num != 0 {
						stack = append(stack, pos*(-num))
					}
					opStack = append(opStack,1)
				}

				if preFlag == '-' {
					minusCount++
				}
				preFlag = '+'
				continue
			}
			if char == ')' {
				if preFlag == '+' {
					stack = append(stack, pos*num)
				} else if preFlag == '-' {
					stack = append(stack, pos*(-num))
				}
				opNum := opStack[len(opStack)-1]
				opStack = opStack[:len(opStack)-1]
				// 减号时，minusCount--
				if opNum == 1 && minusCount > 0 {
					minusCount--
				}
				preFlag = char
				continue
			}

			if preFlag == '+' {
				stack = append(stack, pos*num)
			} else if preFlag == '-' {
				stack = append(stack, pos*(-num))
			}
			preFlag = char
			num = 0
		}
	}
	res := 0
	for _, i := range stack {
		res += i
	}
	return res
}

/**
https://leetcode-cn.com/problems/basic-calculator/
224. 基本计算器 给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
//(1-2-(4+5-2))-(6+8)
*/
func calculate3I(s string) int {
	sign :=1
	n := len(s)
	stack := []int{1}
	res := 0
	for i:=0;i<n; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = stack[len(stack)-1]
			i++
		case '-':
			sign = -stack[len(stack)-1]
			i++
		case '(':
			stack = append(stack,sign)
			i++
		case ')':
			stack = stack[:len(stack)-1]
			i++
		default:
			num :=0
			for ;i<n && s[i]>='0' && s[i]<='9';i++ {
				num = num*10+ int(s[i]-'0')
			}
			res += sign*num
		}
	}
	return res
}

/**
给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数。返回这个结果。
https://leetcode-cn.com/problems/add-digits/
 */
func addDigits(num int) int {
	for  num >= 10{
		res := 0
		for num > 0{
			res += num%10
			num/=10
		}
		num = res
	}
	return num
}