// 由res2go自动生成。
package main

import (
    _ "github.com/ying32/govcl/pkgs/winappres"
    "github.com/ying32/govcl/vcl"
    _ "net/http/pprof"
)

func main() {
    //go func() {
    //   http.ListenAndServe("localhost:6060", nil)
    //   fmt.Println("启动")
    //}()
   vcl.Application.SetFormScaled(true)
   vcl.Application.Initialize()
   vcl.Application.SetMainFormOnTaskBar(true)
    vcl.Application.CreateForm(&FMain)
   vcl.Application.Run()
}


