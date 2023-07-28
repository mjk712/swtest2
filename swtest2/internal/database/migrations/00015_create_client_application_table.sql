-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
create table if not exists ClientApplication (
    id                  BIGSERIAL primary key,
    ClientId			int not null,
    ApplicationId			int not null,

    foreign key (ApplicationId) references Application(id),
    foreign key (ClientId) references Client(id)
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
drop table if exists ClientApplication;