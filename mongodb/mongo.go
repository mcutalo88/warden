package mongodb

import (
    mong "gopkg.in/mgo.v2"
    config "warden/config"
)

var DB *mong.Database

// Takes in a connection as a string
func Connect(connString string) {
  session, err := mong.Dial(connString)
  if err != nil {
      panic(err)
  }

  DB = session.DB(config.Get().Database)
}
