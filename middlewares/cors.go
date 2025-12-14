package middlewares

import (
    "strings"

    "github.com/gofiber/fiber/v2/middleware/cors"
)

var origins = []string{
    "https://nekowawolf.xyz",
    "https://www.nekowawolf.xyz",
    "https://cmty.nekowawolf.xyz",
    "https://admin.nekowawolf.xyz",
    "https://nekowawolf.github.io",
    "https://airdrop.nekowawolf.xyz",
    "https://portfolio.nekowawolf.xyz",
    "https://airdropv2-src.vercel.app",
    "https://nekowawolfv2-src.vercel.app",
    "https://admin-airdropv2-src.vercel.app",
    "https://portfolio-src-orpin.vercel.app",
    "https://crypto-communityv2-src.vercel.app",
}

var Cors = cors.Config{
    AllowOrigins:     strings.Join(origins[:], ","),
    AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
    ExposeHeaders:    "Content-Length",
    AllowCredentials: true,
}