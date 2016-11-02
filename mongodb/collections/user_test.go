package collections

import (
    "testing"
)

func TestCheckToBan(t *testing.T) {

    t.Run("Banned user true", func(t *testing.T) {
        testUser:=User{Banned:true}
        user :=CheckToBan(testUser)
        if testUser.Banned != user.Banned {
            t.Error("Expected true , got ", user.Banned)
        }
    })

    t.Run("Banned user false", func(t *testing.T) {
        testUser:=User{}
        user :=CheckToBan(testUser)
        if testUser.Banned != user.Banned {
            t.Error("Expected false , got ", user.Banned)
        }
    })
}
