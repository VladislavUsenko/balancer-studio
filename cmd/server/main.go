package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// @title           Balancer Studio API
// @version         1.0
// @description     Professional Nginx management platform with beautiful UI and powerful API
// @termsOfService  http://swagger.io/terms/

// @contact.name   Balancer Studio Support
// @contact.email  support@balancer.studio

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:3000
// @BasePath  /api/v1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Balancer Studio v1.0",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Serve Scalar API Documentation
	app.Get("/docs", func(c *fiber.Ctx) error {
		return c.Type("html").SendString(`
<!doctype html>
<html>
<head>
    <title>Balancer Studio API Documentation</title>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <style>
        body { margin: 0; }
    </style>
</head>
<body>
    <script 
        id="api-reference" 
        data-url="/swagger.json"
        data-configuration='{"theme":"purple","layout":"modern","darkMode":true}'
    ></script>
    <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
</body>
</html>
		`)
	})

	// Serve OpenAPI JSON
	app.Get("/swagger.json", func(c *fiber.Ctx) error {
		return c.JSON(getOpenAPISpec())
	})

	// API Routes
	api := app.Group("/api/v1")

	// Health check
	api.Get("/health", HealthCheck)

	// Proxy Hosts routes
	proxyHosts := api.Group("/proxy-hosts")
	proxyHosts.Get("/", ListProxyHosts)
	proxyHosts.Post("/", CreateProxyHost)
	proxyHosts.Get("/:id", GetProxyHost)
	proxyHosts.Put("/:id", UpdateProxyHost)
	proxyHosts.Delete("/:id", DeleteProxyHost)

	// SSL Certificates routes
	certificates := api.Group("/certificates")
	certificates.Get("/", ListCertificates)
	certificates.Post("/", CreateCertificate)

	// Nginx control
	nginx := api.Group("/nginx")
	nginx.Post("/reload", ReloadNginx)
	nginx.Post("/test", TestNginxConfig)
	nginx.Get("/status", GetNginxStatus)

	// Upstream servers management
	upstreams := api.Group("/upstreams")
	upstreams.Get("/", ListUpstreams)
	upstreams.Post("/", CreateUpstream)
	upstreams.Get("/:id/servers", ListUpstreamServers)
	upstreams.Post("/:id/servers", AddUpstreamServer)

	log.Println("🚀 Balancer Studio starting on http://localhost:3000")
	log.Println("📚 API Documentation: http://localhost:3000/docs")
	log.Fatal(app.Listen(":3000"))
}

// HealthCheck godoc
// @Summary      Health check
// @Description  Check if Balancer Studio API is running
// @Tags         system
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Router       /health [get]
func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "ok",
		"service": "balancer-studio",
		"version": "1.0.0",
	})
}

// ProxyHost represents a proxy host configuration
type ProxyHost struct {
	ID          int      `json:"id" example:"1"`
	DomainNames []string `json:"domain_names" example:"example.com,www.example.com"`
	ForwardHost string   `json:"forward_host" example:"192.168.1.100"`
	ForwardPort int      `json:"forward_port" example:"8080"`
	SSLEnabled  bool     `json:"ssl_enabled" example:"true"`
	SSLCertID   *int     `json:"ssl_cert_id,omitempty" example:"1"`
	Enabled     bool     `json:"enabled" example:"true"`
	CreatedAt   string   `json:"created_at" example:"2025-12-08T10:00:00Z"`
}

// ProxyHostRequest represents the request body for creating/updating proxy hosts
type ProxyHostRequest struct {
	DomainNames []string `json:"domain_names" binding:"required" example:"example.com"`
	ForwardHost string   `json:"forward_host" binding:"required" example:"192.168.1.100"`
	ForwardPort int      `json:"forward_port" binding:"required" example:"8080"`
	SSLEnabled  bool     `json:"ssl_enabled" example:"false"`
	SSLCertID   *int     `json:"ssl_cert_id,omitempty" example:"1"`
}

// Certificate represents an SSL certificate
type Certificate struct {
	ID         int    `json:"id" example:"1"`
	Name       string `json:"name" example:"example.com SSL"`
	Provider   string `json:"provider" example:"letsencrypt"`
	DomainName string `json:"domain_name" example:"example.com"`
	ExpiresAt  string `json:"expires_at" example:"2025-12-31T23:59:59Z"`
	Status     string `json:"status" example:"active"`
}

// Upstream represents an upstream server group
type Upstream struct {
	ID          int    `json:"id" example:"1"`
	Name        string `json:"name" example:"backend"`
	Algorithm   string `json:"algorithm" example:"round_robin"`
	Description string `json:"description" example:"Backend application servers"`
}

// UpstreamServer represents a server in an upstream group
type UpstreamServer struct {
	ID       int    `json:"id" example:"1"`
	Host     string `json:"host" example:"192.168.1.100"`
	Port     int    `json:"port" example:"8080"`
	Weight   int    `json:"weight" example:"1"`
	MaxFails int    `json:"max_fails" example:"3"`
	Status   string `json:"status" example:"up"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error" example:"Invalid request"`
	Message string `json:"message" example:"Domain names are required"`
}

// ListProxyHosts godoc
// @Summary      List all proxy hosts
// @Description  Get a list of all configured proxy hosts
// @Tags         proxy-hosts
// @Produce      json
// @Success      200 {array} ProxyHost
// @Router       /proxy-hosts [get]
func ListProxyHosts(c *fiber.Ctx) error {
	// Mock data - replace with database query
	hosts := []ProxyHost{
		{
			ID:          1,
			DomainNames: []string{"example.com", "www.example.com"},
			ForwardHost: "192.168.1.100",
			ForwardPort: 8080,
			SSLEnabled:  true,
			Enabled:     true,
			CreatedAt:   "2025-12-08T10:00:00Z",
		},
		{
			ID:          2,
			DomainNames: []string{"api.example.com"},
			ForwardHost: "192.168.1.101",
			ForwardPort: 3000,
			SSLEnabled:  true,
			Enabled:     true,
			CreatedAt:   "2025-12-08T11:00:00Z",
		},
	}
	return c.JSON(hosts)
}

// CreateProxyHost godoc
// @Summary      Create a new proxy host
// @Description  Create a new proxy host configuration
// @Tags         proxy-hosts
// @Accept       json
// @Produce      json
// @Param        host body ProxyHostRequest true "Proxy Host Configuration"
// @Success      201 {object} ProxyHost
// @Failure      400 {object} ErrorResponse
// @Router       /proxy-hosts [post]
func CreateProxyHost(c *fiber.Ctx) error {
	var req ProxyHostRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(ErrorResponse{
			Error:   "Invalid request",
			Message: err.Error(),
		})
	}

	// Mock response - replace with actual logic
	host := ProxyHost{
		ID:          3,
		DomainNames: req.DomainNames,
		ForwardHost: req.ForwardHost,
		ForwardPort: req.ForwardPort,
		SSLEnabled:  req.SSLEnabled,
		SSLCertID:   req.SSLCertID,
		Enabled:     true,
		CreatedAt:   "2025-12-08T12:00:00Z",
	}

	return c.Status(201).JSON(host)
}

// GetProxyHost godoc
// @Summary      Get a proxy host
// @Description  Get a specific proxy host by ID
// @Tags         proxy-hosts
// @Produce      json
// @Param        id path int true "Proxy Host ID"
// @Success      200 {object} ProxyHost
// @Failure      404 {object} ErrorResponse
// @Router       /proxy-hosts/{id} [get]
func GetProxyHost(c *fiber.Ctx) error {
	_ = c.Params("id")

	// Mock response
	host := ProxyHost{
		ID:          1,
		DomainNames: []string{"example.com"},
		ForwardHost: "192.168.1.100",
		ForwardPort: 8080,
		SSLEnabled:  true,
		Enabled:     true,
		CreatedAt:   "2025-12-08T10:00:00Z",
	}

	return c.JSON(host)
}

// UpdateProxyHost godoc
// @Summary      Update a proxy host
// @Description  Update an existing proxy host configuration
// @Tags         proxy-hosts
// @Accept       json
// @Produce      json
// @Param        id path int true "Proxy Host ID"
// @Param        host body ProxyHostRequest true "Updated Proxy Host Configuration"
// @Success      200 {object} ProxyHost
// @Failure      400 {object} ErrorResponse
// @Failure      404 {object} ErrorResponse
// @Router       /proxy-hosts/{id} [put]
func UpdateProxyHost(c *fiber.Ctx) error {
	_ = c.Params("id")

	var req ProxyHostRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(ErrorResponse{
			Error:   "Invalid request",
			Message: err.Error(),
		})
	}

	// Mock response
	host := ProxyHost{
		ID:          1,
		DomainNames: req.DomainNames,
		ForwardHost: req.ForwardHost,
		ForwardPort: req.ForwardPort,
		SSLEnabled:  req.SSLEnabled,
		SSLCertID:   req.SSLCertID,
		Enabled:     true,
		CreatedAt:   "2025-12-08T10:00:00Z",
	}

	return c.JSON(host)
}

// DeleteProxyHost godoc
// @Summary      Delete a proxy host
// @Description  Delete a proxy host configuration
// @Tags         proxy-hosts
// @Produce      json
// @Param        id path int true "Proxy Host ID"
// @Success      200 {object} map[string]interface{}
// @Failure      404 {object} ErrorResponse
// @Router       /proxy-hosts/{id} [delete]
func DeleteProxyHost(c *fiber.Ctx) error {
	id := c.Params("id")

	return c.JSON(fiber.Map{
		"message": "Proxy host deleted successfully",
		"id":      id,
	})
}

// ListCertificates godoc
// @Summary      List all SSL certificates
// @Description  Get a list of all SSL certificates
// @Tags         certificates
// @Produce      json
// @Success      200 {array} Certificate
// @Router       /certificates [get]
func ListCertificates(c *fiber.Ctx) error {
	certs := []Certificate{
		{
			ID:         1,
			Name:       "example.com SSL",
			Provider:   "letsencrypt",
			DomainName: "example.com",
			ExpiresAt:  "2025-12-31T23:59:59Z",
			Status:     "active",
		},
		{
			ID:         2,
			Name:       "api.example.com SSL",
			Provider:   "letsencrypt",
			DomainName: "api.example.com",
			ExpiresAt:  "2026-01-15T23:59:59Z",
			Status:     "active",
		},
	}
	return c.JSON(certs)
}

// CreateCertificate godoc
// @Summary      Create a new SSL certificate
// @Description  Request a new SSL certificate from Let's Encrypt
// @Tags         certificates
// @Accept       json
// @Produce      json
// @Param        cert body map[string]interface{} true "Certificate Request"
// @Success      201 {object} Certificate
// @Failure      400 {object} ErrorResponse
// @Router       /certificates [post]
func CreateCertificate(c *fiber.Ctx) error {
	return c.Status(201).JSON(Certificate{
		ID:         3,
		Name:       "new-domain.com SSL",
		Provider:   "letsencrypt",
		DomainName: "new-domain.com",
		ExpiresAt:  "2026-12-31T23:59:59Z",
		Status:     "pending",
	})
}

// ListUpstreams godoc
// @Summary      List all upstream groups
// @Description  Get a list of all configured upstream server groups
// @Tags         upstreams
// @Produce      json
// @Success      200 {array} Upstream
// @Router       /upstreams [get]
func ListUpstreams(c *fiber.Ctx) error {
	upstreams := []Upstream{
		{
			ID:          1,
			Name:        "backend",
			Algorithm:   "round_robin",
			Description: "Backend application servers",
		},
		{
			ID:          2,
			Name:        "api_servers",
			Algorithm:   "least_conn",
			Description: "API server pool",
		},
	}
	return c.JSON(upstreams)
}

// CreateUpstream godoc
// @Summary      Create a new upstream group
// @Description  Create a new upstream server group
// @Tags         upstreams
// @Accept       json
// @Produce      json
// @Success      201 {object} Upstream
// @Router       /upstreams [post]
func CreateUpstream(c *fiber.Ctx) error {
	return c.Status(201).JSON(Upstream{
		ID:          3,
		Name:        "new_upstream",
		Algorithm:   "round_robin",
		Description: "New upstream group",
	})
}

// ListUpstreamServers godoc
// @Summary      List servers in an upstream group
// @Description  Get all servers in a specific upstream group
// @Tags         upstreams
// @Produce      json
// @Param        id path int true "Upstream ID"
// @Success      200 {array} UpstreamServer
// @Router       /upstreams/{id}/servers [get]
func ListUpstreamServers(c *fiber.Ctx) error {
	servers := []UpstreamServer{
		{
			ID:       1,
			Host:     "192.168.1.100",
			Port:     8080,
			Weight:   1,
			MaxFails: 3,
			Status:   "up",
		},
		{
			ID:       2,
			Host:     "192.168.1.101",
			Port:     8080,
			Weight:   1,
			MaxFails: 3,
			Status:   "up",
		},
	}
	return c.JSON(servers)
}

// AddUpstreamServer godoc
// @Summary      Add server to upstream group
// @Description  Add a new server to an upstream group
// @Tags         upstreams
// @Accept       json
// @Produce      json
// @Param        id path int true "Upstream ID"
// @Success      201 {object} UpstreamServer
// @Router       /upstreams/{id}/servers [post]
func AddUpstreamServer(c *fiber.Ctx) error {
	return c.Status(201).JSON(UpstreamServer{
		ID:       3,
		Host:     "192.168.1.102",
		Port:     8080,
		Weight:   1,
		MaxFails: 3,
		Status:   "up",
	})
}

// ReloadNginx godoc
// @Summary      Reload Nginx
// @Description  Reload Nginx configuration without downtime
// @Tags         nginx
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      500 {object} ErrorResponse
// @Router       /nginx/reload [post]
func ReloadNginx(c *fiber.Ctx) error {
	// Execute: nginx -s reload
	return c.JSON(fiber.Map{
		"message": "Nginx reloaded successfully",
		"status":  "ok",
	})
}

// TestNginxConfig godoc
// @Summary      Test Nginx configuration
// @Description  Test Nginx configuration for syntax errors
// @Tags         nginx
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} ErrorResponse
// @Router       /nginx/test [post]
func TestNginxConfig(c *fiber.Ctx) error {
	// Execute: nginx -t
	return c.JSON(fiber.Map{
		"message": "Configuration is valid",
		"status":  "ok",
		"output":  "nginx: configuration file /etc/nginx/nginx.conf test is successful",
	})
}

// GetNginxStatus godoc
// @Summary      Get Nginx status
// @Description  Get current Nginx status and metrics
// @Tags         nginx
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Router       /nginx/status [get]
func GetNginxStatus(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"active_connections": 42,
		"accepts":            1234,
		"handled":            1234,
		"requests":           5678,
		"reading":            0,
		"writing":            1,
		"waiting":            41,
		"uptime":             "5 days, 3 hours",
	})
}

// getOpenAPISpec returns the OpenAPI specification
func getOpenAPISpec() map[string]interface{} {
	return map[string]interface{}{
		"openapi": "3.0.0",
		"info": map[string]interface{}{
			"title":       "Balancer Studio API",
			"description": "Professional Nginx management platform with beautiful UI and powerful API. Modern alternative to Nginx Plus.",
			"version":     "1.0.0",
			"contact": map[string]string{
				"name":  "Balancer Studio Support",
				"email": "support@balancer.studio",
			},
			"license": map[string]string{
				"name": "MIT",
				"url":  "https://opensource.org/licenses/MIT",
			},
		},
		"servers": []map[string]string{
			{"url": "http://localhost:3000/api/v1", "description": "Development server"},
		},
		"tags": []map[string]string{
			{"name": "system", "description": "System operations"},
			{"name": "proxy-hosts", "description": "Proxy host management"},
			{"name": "certificates", "description": "SSL certificate management"},
			{"name": "upstreams", "description": "Upstream server management"},
			{"name": "nginx", "description": "Nginx control operations"},
		},
		"paths": map[string]interface{}{
			"/health": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Health check",
					"description": "Check if Balancer Studio API is running",
					"tags":        []string{"system"},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Service is healthy",
							"content": map[string]interface{}{
								"application/json": map[string]interface{}{
									"schema": map[string]interface{}{
										"type": "object",
										"properties": map[string]interface{}{
											"status":  map[string]string{"type": "string", "example": "ok"},
											"service": map[string]string{"type": "string", "example": "balancer-studio"},
											"version": map[string]string{"type": "string", "example": "1.0.0"},
										},
									},
								},
							},
						},
					},
				},
			},
			"/proxy-hosts": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "List all proxy hosts",
					"description": "Get a list of all configured proxy hosts",
					"tags":        []string{"proxy-hosts"},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "List of proxy hosts",
							"content": map[string]interface{}{
								"application/json": map[string]interface{}{
									"schema": map[string]interface{}{
										"type": "array",
										"items": map[string]interface{}{
											"$ref": "#/components/schemas/ProxyHost",
										},
									},
								},
							},
						},
					},
				},
				"post": map[string]interface{}{
					"summary":     "Create a new proxy host",
					"description": "Create a new proxy host configuration",
					"tags":        []string{"proxy-hosts"},
					"requestBody": map[string]interface{}{
						"required": true,
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/ProxyHostRequest",
								},
							},
						},
					},
					"responses": map[string]interface{}{
						"201": map[string]interface{}{
							"description": "Proxy host created",
							"content": map[string]interface{}{
								"application/json": map[string]interface{}{
									"schema": map[string]interface{}{
										"$ref": "#/components/schemas/ProxyHost",
									},
								},
							},
						},
					},
				},
			},
		},
		"components": map[string]interface{}{
			"schemas": map[string]interface{}{
				"ProxyHost": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"id":           map[string]interface{}{"type": "integer", "example": 1},
						"domain_names": map[string]interface{}{"type": "array", "items": map[string]string{"type": "string"}, "example": []string{"example.com"}},
						"forward_host": map[string]interface{}{"type": "string", "example": "192.168.1.100"},
						"forward_port": map[string]interface{}{"type": "integer", "example": 8080},
						"ssl_enabled":  map[string]interface{}{"type": "boolean", "example": true},
						"enabled":      map[string]interface{}{"type": "boolean", "example": true},
						"created_at":   map[string]interface{}{"type": "string", "example": "2025-12-08T10:00:00Z"},
					},
				},
				"ProxyHostRequest": map[string]interface{}{
					"type":     "object",
					"required": []string{"domain_names", "forward_host", "forward_port"},
					"properties": map[string]interface{}{
						"domain_names": map[string]interface{}{"type": "array", "items": map[string]string{"type": "string"}, "example": []string{"example.com"}},
						"forward_host": map[string]interface{}{"type": "string", "example": "192.168.1.100"},
						"forward_port": map[string]interface{}{"type": "integer", "example": 8080},
						"ssl_enabled":  map[string]interface{}{"type": "boolean", "example": false},
					},
				},
			},
		},
	}
}
