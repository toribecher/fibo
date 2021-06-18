# fibo
A fibonacci fun house

This application exposes three endpoints
It runs using postgresql => https://www.postgresql.org/docs/

Prerequisites: 
1. You will need to set your environment variables, I recommend https://direnv.net/

Environment variables to set in .envrc:
    export HOST={YOURHOST}
    export PORT={YOURPORT}
    export PASSWORD={YOURPASSWORD}
    export USER={YOURUSER}
    export DBNAME={YOURDB}

2. Remember to clone repo in your GOPATH :p

To See App in Action

Step 1. run tests `make test`
Step 2. run application `make run`

3 endpoints are exposed
GET localhost:8000/fibonacci/{number}
GET localhost:8000/memoizedresults/{memoNumber}
GET localhost:8000/deleteall

Note: Delete is usually a delete request but I made it a Get for simplicity (no curl required)

Thanks for the challenge!
