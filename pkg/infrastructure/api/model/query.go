package model

type Select struct {
  Query string `json:"query"`
}

type Profile struct {
  Name string `json:"name"`
  Email string `json:"email"`
}
