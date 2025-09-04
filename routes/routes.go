package routes

import (
	"github.com/nekowawolf/airdropv2/controllers"
	"github.com/nekowawolf/airdropv2/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/airdrop")

	api.Post("/login", controllers.LoginAdminHandler)
	
	api.Get("/freeairdrop", controllers.GetAirdropFreeHandler)
	api.Get("/paidairdrop", controllers.GetAirdropPaidHandler)
	api.Get("/allairdrop", controllers.GetAllAirdropHandler)
	api.Get("/allairdrop/search/:name", controllers.GetAllAirdropByNameHandler)
	api.Get("/freeairdrop/search/:name", controllers.GetAirdropFreeByNameHandler)
	api.Get("/paidairdrop/search/:name", controllers.GetAirdropPaidByNameHandler)

	api.Get("/cryptocommunity", controllers.GetAllCryptoCommunity)
	api.Get("/cryptocommunity/search/:name", controllers.GetCryptoCommunityByName)
	
	api.Get("/price", controllers.PriceHandler)

	protected := api.Group("/", middlewares.AdminMiddleware())

	protected.Get("/freeairdrop/:id", controllers.GetAirdropFreeByIDHandler)
	protected.Get("/paidairdrop/:id", controllers.GetAirdropPaidByIDHandler)
	protected.Post("/freeairdrop", controllers.InsertAirdropFreeHandler)
	protected.Post("/paidairdrop", controllers.InsertAirdropPaidHandler)
	protected.Put("/allairdrop/:id", controllers.UpdateAllAirdropByIDHandler)
	protected.Put("/freeairdrop/:id", controllers.UpdateAirdropFreeByIDHandler)
    protected.Put("/paidairdrop/:id", controllers.UpdateAirdropPaidByIDHandler)
	protected.Delete("/allairdrop/:id", controllers.DeleteAllAirdropByIDHandler)
	protected.Delete("/freeairdrop/:id", controllers.DeleteAirdropFreeByIDHandler)
    protected.Delete("/paidairdrop/:id", controllers.DeleteAirdropPaidByIDHandler)

	protected.Get("/notes", controllers.GetAllNotes)
	protected.Get("/notes/:id", controllers.GetNotesByID)
	protected.Post("/notes", controllers.InsertNotes)
	protected.Put("/notes/:id", controllers.UpdateNotesByID)
	protected.Delete("/notes/:id", controllers.DeleteNotesByID)

	protected.Get("/cryptocommunity/:id", controllers.GetCryptoCommunityByID)
	protected.Post("/cryptocommunity", controllers.InsertCryptoCommunity)
	protected.Put("/cryptocommunity/:id", controllers.UpdateCryptoCommunityByID)
	protected.Delete("/cryptocommunity/:id", controllers.DeleteCryptoCommunityByID)

	protected.Get("/admin/info", controllers.GetAdminInfo)
	protected.Post("/logout", controllers.LoginAdminHandler)
}