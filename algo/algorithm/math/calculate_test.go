package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculate(t *testing.T) {
	ass:= assert.New(t)
	ass.Equal(10, calculate("2*3+4"))
	ass.Equal(0, calculate("0"))
	ass.Equal(5, calculate(" 3+5 / 2 "))
	ass.Equal(1, calculate(" 3/2 "))
	ass.Equal(1, calculate("1-1+1"))
}


func TestCalculate2(t *testing.T) {
	ass:= assert.New(t)
	ass.Equal(6, calculate3("2+4 "))
	ass.Equal(3, calculate3("2-1 + 2 "))
	ass.Equal(23, calculate3("(1+(4+5+2)-3)+(6+8)"))
	ass.Equal(-20, calculate3("(1-(4+5-2))-(6+8)"))
	ass.Equal(9, calculate3("(1-2+3-(4+5-2))+(6+8)"))
	ass.Equal(-15, calculate3("1-(3+5-2+(3+19-(3-1-4+(9-4-(4-(1+(3)-2)-5)+8-(3-5)-1)-4)-5)-4+3-9)-4-(3+2-5)-10"))
}
