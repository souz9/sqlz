# sqlz

SQL query helper.

# Example

```go
type User struct {
	ID   int
	Name string
}

func QueryUsers(ctx context.Context, db sql.DB) ([]User, error) {
	var id int
	var name string
	var users []User
	err := sqlz.For(&id, &name).In(db.QueryContext(ctx, `
		SELECT id, name FROM user
	`)).EachRow(func() error {
		users = append(users, User{ID: id, Name: name})
		return nil
	})
	return users, err
}
```
