// Package api serves as the entrypoint for the serverless API using the Vercel Go Runtime.
//
// # Vercel Go Runtime
//
// Vercel automatically attaches an endpoint based on the filename and folder structure much like a
// typical web framework, such that api/index.go will be attached to '<APP_NAME>.vercel.app/api/'.
// Within the api folder, any go source files, with names not beginning in '_', and including an
// exported function that matches the http.HandlerFunc signature will be handled as entrypoints.
//
// # Config
//
// The vercel.json file in the root of the project should include a 'rewrites' section to forward
// all requests to this file. This allows the use of an external router/mux to handle all /api/*.
//
//	"rewrites": [
//	  {
//	    "source": "/api/(.*)",
//	    "destination": "/api/index.go"
//	  }
//	]
//
// # Reference
//
// https://github.com/vercel/vercel/tree/main/packages/go
package api

import (
	"net/http"

	"github.com/cybercapsicum/cycap/internal/app"
)

// Handler is exposed as the entrypoint for the serverless function.
func Handler(w http.ResponseWriter, r *http.Request) {
	app.Router.ServeHTTP(w, r)
}
