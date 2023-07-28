-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO CompanyClients(CompanyId,ClientId)
VALUES(1,1),
(1,2),
(1,3),
(2,4),
(2,5),
(2,6),
(2,7),
(2,8),
(2,9);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
