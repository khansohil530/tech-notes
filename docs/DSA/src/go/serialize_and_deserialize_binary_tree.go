package main

import (
	"strconv"
	"strings"
)

type Codec struct {
	Delimiter string
}

func ConstructorBT() Codec {
	return Codec{Delimiter: "$"}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	res := make([]string, 0)
	stk := make([]*TreeNode, 0)
	stk = append(stk, root)
	for len(stk) > 0 {
		node := stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if node != nil {
			res = append(res, strconv.Itoa(node.Val))
			stk = append(stk, node.Right)
			stk = append(stk, node.Left)
		} else {
			res = append(res, "N")
		}
	}
	return strings.Join(res, this.Delimiter)
}

type Pair struct {
	node  *TreeNode
	state bool
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {

	if len(data) == 0 {
		return nil
	}

	res := strings.Split(data, this.Delimiter)
	val, _ := strconv.Atoi(res[0])
	root := &TreeNode{Val: val}
	stk := make([]Pair, 0)
	stk = append(stk, Pair{node: root, state: false})
	i := 1
	for len(stk) > 0 && i < len(res) {
		pair := stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		node, processed := pair.node, pair.state
		if !processed {
			if res[i] != "N" {
				val, _ = strconv.Atoi(res[i])
				node.Left = &TreeNode{Val: val}
				stk = append(stk, Pair{node: node, state: true})
				stk = append(stk, Pair{node: node.Left, state: false})
			} else {
				node.Left = nil
				stk = append(stk, Pair{node: node, state: true})
			}
			i++
		} else {
			if res[i] != "N" {
				val, _ = strconv.Atoi(res[i])
				node.Right = &TreeNode{Val: val}
				stk = append(stk, Pair{node: node.Right, state: false})
			} else {
				node.Right = nil
			}
			i++
		}
	}
	return root
}
