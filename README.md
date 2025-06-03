# User & Actions Sample API (Go)


## Live Demo (deployed with Railway ❤️): https://useractionssamplego-production.up.railway.app/

---

## Features

- **User Management**:
    - Retrieve user details by ID.
    - Count actions performed by a user.

- **Action Management**:
    - Retrieve probabilities for the next action based on type.

- **Referral Index**:
    - Calculate referral index based on user data.

---

## Project Structure

The project is organized into the following packages:

- **`handlers`**:
  Contains HTTP handlers for processing API requests.

- **`middleware`**:
  Implements http middlewares for better logging and JSON response headers.

- **`models`**:
  Defines the data structures for users and actions.

- **`repository/memory`**:
  Implements in-memory repositories for users and actions.

- **`utils`**:
  Contains utility functions for tasks like computing the referral index and the next action probabilities.
---

## Endpoints

### 1. **Get User**
- **URL**: `/users/{id}`
- **Method**: `GET`
- **Description**: Retrieves user details by ID.
- **Response**:
    - `200 OK`: User details.
    - `404 Not Found`: User not found.

### 2. **Get User Action Count**
- **URL**: `/users/{id}/actions/count`
- **Method**: `GET`
- **Description**: Retrieves the count of actions performed by a user.
- **Response**:
    - `200 OK`: Action count.
    - `404 Not Found`: User not found.

### 3. **Get Next Action Probabilities**
- **URL**: `/actions/next/{type}`
- **Method**: `GET`
- **Description**: Retrieves probabilities for the next action based on type.
- **Response**:
    - `200 OK`: Probabilities.
    - `400 Bad Request`: Invalid action type.

### 4. **Get Referral Index**
- **URL**: `/referral_index`
- **Method**: `GET`
- **Description**: Calculates referral index based on user data.
- **Response**:
    - `200 OK`: Referral index.

---

## Installation

1. Clone the repository:
```bash
git clone https://github.com/rodrigobressan/surfe_assignment.git
cd surfe_assignment
```
   
2. Install dependencies:  
> go mod tidy

3. Run the server:  
> go run main.go

## Improvements:

- Using Swagger for API documentation (https://github.com/swaggo/swag)
- Add more unit tests (for handlers and repositories)
- Improve error handling