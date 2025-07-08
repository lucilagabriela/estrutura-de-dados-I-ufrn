package main

import (
	"fmt"
)

type BSTNode struct {
	Key int // valor armazenado no nó
	left *BSTNode // filho esquerdo
	Right *BSTNode // filho direito
	Height int // altura do nó
	BalanceFactor int // fator de balanceamento (alt(direita) - alt(esquerda))
}

func (node *BSTNode) Add(key int) {
	if node == nil { return	}
	if key < node.Key {
		if node.left == nil {
			node.left = createNode(key)
		} else {
			node.left.Add(key)
		}
	} else {
		if node.Right == nil {
			node.Right = createNode(key)
		} else {
			node.Right.Add(key)
		}
	}
	node.UpdateProperties(); node.Rebalanced()
}

func main() {
	fmt.Println("Hello, World!")
}
