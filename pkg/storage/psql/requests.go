package psql

import "context"

func (r *req) SelectData(
	ctx context.Context,
	query string,
) ([][]interface{}, error) {

	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result [][]interface{}
	col := len(rows.FieldDescriptions())

	for rows.Next() {
		line := make([]interface{}, col)
		for i := range line {
			var val interface{}
			line[i] = &val
		}

		if err := rows.Scan(line...); err != nil {
			return nil, err
		}

		for i, col := range line {
			line[i] = *col.(*interface{})
		}

		result = append(result, line)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return result, nil
}

func (r *req) ExecQuery(
	ctx context.Context,
	query string,
) error {

	if ctx.Err() != nil {
		return ctx.Err()
	}

	_, err := r.pool.Exec(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
