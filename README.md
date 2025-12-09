# Balancer Studio

> Professional Nginx management platform with beautiful UI and powerful API

**Balancer Studio** is a modern alternative to Nginx Plus, providing an intuitive web interface, comprehensive REST API, and advanced load balancing capabilities.

## ✨ Features

- 🎨 **Beautiful Web Interface** — Intuitive dashboard for management
- 🚀 **Powerful REST API** — Full automation through API
- 📚 **Excellent Documentation** — Interactive API docs via Scalar UI
- 🔒 **SSL/TLS Management** — Let's Encrypt integration
- 📊 **Real-time Monitoring** — Live metrics and statistics
- ⚡ **Fast & Lightweight** — Written in Go
- 🐳 **Docker Ready** — Easy container deployment

## 🚀 Quick Start

### Prerequisites

- Go 1.21 or higher
- Nginx (optional for testing)

### Installation

```bash
# Clone the repository
git clone https://github.com/VladislavUsenko/balancer-studio.git
cd balancer-studio

# Install dependencies
go mod download

# Run the server
go run main.go
```

Server will start on `http://localhost:3000`

### 📚 API Documentation

**Scalar UI (Interactive documentation):**
```
http://localhost:3000/docs
```

**OpenAPI Specification:**
```
http://localhost:3000/swagger.json
```

## 🎯 Features

### ✅ Implemented (v1.0)

- [x] REST API with Fiber
- [x] Interactive documentation (Scalar UI)
- [x] CRUD for Proxy Hosts
- [x] SSL Certificate management
- [x] Upstream server management
- [x] Nginx control (reload, test, status)
- [x] Health check monitoring

### 🔨 In Development

- [ ] PostgreSQL integration
- [ ] JWT authentication and RBAC
- [ ] Nginx config parsing and generation
- [ ] Let's Encrypt automation
- [ ] Real-time metrics and charts
- [ ] React web interface
- [ ] Rate limiting
- [ ] Access control lists
- [ ] WebSocket for real-time updates

## 📖 API Endpoints

### System
- `GET /api/v1/health` - Health check

### Proxy Hosts
- `GET /api/v1/proxy-hosts` - List proxy hosts
- `POST /api/v1/proxy-hosts` - Create proxy host
- `GET /api/v1/proxy-hosts/:id` - Get proxy host
- `PUT /api/v1/proxy-hosts/:id` - Update proxy host
- `DELETE /api/v1/proxy-hosts/:id` - Delete proxy host

### SSL Certificates
- `GET /api/v1/certificates` - List certificates
- `POST /api/v1/certificates` - Create certificate

### Upstream Servers
- `GET /api/v1/upstreams` - List upstream groups
- `POST /api/v1/upstreams` - Create upstream group
- `GET /api/v1/upstreams/:id/servers` - List servers in group
- `POST /api/v1/upstreams/:id/servers` - Add server to group

### Nginx Control
- `POST /api/v1/nginx/reload` - Reload Nginx
- `POST /api/v1/nginx/test` - Test configuration
- `GET /api/v1/nginx/status` - Get status and metrics

## 🧪 Usage Examples

### Create Proxy Host

```bash
curl -X POST http://localhost:3000/api/v1/proxy-hosts \
  -H "Content-Type: application/json" \
  -d '{
    "domain_names": ["example.com", "www.example.com"],
    "forward_host": "192.168.1.100",
    "forward_port": 8080,
    "ssl_enabled": true
  }'
```

### Add Server to Upstream

```bash
curl -X POST http://localhost:3000/api/v1/upstreams/1/servers \
  -H "Content-Type: application/json" \
  -d '{
    "host": "192.168.1.102",
    "port": 8080,
    "weight": 1
  }'
```

### Get Nginx Status

```bash
curl http://localhost:3000/api/v1/nginx/status
```

## 🏗️ Project Structure

```
balancer-studio/
├── main.go              # Entry point with API handlers
├── go.mod               # Go dependencies
├── README.md            # Documentation
├── internal/            # Private code (planned)
│   ├── api/            # HTTP handlers
│   ├── config/         # Nginx config management
│   ├── models/         # Data models
│   ├── repository/     # Database layer
│   └── service/        # Business logic
├── pkg/                # Public libraries (planned)
├── web/                # React UI (planned)
├── migrations/         # DB migrations (planned)
└── docker/             # Docker configs (planned)
```

## 🔧 Tech Stack

**Backend:**
- Go 1.21+
- Fiber (web framework)
- GORM (ORM, planned)
- PostgreSQL (planned)

**Frontend (planned):**
- React 18 + TypeScript
- Tailwind CSS + shadcn/ui
- TanStack Query
- Vite

**Documentation:**
- Scalar UI
- OpenAPI 3.0

## 💡 Development

### Run in dev mode

```bash
# With hot reload
go install github.com/cosmtrek/air@latest
air
```

### Build for production

```bash
go build -o balancer-studio main.go
./balancer-studio
```

### Docker (planned)

```bash
docker build -t balancer-studio .
docker run -p 3000:3000 balancer-studio
```

## 🆚 Comparison

| Feature | Balancer Studio | Nginx Plus | NginxProxyManager |
|---------|-----------------|------------|-------------------|
| Price | Free/Open Source | $2500+/year | Free |
| API | ✅ REST + Docs | ✅ | ⚠️ Incomplete |
| UI | 🔨 In Development | ❌ | ✅ |
| API Documentation | ✅ Scalar | ✅ | ⚠️ Incomplete |
| Support | Community | Enterprise | Community |
| Open Source | ✅ | ❌ | ✅ |

## 🤝 Contributing

We welcome contributions! Please:

1. Fork the project
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

MIT License - free to use in commercial projects

## 🌟 Roadmap

### Q1 2026
- [ ] PostgreSQL integration
- [ ] JWT authentication
- [ ] Nginx config parsing
- [ ] React UI (MVP)

### Q2 2026
- [ ] Let's Encrypt integration
- [ ] Real-time monitoring
- [ ] Docker images
- [ ] Kubernetes operator

### Q3 2026
- [ ] Multi-server support
- [ ] Backup/restore
- [ ] Advanced analytics
- [ ] Plugin system

## 📧 Contact

- **Website:** https://balancer.studio (planned)
- **Email:** support@balancer.studio
- **GitHub:** [@yourusername/balancer-studio](https://github.com/VladislavUsenko/balancer-studio)

---

**Balancer Studio** — Manage Nginx like a pro 🚀

Made with ❤️ and Go