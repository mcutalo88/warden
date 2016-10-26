package mongodb

import (
    mong "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    config "warden/config"
)

type User struct {
    Userid     string // fix this Im liking int better
    Username   string
    Warningnum int
    Banned     bool
    Bandate    string //This will need to be changed to time ojbect?
}

// Takes in a connection as a string
// Returns session
func Connect(connString string) *mong.Session {

    session, err := mong.Dial(connString)
    if err != nil {
        panic(err)
    }
    return session
    //This closes the session after surrounding function completes pretty neat
    //defer session.Close()
}

// This function checks if the user is in mongodb.
func IsUserInMongo(clientId string, userName string) User {

    var user = User{}
    session := Connect(config.GetIp())
    userCollect := session.DB(config.GetDatabase()).C("user")
    getUser := User{}
    userCollect.Find(bson.M{"userid":clientId}).One(&getUser)

    if((User{}) != getUser) {
        user = CheckToBan(getUser)
        userCollect.Update(bson.M{"userid":clientId},bson.M{"$set": bson.M{"warningnum":user.Warningnum,"banned":user.Banned}})
    } else {
        userCollect.Insert(&User{clientId,userName,0,false,""})
    }
    session.Close()
    return user
}

//This functions check if user should now be banned from text channels
func CheckToBan(user User) User {

    user.Warningnum = user.Warningnum + 1
    configWarnings:=config.GetNumberOfWarnings()

    if(user.Warningnum == configWarnings) {
        user.Banned = true
    }
    return user
}
