# Hotel Reservation backend

## Project Outline
- Users -> book room from hotel
- Admins -> going to check reservations/bookings
- Authentication/Authorization -> JWT 
- Hotels -> CRUD API -> JSON
- Admins -> CRUD API -> JSON
- Scripts -> DBMS -> seeding, migration

### Project Environment Variables
```
HTTP_LISTEN_ADDRESS=:8000
JWT_SECRET=somethingsupersecretthatnobodyknows
MONGO_DB_NAME=hotel-reservation
MONGO_DB_URL=mongodb://localhost:27017
MONGO_DB_URL_TEST=mongodb://localhost:27017
```

## Resources
### MongoDB Driver ->
``` 
https://mongodb.com/docs/drivers/go/current/quick-start
```

#### Installing MongoDB client
```
go get go.mongodb.org/mongo-driver/mongo
```           

#### Start MongoDB 
```
sudo systemctl start mongod
```

### Go Fiber ->
```
https://gofiber.io
```

#### Installing Go Fiber
```
go get github.com/gofiber/fiber/v2
```

### Docker ->
#### Installing MongoDB as a Docker container
```
docker run --name mongodb -d mongo:latest -p 27017:27017
```

