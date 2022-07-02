package api

import (
	"github.com/go-chi/chi"
	"test/internal/pkg"
)

type Controller struct {
	Router      chi.Router
	FileService *pkg.FileService
	KeyService  *pkg.KeyService
}

func CreateNewController(fs *pkg.FileService, ks *pkg.KeyService) *Controller {
	c := &Controller{chi.NewMux(), fs, ks}
	c.constructRoutes()
	return c
}

func (c *Controller) constructRoutes() {
	c.Router.Route("/key", func(r chi.Router) {
		r.Post("/", c.savePublicKey)
		r.Get("/", c.getPublicKey)
	})
	c.Router.Route("/file", func(r chi.Router) {
		r.Get("/", c.getFiles)
		r.Post("/", c.saveFile)
		r.Get("/{fileId}", c.getFile)
	})
}
