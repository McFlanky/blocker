package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/McFlanky/hotel-reservations-api/api/middleware"
	"github.com/McFlanky/hotel-reservations-api/db/fixtures"
	"github.com/McFlanky/hotel-reservations-api/types"
	"github.com/gofiber/fiber/v2"
)

func TestUserGetBooking(t *testing.T) {
	db := setup(t)
	defer db.teardown(t)
	var (
		nonAuthUser    = fixtures.AddUser(db.Store, "jimmy", "watercooler", false)
		user           = fixtures.AddUser(db.Store, "john", "doe", false)
		hotel          = fixtures.AddHotel(db.Store, "bar hotel", "a", 2, nil)
		room           = fixtures.AddRoom(db.Store, "small", true, 50, hotel.ID)
		from           = time.Now()
		till           = from.AddDate(0, 0, 2)
		booking        = fixtures.AddBooking(db.Store, user.ID, room.ID, from, till)
		app            = fiber.New()
		route          = app.Group("/", middleware.JWTAuthentication(db.User))
		bookingHandler = NewBookingHandler(db.Store)
	)
	route.Get("/:id", bookingHandler.HandleGetBooking)
	req := httptest.NewRequest("GET", fmt.Sprintf("/%s", booking.ID.Hex()), nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(user))
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("non 200 code got %d", resp.StatusCode)
	}
	var bookingResp *types.Booking
	if err := json.NewDecoder(resp.Body).Decode(&bookingResp); err != nil {
		log.Fatal(err)
	}
	if bookingResp.ID != booking.ID {
		t.Fatalf("expected %s but got %s", booking.ID, bookingResp.ID)
	}
	if bookingResp.UserID != booking.UserID {
		t.Fatalf("expected %s but got %s", booking.UserID, bookingResp.UserID)
	}

	// test non-auth user cannot access the booking
	req = httptest.NewRequest("GET", fmt.Sprintf("/%s", booking.ID.Hex()), nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(nonAuthUser))
	resp, err = app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected a non 200 status code but got %d", resp.StatusCode)
	}

}

func TestAdminGetBookings(t *testing.T) {
	db := setup(t)
	defer db.teardown(t)
	var (
		adminUser      = fixtures.AddUser(db.Store, "admin", "admin", true)
		user           = fixtures.AddUser(db.Store, "john", "doe", false)
		hotel          = fixtures.AddHotel(db.Store, "bar hotel", "a", 2, nil)
		room           = fixtures.AddRoom(db.Store, "small", true, 50, hotel.ID)
		from           = time.Now()
		till           = from.AddDate(0, 0, 2)
		booking        = fixtures.AddBooking(db.Store, user.ID, room.ID, from, till)
		app            = fiber.New()
		admin          = app.Group("/", middleware.JWTAuthentication(db.User), middleware.AdminAuth)
		bookingHandler = NewBookingHandler(db.Store)
	)
	admin.Get("/", bookingHandler.HandleGetBookings)
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(adminUser))
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("non 200 response got: %d", resp.StatusCode)
	}
	var bookings []*types.Booking
	if err := json.NewDecoder(resp.Body).Decode(&bookings); err != nil {
		log.Fatal(err)
	}
	if len(bookings) != 1 {
		t.Fatalf("expected 1 booking but got %d", len(bookings))
	}
	have := bookings[0]
	if have.ID != booking.ID {
		t.Fatalf("expected %s but got %s", booking.ID, have.ID)
	}
	if have.UserID != booking.UserID {
		t.Fatalf("expected %s but got %s", booking.UserID, have.UserID)
	}

	// test non-admin cannot access the bookings
	req = httptest.NewRequest("GET", "/", nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(user))
	resp, err = app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected a non 200 status code but got: %d", resp.StatusCode)
	}
}
