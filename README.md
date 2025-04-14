# BackendDbProject
Personal Blog REST API
Users can register, obtain JWT tokens for authentication, and perform CRUD operations on posts.
MySQL is used for storage, with included migration scripts for setup.

Setup
Clone repository
git clone 
cd personal_blog
Install Dependencies
go mod tidy
Database Configuration
Set up MySQL database.
Configure connection details in internal/config/env.go
Run Migrations
make migrate-up
Start application
make run
API Endpoints
POST /api/v1/register - Register a new user.
POST /api/v1/login - Obtain a JWT token for authentication.

