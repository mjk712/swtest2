-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
create table if not exists ClientVacations (
    id                  BIGSERIAL primary key,
    ClientId			int not null,
    VacationId			int not null,

    foreign key (VacationId) references Vacation(id),
    foreign key (ClientId) references Client(id)
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
drop table if exists ClientVacations;