package network

import (
	"net/url"

	"github.com/containers/podman/v2/pkg/bindings/util"
)

/*
This file is generated automatically by go generate.  Do not edit.
*/

// Changed
func (o *PruneOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams
func (o *PruneOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}
