package main

import "fmt"


type Node struct {
  Facing *Node
  G, H int
}
func (n Node) F() int {
  return n.G + n.H
}
func astar() {

  start := Node{G: 2, H: 10}
  node1 := Node{}
  node1.Facing = &start
  fmt.Printf("Node F is facing ", node1.Facing)
}
