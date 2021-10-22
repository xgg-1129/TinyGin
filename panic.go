package TinyGin

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

//添加一个宕机恢复的中间件

func PanicRecover() HandleFun {
	return func(ctx *HttpContext) {
		defer func() {
			if err:=recover();err!=nil{
				//如果发生了宕机,恢复它并且答应出堆栈的信息，打印堆栈需要用系统调用trace函数【好像是叫这个】
				//打印panic信息
				log.Println(fmt.Sprintf("%s",err))
				//打印堆栈信息
				log.Println(trace())
				//关闭此http连接
				ctx.SendString(500,"server Error")
			}
		}()


		ctx.doAllNext()
	}
}

func trace() string{
	//strings.Builder 可以减少内存拷贝
	var str strings.Builder
	str.WriteString("TraceBack:")
	//最大堆栈追踪长度为32
	var pcs [32]uintptr

	//跳出 自己+trace+defer
	n := runtime.Callers(3, pcs[:])
	for _ ,pc := range pcs[0:n]{
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(fn.Entry())
		str.WriteString(fmt.Sprintf("\n\t[%s]:%s :%d",fn.Name(),file,line))
	}
	return str.String()
}
