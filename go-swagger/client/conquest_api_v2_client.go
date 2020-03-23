// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/ConquestSolutions/apiv2.examples/go-swagger/client/action_service"
	"github.com/ConquestSolutions/apiv2.examples/go-swagger/client/asset_service"
	"github.com/ConquestSolutions/apiv2.examples/go-swagger/client/defect_service"
	"github.com/ConquestSolutions/apiv2.examples/go-swagger/client/document_service"
	"github.com/ConquestSolutions/apiv2.examples/go-swagger/client/jobs_service"
	"github.com/ConquestSolutions/apiv2.examples/go-swagger/client/knowledge_base_service"
	"github.com/ConquestSolutions/apiv2.examples/go-swagger/client/request_service"
	"github.com/ConquestSolutions/apiv2.examples/go-swagger/client/system_service"
	"github.com/ConquestSolutions/apiv2.examples/go-swagger/client/view_service"
)

// Default conquest API v2 HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new conquest API v2 HTTP client.
func NewHTTPClient(formats strfmt.Registry) *ConquestAPIV2 {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new conquest API v2 HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *ConquestAPIV2 {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new conquest API v2 client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *ConquestAPIV2 {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(ConquestAPIV2)
	cli.Transport = transport
	cli.ActionService = action_service.New(transport, formats)
	cli.AssetService = asset_service.New(transport, formats)
	cli.DefectService = defect_service.New(transport, formats)
	cli.DocumentService = document_service.New(transport, formats)
	cli.JobsService = jobs_service.New(transport, formats)
	cli.KnowledgeBaseService = knowledge_base_service.New(transport, formats)
	cli.RequestService = request_service.New(transport, formats)
	cli.SystemService = system_service.New(transport, formats)
	cli.ViewService = view_service.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// ConquestAPIV2 is a client for conquest API v2
type ConquestAPIV2 struct {
	ActionService action_service.ClientService

	AssetService asset_service.ClientService

	DefectService defect_service.ClientService

	DocumentService document_service.ClientService

	JobsService jobs_service.ClientService

	KnowledgeBaseService knowledge_base_service.ClientService

	RequestService request_service.ClientService

	SystemService system_service.ClientService

	ViewService view_service.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *ConquestAPIV2) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.ActionService.SetTransport(transport)
	c.AssetService.SetTransport(transport)
	c.DefectService.SetTransport(transport)
	c.DocumentService.SetTransport(transport)
	c.JobsService.SetTransport(transport)
	c.KnowledgeBaseService.SetTransport(transport)
	c.RequestService.SetTransport(transport)
	c.SystemService.SetTransport(transport)
	c.ViewService.SetTransport(transport)
}