-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
create table if not exists Companyclients (
    id                  BIGSERIAL primary key,
    CompanyId			int not null,
    ClientId			int not null,

    foreign key (CompanyId) references Company(id),
    foreign key (ClientId) references Client(id)
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
drop table if exists Companyclients;