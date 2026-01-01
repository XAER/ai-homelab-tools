package httpserver

import (
	"ai-homelab-tools/internal/clients"
	"os"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:               "ai-homelab-tools",
		DisableStartupMessage: true,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
	})

	// Health
	app.Get("/health", func(c *fiber.Ctx) error { return c.SendString("ok") })

	// Clients
	ha := clients.NewHA(os.Getenv("HA_BASE"), os.Getenv("HA_TOKEN"))
	portainer := clients.NewPortainer(os.Getenv("PORTAINER_BASE"), os.Getenv("PORTAINER_USER"), os.Getenv("PORTAINER_PASS"), os.Getenv("PORTAINER_APIKEY"))
	proxmox := clients.NewProxmox(os.Getenv("PROXMOX_BASE"), os.Getenv("PROXMOX_TOKEN_ID"), os.Getenv("PROXMOX_TOKEN_SECRET"))
	uk := clients.NewUptimeKuma(os.Getenv("UPTIMEKUMA_BASE"), os.Getenv("UPTIMEKUMA_API_KEY"))
	ag := clients.NewAdGuard(os.Getenv("ADGUARD_BASE"), os.Getenv("ADGUARD_USER"), os.Getenv("ADGUARD_PASS"))

	// Routes
	app.Get("/ha/entities", with(ha.ListEntities))
	app.Get("/portainer/stacks", with(portainer.ListStacks))
	app.Get("/proxmox/nodes", with(proxmox.ListNodes))
	app.Get("/uptimekuma/status", with(uk.Status))
	app.Get("/adguard/stats", with(ag.Stats))

	return app
}

func with(fn func(*fiber.Ctx) error) fiber.Handler { return fn }
