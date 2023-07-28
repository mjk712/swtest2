-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO Client(fio,passport,email_telephone,loginpassword)
VALUES('Загонкин Евгений Крутович' ,'5612-654732','jekaflex@mail.ru/8-966-768-87-66', 'evgeha18/master13'),
('Александров Тимофей Дмитриевич' ,'4345-876293','timosha2014@yandex.ru/8-934-256-84-23', 'timosha14/timohabarboskin'),
('Дмитриев Матвей Михайлович' ,'5873-950393','matwei1best@gmail.com/8-978-098-34-54', 'matusha43/zxcghoul12'),
('Скворцова Николь Матвеевна' ,'5235-858353','nikolalko@rambler.ru/8-901-292-90-12', 'nikolaTesla/osumaster'),
('Орехова Дарья Ильинична' ,'4098-846293','dashkarar@mail.ru/8-983-567-01-65', 'dashathetravel/iloveparis'),
('Соколова Арина Матвеевна' ,'5783-524145','arinmat@gmail.ru/8-908-287-23-34', 'nearinamat/yoyo54'),
('Громов Вячеслав Львович' ,'5884-198345','gromviacheslav@mail.ru/8-955-432-56-22', 'gromisdom/waitwaitwait'),
('Кузнецова Арина Артёмовна' ,'5654-983850','kuznecova2003@yandex.ru/8-932-190-92-90', 'kuzneckiibank/69mark32'),
('Мельников Иван Русланович' ,'5621-769354','ivanmelnikov@yahoo.com/8-912-232-31-07', 'ivanushka984/qw12er34ty56'),
('Фомина Ксения Михайловна' ,'4430-423155','yeunasacoupe-2267@yopmail.com/8-308-658-96-28','fominaksu/JZ0zUkpd'),
('Короткова Аиша Кирилловна' ,'4420-948903','quammuddoutrippe-7836@yopmail.com/8-717-950-93-58','aishakruta2/l4xGeGrh'),
('Симонов Арсен Павлович' ,'4161-195446','zilannillaffe-8656@yopmail.com/8-190-912-77-67','arsenmarsen3/aUat7B30'),
('Максимов Дмитрий Дмитриевич' ,'5074-110335','grosauzahikau-2760@yopmail.com/8-002-849-54-94','dimo4ka132/Z16Ui3Wl'),
('Белякова Арина Демидовна' ,'4461-214211','q3eitbw4096k@gmail.com/8-473-824-22-55','beluga140/FCo5MmHJ'),
('Майорова Полина Захаровна' ,'4150-264985','umsa0te2h3nw@gmail.com/8-361-610-40-46','polyamashina/fJs2A72C'),
('Мартынова София Константиновна' ,'4062-952841','heyix70579@weizixu.com/8-732-395-77-81','martinaltin/h01vHIG5'),
('Мухина Мирослава Данииловна' ,'5524-341320','sogtov@hottempmail.cc/8-769-097-16-33','muhahaha12/zgFkBSRZ'),
('Алексеев Илья Даниилович' ,'4898-882777','iluaalex43@rambler.ru/8-410-583-71-77','ilushaden53/n975DE4F'),
('Федоров Максим Степанович' ,'4410-446869','maxfedstep13@main.go/8-556-642-74-14','fed2fed4fed6/GVqoDc0u'),
('Царева Есения Ильинична' ,'5883-642728','charwodvorcha87@gogogo.com/8-746-585-54-42','eseninjiv532/7L58hlN6'),
('Семенова София Глебовна' ,'4953-928932','semenbogomol@iche.com/8-692-488-04-21','sonka32gleb/90OzFkOM'),
('Григорьев Алексей Давидович' ,'5246-624732','daviddavit52@yandex.com/8-341-460-91-97','aleshagrishka4/fF782Hqd'),
('Горбунова Софья Матвеевна' ,'4637-498729','sonyaxo2023@elmeh.org/8-965-358-49-70','matvgorbun134/2gw04156'),
('Соловьева Екатерина Львовна' ,'5611-971586','katuhapruha1290@katua.com/8-502-269-68-44','katkalvovska32/49aKrVWq'),
('Григорьев Мирон Иванович' ,'4693-855737','mironmir666@dadada.ru/8-163-476-03-23','ivanmiron78/QGBwS6pX');
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
