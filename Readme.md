# How to run
Make sure that docker already installed on your local computer

## Backend App:
Go to `backend\` folder
```sh
go mod tidy
go run cmd/main.go
```

## Frontend App
Go to `frontend\` folder
```sh
docker build -t imam-bookcabin-frontend .
docker run -p 5173:5173 imam-bookcabin-frontend
```