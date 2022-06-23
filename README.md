
# <p align="center">Payment Service Emulator

<p align="center">
<img src="https://media.giphy.com/media/rLpgNiviGbsw2EBa4G/giphy.gif " width="20%"></p>


## About The Project
________


The Payment Service Emulator creates, changes, cancels payments in the database and gives information about payments. Service implemented on Go.

## Architecture
______
REST API architectural style. The service consists of an API and a database. 

## Used Technologies and protocols

        • Docker
        • HTTP
        • PostgreSQL
        • REST API


## Commands (if you want to run the program)
__________
### Deploy database to Postgres in a Docker container:
``` 
sudo docker run --name=emulator.db -e POSTGRES_PASSWORD='1234' -p 5432:5432 -d postgres
CREATE DATABASE payment
```

### Transfer Database (migration):
```
migration CD
goose postgres "postgres://postgres:123@localhost:5432/payment?sslmode=disable" up
```

### Run Makefile to start the service:
```
make run 
```

### Stop the service:
```
make down 
```

## Description of items
_________
### User
User is stored outside of the database.
* ID
* Email

### Payment 
Payment is stored in the database according to the following model:
* ID
* UserID
* UserEmail
* Amount
* Currency
* TimeCreation
* TimeChange
* Status

Status of payment can take one of the following values: 
* StatusNew "new" means a payment is new
* StatusSuccess "success" informs that a payment is paid
* StatusFailed "failed" means a payment has failed
* StatusError "error" means that something went wrong at the time of creating the payment
* StatusCancelled "cancelled" means a payment is cancelled

The payment life cycle is as follows: the user creates a payment, it is created in the NEW status. After the payment service notifies us whether the payment transfered through on its side, after which we change the status in our database.

SUCCESS and FAIL statuses are terminal - if the payment is in them, its status should be impossible to change. The transition to the SUCCESS and FAIL statuses should be carried out only after receiving the request StatusUpdate. 
ERROR is a status when something went wrong at the time of creating the payment.
CANCELED - this is the status that the payment goes to when the cancel action is called.

## Description of methods (requests to the server)
_______

### CreatePayment 
Creates payment in database.
```
Path "/payment", HTTP method POST. 
```

### UpdateStatus
Updates payment status in database.
```
Path "/payment/{id}", HTTP method PUT. 
```
### GetStatus
Receives status of payment by payment ID from database.
```
Path "/payment/{id}", HTTP method GET. 
```

### GetPaymentsByUserId 
Receives payments by user ID from database.
```
Path "/payment/{userID}", HTTP method GET. 
```
### GetPaymentsByUserEmail 
Receives payments by user email from database.
```
Path "/payment/{userEmail}", HTTP method GET. 
```

### CancelPayment 
Cancels payment by ID and deletes it from database.
```
Path "/payment/{id}", HTTP method DELETE. 
```

## License
_________
All source code is licensed under the [MIT License](https://choosealicense.com/licenses/mit/).