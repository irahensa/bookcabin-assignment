# How to run

## Backend App:
Make sure your [go version is at least 1.24.5](https://go.dev/dl/)
Go to `backend\` folder
```sh
go mod tidy
go run cmd/main.go
```

## Frontend App
Make sure that docker already installed on your local computer
Go to `frontend\` folder
```sh
docker build -t imam-bookcabin-frontend .
docker run -p 5173:5173 imam-bookcabin-frontend
```