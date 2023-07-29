package library

import (
	"fmt"
	"github.com/bxcodec/faker/v4"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func NewConnection(config Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {

		return db, err
	}

	return db, nil
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]Books{}

	err := r.DB.Find(bookModels).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get books"})
		return err
	}

	err = context.Render("index", fiber.Map{
		"title": "Hello, World!",
		"data":  bookModels,
	})

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not render index page"})
		return err
	}

	return nil
}

func (r *Repository) randomFillDB() {
	for i := 0; i < 10; i++ {
		book := Book{
			Title:     faker.Name(),
			Author:    faker.Name(),
			Publisher: faker.Name(),
		}

		err := r.CreateBook(book)
		if err != nil {
			log.Fatal(fmt.Sprintf("DB fill [%d] -> %s\n", i, err))
		}
	}
}

func (r *Repository) CreateBook(book Book) error {
	databaseValidator := validator.New()
	err := databaseValidator.Struct(book)

	if err != nil {
		log.Fatal(err)
		return err
	}

	// Adicione validacao aqui

	err = r.DB.Create(&book).Error

	if err != nil {
		log.Fatal("could not create book", err)
		return err
	}

	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/books", r.GetBooks)

}
