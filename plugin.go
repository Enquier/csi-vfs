// +build linux,plugin

package main

import "C"

import (
	"github.com/enquier/csi-vfs/provider"
	"github.com/enquier/csi-vfs/service"
)

////////////////////////////////////////////////////////////////////////////////
//                              Go Plug-in                                    //
////////////////////////////////////////////////////////////////////////////////

// SupportedVersions is a space-delimited list of supported CSI versions.
var SupportedVersions = service.SupportedVersions

// ServiceProviders is an exported symbol that provides a host program
// with a map of the service provider names and constructors.
var ServiceProviders = map[string]func() interface{}{
	service.Name: func() interface{} { return provider.New() },
}
