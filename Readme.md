# How to run

## Backend App

### With Docker:
- Make sure that docker is already installed on your local computer
- Go to `backend\` folder
```sh
docker build -t imam-bookcabin-backend .
docker run -p 8080:8080 imam-bookcabin-backend
```

### Without docker:
- Make sure your [go version is at least 1.24.5](https://go.dev/dl/)
- Go to `backend\` folder
```sh
go mod tidy
go run cmd/main.go
```

## Frontend App

### With docker
- Make sure that docker is already installed on your local computer
- Go to `frontend\` folder
```sh
docker build -t imam-bookcabin-frontend .
docker run -p 5174:5174 imam-bookcabin-frontend
```

### Without docker
- Make sure your NPM version is at least 10.9.3
- Go to `frontend\` folder
```sh
npm install
npm run dev
```

