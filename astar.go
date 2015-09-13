package main

import "fmt"



/* Node
  G = movement cost from current node to start point, following Pathp
  H = estimated movement cost from current node to goal
  F = G + H
*/
type Node struct {
  Tile *Tile
  Parent *Node
  G, H int
}
func (n Node) F() int {
  return n.G + n.H
}
func astar(start *Tile, goal *Tile) {
  openList := make(map[*Node]*Node)
  startNode := Node{Tile: start}
  goalNode := Node{Tile: goal}
  openList[&startNode] = &startNode // TODO: Not a good solution

  // TODO: Make container for open/closed list
  startNodeNeighbourTiles := startNode.Tile.getWalkableNeighbours()
  for _, neighbourTile := range startNodeNeighbourTiles {
    neighbourNode := Node{Tile: neighbourTile, Parent: &startNode}
    neighbourNode.G = neighbourNode.Tile.calcDistance(startNode.Tile)
    neighbourNode.H = neighbourNode.Tile.calcDistance(goalNode.Tile)
    fmt.Println(neighbourNode.Tile, goalNode.Tile,  neighbourNode)
    openList[&neighbourNode] = &neighbourNode // TODO: Not a good solution
  }



  fmt.Println(openList)

}
