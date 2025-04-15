# BackendDbProject
Users can register, obtain JWT tokens for authentication, and perform CRUD operations on posts.  
MySQL is used for storage, with included migration scripts for setup.
## Setup
1. Clone repository
```sh
git clone 
cd personal_blog
```
2. Install Dependencies
 ```sh
go mod tidy
```
3. Database Configuration  
* Set up MySQL database.  
* Configure connection details in internal/config/env.go  
4. Run Migrations
 ```sh
make migrate-up
```
5. Start application
 ```sh
make run
```
## API Endpoints
POST /api/v1/register - Register a new user.  
POST /api/v1/login - Obtain a JWT token for authentication.
