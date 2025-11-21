package utils

import "context"

type ctxKey string

const (
	CtxUserID ctxKey = "user_id"
	CtxRoles  ctxKey = "roles"
)

// FromContext helpers
func UserIDFromContext(ctx context.Context) (string, bool) {
	v := ctx.Value(CtxUserID)
	if v == nil {
		return "", false
	}
	id, ok := v.(string)
	return id, ok
}
func RolesFromContext(ctx context.Context) ([]string, bool) {
	v := ctx.Value(CtxRoles)
	if v == nil {
		return nil, false
	}
	r, ok := v.([]string)
	return r, ok
}
