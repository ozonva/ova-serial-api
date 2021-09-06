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

COMMENT ON TABLE serial IS 'table with user serials';
COMMENT ON COLUMN serial.user_id IS 'owner id';
COMMENT ON COLUMN serial.title IS 'serial title';
COMMENT ON COLUMN serial.genre IS 'serial genre';
COMMENT ON COLUMN serial.year IS 'the year first series came out';
COMMENT ON COLUMN serial.seasons IS 'number of seasons';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table serial;
-- +goose StatementEnd
