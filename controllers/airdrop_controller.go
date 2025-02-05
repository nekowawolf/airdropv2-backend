package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nekowawolf/airdropv2/module"
	"github.com/nekowawolf/airdropv2/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
)

func GetAirdropFreeHandler(c *fiber.Ctx) error {
	data, err := module.GetAllAirdropFree() 
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve AirdropFree data",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    data,
	})
}

func GetAirdropPaidHandler(c *fiber.Ctx) error {
	data, err := module.GetAllAirdropPaid() 
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve AirdropFree data",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    data,
	})
}

func GetAirdropFreeByIDHandler(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	data, err := module.GetAirdropFreeByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve AirdropFree by ID",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    data,
	})
}

func GetAirdropPaidByIDHandler(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	data, err := module.GetAirdropPaidByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve AirdropPaid by ID",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    data,
	})
}

func GetAllAirdropHandler(c *fiber.Ctx) error {
	data, err := module.GetAllAirdrop()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve all Airdrop data",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    data,
	})
}

func GetAllAirdropByNameHandler(c *fiber.Ctx) error {
    name := c.Params("name")
    
    data, err := module.GetAllAirdropByName(name)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to retrieve Airdrop by Name",
        })
    }
    
    if len(data) == 0 {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "No Airdrop found with the specified name",
        })
    }
    
    return c.JSON(fiber.Map{
        "message": "Data retrieved successfully",
        "data":    data,
    })
}

func GetAirdropFreeByNameHandler(c *fiber.Ctx) error {
    name := c.Params("name")
    
    data, err := module.GetAirdropFreeByName(name)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to retrieve AirdropFree by Name",
        })
    }
    
    if len(data) == 0 {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "No AirdropFree found with the specified name",
        })
    }
    
    return c.JSON(fiber.Map{
        "message": "Data retrieved successfully",
        "data":    data,
    })
}

func GetAirdropPaidByNameHandler(c *fiber.Ctx) error {
    name := c.Params("name")
    
    data, err := module.GetAirdropPaidByName(name)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to retrieve AirdropPaid by Name",
        })
    }
    
    if len(data) == 0 {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "No AirdropPaid found with the specified name",
        })
    }
    
    return c.JSON(fiber.Map{
        "message": "Data retrieved successfully",
        "data":    data,
    })
}

func InsertAirdropFreeHandler(c *fiber.Ctx) error {
	var reqAirdrop models.AirdropFree

	if err := c.BodyParser(&reqAirdrop); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	insertedID, err := module.InsertAirdropFree(
		reqAirdrop.Name,
		reqAirdrop.Task,
		reqAirdrop.Link,
		reqAirdrop.Level,
		reqAirdrop.Status,
		reqAirdrop.Backed,
		reqAirdrop.Funds,
		reqAirdrop.Supply,
		reqAirdrop.MarketCap,
		reqAirdrop.Vesting,
		reqAirdrop.LinkClaim,
		reqAirdrop.Price,
		reqAirdrop.USDIncome,
		
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert AirdropFree",
		})
	}

	if objectID, ok := insertedID.(primitive.ObjectID); ok {
		return c.JSON(fiber.Map{
			"message":     "AirdropFree inserted successfully",
			"inserted_id": objectID.Hex(),
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "Failed to retrieve inserted ID",
	})
}

func InsertAirdropPaidHandler(c *fiber.Ctx) error {
	var reqAirdrop models.AirdropPaid

	if err := c.BodyParser(&reqAirdrop); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	insertedID, err := module.InsertAirdropPaid(
		reqAirdrop.Name,
		reqAirdrop.Task,
		reqAirdrop.Link,
		reqAirdrop.Level,
		reqAirdrop.Status,
		reqAirdrop.Backed,
		reqAirdrop.Funds,
		reqAirdrop.Supply,
		reqAirdrop.MarketCap,
		reqAirdrop.Vesting,
		reqAirdrop.LinkClaim,
		reqAirdrop.Price,
		reqAirdrop.USDIncome,

	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert AirdropPaid",
		})
	}

	if objectID, ok := insertedID.(primitive.ObjectID); ok {
		return c.JSON(fiber.Map{
			"message":     "AirdropPaid inserted successfully",
			"inserted_id": objectID.Hex(),
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "Failed to retrieve inserted ID",
	})
}

func UpdateAirdropFreeByIDHandler(c *fiber.Ctx) error {
    idParam := c.Params("id")
    id, err := primitive.ObjectIDFromHex(idParam)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid ID format",
        })
    }

    var updateData struct {
        Name      string  `json:"name"`
        Task      string  `json:"task"`
        Link      string  `json:"link"`
        Level     string  `json:"level"`
        Status    string  `json:"status"`
        Backed    string  `json:"backed"`
        Funds     string  `json:"funds"`
		Supply    string  `json:"Supply"`
        MarketCap string  `json:"market_cap"`
        Vesting   string  `json:"vesting"`
		LinkClaim string  `json:"link_claim"`
		Price     float64 `json:"price"`
        USDIncome int     `json:"usd_income"`

    }

    if err := c.BodyParser(&updateData); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Failed to parse request body",
        })
    }

    err = module.UpdateAirdropFreeByID(
        id,
        updateData.Name,
        updateData.Task,
        updateData.Link,
        updateData.Level,
        updateData.Status,
        updateData.Backed,
        updateData.Funds,
		updateData.Supply,
        updateData.MarketCap,
        updateData.Vesting,
		updateData.LinkClaim,
		updateData.Price,
        updateData.USDIncome,
        
    )
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update AirdropFree by ID",
        })
    }

    return c.JSON(fiber.Map{
        "message": "AirdropFree updated successfully",
    })
}

func UpdateAirdropPaidByIDHandler(c *fiber.Ctx) error {
    idParam := c.Params("id")
    id, err := primitive.ObjectIDFromHex(idParam)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid ID format",
        })
    }

    var updateData struct {
        Name      string  `json:"name"`
        Task      string  `json:"task"`
        Link      string  `json:"link"`
        Level     string  `json:"level"`
        Status    string  `json:"status"`
        Backed    string  `json:"backed"`
        Funds     string  `json:"funds"`
		Supply    string  `json:"Supply"`
        MarketCap string  `json:"market_cap"`
        Vesting   string  `json:"vesting"`
        LinkClaim string  `json:"link_claim"`
        Price     float64 `json:"price"`
        USDIncome int     `json:"usd_income"`
    }

    if err := c.BodyParser(&updateData); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Failed to parse request body",
        })
    }

    err = module.UpdateAirdropPaidByID(
        id,
        updateData.Name,
        updateData.Task,
        updateData.Link,
        updateData.Level,
        updateData.Status,
        updateData.Backed,
        updateData.Funds,
		updateData.Supply,
        updateData.MarketCap,
        updateData.Vesting,
        updateData.LinkClaim,
        updateData.Price,
        updateData.USDIncome,
    )
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update AirdropPaid by ID",
        })
    }

    return c.JSON(fiber.Map{
        "message": "AirdropPaid updated successfully",
    })
}

func DeleteAirdropFreeByIDHandler(c *fiber.Ctx) error {
    idParam := c.Params("id")
    id, err := primitive.ObjectIDFromHex(idParam)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid ID format",
        })
    }

    err = module.DeleteAirdropFreeByID(id)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete AirdropFree by ID",
        })
    }

    return c.JSON(fiber.Map{
        "message": "AirdropFree deleted successfully",
    })
}

func DeleteAirdropPaidByIDHandler(c *fiber.Ctx) error {
    idParam := c.Params("id")
    id, err := primitive.ObjectIDFromHex(idParam)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid ID format",
        })
    }

    err = module.DeleteAirdropPaidByID(id)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete AirdropPaid by ID",
        })
    }

    return c.JSON(fiber.Map{
        "message": "AirdropPaid deleted successfully",
    })
} 	