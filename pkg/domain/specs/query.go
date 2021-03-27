package specs

import (
  "github.com/mlambda-net/net/pkg/spec"
)

func ById(id string) spec.Expression {
	return spec.NewEval("id", id, "=")
}

func ByName(name string, lastName string) spec.Expression {
	return spec.And(
    spec.NewEval("name", name, "="),
    spec.NewEval("last_name", lastName, "="))
}

func ByEmail(email string) spec.Expression {
	return spec.NewEval("email", email, "=")
}

func ByRole(filter string) spec.Expression {
  if filter == "" {
    return spec.Empty()
  }
  return spec.Or(spec.NewEval("r.name", filter, "="), spec.NewEval("r.description", filter, "="))
}

func ByRoleId(id string) spec.Expression  {
  return spec.NewEval("r.id", id, "=")
}

func ByUser(filter string) spec.Expression {
  if filter == "" {
    return spec.Empty()
  }
  return spec.Or(
    spec.Or(spec.NewEval("name", filter, "like"), spec.NewEval("last_name", filter, "like")),
    spec.NewEval("email", filter, "like"))
}
