package apputils

import (
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/tealeg/xlsx"
	"math"
)

func ReadXlsx(path string) {
	// Open the xlsx file
	xlFile, err := xlsx.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	//cols := []int{7,8,9}
	// Iterate through each sheet
	for _, sheet := range xlFile.Sheets {
		fmt.Printf("Sheet: %s\n", sheet.Name)

		// Iterate through each row in the sheet
		for rowIndex, row := range sheet.Rows {
			if rowIndex <= 2 {
				continue
			}
			fmt.Println(fmt.Sprintf("第%d行", rowIndex))
			fmt.Println(fmt.Sprintf("第%d行", len(row.Cells)))
			// Iterate through each cell in the row
			L := row.Cells[7].Value
			a := row.Cells[8].Value
			b := row.Cells[9].Value
			html := Post(L, a, b)
			x, y := GetFrom(html)
			columnIndex := 17
			if len(row.Cells) < columnIndex {
				// If the column doesn't exist, create it
				for j := len(row.Cells); j < columnIndex; j++ {
					row.AddCell()
				}
			}
			row.Cells[15].SetValue(x)
			row.Cells[16].SetValue(y)
		}
	}
	xlFile.Save("dalian_modified.xlsx")
}

// RGB2Lab 将RGB颜色转换为LAB颜色
func RGB2Lab(r, g, b float64) (float64, float64, float64) {

	c := colorful.Color{R: float64(r) / 255, G: float64(g) / 255, B: float64(b) / 255}
	l, a, b := c.Lab()
	return l, a, b
}

func RGB2XYZ(ri, gi, bi float64) {
	// Convert CIE RGB to XYZ
	var x, y, z float64

	// Normalize RGB values
	rn := ri / (ri + gi + bi)
	gn := gi / (ri + gi + bi)
	bn := bi / (ri + gi + bi)

	// Convert normalized RGB to linear RGB
	var r, g, b float64
	if rn <= 0.04045 {
		r = rn / 12.92
	} else {
		r = math.Pow((rn+0.055)/1.055, 2.4)
	}
	if gn <= 0.04045 {
		g = gn / 12.92
	} else {
		g = math.Pow((gn+0.055)/1.055, 2.4)
	}
	if bn <= 0.04045 {
		b = bn / 12.92
	} else {
		b = math.Pow((bn+0.055)/1.055, 2.4)
	}

	// Convert linear RGB to XYZ
	x = r*0.4124564 + g*0.3575761 + b*0.1804375
	y = r*0.2126729 + g*0.7151522 + b*0.0721750
	z = r*0.0193339 + g*0.1191920 + b*0.9503041

	// Output XYZ color coordinates
	fmt.Printf("XYZ: %v %v %v", x, y, z)
}
