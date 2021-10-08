package TinyGin

import (
	"errors"
	"fmt"
)

type HandleFun func(ctx *HttpContext)

type Route struct {
	HandleFunsMap map[string]HandleFun
}

func  NewRoute() *Route{
	r:=&Route{HandleFunsMap: make(map[string]HandleFun)}
	return r
}

func (r *Route) addRoute(method string,path string,fun HandleFun)error{
		key:=method+"-"+path
		handleFun := r.HandleFunsMap[key]
		if handleFun!=nil{
			return errors.New("the route has been registered")
		}
		r.HandleFunsMap[key]=fun
		return nil
}
func (r *Route) AddGet(path string,fun HandleFun) error{
	return r.addRoute("GET",path,fun)
}
func (r *Route) AddPost(path string,fun HandleFun) error{
	return r.addRoute("POST",path,fun)
}
func (r Route) handle(ctx *HttpContext) {
	key:=ctx.method+"-"+ctx.path
	fun:=r.HandleFunsMap[key]
	if fun != nil{
		fun(ctx)
	}else{
		ctx.W.WriteHeader(404)
		fmt.Fprintf(ctx.W,"404 not found")
	}
}





