package graph

import "github.com/sumit-dhull-97/assignment/auth/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Service service.Auth
}
