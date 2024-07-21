package main

import (
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ccadden/htmx-go/web_1.0/models"
	"github.com/ccadden/htmx-go/web_1.0/views"
)

var db *gorm.DB

func main() {
	initDB()

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/contacts", contactsHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Migrator().DropTable(&models.Contact{})
	db.AutoMigrate(&models.Contact{})
	db.Create(&models.Contact{Name: "MC"})
	db.Create(&models.Contact{Name: "CC"})
	db.Create(&models.Contact{Name: "CR"})
}

func contactsHandler(w http.ResponseWriter, r *http.Request) {
	var contacts []models.Contact

	q := r.URL.Query().Get("q")

	if q == "" {
		result := db.Find(&contacts)
		if result.Error != nil {
			log.Fatal("Failed to get contacts")
		}
	} else {
		q = "%" + q + "%"
		result := db.Where("Name LIKE ?", q).Find(&contacts)
		if result.Error != nil {
			log.Fatal("Failed to get contacts")
		}

	}

	if err := views.ContactsList(contacts).Render(r.Context(), w); err != nil {
		log.Fatal("Failed to render contacts page")

	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if err := views.Hello("World").Render(r.Context(), w); err != nil {
		log.Fatal("Failed to render hello page")

	}
}
