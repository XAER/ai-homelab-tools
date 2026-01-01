package clients

import "github.com/gofiber/fiber/v2"

type Portainer struct {
	base, user, pass, apiKey string
}
func NewPortainer(base, user, pass, apiKey string) *Portainer {
	return &Portainer{base: base, user: user, pass: pass, apiKey: apiKey}
}

func (p *Portainer) ListStacks(c *fiber.Ctx) error {
	req := fiber.AcquireAgent(); defer fiber.ReleaseAgent(req)
	req.Request().Header.SetMethod(fiber.MethodGet)
	req.Request().SetRequestURI(p.base + "/stacks")
	if p.apiKey != "" {
		req.Request().Header.Set("X-API-Key", p.apiKey)
	} else {
		req.Request().Header.Set("Authorization", "Basic "+basic(p.user, p.pass))
	}
	if err := req.Parse(); err != nil { return err }
	code, body, errs := req.Bytes()
	if len(errs) > 0 { return errs[0] }
	return c.Status(code).Send(body)
}
func basic(u, p string) string { return fiberutilsBasic(u, p) }

