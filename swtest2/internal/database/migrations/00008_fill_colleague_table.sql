-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO Colleague(fio,loginpassword,status)
VALUES('Герасимов Анатолий Викторович' , 'gera12/1234', 'Открыт'),
('Пак Си Хон' , 'xiaomibest/redmi1234', 'Открыт'),
('Джугашвили Джузеппе Генадьевна' , 'lafrance/kruason18', 'Открыт'),
('Марьинская Мария Антонова' , 'mashkaXd18/lolkekxdxd', 'Открыт');
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
