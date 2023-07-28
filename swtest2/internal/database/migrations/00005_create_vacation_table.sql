-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
create table if not exists Vacation (
    id                  BIGSERIAL primary key,
    start_date             date not null,
    end_date               date not null,
    total_cost 		    varchar not null,
    name               varchar not null,
    ColleagueId			int not null,

    foreign key (ColleagueId) references Colleague(id)
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
drop table if exists Vacation;