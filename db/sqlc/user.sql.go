// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
insert into users (
    login,
    name,
    surname
) values (
    $1,
    $2,
    $3
) returning id, login, name, surname
`

type CreateUserParams struct {
	Login   string `json:"login"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Login, arg.Name, arg.Surname)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Login,
		&i.Name,
		&i.Surname,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
delete from users where users.id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getAll = `-- name: GetAll :many
select id, login, name, surname from users order by users.id limit $1 offset $2
`

type GetAllParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetAll(ctx context.Context, arg GetAllParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getAll, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Login,
			&i.Name,
			&i.Surname,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
select id, login, name, surname from users where users.id = $1 limit 1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Login,
		&i.Name,
		&i.Surname,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
update users set name = $1, surname = $2 where id = $3 returning id, login, name, surname
`

type UpdateUserParams struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	ID      int32  `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.Name, arg.Surname, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Login,
		&i.Name,
		&i.Surname,
	)
	return i, err
}
