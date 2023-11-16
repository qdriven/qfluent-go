curl -XPOST localhost:9090/auth/setup

curl -XPOST localhost:9090/auth/register -d '{"username": "hello", "password":"world"}'

curl -i -XPOST localhost:9090/auth/login -d '{"username": "hello", "password":"world"}

{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDEzMTYwOTAsImlzX2FkbWluIjpmYWxzZSwidXNlcl9pZCI6Mn0.PzWGhUALPfO_ZNmXfWX1lJMJHlJbotk5wSWorBK8K1I"}

curl "localhost:9090/products" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDEzMTYwOTAsImlzX2FkbWluIjpmYWxzZSwidXNlcl9pZCI6Mn0.PzWGhUALPfO_ZNmXfWX1lJMJHlJbotk5wSWorBK8K1I"

curl  -XPOST "localhost:3000/auth/logout"

curl "localhost:9090/auth_policies" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDEzMTYwOTAsImlzX2FkbWluIjpmYWxzZSwidXNlcl9pZCI6Mn0.PzWGhUALPfO_ZNmXfWX1lJMJHlJbotk5wSWorBK8K1I"
