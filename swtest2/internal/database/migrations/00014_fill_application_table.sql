-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO Application(VacationId,status,client_message)
VALUES(3,'В работе','Хочется трипл ипу и тульского пряника'),
(6,'Оплачена','Надеемся на целебный отпуск'),
(2,'Оплачена','Побольше моря хотелось бы'),
(1,'Новая','Хочется также затронуть Костромскую область'),
(7,'В работе','Экскурсовода женщину желательно'),
(4,'Согласована','А сахарная вата будет?'),
(8,'Оплачена','Добавим солевые ванны и прогулки под мостом'),
(5,'В работе','Побольше купаться'),
(2,'Оплачена','Нальчик Нальчик, парень работящий'),
(10,'В работе','Побольше Гастрономических мест'),
(15,'В работе','코카콜라 홀리데이 아이스크림'),
(9,'В работе','Симфоний любви, а также джазовых концертов'),
(5,'Согласована','Добавьте экскурсию в краевеческий музей Кубани'),
(4,'Оплачена','Также хочется посетить Липецкий Металлургический Завод'),
(7,'Новая','Добавьте 2й завтрак пожалуйста'),
(11,'В работе','Слишком долгая дорога, давайте вертолётом'),
(6,'В работе','Грязевые ванны желательны 2 раза в день'),
(14,'В работе','Добавьте экскурсию к дому пастуха'),
(1,'Новая','Интересны подробности третьего дня'),
(12,'Оплачена','Нарзанный поход включён?'),
(3,'В работе','Мы едем с детьми, можно исключить пивные экскурсии?'),
(2,'В работе','Аэропортнаяя ночь в нальчике обязательна?'),
(7,'Согласована','Прошу добавить экскурсию на Мурманский Ликёроводочный завод'),
(4,'Новая','А почему из программы убрали посещение музея матрацев'),
(13,'В работе','Побольше дегустаций чая улун и дайвинга');
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
