package main

import "github.com/jeffail/gabs"
import "fmt"
import "io/ioutil"
import "log"

// Two dimensional map, consisting of Tiles
type Map map[int]map[int]*Tile
type Tile struct {
  X, Y int
  M Map
}

func (m Map) addTile(t *Tile, x, y int) {
  m[x][y] = t
  t.X = x
  t.Y = y
  t.M = m
}

func calcTileCoords(index, mapWidth int) (int, int) {
  x := index / mapWidth
  y := index - x * mapWidth
  return x, y
}

func parseMapFile() string {
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
  for index, tileType := range tileTypes {
    x, y := calcTileCoords(index, mapWidth)
      fmt.Println(tileType.Data(), x, y)
  }
  return "value"
}

func parseMapData() {}

func main() {

  parsedMap := parseMapFile()


  // value == 10.0, ok == true
  fmt.Println(parsedMap)
}
