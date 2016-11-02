package collections

import (
  // mgo "warden/mongodb"
)

type Channel struct {
  _Id   string
  Name  string
  Rules []Rule
}
