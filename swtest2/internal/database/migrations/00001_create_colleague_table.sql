-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
create table if not exists Colleague (
    id                  BIGSERIAL primary key,
    fio             varchar not null,
    loginpassword               varchar not null,
    status varchar not null
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
drop table if exists Colleague;