package main

func main() {
  //parsedMap is a 2 dim map containing Tiles
  worldMap := parseMapFile("assets/map.json")
  startTile := worldMap.getStartTile()
  goalTile := worldMap.getGoalTile()
  if (startTile != nil && goalTile != nil) {
      astar(startTile, goalTile)
  }



  /*
  for _, row := range parsedMap {
    for _, tile := range row {
      fmt.Println(tile.X, tile.Y, TypeNames[tile.Type])
    }
  }*/
}
