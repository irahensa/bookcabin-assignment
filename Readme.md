# How to run
Make sure that docker already installed on your local computer

### Backend App:
Go to `backend\` folder
```sh
docker build -t imam-bookcabin-backend .
docker run -p 8080:8080 imam-bookcabin-backend
```

### Frontend App
Go to `frontend\` folder
```sh
docker build -t imam-bookcabin-frontend .
docker run -p 5174:5174 imam-bookcabin-frontend
```