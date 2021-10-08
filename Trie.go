package TinyGin

import "strings"

type Node struct {
	path string
	part string
	dim bool
	children []*Node
}

func (root *Node)Insert(path string,parts []string,height int) {
	length:=len(parts)
	if length == height{
		root.path=path
		return
	}
	part:=parts[height]
	child:=root.matchChild(part)
	if child == nil{
		child=&Node{
			part:    part,
			dim:      part[0]==':'||part[0]=='*',
		}
		root.children=append(root.children,child)
	}
	child.Insert(path,parts,height+1)
}
func (root *Node)Search(path string,parts []string,height int)*Node{
	length:=len(parts)
	if length == height || strings.HasPrefix(root.part,"*"){
		if root.path==""{
			return nil
		}
		return root
	}
	part:=parts[height]
	childs:=root.matchChildren(part)
	for _,child :=range childs{
		if search := child.Search(path, parts, height+1);search!=nil{
			return search
		}
	}
	return nil
}
func (root *Node) matchChild(part string)*Node{
	for _,child := range root.children{
		if child.part == part || child.dim{
			return child
		}
	}
	return nil
}

func (root *Node) matchChildren(part string) []*Node{
	res:=make([]*Node,0)
	for _,child := range root.children{
		if child.part==part||child.dim{
			res=append(res,child)
		}
	}
	return res
}