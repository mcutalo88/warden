package collections

import (
 "fmt"
  // mgo "warden/mongodb"
)

type Rule struct {
  Regex      string
  allowImage bool
}

func (r Rule) Create() {
  fmt.Println("create from rule")
}

func (r Rule) Read() {

}

func (r Rule) Update() {

}

func (r Rule) Delete() {

}
