# HaDiBar #
The goal is to manage the accountings for our beverages on our floor of a dorm conveniently

This is a Server for the HaDiBar-API. It exposes functionality to manage accounts and beverages.

The definiton of api can be found at [Hadibar-API](https://github.com/killingspark/Hadibar-api)
a reference implementation of a webapp that uses said API can be found at [Hadibar-Webapp](https://github.com/killingspark/Hadibar-Webapp).
The Webapp also contains some conveniance functions capsulating the ajax-calls and session management to use the api from javascript 

Webapp made with Vuejs, JQuery and Bootstrap
Accounts/Beverages/Users are stored in bolt key-value stores


## General usecase
This is meant for one person in a group of people to manage the accounts of all of them (in most cases that will be the one that manages the physical beverages as well). It is more of a convenient way of book keeping, not a way so that everyone is managing their own accounts.

## Users ##
There is no explicit user management right now. Usernames are aquired by logging in with the name and the password for the first time.
This will hopefully be improved in the future (with password-resets/registering with an email etc,etc)

## Test the server without the webapp ##
Make calls with curl like those:

Testlogin: 
```
    export SES="$(curl -X GET 127.0.0.1:8080/api/session/getid)"
    curl -X POST 127.0.0.1:8080/api/session/login -H "sessionID: "$SES --form "name=Moritz" --form "password=test"
```

## Admin Server
Besides the rest-api for the normal users there is an additional admin-server listening on a unix socket.

The admin-server uses a JSON based PRC that allows to manage all things. There is support for 
1. listing/removing users
1. perform cleanup after deleting users
1. listing accounts
1. listing beverages
1. perform backups

The most comfortable way is probably to use the admin-client in the cmd directory

Example for performing backups and put them in a directory with the current date

```
go run src/cmd/admin-client/main.go -s sockets/control.socket backup "backups/$(date -u +"%Y-%m-%d__%H:%M:%S")"
```
