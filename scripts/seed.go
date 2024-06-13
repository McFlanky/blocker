package main

import (
	"context"
	"fmt"
	"log"

	"github.com/McFlanky/hotel-reservations-api/db"
	"github.com/McFlanky/hotel-reservations-api/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)

	hotel := types.Hotel{
		Name:     "Hotel California",
		Location: "California",
	}

	room := types.Room{
		Type:      types.SingleRoomType,
		BasePrice: 100.9,
	}
	_ = room

	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	room.HotelID = insertedHotel.ID

	fmt.Println(insertedHotel)

}
