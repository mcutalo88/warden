package collections

import (
  "fmt"
  "strings"
  // mgo "warden/mongodb"
)

type Rule struct {
  _Id        string
  Name       string
  Regex      string
  AllowImage bool
}

func parseOp(rule string) string {
  fmt.print(strings.Split(rule,"-"))
}

func Execute(ruleCreate string) {
  idk := parseOp(ruleCreate)
}
