package mongodb

import (
    "gopkg.in/mgo.v2/bson"
     mgo "gopkg.in/mgo.v2"
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

    var user = User {
        Userid:     clientId,
        Username:   userName,
        Warningnum: 0,
        Banned:     false,
        Bandate:    "",
    }
    
    change := mgo.Change {
        Update: bson.M{"$inc": bson.M{"warningnum": 1}, "$setOnInsert": bson.M{"username":userName, "banned":false, "bandate":""}},
        Upsert:    true,
        ReturnNew: true,
    }

    userCollect := Session.DB(config.Get().Database).C("user")
    userCollect.Find(bson.M{"userid":clientId}).Apply(change, &user)

    return user
}

//This functions check if user should now be banned from text channels
func CheckToBan(user User) User {
    if (user.Warningnum > config.Get().NumberOfWarnings) {
        user.Banned = true
    }
    return user
}


// bson.M{
//     {"$match": bson.M{"clktime": bson.M{"gt": 1425289561}}},
//     {"$group": bson.M{"_id": bson.M{"$subtract": []interface{}{"$clktime", bson.M{"$mod": []interface{}{"$clktime", 60 * 5}}}}, "count": bson.M{"$sum": 1}}}
// }

// bson.M{
//     {"$setOnInsert": bson.M{"username": userName, "warningnum":0, "banned":false, "bandate":""}},
//     {"$inc": bson.M{"warningnum": 1}}
// },

//bson.M{"$inc": bson.M{"n": 1}, "$set": bson.M{"name": myStruct.Name}},
