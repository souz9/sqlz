package sqlz

type Rows interface {
	Next() bool
	Scan(...interface{}) error
	Err() error
	Close() error
}

type forFields []interface{}

type inRows struct {
	forFields
	rows Rows
	err  error
}

func For(fields ...interface{}) forFields {
	return fields
}

func (f forFields) In(rows Rows, err error) inRows {
	return inRows{f, rows, err}
}

func (r inRows) EachRow(f func() error) error {
	if r.err != nil {
		return r.err
	}
	return eachRow(r.rows, r.forFields, f)
}

func eachRow(rows Rows, fields []interface{}, f func() error) error {
	defer rows.Close()
	for rows.Next() {
		if len(fields) > 0 {
			if err := rows.Scan(fields...); err != nil {
				return err
			}
		}
		if err := f(); err != nil {
			return err
		}
	}
	return rows.Err()
}

func (r inRows) Scan() error { return r.EachRow(noop) }

func noop() error { return nil }
