package mongodb

import (
    mong "gopkg.in/mgo.v2"
)

var Session *mong.Session

// Takes in a connection as a string
func Connect(connString string) {

        session, err := mong.Dial(connString)
        if err != nil {
            panic(err)
        }

        Session = session
}

func getSession() *mong.Session {
    return Session
}
