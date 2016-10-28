package mongodb

import (
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

// This function checks if the user is in mongodb.
func IsUserInMongo(clientId string, userName string) User {

    var user = User{}
    userCollect := Session.DB(config.GetDatabase()).C("user")
    getUser := User{}
    userCollect.Find(bson.M{"userid":clientId}).One(&getUser)

    if((User{}) != getUser) {
        user = CheckToBan(getUser)
        userCollect.Update(bson.M{"userid":clientId},bson.M{"$set": bson.M{"warningnum":user.Warningnum,"banned":user.Banned}})
    } else {
        userCollect.Insert(&User{clientId,userName,0,false,""})
    }
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

func UserOb() User {
    var user = User{}
    return user
}
