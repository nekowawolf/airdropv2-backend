package routes

import (
	"github.com/nekowawolf/airdropv2/controllers"
	"github.com/nekowawolf/airdropv2/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/airdrop")

	api.Post("/login", controllers.LoginAdminHandler)
	api.Post("/refresh", controllers.RefreshTokenHandler)
	api.Post("/logout", controllers.LogoutHandler)

	api.Get("/portfolio", controllers.GetPortfolio)
	
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

	protected.Get("/allairdrop/:id", controllers.GetAllAirdropByIDHandler)
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

	protected.Put("/portfolio", controllers.UpdatePortfolio)

	protected.Post("/portfolio/certificates", controllers.AddCertificate)
	protected.Post("/portfolio/designs", controllers.AddDesign)
	protected.Post("/portfolio/projects", controllers.AddProject)
	protected.Post("/portfolio/experience", controllers.AddExperience)
	protected.Post("/portfolio/education", controllers.AddEducation)
	protected.Post("/portfolio/skills/tech", controllers.AddTechSkill)

	protected.Delete("/portfolio/certificates/:id", controllers.DeleteCertificate)
	protected.Delete("/portfolio/designs/:id", controllers.DeleteDesign)
	protected.Delete("/portfolio/projects/:id", controllers.DeleteProject)
	protected.Delete("/portfolio/experience/:id", controllers.DeleteExperience)
	protected.Delete("/portfolio/education/:id", controllers.DeleteEducation)
	protected.Delete("/portfolio/skills/tech/:id", controllers.DeleteTechSkill)

	protected.Post("/images", controllers.UploadImageHandler)
	protected.Get("/images", controllers.GetAllImages)
	protected.Delete("/images/:id", controllers.DeleteImage)
}