package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
    clientID     = os.Getenv("clientID")
    clientSecret = os.Getenv("clientSecret")
    redirectURL  = os.Getenv("redirectURL")
    oauthConfig  = &oauth2.Config{
        ClientID:     clientID,
        ClientSecret: clientSecret,
        RedirectURL:  redirectURL,
        Scopes: []string{
            "https://www.googleapis.com/auth/blogger",
        },
        Endpoint: google.Endpoint,
    }
)

var refreshTokenStore = make(map[string]string)

func main() {
    app := fiber.New()

    app.Get("/auth", func(c *fiber.Ctx) error {
        url := oauthConfig.AuthCodeURL("state-token-123", oauth2.AccessTypeOffline)
        return c.Redirect(url)
    })

    app.Get("/accesstoken", func(c *fiber.Ctx) error {
        code := c.Query("code")
        if code == "" {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "Missing authorization code",
            })
        }

        ctx := context.Background()
        token, err := oauthConfig.Exchange(ctx, code)
        if err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": fmt.Sprintf("Failed to exchange token: %v", err),
            })
        }

        if token.RefreshToken != "" {
            refreshTokenStore["current-user"] = token.RefreshToken
        }

        return c.JSON(fiber.Map{
            "access_token":  token.AccessToken,
            "refresh_token": token.RefreshToken,
            "expiry":        token.Expiry.Format(time.RFC3339),
            "token_type":    token.TokenType,
        })
    })

    app.Get("/refreshtoken", func(c *fiber.Ctx) error {
        rToken := c.Query("refresh_token")
        if rToken == "" {
            rToken = refreshTokenStore["current-user"]
            if rToken == "" {
                return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                    "error": "Missing refresh_token parameter",
                })
            }
        }

        ctx := context.Background()
        tokenSource := oauthConfig.TokenSource(ctx, &oauth2.Token{RefreshToken: rToken})

        newToken, err := tokenSource.Token()
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": fmt.Sprintf("Failed to refresh token: %v", err),
            })
        }

        if newToken.RefreshToken != "" {
            refreshTokenStore["current-user"] = newToken.RefreshToken
        }

        return c.JSON(fiber.Map{
            "access_token":  newToken.AccessToken,
            "refresh_token": newToken.RefreshToken,
            "expiry":        newToken.Expiry,
            "token_type":    newToken.TokenType,
        })
    })


    log.Fatal(app.Listen(":3000"))
}
