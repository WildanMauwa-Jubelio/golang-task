# Golang Task Project

Projek pembelajaran bahasa Go dengan implementasi 5 tasks yang mencakup web server, cron jobs, inventory management, programming fundamentals, dan concurrent programming.

## 🏗️ Struktur Project

```
golang-task/
├── cmd/                     # Main applications (executable commands)
│   ├── webserver/           # Task 1: Fiber web server
│   │   └── main.go
│   ├── cronjob/             # Task 2: Cron job scheduler  
│   │   └── main.go
│   ├── inventory/           # Task 3: Inventory management demo
│   │   └── main.go
│   ├── basics/              # Task 4: Programming basics demo
│   │   └── main.go
│   └── aggregator/          # Task 5: Concurrent product aggregator
│       └── main.go
├── pkg/                     # Library packages (reusable code)
│   ├── inventory/           # Inventory management library
│   │   └── inventory.go
│   └── fetcher/             # Product fetcher library
│       └── fetcher.go
├── docs/                    # Documentation
│   └── dokumentasi-pengerjaan-tugas.md
├── internal/                # Private application code
├── main.go                  # Task launcher/menu system
├── go.mod                   # Go module declaration
├── go.sum                   # Go dependencies checksums
└── README.md                # This file
```

## 🚀 Quick Start

### Prerequisites
- Go 1.24.0 atau lebih tinggi
- Internet connection (untuk Task 5)

### Installation & Setup

1. **Clone/Download project**
   ```bash
   cd golang-task
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Run task launcher**
   ```bash
   go run main.go
   ```

## 📋 Available Tasks

### 1. Web Server (Fiber) - `cmd/webserver/`
- **Deskripsi**: Web server menggunakan Fiber framework dengan health endpoint
- **Port**: 9200
- **Endpoints**:
  - `GET /` - Root endpoint
  - `GET /health` - Health check endpoint
- **Cara menjalankan**:
  ```bash
  cd cmd/webserver
  go run main.go
  ```

### 2. Cron Job Scheduler - `cmd/cronjob/`
- **Deskripsi**: Background job scheduler yang berjalan setiap 10 detik
- **Features**: Graceful shutdown dengan Ctrl+C
- **Cara menjalankan**:
  ```bash
  cd cmd/cronjob
  go run main.go
  ```

### 3. Inventory Management System - `cmd/inventory/`
- **Deskripsi**: Sistem manajemen inventori dengan struct dan methods
- **Features**: 
  - Add/update items
  - Calculate total value
  - Print inventory reports
- **Package**: `pkg/inventory/`
- **Cara menjalankan**:
  ```bash
  cd cmd/inventory
  go run main.go
  ```

### 4. Programming Basics - `cmd/basics/`
- **Deskripsi**: Demo konsep dasar Go programming
- **Topics**:
  - **Loops**: FizzBuzz, slice summation
  - **Conditionals**: Switch statements, type inspection
  - **Functions**: Multiple returns, variadic functions
- **Cara menjalankan**:
  ```bash
  cd cmd/basics
  go run main.go
  ```

### 5. Concurrent Product Aggregator - `cmd/aggregator/`
- **Deskripsi**: Concurrent data fetching menggunakan goroutines dan channels
- **Features**:
  - Fetch data dari dummyjson.com API
  - Concurrent processing dengan goroutines
  - Error handling dan performance metrics
- **Package**: `pkg/fetcher/`
- **Cara menjalankan**:
  ```bash
  cd cmd/aggregator
  go run main.go
  ```

## 🏛️ Architecture & Design Patterns

### Project Structure Pattern
Mengikuti **Go Standard Project Layout**:

- **`cmd/`**: Main applications untuk project ini. Directory name untuk setiap aplikasi sesuai dengan nama executable
- **`pkg/`**: Library code yang bisa digunakan oleh aplikasi eksternal
- **`internal/`**: Private application dan library code
- **`docs/`**: Design dan user documents

### Code Organization
- **Separation of Concerns**: Setiap task dipisah ke dalam module tersendiri
- **Reusable Packages**: Logic bisnis dipisah ke dalam `pkg/` untuk reusability
- **Single Responsibility**: Setiap file/package memiliki tanggung jawab spesifik

## 🔧 Dependencies

```go
require (
    github.com/gofiber/fiber/v2 v2.52.9    // Web framework
    github.com/robfig/cron/v3 v3.0.1       // Cron job scheduler
)
```

## 🏃‍♂️ Running Individual Tasks

Setiap task dapat dijalankan secara independen:

```bash
# Task 1: Web Server
cd cmd/webserver && go run main.go

# Task 2: Cron Job  
cd cmd/cronjob && go run main.go

# Task 3: Inventory Management
cd cmd/inventory && go run main.go

# Task 4: Programming Basics
cd cmd/basics && go run main.go

# Task 5: Concurrent Aggregator
cd cmd/aggregator && go run main.go
```

## 📊 Build & Deployment

### Build All Applications
```bash
# Build semua applications
for dir in cmd/*/; do
    cd "$dir"
    app_name=$(basename "$dir")
    go build -o "../../bin/$app_name" .
    cd ../..
done
```

### Build Specific Task
```bash
# Contoh: Build web server
cd cmd/webserver
go build -o webserver .
```

## 🧪 Testing

```bash
# Run tests untuk semua packages
go test ./...

# Test specific package
go test ./pkg/inventory
go test ./pkg/fetcher
```

## 📚 Learning Objectives

Project ini dirancang untuk mempelajari:

1. **Go Basics**: Variables, functions, structs, methods
2. **Web Development**: HTTP servers dengan Fiber framework
3. **Concurrency**: Goroutines, channels, sync.WaitGroup
4. **Package Management**: Go modules, dependencies
5. **Project Organization**: Standard Go project layout
6. **Error Handling**: Proper error handling patterns
7. **JSON Processing**: Encoding/decoding JSON
8. **Background Jobs**: Cron scheduling
9. **API Integration**: HTTP clients, external API calls

## 🛠️ Development

### Adding New Tasks
1. Create new directory in `cmd/`
2. Add main.go with your implementation
3. Update task list in root `main.go`
4. Add documentation

### Adding Reusable Packages
1. Create new package in `pkg/`
2. Implement your library code
3. Use in cmd applications

## 📝 Notes

- Semua emoji telah dihapus dari code untuk tampilan yang lebih professional
- Project menggunakan Go 1.24.0 
- Mengikuti Go coding standards dan best practices
- Mendukung graceful shutdown untuk long-running services
- Error handling yang comprehensive di semua components

## 📖 Additional Documentation

Lihat `docs/dokumentasi-pengerjaan-tugas.md` untuk dokumentasi lengkap implementasi setiap task.