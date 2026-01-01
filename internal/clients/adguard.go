package clients

import "github.com/gofiber/fiber/v2"

type AdGuard struct {
	base, user, pass string
}
func NewAdGuard(base, user, pass string) *AdGuard { return &AdGuard{base: base, user: user, pass: pass} }

func (a *AdGuard) Stats(c *fiber.Ctx) error {
	// AdGuard JSON API: /control/stats
	req := fiber.AcquireAgent(); defer fiber.ReleaseAgent(req)
	req.Request().Header.SetMethod(fiber.MethodGet)
	req.Request().SetRequestURI(a.base + "/control/stats")
	req.BasicAuth(a.user, a.pass)
	if err := req.Parse(); err != nil { return err }
	code, body, errs := req.Bytes()
	if len(errs) > 0 { return errs[0] }
	return c.Status(code).Send(body)
}

