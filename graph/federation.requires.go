package graph

import (
	"context"
	"fmt"
)

// PopulateUserRequires is the requires populator for the User entity.
func (ec *executionContext) PopulateUserRequires(ctx context.Context, entity *graph.User, reps map[string]interface{}) error {
	panic(fmt.Errorf("not implemented: PopulateUserRequires"))
}
