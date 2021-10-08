package TinyGin

type Group struct {
	prefix string

	//通过Server访问路由器
	Server *Server
}

func (g *Group) AddGet(path string ,fun HandleFun)error{
	return g.Server.AddGet(g.prefix+path,fun)
}
func (g *Group) AddPost(path string,fun HandleFun)error{
	return g.Server.AddPost(g.prefix+path,fun)
}

