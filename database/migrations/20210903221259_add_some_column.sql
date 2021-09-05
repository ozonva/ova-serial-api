-- +goose Up
-- +goose StatementBegin
CREATE table serial
(
    id      serial primary key,
    user_id int          not null, -- owner id
    title   varchar(255) not null, -- serial title
    genre   varchar(255) not null, -- serial genre
    year    int          not null, -- the year first series came out
    seasons int          not null  -- number of seasons
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table serial;
-- +goose StatementEnd
