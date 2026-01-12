# KV-Go Storage Engine 🏎️

A simple and reliable key-value database written in **Go**. This project was created as part of an in-depth study of Go architecture, multithreading, and working with file systems.

## 🚀 Features

* **Thread-Safety**: Full support for multithreading using `sync.RWMutex`. Safe reading and writing from different goroutines.
* **Performance**: Data is cached in RAM (Map), which provides instant access via the `GET` command.
* **DevOps-Ready**: Configuration via environment variables (`.env`) and Docker support (planned).
* **High Coverage**: Robust reliability with **92.5% unit test coverage**.
* **Tombstone Deletion**: Professional-grade deletion logic. Instead of slow file rewrites, it uses fast append-only "tombstone" markers.

## 🛠 Commands

| Command | Description | Example |
| :--- | :--- | :--- |
| **SET** | Save the value by key | `SET user_1 ivan` |
| **GET** | Get value by key | `GET user_1` |
| **LIST** | Show all entries in the database | `LIST` |
| **DELETE** | Remove key (Append-only) | `DELETE user_1` |
| **EXIT** | Finish work | `EXIT` |

## 📦 How to start

1. **Clone the repository:**
   ```bash
   git clone [https://github.com/Zapi-web/kv-go-db.git](https://github.com/Zapi-web/kv-go-db.git)

2. **Configure environment variables:** Create an .env file and specify the path to the database:
    ```Plaintext
    FILEPATH=data.db
3. **Start the app**
    ```bash
    go run main.go

## 👨‍💻 About me

I am 16 years old, I am from Kharkiv, and I currently live in Vienna. My goal is to become a top DevOps engineer and enroll at TU Wien. This project is part of my journey to master Go and system programming. I believe in discipline and continuous development.