-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
create table if not exists Company (
    id                  BIGSERIAL primary key,
    name             varchar not null,
    inn               varchar not null,
    legal_address varchar not null
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
drop table if exists Company;