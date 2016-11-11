package collections

type Channel struct {
  _Id   string
  Name  string
  Rules []Rule
}

type Rule struct {
  _Id        string
  Name       string
  Regex      string
  AllowImage bool
}

type User struct {
  UserId   string // fix this Im liking int better
  UserName string
  Warnings int
  Banned   bool
  BanDate  string //This will need to be changed to time ojbect?
}
