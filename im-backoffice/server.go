package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"im-backoffice/auth"
	"im-backoffice/config"
	db "im-backoffice/db/sqlc"
	"im-backoffice/docs"
	"im-backoffice/inventory"
	"im-backoffice/inventoryCategory"
	"im-backoffice/keycloak"
	"im-backoffice/location"
	"im-backoffice/logisticCenter"
	"im-backoffice/middleware"
	"im-backoffice/user"
	"im-backoffice/warehouse"
	"log"
)

// @title Inventory Management API
// @version 2.0
// @description This is an inventory management server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1
// @schemes http https
// @securityDefinitions.apikey ApiKeyAuth
// @name Authorization
// @in header
func main() {
	app, err := NewServer()

	if err != nil {
		log.Fatalln(err)
	}

	err = app.Listen(fmt.Sprintf("%v:%v", config.Configuration.App.Host, config.Configuration.App.Port))

	if err != nil {
		log.Fatalln(err)
	}
}

func NewServer() (*fiber.App, error) {
	connection, err := config.DBConn()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	runDbMigrations("file://db/migration", connection)

	// make fiber faster
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{AllowCredentials: true}))
	app.Use(middleware.Common)
	v1 := app.Group("/api/v1")

	// repo & services
	repository := db.NewRepository(connection)
	keycloakClient := keycloak.NewKeycloak()

	docs.SwaggerInfo.Host = config.Configuration.Swagger.Host
	docs.SwaggerInfo.Schemes = []string{config.Configuration.Swagger.Scheme}
	v1.Get("/swagger/*", swagger.HandlerDefault)

	authService := auth.GetNewAuthService(keycloakClient)
	auth.RouterAuth(v1, authService)

	userService := user.GetNewUserService(keycloakClient)
	user.RouterUser(v1, userService)

	locationService := location.GetNewLocationService(repository)
	location.RouterLocation(v1, locationService)

	inventoryCategoryService := inventoryCategory.GetNewInventoryCategoryService(repository)
	inventoryCategory.RouterInventoryCategory(v1, inventoryCategoryService)

	warehouseService := warehouse.GetNewWarehouseService(repository)
	warehouse.RouterWarehouse(v1, warehouseService)

	logisticCenterService := logisticCenter.GetNewLogisticCenterService(repository)
	logisticCenter.RouterLogisticCenter(v1, logisticCenterService)

	inventoryService := inventory.GetNewInventoryService(repository)
	inventory.RouterInventory(v1, inventoryService)

	return app, nil
}

func runDbMigrations(migrationURL string, db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		log.Fatalln("cannot create instance: ", err)
	}

	m, err := migrate.NewWithDatabaseInstance(migrationURL, config.Configuration.Database.DBName, driver)

	if err != nil {
		log.Fatalln("cannot up migrate: ", err)
	}

	err = m.Up()

	if err != nil && err != migrate.ErrNoChange {
		log.Fatalln("cannot up migrate: ", err)
	}

	log.Println("db migrated successfully")
}
