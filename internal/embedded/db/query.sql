-- name: GetAuthor :one
select * 
from authors
where id = ? limit 1;

-- name: ListAuthors :many
select *
from authors
order by name;

-- name: CreateAuthor :one
insert into authors (
    name,
    bio
) values (
    ?,
    ?
)
returning *;

-- name: UpdateAuthor :exec
update authors
set name = ?,
    bio = ?
where id = ?;

-- name: DeleteAuthor :exec
delete from authors
where id = ?;
