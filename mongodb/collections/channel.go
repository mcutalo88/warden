package collections

import (
  "fmt"
  // mgo "warden/mongodb"
)

type Channel struct {
  Id    string
  name  string
  rules []Rule
}

const COLLECTION string = "channel"

/**
 * Atomic Creation
 */

func (c Channel) Create() {
  fmt.Print(c)
  fmt.Println("create from channel")
}

func (c Channel) Read() {

}

func (c Channel) Update() {

}

func (c Channel) Delete() {

}
