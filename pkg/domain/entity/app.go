package entity

import "github.com/google/uuid"

type App struct {
  ID   uuid.UUID
  Name string
  Description string
}
