-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO Company(name,inn,legal_address)
VALUES('Aльфа' , '74893847', '039194, Новгородская область, город Павловский Посад, бульвар Балканская, 25'),
('Бета' , '97369237', '039194, 699064, Саратовская область, город Наро-Фоминск, ул. Сталина, 58');
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
