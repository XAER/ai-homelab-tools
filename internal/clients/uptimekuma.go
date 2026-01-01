package clients

import "github.com/gofiber/fiber/v2"

type UptimeKuma struct {
	base, api string
}
func NewUptimeKuma(base, api string) *UptimeKuma { return &UptimeKuma{base: base, api: api} }

func (u *UptimeKuma) Status(c *fiber.Ctx) error {
	// If public status page JSON is available:
	req := fiber.AcquireAgent(); defer fiber.ReleaseAgent(req)
	req.Request().Header.SetMethod(fiber.MethodGet)
	req.Request().SetRequestURI(u.base + "/api/status-page/homepage")
	if u.api != "" {
		req.Request().Header.Set("Authorization", "Bearer "+u.api)
	}
	if err := req.Parse(); err != nil { return err }
	code, body, errs := req.Bytes()
	if len(errs) > 0 { return errs[0] }
	return c.Status(code).Send(body)
}

