package utils

import (
  "testing"
)

func TestHelp(t *testing.T) {
  expected := CODE_BLOCK +
    "./warden \n" +
    "\t-h :\tprint help menu \n" +
    CODE_BLOCK
   actual := Help()

   if actual != expected {
      t.Errorf("Help(): expected %d, actual %d", expected, actual)
   }
}
