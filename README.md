# **API Documentation: Simple User Management API**

## **1. Overview**
This API allows users to **store** and **retrieve** user information using a **unique ID**. The data is stored in a JSON file, ensuring persistence across API restarts.

- **Base URL:** `http://localhost:8080`
- **Content-Type:** `application/json`
- **Authentication:** None (Open API)
- **Data Persistence:** Stored in `data.json`

---

## **2. Endpoints**

### **🔹 1. Create a User (POST /post)**
#### **Request**
- **Method:** `POST`
- **URL:** `/post`
- **Headers:**
  - `Content-Type: application/json`
- **Body:**
  ```json
  {
    "name": "Sheila",
    "age": 25
  }
  ```

#### **Response**
- **Success (`200 OK`)**
  ```json
  {
    "id": "550e8400-e29b-41d4-a716-446655440000"
  }
  ```
- **Failure (`400 Bad Request`)** – If invalid JSON is sent.
  ```json
  {
    "error": "Invalid JSON"
  }
  ```

#### **Example cURL Request**
```sh
curl -X POST http://localhost:8080/post \
     -H "Content-Type: application/json" \
     -d '{"name": "Sheila", "age": 25}'
```

---

### **🔹 2. Retrieve a User (GET /get/:id)**
#### **Request**
- **Method:** `GET`
- **URL:** `/get/:id`
- **Headers:** None

#### **Response**
- **Success (`200 OK`)**
  ```json
  {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "Sheila",
    "age": 25
  }
  ```
- **Failure (`404 Not Found`)** – If ID does not exist.
  ```json
  {
    "error": "Entry not found"
  }
  ```

#### **Example cURL Request**
```sh
curl -X GET http://localhost:8080/get/550e8400-e29b-41d4-a716-446655440000
```

---

## **3. Running the API**
### **🔹 1. Install Dependencies**
Ensure you have **Go** installed (version 1.19+).
```sh
go mod tidy
```

### **🔹 2. Run the API**
```sh
go run main.go
```

### **🔹 3. Run Tests**
#### **Go Unit Tests**
```sh
go test -v
```

#### **Automated API Tests**
```sh
chmod +x test_api.sh
./test_api.sh
```

#### **Using Makefile**
```sh
make run   # Start the API
make test  # Run Go unit tests
make e2e   # Run end-to-end API tests
```

---

## **4. Error Handling**
| Error Code | Reason |
|------------|--------|
| `400` Bad Request | Invalid JSON format |
| `404` Not Found | User ID does not exist |
| `500` Internal Server Error | Server-side issue |

---

## **5. File Storage (`data.json`)**
The API saves user data persistently in `data.json` in the following format:
```json
{
    "550e8400-e29b-41d4-a716-446655440000": {
        "id": "550e8400-e29b-41d4-a716-446655440000",
        "name": "Sheila",
        "age": 25
    }
}
```

---

## **6. Notes**
- The API is **stateless** except for the JSON file storage.
- If the `data.json` file is deleted, all stored user data is lost.
- For **Windows users**, run the API inside **WSL** for best results.

---

### ** Now You’re Ready to Use the API! **
