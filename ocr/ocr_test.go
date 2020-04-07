package ocr

import (
	"fmt"
	"testing"
)

func Test_run_ocr(t *testing.T){
	//fmt.Println(os.)
	img := "name3.png"
	s,_ :=GetText(img,"zh")
	fmt.Println(s)
}