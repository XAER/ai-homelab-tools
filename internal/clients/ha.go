package clients

import "github.com/gofiber/fiber/v2"

type HA struct {
	base, token string
}
func NewHA(base, token string) *HA { return &HA{base: base, token: token} }

func (h *HA) ListEntities(c *fiber.Ctx) error {
	req := fiber.AcquireAgent()
	defer fiber.ReleaseAgent(req)
	req.Request().Header.SetMethod(fiber.MethodGet)
	req.Request().Header.Set("Authorization", "Bearer "+h.token)
	req.Request().SetRequestURI(h.base + "/api/states")
	if err := req.Parse(); err != nil { return err }
	code, body, errs := req.Bytes()
	if len(errs) > 0 { return errs[0] }
	return c.Status(code).Send(body)
}

