package action

import (
	"fmt"
	"reflect"
	"testing"
)
import "github.com/go-vgo/robotgo"

func Test_ac1(t *testing.T)  {
	x, y := robotgo.GetMousePos()
	fmt.Println("pos:", x, y)
	robotgo.MouseClick()
}

func Test_ac(t *testing.T)  {
	abool := robotgo.ShowAlert("test", "robotgo")
	if abool == 0 {
		fmt.Println("ok@@@", "ok")
	}

	x, y := robotgo.GetMousePos()
	fmt.Println("pos:", x, y)

	robotgo.MoveMouse(x, y)
	robotgo.MoveMouse(100, 200)

	robotgo.MouseToggle("up")

	for i := 0; i < 1080; i += 1000 {
		fmt.Println(i)
		robotgo.MoveMouse(800, i)
	}

	fmt.Println(robotgo.GetPixelColor(x, y))

	color := robotgo.GetPixelColor(100, 200)
	fmt.Println("color@@@", color)

	robotgo.TypeString("Hello World")
	// robotgo.KeyTap("a", "control")
	robotgo.KeyTap("f1", "control")
	// robotgo.KeyTap("enter")
	// robotgo.KeyToggle("enter", "down")
	robotgo.TypeString("en")

	abitmap := robotgo.CaptureScreen()
	fmt.Println("all...", abitmap)

	bitmap := robotgo.CaptureScreen(10, 20, 30, 40)
	fmt.Println("...", bitmap)

	fx, fy := robotgo.FindBitmap(bitmap)
	fmt.Println("FindBitmap------", fx, fy)

	robotgo.SaveBitmap(bitmap, "test.png", 1)

	var bitmapTest robotgo.Bitmap
	bitTest := robotgo.OpenBitmap("test.png")
	bitmapTest = robotgo.ToBitmap(bitTest)
	fmt.Println("...type", reflect.TypeOf(bitTest), reflect.TypeOf(bitmapTest))

	// robotgo.MouseClick()
	robotgo.ScrollMouse(10, "up")
}