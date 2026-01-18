# KV-Go Storage Engine 🏎️

A simple and reliable key-value database written in **Go**. This project was created as part of an in-depth study of Go architecture, multithreading, and working with file systems.

## 🚀 Features

* **Thread-Safety**: Full support for multithreading using `sync.RWMutex`. Safe reading and writing from different goroutines.
* **Optimized I/O**: High performance achieved by maintaining an open file descriptor.
* **DevOps-Ready**: Configuration via environment variables (`.env`) and Docker support (planned).
* **High Coverage**: Robust reliability with **89% unit test coverage**.
* **Observability**: Structured JSON logging with `zap` and latency monitoring.

## 📈 Performance (Benchmark Results)
The engine provides a clear trade-off between reliability and speed (Tested on AMD Ryzen 5 5600.):
- **Durable writes**: 1.5k ops/sec.
- **Async writes**: 300k+ ops/sec.

## 🛠 Commands

| Command | Description | Example |
| :--- | :--- | :--- |
| **set** | Save the value by key | `./db set user_1 ivan` |
| **get** | Get value by key | `./db get user_1` |
| **list** | Show all entries in the database | `./db list` |
| **delete** | Remove key (Append-only) | `./db delete user_1` |

## 📦 How to start

1. **Clone the repository:**
   ```bash
    git clone https://github.com/Zapi-web/kv-go-db.git
    cd kv-go-db
    go build -o db

2. **Configure environment variables:** Create an .env file and specify the path to the database:
    ```Plaintext
    FILEPATH=data.db
3. **Usage:** You can now run commands directly from your terminal:
    ```bash
    # Save data
    ./db set mykey "Hello World"

    # Retrieve data
    ./db get mykey

    # List all records
    ./db list

## 👨‍💻 About me

I am 16 years old, I am from Kharkiv, and I currently live in Vienna. My goal is to become a top DevOps engineer and enroll at TU Wien. This project is part of my journey to master Go and system programming. I believe in discipline and continuous development.