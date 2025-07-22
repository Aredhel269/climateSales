# 🌱 EcoHort

**EcoHort** is a desktop application developed in Go to query and store weather data for organic gardens. It uses the [Fyne](https://fyne.io/) library for the graphical interface, connects to the AEMET API to obtain weather data, and stores information locally using SQLite.

This project was created as part of the course **"Go for programmers: build your own executable application"** taught by **Oriol Tinoco at Cibernàrium – Barcelona Activa**.

---

## ✨ Main Features

- Query the **weather forecast** by municipality.
- Integration with the **AEMET API** (API key and municipality management).
- Local storage with **SQLite**.
- Add, view, update, and delete manual weather records (date, precipitation, max/min temperature, humidity).
- Display records in a **table** and a **graph**.
- Persistent **user preferences** configuration for municipality and API key.
- Fixed-size graphical interface (900x700 pixels).

---

## 🧰 Technologies Used

| Technology          | Purpose                                  |
|---------------------|------------------------------------------|
| Go 1.19+            | Main programming language                |
| [Fyne v2.3.0](https://fyne.io/) | GUI library for desktop applications    |
| SQLite              | Local database                           |
| AEMET API           | Official source of meteorological data  |

---

## 🚀 How to Run the Project

### Prerequisites

- Go installed (version 1.19 or higher)  
- Free API key from the [AEMET API](https://opendata.aemet.es/centrodedescargas/inicio)

### Steps

```bash
git clone https://github.com/your-username/ecohort.git
cd ecohort
go mod tidy
go run main.go
````

To build the executable:

```bash
go build -o ecohort.exe
```

---

## ⚙️ Configuration

The application stores the municipality and API key through Fyne’s preference system. By default, the municipality is `08001`. You can modify these preferences inside the app.

To change the location of the SQLite database file, set the environment variable:

```bash
export DB_PATH=/path/to/sql.db    # Linux/macOS
$env:DB_PATH="C:\path\to\sql.db" # Windows PowerShell
```

---

## 🗂️ Project Structure

```
ecohort/
│
├── main.go                        # Entry point and general configuration
├── go.mod / go.sum                # Dependency management
│
├── repository/
│   ├── repository.go              # Interface and data model
│   └── db-sqlite.go               # SQLite implementation (migration, CRUD)
```

---

## 📦 Data Model and Access Layer

### `Registres` Model

Weather records have the following structure:

| Field          | Type        | Description           |
| -------------- | ----------- | --------------------- |
| `ID`           | `int64`     | Unique identifier     |
| `Data`         | `time.Time` | Record date           |
| `Precipitacio` | `int`       | Precipitation (mm)    |
| `TempMaxima`   | `int`       | Max temperature (°C)  |
| `TempMinima`   | `int`       | Min temperature (°C)  |
| `Humitat`      | `int`       | Relative humidity (%) |

### `Repository` Interface

Defines CRUD operations for weather records:

* `Migrate() error` – Creates the `registres` table.
* `InsertRegistre(Registres) (*Registres, error)` – Inserts a new record.
* `ObtenirTotsRegistres() ([]Registres, error)` – Returns all records.
* `ObtenirRegistrePerID(id int) (*Registres, error)` – Returns a record by ID.
* `ActualitzarRegistre(id int64, Registres) error` – Updates a record by ID.
* `BorrarRegistre(id int64) error` – Deletes a record by ID.

### `SQLiteRepository` Implementation

Implements the `Repository` interface using SQLite with these characteristics:

* Stores dates as Unix timestamps (`int64`) in the database.
* Handles errors and checks affected rows on updates and deletes.
* Performs automatic migrations to create the table if it does not exist.

---

## 👩‍🏫 Credits

Created as part of the course
**"Go for programmers: build your own executable application"**
📍 Cibernàrium – Barcelona Activa, October 2024

---

## 📄 License

This project is licensed under the **MIT License**. See the `LICENSE` file for details.
