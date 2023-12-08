How to start server:
go run main.go

APIS:

/api/accounts :
curl --location 'http://localhost:8080/api/accounts' \
--header 'Auth: Trupe'


/api/v1/transfer :(the jwt token will be recieved in /api/accounts api response header / header name token)
curl --location 'http://localhost:8080/api/v1/transfer' \
--header 'Auth: Trupe' \
--header 'Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MDE5NjEwNjYsInVzZXJfaWQiOiI0ZjNjYzIyZi04YzA1LTQ3NTAtODJmZi1jZmY4NjNiMWM4MTYifQ.LyxbLFTaPvNyJpX8dDJ85jx9PWRlc-pyp5Enx0uiEwk' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MDIwMTg1MzUsInVzZXJfaWQiOiI0ZjNjYzIyZi04YzA1LTQ3NTAtODJmZi1jZmY4NjNiMWM4MTYifQ.IlqFScbBIdO7lfYrAl964fjBLZNgNPT8SQvYwIdheCU' \
--data '{"from_account_id":"b2fd5a7d-d46c-4b57-a167-b802d4d37ffe",
"to_account_id":"b4433e5b-af63-4aa8-9887-3dd98c2d7968",
"amount":"2000"
}'



