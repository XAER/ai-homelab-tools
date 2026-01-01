package clients

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
)

type Proxmox struct {
	base, tokenID, tokenSecret string
}
func NewProxmox(base, id, secret string) *Proxmox {
	return &Proxmox{base: base, tokenID: id, tokenSecret: secret}
}

func (p *Proxmox) ListNodes(c *fiber.Ctx) error {
	req := fiber.AcquireAgent(); defer fiber.ReleaseAgent(req)
	req.Request().Header.SetMethod(fiber.MethodGet)
	req.Request().Header.Set("Authorization", "PVEAPIToken="+p.tokenID+"="+p.tokenSecret)
	req.Request().SetRequestURI(p.base + "/nodes")
	if err := req.Parse(); err != nil { return err }
	code, body, errs := req.Bytes()
	if len(errs) > 0 { return errs[0] }
	// Proxmox wraps data {data: ...}; return as-is
	return c.Status(code).Send(body)
}

// (Optional) helpers for VM/CT actions later…
func urlJoin(base string, path string) string {
	u, _ := url.Parse(base); u.Path = path; return u.String()
}

