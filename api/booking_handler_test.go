package api

import (
	"fmt"
	"testing"
	"time"

	"github.com/McFlanky/hotel-reservations-api/db/fixtures"
)

func TestAdminGetBookings(t *testing.T) {
	db := setup(t)
	defer db.teardown(t)

	user := fixtures.AddUser(db.Store, "john", "doe", false)
	hotel := fixtures.AddHotel(db.Store, "bar hotel", "a", 2, nil)
	room := fixtures.AddRoom(db.Store, "small", true, 50, hotel.ID)

	from := time.Now()
	till := from.AddDate(0, 0, 2)
	booking := fixtures.AddBooking(db.Store, user.ID, room.ID, from, till)
	fmt.Println(booking)
}
