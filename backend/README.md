# fasms

# Project Name

This project consists of a backend HTTP server written in Go (Golang) and a frontend application built with ReactJS. The frontend is still a work in progress, while the backend API is fully functional.

## Folder Structure

- **`backend/`**: Contains the Go backend server and all associated logic.
- **`frontend/`**: Contains the ReactJS frontend application (work in progress).

## Prerequisites

1. **Go 1.22.7** (specifically used for this project)
2. **MySQL/PostgreSQL** or any other supported database.

## Backend Setup

### 1. Download Go 1.22.7
Make sure you have Go version `1.22.7` installed on your machine. You can download it from the official Go website:

- [Go 1.22.7 Download](https://go.dev/dl/go1.22.7.linux-amd64.tar.gz) *(Adjust for your operating system)*

After downloading, follow the installation instructions specific to your operating system.

### 2. Set Up the Database

1. **Run `initTable.sql`**:
   You need to load some initial data into your local database. The `initTable.sql` script is provided in the project. Run this SQL script to set up your database schema and seed the necessary tables.

2. **Configure Database Connection**:
   - Open the file `backend/internal/db/db.go`.
   - Locate the following lines:

   ```go
   dsn, err := utilities.GetDatabaseUrl()
   if err != nil {
       log.Fatal("Unable to get dsn", err)
   }
    ```

   - Replace the `dsn` with your own database connection string in the following manner:

     - For MySQL:

       ```go
       dsn := "user:password@tcp(localhost:3306)/dbname"
       ```

     - For PostgreSQL:

       ```go
       dsn := "user=youruser password=yourpassword dbname=yourdbname sslmode=disable"
       ```

     Ensure to replace `user`, `password`, `localhost:3306`, and `dbname` with your actual database details.

### 3. Run the Backend

Once the database is set up and your connection string is updated:

1. Open a terminal and navigate to the `backend/` directory.
2. Run the following command to start the server:

   ```bash
   go run ./cmd/main.go
   ```
3. The backend server should now be running on http://localhost:8080.

#### 4. Access Swagger documentation

To interact with and test the backend API, you can access the Swagger page by navigating to:

    ```
    http://localhost:8080/swagger/
    ```
This will provide you with a user interface for the API, including details about each endpoint 
and the ability to make requests directly from the browser.