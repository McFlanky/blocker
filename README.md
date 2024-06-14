# Hotel Reservation backend

## Project Outline
- Users -> book room from hotel
- Admins -> going to check reservations/bookings
- Authentication/Authorization -> JWT 
- Hotels -> CRUD API -> JSON
- Admins -> CRUD API -> JSON
- Scripts -> DBMS -> seeding, migration

## Resources
### MongoDB Driver
``` 
https://mongodb.com/docs/drivers/go/current/quick-start
```

Installing mongodb client
```
go get go.mongodb.org/mongo-driver/mongo
```           

Start MongoDB
```
sudo systemctl start mongod
```

### Go Fiber
```
https://gofiber.io
```

Installing gofiber
```
go get github.com/gofiber/fiber/v2
```

## Docker
### Installing mongo db as a Docker container
```
docker run --name mongodb -d mongo:latest -p 27017:27017
```

