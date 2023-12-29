package apputils

import (
	"fmt"
	"testing"
)

func TestReadXlsx(t *testing.T) {
	file := "/Users/catwang/Downloads/liaobo.xlsx"
	ReadXlsx(file)
}

func TestRGB(t *testing.T) {
	l, a, b := RGB2Lab(43.45, 18.76, 23.2)
	fmt.Println(fmt.Sprintf("%f,%f,%f", l, a, b))
}

func TestRGB2Lab(t *testing.T) {
	RGB2XYZ(43.45, 18.76, 23.2)
}

func TestUrl(t *testing.T) {
	//str := Post()
	//GetFrom(str)
}
