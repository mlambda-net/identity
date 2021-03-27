package entity

import "github.com/google/uuid"

type Role struct {
  ID   uuid.UUID
  Name string
  App         App
  Description string
}
