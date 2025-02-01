# Web Forum Project

## Introduction
This project is a fully functional web forum that allows users to communicate, share posts, and interact with each other. The forum supports user authentication, post categorization, likes/dislikes, and filtering mechanisms. We utilized SQLite for database management and implemented Docker for containerization.


## Project Structure
```
FORUM
 ┣ app
 ┃  ┣ http
 ┃  ┃  ┣ controllers
 ┃  ┃  ┃  ┣ commentController.go
 ┃  ┃  ┃  ┣ cssJs.go
 ┃  ┃  ┃  ┣ errController.go
 ┃  ┃  ┃  ┣ homeController.go
 ┃  ┃  ┃  ┣ loginController.go
 ┃  ┃  ┃  ┣ logout.go
 ┃  ┃  ┃  ┣ parseFileController.go
 ┃  ┃  ┃  ┣ postController.go
 ┃  ┃  ┃  ┣ reactController.go
 ┃  ┃  ┃  ┗ registerController.go
 ┃  ┃  ┗ middleware
 ┃  ┃     ┣ authorization.go
 ┃  ┃     ┗ ratelimited.go
 ┃  ┗ models
 ┃     ┣ category.go
 ┃     ┣ comment.go
 ┃     ┣ data.go
 ┃     ┣ database.go
 ┃     ┣ error.go
 ┃     ┣ post.go
 ┃     ┣ react.go
 ┃     ┣ session.go
 ┃     ┗ user.go
 ┣ database
 ┃  ┣ schema
 ┃  ┃  ┗ schema.sql
 ┃  ┗ database.db
 ┣ resources
 ┃  ┣ css
 ┃  ┣ imgs
 ┃  ┣ js
 ┃  ┗ views
 ┣ routes
 ┃  ┣ api.go
 ┃  ┗ web.go
 ┣ utils
 ┃  ┗ validation.go
 ┣ Dockerfile
 ┣ go.mod
 ┣ go.sum
 ┣ main.go
 ┗ README.md
```


## Features
### 1. User Authentication
- Users can register with an email, username, and password.
- Unique email validation to prevent duplicates.
- Password encryption (Bonus feature).
- Secure session management using cookies.
- Users can log in and log out securely.

### 2. Posts & Comments
- Registered users can create posts.
- Posts can be categorized under different topics.
- Users can comment on posts.
- All posts and comments are publicly visible.

### 3. Reactions (Likes & Dislikes)
- Registered users can like or dislike posts and comments.
- Like/dislike counts are visible to all users.

### 4. Post Filtering
- Users can filter posts based on:
  - Categories
  - Created posts (for logged-in users)
  - Liked posts (for logged-in users)

### 5. Database
- SQLite is used for data storage.
- The database schema includes tables for users, posts, comments, categories, and reactions.
- The implementation includes at least one `SELECT`, `CREATE`, and `INSERT` query.

### 6. Error Handling
- Proper handling of HTTP errors.
- Input validation for user credentials and form submissions.
- Technical errors are managed gracefully.

### 7. Docker Integration
- The application is containerized using Docker.
- Ensures environment consistency across deployments.

## Technologies Used
- **Go** (Backend Development)
- **HTML, CSS, JavaScript** (Frontend)
- **SQLite** (Database Management)
- **Docker** (Containerization)
- **bcrypt** (Password Encryption - Bonus)
- **UUID** (Unique User Sessions - Bonus)

## Installation & Setup
### Prerequisites
- Go installed on your system
- Docker installed

### Running the Project
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd forum_project
   ```
2. Build and run the Docker container:
   ```bash
   docker build -t forum .
   docker run -p 8080:8080 forum
   ```
3. Access the forum at `http://localhost:8080`

## Future Enhancements
- Implement user profiles.
- Add private messaging between users.
- Enhance UI/UX for better usability.
- Implement additional security features.

## Contributors

[@fahdaguenouz]
[@mohamedelfarssi]
[@mohamedseffine]
[@hatimtahiri]
[@hamzakoki]



##
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes
