package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ccadden/htmx-go/web_1.0/models"
	"github.com/ccadden/htmx-go/web_1.0/views"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Contact{})
	db.Create(&models.Contact{Name: "MC"})
	db.Create(&models.Contact{Name: "CC"})
	db.Create(&models.Contact{Name: "CR"})

	helloComponent := views.Hello("World")
	http.Handle("/", templ.Handler(helloComponent))

	var contacts []models.Contact
	result := db.Find(&contacts)

	if result.Error != nil {
		log.Fatal("Failed to get contacts")
	}

	contactsListComponent := views.ContactsList(contacts)
	http.Handle("/contacts", templ.Handler(contactsListComponent))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
