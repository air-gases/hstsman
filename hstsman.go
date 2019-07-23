package hstsman

import (
	"fmt"
	"strings"

	"github.com/aofei/air"
)

// GasConfig is a set of configurations for the `Gas`.
type GasConfig struct {
	MaxAge            int
	IncludeSubDomains bool

	Skippable func(*air.Request, *air.Response) bool
}

// Gas returns an `air.Gas` that is used to manage the Strict-Transport-Security
// header based on the gc.
func Gas(gc GasConfig) air.Gas {
	ds := []string{}
	if gc.MaxAge >= 0 {
		ds = append(ds, fmt.Sprintf("max-age=%d", gc.MaxAge))
	}

	if gc.IncludeSubDomains {
		ds = append(ds, "includeSubDomains")
	}

	directives := strings.Join(ds, "; ")

	return func(next air.Handler) air.Handler {
		return func(req *air.Request, res *air.Response) error {
			if gc.Skippable != nil && gc.Skippable(req, res) {
				return next(req, res)
			}

			if req.Air.HTTPSEnforced {
				res.Header.Set(
					"Strict-Transport-Security",
					directives,
				)
			}

			return next(req, res)
		}
	}
}
