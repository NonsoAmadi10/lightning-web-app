# LOMO

Lomo(lightning money) is a payment app that allows users to make payments using their lightning url (lnurl), and enables them to withdraw(cash back) from the app to any LN wallet. It will be built using the golang lnrpc library.

Users can be individuals, businesses, shops, or online vendors. The app will enable users to generate LN URL, and track payments made to that url.


### Functionalities

Generate LN URL for user customers..
Notify and update user balance after an invoice is paid.
Enable users to withdraw their balance to any LN wallet using their LN url.
Enable users to view their transaction history for paid payments and withdrawals.
Authenticate, authorize, and identify users.

### Goals:

Have an understanding of the golang lnrpc library.
Generate LN invoices via the lnrpc library.
Experiment with and understand encoded and decoded LN invoices.
Track payments to LN node-generated invoices in the network.
Create a send and receive payment flow on the LN network.

### Non-Goals:
LN network graph.
Bitcoin on-chain.
Bitcoin on-chain settlement.

### Tools:

Programming language: Golang.
Backend server: Echo framework.
Database: PostgreSQL
Deployment: Docker, Docker Compose.

### Dependency software:

Lnd Node
LNURL 

### Dependency libraries:

Lnrpc for interacting with lnd.
gorm for database operations.

### Backend:

RESTful API using Echo framework.
Modular file structure.
Endpoints:

`/registration`: Required fields - {email, password}.
`/login`: Required fields - {email, password}.
`/payment-link`: Required fields - {amount, description}.

 The generated invoice will then be saved in the database with the userID until it expires or is paid. If it expires, it will be deleted. If it is paid, a new payment document will be created and the user balance will be incremented.

`/lnwithdraw/code`: Required field - {auth token}.
`/withdraw/callback`: Required field - {payment-request, amount, auth token}. Initiate a payment to the LN network with the invoice via lnrpc, decrement the balance, and create a new withdrawal document with processing status. If it succeeds, update the status to success. If it fails, reverse the balance and update the status to fail.
`/invoice/:id`: Required field - {auth token, invoiceID}.



### Introspection:

Updating successful payments: The current approach is to have a cron job that runs (if there are unexpired payments) after every five seconds on the server making RPC calls to LND with unexpired unpaid invoices to get new payments, create a new payment document, and increment the user balance.

### MVP Milestones:
Backend API completed
