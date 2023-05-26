# README

1. Run prestd server

```shell
prestd 
```

```shell
curl -i -X POST http://127.0.0.1:3000/auth -H "Content-Type: application/json" -d '{"username": "prest", "password": "prest"}'


curl -i -X GET http://127.0.0.1:3000/prest/public/prest_users -H "Accept: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySW5mbyI6eyJpZCI6MSwibWV0YWRhdGEiOm51bGwsIm5hbWUiOiJwcmVzdCIsInVzZXJuYW1lIjoicHJlc3QifSwiZXhwIjoxNjgyMDkzNTk5LCJuYmYiOjE2ODIwNzE5OTl9.OO388BhaKKfYjNmxtrVCuRl4L4OxW3E-Zh-RqjJOJzY"
curl -i -X GET http://127.0.0.1:3000/databases -H "Accept: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySW5mbyI6eyJpZCI6MSwibWV0YWRhdGEiOm51bGwsIm5hbWUiOiJwcmVzdCIsInVzZXJuYW1lIjoicHJlc3QifSwiZXhwIjoxNjgyMDkzNTk5LCJuYmYiOjE2ODIwNzE5OTl9.OO388BhaKKfYjNmxtrVCuRl4L4OxW3E-Zh-RqjJOJzY"
curl -i -X GET http://127.0.0.1:3000/schemas -H "Accept: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySW5mbyI6eyJpZCI6MSwibWV0YWRhdGEiOm51bGwsIm5hbWUiOiJwcmVzdCIsInVzZXJuYW1lIjoicHJlc3QifSwiZXhwIjoxNjgyMDkzNTk5LCJuYmYiOjE2ODIwNzE5OTl9.OO388BhaKKfYjNmxtrVCuRl4L4OxW3E-Zh-RqjJOJzY"
curl -i -X GET http://127.0.0.1:3000/tables -H "Accept: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySW5mbyI6eyJpZCI6MSwibWV0YWRhdGEiOm51bGwsIm5hbWUiOiJwcmVzdCIsInVzZXJuYW1lIjoicHJlc3QifSwiZXhwIjoxNjgyMDkzNTk5LCJuYmYiOjE2ODIwNzE5OTl9.OO388BhaKKfYjNmxtrVCuRl4L4OxW3E-Zh-RqjJOJzY"

curl -i -X GET http://127.0.0.1:3000/show/prest/public/prest_users -H "Accept: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySW5mbyI6eyJpZCI6MSwibWV0YWRhdGEiOm51bGwsIm5hbWUiOiJwcmVzdCIsInVzZXJuYW1lIjoicHJlc3QifSwiZXhwIjoxNjgyMDkzNTk5LCJuYmYiOjE2ODIwNzE5OTl9.OO388BhaKKfYjNmxtrVCuRl4L4OxW3E-Zh-RqjJOJzY"


```

```shell
{
    "FIELD1": "string value",
    "FIELD2": 1234567890
}

Using query string to make filter (WHERE), example:

/{DATABASE}/{SCHEMA}/{TABLE}?{FIELD NAME}={VALUE}
JSON DATA:

{
    "FIELD1": "string value",
    "FIELD2": 1234567890,
    "ARRAYFIELD": ["value 1","value 2"]
}
```

## Defined A Table
Table to API:

```shell
PostgreSQL ➕ REST, low-code, simplify and accelerate development, ⚡ instant, realtime, high-performance on any Postgres application, existing or new
```
