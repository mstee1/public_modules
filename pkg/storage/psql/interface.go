package psql

import "context"

type Requests interface {
	SelectData(ctx context.Context, query string) ([][]interface{}, error)
	ExecQuery(ctx context.Context, query string) error
}
