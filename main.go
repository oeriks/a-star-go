package main

import (
  "fmt"
)

func main() {

  parsedMap := parseMapFile()

  for _, row := range parsedMap {
    for _, tile := range row {
      fmt.Println(tile.X, tile.Y, TypeNames[tile.Type])
    }
  }

}
