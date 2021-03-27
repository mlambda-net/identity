package repository



type IdentityCache interface {
  Get(key string, data interface{})
  Set(key string, value interface{})
}

