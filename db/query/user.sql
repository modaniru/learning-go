-- name: CreateUser :one
insert into users (
    login,
    name,
    surname
) values (
    $1,
    $2,
    $3
) returning *;

-- name: GetUser :one
select * from users where users.id = $1 limit 1;

-- name: GetAll :many
select * from users order by users.id limit $1 offset $2;

-- name: UpdateUser :one
update users set name = $1, surname = $2 where id = $3 returning *;

-- name: DeleteUser :exec
delete from users where users.id = $1;