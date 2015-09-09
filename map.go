package main

import "github.com/jeffail/gabs"
import "fmt"
import "io/ioutil"
import "log"


type Map struct {
  width, height int
  tileWidth, tileHeight int
  data []int
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


  children, _ := jsonParsed.Search("layers", "data").Children()
  for _, child := range children {
      fmt.Println(child.Data())
  }
  return "value"
}
func main() {

  parsedMap := parseMapFile()


  // value == 10.0, ok == true
  fmt.Println(parsedMap)
}
