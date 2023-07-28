-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
create table if not exists Client (
    id                  BIGSERIAL primary key,
    fio             varchar not null,
    passport               varchar not null,
    email_telephone varchar not null,
    loginpassword               varchar not null
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
drop table if exists Client;