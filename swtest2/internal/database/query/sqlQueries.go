package query

var AddColleague = `INSERT INTO Colleague(fio,loginpassword,status)
VALUES(:fio, :loginpassword, :status)`

var BlockColleague = `UPDATE Colleague SET status = 'Заблокирован'
WHERE id = $1`

var AddClient = `INSERT INTO Client(fio,passport,email_telephone,loginpassword)
VALUES($1,$2,$3,$4)`

var ShowClientInfo = `SELECT c.fio AS clientFullname, c.passport AS clientPassport,c.email_telephone AS clientEmailTelephone, v.name AS vacationName
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

var ShowClientsApplication = `SELECT c.fio AS clientFullname,a.status AS applicationStatus, a.id AS applicationId, a.client_message AS clientMessage
FROM client c
LEFT JOIN clientapplication ca ON ca.clientid = c.id
LEFT JOIN application a ON a.id = ca.applicationid
WHERE a.id IS NOT NULL;`

var ChangeApplicationStatus = `UPDATE application SET status = $1
WHERE id = $2`

var InsertCompany = `INSERT INTO Company (name,inn,legal_address)
VALUES($1,$2,$3)`

var InsertCompanyClients = `INSERT INTO CompanyClients(CompanyId,ClientId)
VALUES($1,$2)`

var ChangeCompanyName = `UPDATE Company SET name = $1
WHERE id = $2`

var ChangeCompanyInn = `UPDATE Company SET inn = $1
WHERE id = $2`

var ChangeCompanyLegalAddress = `UPDATE Company SET legal_address = $1
WHERE id = $2`

var DeleteCompanyClients = `DELETE FROM companyclients WHERE companyid = $1`

var ShowClientCompanysInfo = `SELECT c.fio AS clientFullname, v.name AS companyName, v.inn AS companyInn, v.legal_address AS companyLegalAddress
FROM client c
LEFT JOIN companyclients cv ON cv.clientid = c.id
LEFT JOIN company v ON v.id = cv.companyid
WHERE c.id = $1;`

var AddApplication = `INSERT INTO Application(vacationid,status,client_message)
VALUES($1,$2,$3)`

var InsertIntoClientApplication = `INSERT INTO ClientApplication(ClientId,ApplicationId)
VALUES($1,$2);`

var InsertIntoClientVacation = `INSERT INTO ClientVacations(ClientId,VacationId)
VALUES($1,$2);`

var GetClientByMessage = `SELECT id, vacationId, status, client_message AS clientMessage FROM application WHERE client_message = $1`

var GetClientIdFromClientVacations = `SELECT ClientId AS id FROM ClientVacations WHERE ClientId = $1 AND VacationId = $2`

var GetClientById = `SELECT id,fio,passport,email_telephone AS emailTelephone,loginpassword FROM client WHERE id = $1`

var GetCompanyIdByName = `SELECT id FROM company WHERE name = $1`

var GetClientIdByFio = `SELECT id FROM client WHERE fio = $1`

var GetCompanyById = `SELECT id,name,inn,legal_address AS legalAddress FROM company WHERE id = $1`
