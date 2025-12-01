package utils

import "context"

type ctxKey string

const (
	CtxUserID ctxKey = "user_id"
	CtxRole   ctxKey = "role"
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
func RolesFromContext(ctx context.Context) (string, bool) {
	v := ctx.Value(CtxRole)
	if v == nil {
		return "", false
	}
	r, ok := v.(string)
	return r, ok
}
