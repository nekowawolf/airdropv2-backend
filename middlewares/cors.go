package middlewares

import (
    "strings"

    "github.com/gofiber/fiber/v2/middleware/cors"
)

var origins = []string{
    "https://nekowawolf.xyz",
    "https://www.nekowawolf.xyz",
    "https://nekowawolf.github.io",
    "https://nekowawolfv2-src.vercel.app",
}

var Cors = cors.Config{
    AllowOrigins:     strings.Join(origins[:], ","),
    AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
    ExposeHeaders:    "Content-Length",
    AllowCredentials: true,
}