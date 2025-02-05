package routes

import (
	"github.com/nekowawolf/airdropv2/controllers"
	"github.com/nekowawolf/airdropv2/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/airdrop")

	api.Post("/login", controllers.InsertAdminHandler)
	
	api.Get("/freeairdrop", controllers.GetAirdropFreeHandler)
	api.Get("/paidairdrop", controllers.GetAirdropPaidHandler)
	api.Get("/allairdrop", controllers.GetAllAirdropHandler)
	api.Get("/freeairdrop/:id", controllers.GetAirdropFreeByIDHandler)
	api.Get("/paidairdrop/:id", controllers.GetAirdropPaidByIDHandler)
	api.Get("/allairdrop/search/:name", controllers.GetAllAirdropByNameHandler)
	api.Get("/freeairdrop/search/:name", controllers.GetAirdropFreeByNameHandler)
	api.Get("/paidairdrop/search/:name", controllers.GetAirdropPaidByNameHandler)

	api.Get("/cryptocommunity", controllers.GetAllCryptoCommunity)
	api.Get("/cryptocommunity:id", controllers.GetCryptoCommunityByID)
	api.Get("/cryptocommunity/search/:name", controllers.GetCryptoCommunityByName)
	
	protected := api.Group("/", middlewares.AdminMiddleware())

	protected.Post("/freeairdrop", controllers.InsertAirdropFreeHandler)
	protected.Post("/paidairdrop", controllers.InsertAirdropPaidHandler)
	protected.Put("/freeairdrop/:id", controllers.UpdateAirdropFreeByIDHandler)
    protected.Put("/paidairdrop/:id", controllers.UpdateAirdropPaidByIDHandler)
	protected.Delete("/freeairdrop/:id", controllers.DeleteAirdropFreeByIDHandler)
    protected.Delete("/paidairdrop/:id", controllers.DeleteAirdropPaidByIDHandler)

	protected.Get("/notes", controllers.GetAllNotes)
	protected.Get("/notes/:id", controllers.GetNotesByID)
	protected.Post("/notes", controllers.InsertNotes)
	protected.Put("/notes/:id", controllers.UpdateNotesByID)
	protected.Delete("/notes/:id", controllers.DeleteNotesByID)

	protected.Post("/cryptocommunity", controllers.InsertCryptoCommunity)
	protected.Put("/cryptocommunity:id", controllers.UpdateCryptoCommunityByID)
	protected.Delete("/cryptocommunity:id", controllers.DeleteCryptoCommunityByID)

	protected.Get("/admin/info", controllers.GetAdminInfo)
	protected.Post("/logout", controllers.LoginAdminHandler)
}