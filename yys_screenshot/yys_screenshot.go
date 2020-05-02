package yys_screenshot

import (
	"errors"
	win "github.com/lxn/win"
	"image"
	"syscall"
	"unsafe"
	"yys/getyyshwnd"
)

var (
	libUser32, _               = syscall.LoadLibrary("user32.dll")
	funcGetDesktopWindow, _    = syscall.GetProcAddress(syscall.Handle(libUser32), "GetDesktopWindow")
	funcEnumDisplayMonitors, _ = syscall.GetProcAddress(syscall.Handle(libUser32), "EnumDisplayMonitors")
	funcGetMonitorInfo, _      = syscall.GetProcAddress(syscall.Handle(libUser32), "GetMonitorInfoW")
	funcEnumDisplaySettings, _ = syscall.GetProcAddress(syscall.Handle(libUser32), "EnumDisplaySettingsW")
)

//得到一个单个句柄
//var YYS_HWND win.HWND
//得到一个句柄切片
//var list_hwnd Yys_windows_screenshot //接受 遍历的hwnd
//func init(){
//	//var none *uint16
//	var yname *uint16
//	yname= wd.StringToUTF16Ptr("阴阳师-网易游戏")
//	YYS_HWND =win.FindWindow(nil,yname) //获得指定窗口句柄
//}

type Yys_windows_screenshot struct {
	//all_hwnd []win.HWND
}

//获取单个窗口句柄
//func (h *Yys_windows_screenshot)Get_yys_hwnd()win.HWND{
//	var none [20]uint16
//	var yys *uint16
//	retname :=&none[0]//给空指针一个地址接受值
//	yys= wd.StringToUTF16Ptr("阴阳师-网易游戏")
//	YYS_HWND =win.FindWindow(nil,yys) //获得指定窗口句柄
//	me_win32.GetWindowText(YYS_HWND, retname,100)
//	return YYS_HWND
//}

//获取单个窗口句柄
//func (h *Yys_windows_screenshot)Get_yys_hdc()win.HWND{
//	var none [20]uint16
//	var yys *uint16
//	retname :=&none[0]//给空指针一个地址接受值
//	yys= wd.StringToUTF16Ptr("阴阳师-网易游戏")
//	YYS_HWND =win.get(nil,yys) //获得指定窗口句柄
//	me_win32.GetWindowText(YYS_HWND, retname,100)
//	return YYS_HWND
//}


//获取所有同名字窗口的HWND句柄
//func (h *Yys_windows_screenshot)Get_list_hwnd(){
//	//遍历所有顶级窗口,返回需求的句柄
//	me_win32.EnumWindows(wd.NewCallback(h.Get_yys_all_hwnd),0)
//	//fmt.Println(len(list_hwnd.all_hwnd))
//}

//活得句柄列表的一个回调函数
//func (h *Yys_windows_screenshot)Get_yys_all_hwnd(yys_HWND win.HWND)int{
//	var none [200]uint16
//	retname :=&none[0]
//	yys_window_name :="阴阳师-网易游戏"
//	me_win32.GetWindowText(yys_HWND, retname,100)
//	ret_yys_window_name:=win.UTF16PtrToString(retname)
//	if me_win32.IsWindowVisible(yys_HWND) && me_win32.IsWindow(yys_HWND)&&me_win32.IsWindow(yys_HWND)&&
//		ret_yys_window_name ==yys_window_name{
//		list_hwnd.all_hwnd = append(list_hwnd.all_hwnd,yys_HWND)
//		return 2
//	}
//	return 1
//}

func (h *Yys_windows_screenshot)CreateImage(rect image.Rectangle) (img *image.RGBA, e error) {
	//img = new(image.RGBA)
	e = errors.New("Cannot create image.RGBA")
	defer func() {
		err := recover()
		if err == nil {
			e = nil
		}
	}()
	// image.NewRGBA may panic if rect is too large.
	img = image.NewRGBA(rect)
	return img, e
}

//获取句柄窗口图像
func (h *Yys_windows_screenshot)YYS_Capture() (*image.RGBA, error) { //YYS_HWND win.HWND
	x, y, width, height :=0,0,1136,640
	rect := image.Rect(0, 0, width, height)
	img, err := h.CreateImage(rect)
	if err != nil {
		return nil, err
	}
	hd := getyyshwnd.YYSHWND{}
	hwnd := hd.Get_yys_hwnd()
	//fmt.Println("hwnd:",hwnd)
	if hwnd ==0{
		return nil,nil
	}
	hdc := win.GetDC(hwnd)
	if hdc == 0 {
		return nil, errors.New("GetDC failed")
	}
	defer win.ReleaseDC(hwnd, hdc)

	memory_device := win.CreateCompatibleDC(hdc)
	if memory_device == 0 {
		return nil, errors.New("CreateCompatibleDC failed")
	}
	defer win.DeleteDC(memory_device)

	bitmap := win.CreateCompatibleBitmap(hdc, int32(width), int32(height))
	if bitmap == 0 {
		return nil, errors.New("CreateCompatibleBitmap failed")
	}
	defer win.DeleteObject(win.HGDIOBJ(bitmap))

	var header win.BITMAPINFOHEADER
	header.BiSize = uint32(unsafe.Sizeof(header))
	header.BiPlanes = 1
	header.BiBitCount = 32
	header.BiWidth = int32(width)
	header.BiHeight = int32(-height)
	header.BiCompression = win.BI_RGB
	header.BiSizeImage = 0

	// GetDIBits balks at using Go memory on some systems. The MSDN example uses
	// GlobalAlloc, so we'll do that too. See:
	// https://docs.microsoft.com/en-gb/windows/desktop/gdi/capturing-an-image
	bitmapDataSize := uintptr(((int64(width)*int64(header.BiBitCount) + 31) / 32) * 4 * int64(height))
	hmem := win.GlobalAlloc(win.GMEM_MOVEABLE, bitmapDataSize)
	defer win.GlobalFree(hmem)
	memptr := win.GlobalLock(hmem)
	defer win.GlobalUnlock(hmem)

	old := win.SelectObject(memory_device, win.HGDIOBJ(bitmap))
	if old == 0 {
		return nil, errors.New("SelectObject failed")
	}
	defer win.DeleteObject(old)
	defer win.SelectObject(memory_device, old)
	//HFONT oldFont = (HFONT)SelectObject(m_HDC, font);
	//DeleteObject(font);
	//DeleteObject(oldFont);
	if !win.BitBlt(memory_device, 0, 0, int32(width), int32(height), hdc, int32(x), int32(y), win.SRCCOPY) {
		return nil, errors.New("BitBlt failed")
	}

	if win.GetDIBits(hdc, bitmap, 0, uint32(height), (*uint8)(memptr), (*win.BITMAPINFO)(unsafe.Pointer(&header)), win.DIB_RGB_COLORS) == 0 {
		return nil, errors.New("GetDIBits failed")
	}

	i := 0
	src := uintptr(memptr)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			v0 := *(*uint8)(unsafe.Pointer(src))
			v1 := *(*uint8)(unsafe.Pointer(src + 1))
			v2 := *(*uint8)(unsafe.Pointer(src + 2))

			// BGRA => RGBA, and set A to 255
			img.Pix[i], img.Pix[i+1], img.Pix[i+2], img.Pix[i+3] = v2, v1, v0, 255

			i += 4
			src += 4
		}
	}

	return img, nil
}
//获取句柄窗口图像
func (h *Yys_windows_screenshot)YYS_Capture_HWNDs(hwnd win.HWND) (*image.RGBA, error) { //YYS_HWND win.HWND
	x, y, width, height :=0,0,1138,640
	rect := image.Rect(0, 0, width, height)
	img, err := h.CreateImage(rect)
	if err != nil {
		return nil, err
	}

	if hwnd ==0{
		return nil,nil
	}
	hdc := win.GetDC(hwnd)
	if hdc == 0 {
		return nil, errors.New("GetDC failed")
	}
	defer win.ReleaseDC(hwnd, hdc)

	memory_device := win.CreateCompatibleDC(hdc)
	if memory_device == 0 {
		return nil, errors.New("CreateCompatibleDC failed")
	}
	defer win.DeleteDC(memory_device)

	bitmap := win.CreateCompatibleBitmap(hdc, int32(width), int32(height))
	if bitmap == 0 {
		return nil, errors.New("CreateCompatibleBitmap failed")
	}
	defer win.DeleteObject(win.HGDIOBJ(bitmap))

	var header win.BITMAPINFOHEADER
	header.BiSize = uint32(unsafe.Sizeof(header))
	header.BiPlanes = 1
	header.BiBitCount = 32
	header.BiWidth = int32(width)
	header.BiHeight = int32(-height)
	header.BiCompression = win.BI_RGB
	header.BiSizeImage = 0

	// GetDIBits balks at using Go memory on some systems. The MSDN example uses
	// GlobalAlloc, so we'll do that too. See:
	// https://docs.microsoft.com/en-gb/windows/desktop/gdi/capturing-an-image
	bitmapDataSize := uintptr(((int64(width)*int64(header.BiBitCount) + 31) / 32) * 4 * int64(height))
	hmem := win.GlobalAlloc(win.GMEM_MOVEABLE, bitmapDataSize)
	defer win.GlobalFree(hmem)
	memptr := win.GlobalLock(hmem)
	defer win.GlobalUnlock(hmem)

	old := win.SelectObject(memory_device, win.HGDIOBJ(bitmap))
	if old == 0 {
		return nil, errors.New("SelectObject failed")
	}
	defer win.SelectObject(memory_device, old)

	if !win.BitBlt(memory_device, 0, 0, int32(width), int32(height), hdc, int32(x), int32(y), win.SRCCOPY) {
		return nil, errors.New("BitBlt failed")
	}

	if win.GetDIBits(hdc, bitmap, 0, uint32(height), (*uint8)(memptr), (*win.BITMAPINFO)(unsafe.Pointer(&header)), win.DIB_RGB_COLORS) == 0 {
		return nil, errors.New("GetDIBits failed")
	}

	i := 0
	src := uintptr(memptr)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			v0 := *(*uint8)(unsafe.Pointer(src))
			v1 := *(*uint8)(unsafe.Pointer(src + 1))
			v2 := *(*uint8)(unsafe.Pointer(src + 2))

			// BGRA => RGBA, and set A to 255
			img.Pix[i], img.Pix[i+1], img.Pix[i+2], img.Pix[i+3] = v2, v1, v0, 255

			i += 4
			src += 4
		}
	}
	return img, nil
}


func (h *Yys_windows_screenshot)YYS_Capture_HWND(hwnd win.HWND) (*image.RGBA, error) { //YYS_HWND win.HWND
	x, y, width, height :=0,0,1138,640
	rect := image.Rect(0, 0, width, height)
	img, err := h.CreateImage(rect)
	if err != nil {
		return nil, err
	}

	if hwnd ==0{
		return nil,nil
	}
	hdc := win.GetDC(hwnd)
	//***************************
	if hdc == 0 {
		return nil, errors.New("GetDC failed")
	}
	defer win.ReleaseDC(hwnd, hdc)
	memory_device := win.CreateCompatibleDC(hdc)
	if memory_device == 0 {
		return nil, errors.New("CreateCompatibleDC failed")
	}
	defer win.DeleteDC(memory_device)


	bitmap := win.CreateCompatibleBitmap(hdc, int32(width), int32(height))
	if bitmap == 0 {
		return nil, errors.New("CreateCompatibleBitmap failed")
	}
	defer win.DeleteObject(win.HGDIOBJ(bitmap))


	var header win.BITMAPINFOHEADER
	header.BiSize = uint32(unsafe.Sizeof(header))
	header.BiPlanes = 1
	header.BiBitCount = 32
	header.BiWidth = int32(width)
	header.BiHeight = int32(-height)
	header.BiCompression = win.BI_RGB
	header.BiSizeImage = 0

	// GetDIBits balks at using Go memory on some systems. The MSDN example uses
	// GlobalAlloc, so we'll do that too. See:
	// https://docs.microsoft.com/en-gb/windows/desktop/gdi/capturing-an-image
	bitmapDataSize := uintptr(((int64(width)*int64(header.BiBitCount) + 31) / 32) * 4 * int64(height))
	hmem := win.GlobalAlloc(win.GMEM_MOVEABLE, bitmapDataSize)
	defer win.GlobalFree(hmem)
	memptr := win.GlobalLock(hmem)
	defer win.GlobalUnlock(hmem)


	old := win.SelectObject(memory_device, win.HGDIOBJ(bitmap))
	defer win.DeleteObject(old)
	if old == 0 {
		return nil, errors.New("SelectObject failed")
	}
	defer win.SelectObject(memory_device, old)

	if !win.BitBlt(memory_device, 0, 0, int32(width), int32(height), hdc, int32(x), int32(y), win.SRCCOPY) {
		return nil, errors.New("BitBlt failed")
	}

	if win.GetDIBits(hdc, bitmap, 0, uint32(height), (*uint8)(memptr), (*win.BITMAPINFO)(unsafe.Pointer(&header)), win.DIB_RGB_COLORS) == 0 {
		return nil, errors.New("GetDIBits failed")
	}

	i := 0
	src := uintptr(memptr)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			v0 := *(*uint8)(unsafe.Pointer(src))
			v1 := *(*uint8)(unsafe.Pointer(src + 1))
			v2 := *(*uint8)(unsafe.Pointer(src + 2))

			// BGRA => RGBA, and set A to 255
			img.Pix[i], img.Pix[i+1], img.Pix[i+2], img.Pix[i+3] = v2, v1, v0, 255

			i += 4
			src += 4
		}
	}
	imgcopy :=img
	return imgcopy, nil
}