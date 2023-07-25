package query

var AddColleague = `INSERT INTO Colleague(fio,loginpassword,status)
VALUES(:fio, :loginpassword, :status)`

var BlockColleague = `UPDATE Colleague SET status = 'Заблокирован'
WHERE id = $1`

var AddClient = `INSERT INTO Client(fio,passport,email_telephone,loginpassword)
VALUES(:fio,:passport,:email_telephone,:loginpassword)`

var ShowClientInfo = `SELECT c.fio AS client_full_name, c.passport AS client_passport,c.email_telephone AS client_email_telephone, v.name AS vacation_name
FROM client c
LEFT JOIN clientvacations cv ON cv.clientid = c.id
LEFT JOIN vacation v ON v.id = cv.vacationid
WHERE c.id = $1;`

var ChangeClientFio = `UPDATE Client SET fio = $1
WHERE id = $2`

var ChangeClientPassport = `UPDATE Client SET passport = $1
WHERE id = $2`

var ChangeClientEmailTelephone = `UPDATE Client SET email_telephone = $1
WHERE id = $2`

var ChangeClientLoginPassword = `UPDATE Client SET loginpassword = $1
WHERE id = $2`

var ShowClientsApplication = `SELECT c.fio AS client_full_name,a.status AS application_status, a.id AS application_id
FROM client c
LEFT JOIN clientvacations cv ON cv.clientid = c.id
LEFT JOIN application a ON a.id = cv.vacationid
WHERE a.id IS NOT NULL;`

var ChangeApplicationStatus = `UPDATE application SET status = $1
WHERE id = $2`

var InsertCompany = `INSERT INTO Company (name,inn,legal_address)
VALUES(:name,:inn,:legal_address)`

var InsertCompanyClients = `INSERT INTO CompanyClients(CompanyId,ClientId)
VALUES($1,$2)`

var ChangeCompanyName = `UPDATE Company SET name = $1
WHERE id = $2`

var ChangeCompanyInn = `UPDATE Company SET inn = $1
WHERE id = $2`

var ChangeCompanyLegalAddress = `UPDATE Company SET legal_address = $1
WHERE id = $2`
