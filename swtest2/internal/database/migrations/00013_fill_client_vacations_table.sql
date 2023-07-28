-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO ClientVacations(ClientId,VacationId)
VALUES(1,3),
(2,6),
(3,2),
(4,1),
(5,7),
(6,4),
(7,8),
(8,5),
(9,2),
(10,10),
(11,15),
(12,9),
(13,5),
(14,4),
(15,7),
(16,11),
(17,6),
(18,14),
(19,1),
(20,12),
(21,3),
(22,2),
(23,7),
(24,4),
(25,13);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
