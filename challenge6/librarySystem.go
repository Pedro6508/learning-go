package library

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/template/html/v2"
)

type PageData struct {
	Data []Books
}

func RunSystem() *Repository {
	engine := html.New("./challenge6/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	repository := setupDB(app)
	repository.randomFillDB()

	log.Fatal(app.Listen(":3000"))

	return repository
}

func setupDB(app *fiber.App) *Repository {
	config := Config{
		Host:     "localHost",
		Port:     "5432",
		Password: "1234",
		User:     "admin",
		DBName:   "challenge6database",
		SSLMode:  "disable",
	}

	db, err := NewConnection(config)

	if err != nil {
		log.Fatal("could not load database")
	}

	err = MigrateBooks(db)

	if err != nil {
		log.Fatal("could not migrate db")
	}

	r := &Repository{
		DB: db,
	}

	r.SetupRoutes(app)

	return r
}
