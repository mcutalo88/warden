package collections

import (
  // mgo "warden/mongodb"
)

type Channel struct {
  Id    string
  name  string
  Rules []Rule
}

func CreateChannel(c Channel) {

}
