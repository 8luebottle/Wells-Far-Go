# Wells Far-Go
**Bank Account Management System Project in Go**
> Work in progress

<img width="1000" alt="WFgo logo" src="https://user-images.githubusercontent.com/48475824/83945339-f206dd00-a844-11ea-9fe1-f52725fea95d.png">

### Table of contents
* [Documentation](#documentation)
* [DB Table](#db-table)
* [ERD](#erd)
* [Tech Stack](#tech-stack)
  * [Back-end Tech Stack](#back-end-tech-stack)


## Documentation
To view **Wells Far-Go** document, please click on the link below.

[Wells Far-Go Documentation](https://bit.ly/wells-far-Go)  

<img width="400" alt="wfg documentation" src="https://user-images.githubusercontent.com/48475824/83948176-89286080-a856-11ea-936a-1fd141d8c833.png">

[Return to TOC](#table-of-contents)

## DB Table
### Accounts
```sql
   Column    |         Type          |
-------------+-----------------------+
 number      | text                  |
 customer_id | text                  |
 type        | integer               |
 status      | integer               |
 name        | character varying(50) |
 balance     | numeric               |
 opened_at   | bigint                |
 deleted_at  | bigint                |
```

### Addresses
```sql
     Column     |  Type   |
----------------+---------+
 id             | bigint  |
 customer_id    | text    |
 account_number | text    |
 type           | integer |
 state          | text    |
 city           | text    |
 street         | text    |
 zip_code       | bigint  |
 country        | text    | 
```

### Banks
```sql
  Column   | Type |
-----------+------+
 bank_code | text |
 name      | text |
 address   | text |
```

### Customers
```sql
     Column      |  Type   |
-----------------+---------+
 id              | text    |
 password        | text    |
 status          | integer |
 first_name      | text    |
 middle_name     | text    |
 last_name       | text    |
 gender          | integer |
 salary          | numeric |
 country_code    | integer |
 phone           | text    |
 email           | text    |
 dob             | integer |
 registered_date | bigint  |
 updated_at      | bigint  |
 deleted_at      | bigint  |
```

### Transactions
```sql
     Column     |          Type          |
----------------+------------------------+
 id             | uuid                   |
 account_number | text                   |
 customer_id    | text                   |
 type           | text                   |
 status         | text                   |
 amount         | numeric                |
 transfer_date  | bigint                 |
 from           | text                   |
 to             | text                   |
 description    | character varying(150) |
```
[Return to TOC](#table-of-contents)

## ERD
[Wells Far-Go ER Diagram](https://www.erdcloud.com/d/F6f6pKyCn68a4te9z)
* Account
* Admin
* Bank
* Customer
* Deposit
* Withdraw
* Transaction

[Return to TOC](#table-of-contents)

## Tech Stack
### Back-end Tech Stack
#### Programming languages 
* Go
#### Frameworks
* Echo
#### Databases
* Redis
* PostgreSQL
#### Server Providers

[Return to TOC](#table-of-contents)
