**Note API**

A Note Management API with user authentication. 

Users can register, login, and manage their notes (create, read, update, delete) using JWT authentication.

**Features**

    User Authentication
      Register new users
      Login with hashed password
      Receive JWT token on successful login
    
    Notes Management
      Create a note
      Get a note by ID
      Get all notes
      Update a note
      Delete a note
    
    Secure
      	Passwords are hashed in the database
      	JWT token required for all note operations
    
    Configuration
    	  Supports configuration via .env file or config folder

**Project Structure**

    cmd/
      └─ server/
          └─ main.go          # Entry point
    internal/
      ├─ config/             # App configuration
      ├─ database/           # DB connection setup
      ├─ handler/            # HTTP handlers
      ├─ middleware/         # Auth middleware
      ├─ model/              # DB models
      ├─ repository/         # DB queries
      └─ service/            # Business logic
    pkg/
      ├─ jwt/                # JWT generation and validation
      └─ hash/               # Password hashing utilities

**Installation**

Clone the repository
        
    git clone https://github.com/Yogi-1996/notes-api.git
    cd notes-api

Set up your configuration:

  Create a .env file or use the config folder to configure:
  
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=user
    DB_PASSWORD=password
    DB_NAME=notesdb
    JWT_SECRET=your_secret_key
    PORT=8080

Install dependencies:

    go mod tidy

Run the server:

    go run cmd/server/main.go
