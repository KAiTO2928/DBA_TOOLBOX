package Inspection

import (
	"DB_OSInspection/Global"
	"fmt"
	"strconv"

	"github.com/gookit/color"
)

var (
	// color
	Yellow     = color.Yellow.Render
	Cyan       = color.Cyan.Render
	LightGreen = color.Style{color.Green, color.OpBold}.Render
)

func Inspection() {
	fmt.Println("——————————————————提醒和建议[正在开发中...]—————————————————————")
	for k, v := range Global.Table_index_inspection_result {
		i, _ := strconv.Atoi(v)
		if i > 10 {
			fmt.Printf("——您的表:%s索引值太多——", LightGreen(k))

		}

	}
	fmt.Println(" \n ")
	fmt.Println("——————————————————提醒和建议（完）—————————————————————")
}
