# 🏗️ Project Structure (Clean Architecture + Scalable)

Here’s a **battle-tested layout**:

```
url-shortener/
│
├── cmd/                        # Entry points (main apps)
│   └── api/
│       └── main.go
│
├── internal/                   # Private application code
│   ├── config/                # Config loading (env, yaml)
│   │   └── config.go
│   │
│   ├── domain/                # Core business models
│   │   └── url.go
│   │
│   ├── repository/            # DB interactions
│   │   ├── sql/
│   │   │   └── url_repository.go
│   │   └── redis/
│   │       └── cache.go
│   │
│   ├── service/               # Business logic
│   │   └── url_service.go
│   │
│   ├── handler/               # HTTP handlers (controllers)
│   │   └── url_handler.go
│   │
│   ├── middleware/            # Logging, auth, rate limiting
│   │   └── rate_limiter.go
│   │
│   └── utils/                 # Helpers (hashing, encoding)
│       └── shortener.go
│
├── pkg/                       # Reusable packages (optional)
│   └── logger/
│       └── logger.go
│
├── migrations/                # SQL migrations
│   └── 001_create_urls_table.sql
│
├── scripts/                   # Dev scripts
│   └── run.sh
│
├── deployments/               # Docker, K8s, etc.
│   ├── Dockerfile
│   └── docker-compose.yml
│
├── .env
├── go.mod
└── README.md
```

---

# 🔥 Folder Responsibilities (CRYSTAL CLEAR)

### `cmd/api/main.go`

* App entry point
* Wire everything together (DI)

---

### `internal/domain`

Define core models:

```go
type URL struct {
    ID        string
    LongURL   string
    ShortCode string
    CreatedAt time.Time
}
```

---

### `internal/repository`

#### SQL Repo

* Save URLs
* Fetch URLs

#### Redis Repo

* Cache lookup
* TTL management

---

### `internal/service`

This is the **brain** 🧠

```go
type URLService struct {
    repo  URLRepository
    cache Cache
}

func (s *URLService) Shorten(url string) (string, error)
func (s *URLService) Resolve(code string) (string, error)
```

---

### `internal/handler`

HTTP layer:

```go
POST /shorten
GET  /{shortCode}
```

---

### `internal/middleware`

* Rate limiting (Redis 🔥)
* Logging
* Recovery

---

### `internal/utils`

* Base62 encoding (important!)
* Hashing

---

# 🗄️ Database Design (SQL)

### URLs Table

```sql
CREATE TABLE urls (
    id UUID PRIMARY KEY,
    long_url TEXT NOT NULL,
    short_code VARCHAR(10) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

# ⚡ Redis Usage

Use Redis for:

### 1. Caching

```
key: shortCode
value: longURL
```

### 2. Rate Limiting

```
key: rate_limit:user_ip
```

---

# 🔄 Request Flow (IMPORTANT)

### Shorten URL

```
Client → Handler → Service → Repo (SQL)
                         → Cache (Redis)
```

---

### Redirect

```
Client → Handler
        → Redis (fast lookup)
            → hit → return
            → miss → SQL → cache → return
```

---

# 🧪 Tech Choices (Recommended)

* Router: `chi` or `gin` (go with **chi** for clean design)
* ORM: `sqlx` or raw SQL (prefer raw for control)
* Redis: `go-redis`
* Config: `viper` or env-based

---

# 🐳 Docker Setup (Basic)

### docker-compose.yml

```yaml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    depends_on:
      - postgres
      - redis

  postgres:
    image: postgres:15
    environment:
      POSTGRES_DB: url_shortener
      POSTGRES_USER: admin
      POSTGRES_PASSWORD: admin
    ports:
      - "5432:5432"

  redis:
    image: redis:7
    ports:
      - "6379:6379"
```