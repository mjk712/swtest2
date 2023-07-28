-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
create table if not exists Application (
    id                  BIGSERIAL primary key,
    VacationId			int not null,
    status               		varchar not null,
    client_message 	varchar not null,    

    foreign key (VacationId) references Vacation(id)
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
drop table if exists Application;