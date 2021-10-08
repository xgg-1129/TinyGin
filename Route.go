package TinyGin

import (
	"errors"
	"fmt"
	"strings"
)

type HandleFun func(ctx *HttpContext)

type Route struct {
	HandleFunsMap map[string]HandleFun
	Trie map[string]*Node
}

func  NewRoute() *Route{
	r:=&Route{HandleFunsMap: make(map[string]HandleFun)}
	r.Trie=make(map[string]*Node)
	return r
}
func ParsePath(path string) (parts []string){
	vs := strings.Split(path, "/")
	parts = make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return
}

func (r *Route) addRoute(method string,path string,fun HandleFun)error{
	parts:=ParsePath(path)
	key:=method+"-"+path
	if f:=r.HandleFunsMap[key];f!=nil{
		return errors.New("the route has been register")
	}
	if node := r.Trie[method];node ==nil{
		r.Trie[method]=&Node{}
	}
	r.Trie[method].Insert(path,parts,0)
	r.HandleFunsMap[key]=fun
	return nil
}
func (r *Route) GetRoute(method string,path string) *Node{
	root := r.Trie[method]
	if root == nil{
		return  nil
	}
	parts:=ParsePath(path)
	node:= root.Search(path, parts, 0)
	if node == nil{
		return nil
	}
	return node
}
func (r *Route) AddGet(path string,fun HandleFun) error{
	return r.addRoute("GET",path,fun)
}
func (r *Route) AddPost(path string,fun HandleFun) error{
	return r.addRoute("POST",path,fun)
}
func (r Route) handle(ctx *HttpContext) {
	routeNode:=r.GetRoute(ctx.method,ctx.path)
	if routeNode != nil{
		key:=ctx.method+"-"+routeNode.path
		r.HandleFunsMap[key](ctx)
	}else{
		ctx.W.WriteHeader(404)
		fmt.Fprintf(ctx.W,"404 not found")
	}
}





