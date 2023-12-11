package gin

import (
	"fmt"
	"testing"
)

func TestXq(t *testing.T) {
	p := cleanPath("./wuhu/haha/../../a")
	fmt.Println(p)
}

func TestMethodTree(t *testing.T) {
	fmt.Println(countParams("/sdfj:sldkfj"))
	fmt.Println(countParams("/sdfj:sldk*j"))

}

// 测试一下节点优先级
func TestNodePrio(t *testing.T) {
	nodes := node{
		children: []*node{
			{path: "/a", priority: 21},
			{path: "/b", priority: 1, indices: "ab"},
			{path: "/c", priority: 1, indices: "ac"},
		},
	}

	for _, child := range nodes.children {
		fmt.Printf("%d %s\n", child.priority, child.path)
	}

	fmt.Println("---------------------------")

	nodes.incrementChildPrio(2)

	for _, child := range nodes.children {
		fmt.Printf("%d %s\n", child.priority, child.path)
	}
}
