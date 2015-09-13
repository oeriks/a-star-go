package main

import (
  "fmt"
)

func main() {
  //parsedMap is a 2 dim map containing Tiles
  parsedMap := parseMapFile("assets/map.json")

  tile := parsedMap[0][0]

  neighbours := tile.getWalkableNeighbours()
  fmt.Println(neighbours[0].Type)
  fmt.Println(parsedMap.getStartTile())
  fmt.Println(parsedMap.getGoalTile())

  /*
  for _, row := range parsedMap {
    for _, tile := range row {
      fmt.Println(tile.X, tile.Y, TypeNames[tile.Type])
    }
  }*/
}
