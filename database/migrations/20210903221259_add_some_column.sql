-- +goose Up
-- +goose StatementBegin
CREATE table serial
(
    id      serial primary key,
    user_id int          not null,
    title   varchar(255) not null,
    genre   varchar(255) not null,
    year    int          not null,
    seasons int          not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table serial
-- +goose StatementEnd
