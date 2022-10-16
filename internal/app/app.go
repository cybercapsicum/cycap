// Package app is the main
package app

import (
	"net/http"
	"os"
	"strings"

	"github.com/cybercapsicum/cycap/internal/pkg/api"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
	App    *openapi3.T
)

func serverFromEnv() (server *openapi3.Server) {
	var (
		serverURL  = os.Getenv("VERCEL_URL")
		serverDesc = os.Getenv("VERCEL_ENV")
	)

	server = &openapi3.Server{
		URL:         serverURL + "/api",
		Description: strings.Title(serverDesc),
	}

	return
}

func newApp() (app *openapi3.T) {
	app = &openapi3.T{
		OpenAPI: "3.0.2",
	}

	app.Info = &openapi3.Info{
		Title:       "cycap",
		Description: "A serverless API for various Cybersecurity utilities",
		License: &openapi3.License{
			Name: "BSD-3-Clause",
		},
		Version: "0.1.0",
	}

	app.Servers = append(app.Servers, serverFromEnv())

	return
}

func index(c *gin.Context) {
	if App == nil {
		api.ErrNotFound.Responder(c.Writer)
		return
	}

	data, _ := App.MarshalJSON()
	api.DataResponse(data).Responder(c.Writer, http.StatusOK)
}

func ping(c *gin.Context) {
	resp := &api.MsgResponse{
		StatusCode: http.StatusOK,
		Message:    "pong",
	}
	resp.Responder(c.Writer)
}

func addApiRoutes(r *gin.RouterGroup) {
	r.GET("/", index)
	r.GET("/ping", ping)
}

func init() {
	App = newApp()
	Router = gin.Default()

	addApiRoutes(Router.Group("/api"))
}
