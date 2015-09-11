package main

import (
  "github.com/jeffail/gabs"
  "fmt"
  "io/ioutil"
  "log"
)

const (
  TypeStart = iota
  TypeGoal
  TypeGround
  TypeBlock
)
var TypeNames = map[int]string{
  TypeStart: "Start",
  TypeGoal: "Goal",
  TypeGround: "Ground",
  TypeBlock: "Block",
}

// Two dimensional map, consisting of Tiles
type Map map[int]map[int]*Tile
type Tile struct {
  X, Y int
  Type int
}

func (m Map) addTile(t *Tile) {
  if m[t.X] == nil {
		m[t.X] = map[int]*Tile{}
	}
  m[t.X][t.Y] = t
}

func calcTileCoords(index, mapWidth int) (int, int) {
  x := index / mapWidth
  y := index - x * mapWidth
  return x, y
}

func parseMapFile() Map {
  mapData, err := ioutil.ReadFile("map.json")
  if err != nil {
    log.Fatal(err)
  }


  jsonParsed, err := gabs.ParseJSON(mapData)
  if err != nil {
    log.Fatal(err)
  }

  mapWidth := int(jsonParsed.Path("width").Data().(float64))
  // Only first layer from Tiled data is used, rest ignored
  firstMapLayer := jsonParsed.Search("layers", "data").Index(0)
  tileTypes, _ := firstMapLayer.Children()

  m := Map{}
  for index, tileType := range tileTypes {
    x, y := calcTileCoords(index, mapWidth)
    tile := Tile{x, y, TypeGround}
    m.addTile(&tile)
    _ = tileType
  }

  return m
}

func parseMapData() {}

func main() {

  parsedMap := parseMapFile()

  for _, row := range parsedMap {
    for _, tile := range row {
      fmt.Println(tile.X, tile.Y, TypeNames[tile.Type])
    }
  }

}
