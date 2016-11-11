package collections

import (
    "gopkg.in/mgo.v2/bson"
     mongo "warden/mongodb"
     mgo "gopkg.in/mgo.v2"
     config "warden/config"
)

// This function checks if the user is in mongodb.
func IsUserInMongo(clientId string, userName string) User {

    var user = User {
        UserId:   clientId,
        UserName: userName,
        Warnings: 0,
        Banned:   false,
        BanDate:  "",
    }

    change := mgo.Change {
        Update:    bson.M{"$inc": bson.M{"warnings": 1}, "$setOnInsert": bson.M{"username":userName, "banned":false, "bandate":""}},
        Upsert:    true,
        ReturnNew: true,
    }

    userCollect := mongo.DB.C("users")
    userCollect.Find(bson.M{"userid":clientId}).Apply(change, &user)

    return user
}

//This functions check if user should now be banned from text channels
func CheckToBan(user User) User {
    if (user.Warnings > config.Get().NumberOfWarnings) {
        user.Banned = true
    }
    return user
}
