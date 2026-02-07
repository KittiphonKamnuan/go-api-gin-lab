# 🧠 Homework Challenge: REST API Enhancement with Gin and SQLite

⏱️ **Estimated Time:** 2–4 hours  
📌 **Type:** Individual Assignment

---

## 🎯 Objective
This assignment aims to strengthen students’ understanding of:
- REST API design
- Layered Architecture
- Gin Framework
- SQLite database integration
- Clean and maintainable backend code

Students will extend the existing Student API developed in Lab 3.

---

## 🧩 Assignment Overview
You are given a REST API project with the following features:
- Layered Architecture (Handler, Service, Repository, Model)
- Gin Framework
- SQLite database
- Existing APIs:
  - GET /students
  - GET /students/:id
  - POST /students

Your task is to **extend and improve** the system by completing the challenges below.

---

## 🔹 Challenge 1: Update Student (PUT)

### Requirement
Implement an API to update student information.

### Endpoint:
**`PUT /students/:id`**

### Request Body (JSON)
```json
{
  "name": "Updated Name",
  "major": "Updated Major",
  "gpa": 3.50
}
```

### Expected Behavior
- Update student data by ID
- Return the updated student as JSON
- Return HTTP 404 if the student is not found

---

## 🔹 Challenge 2: Delete Student (DELETE)

### Requirement
Implement an API to delete a student by ID.

### Endpoint:
**`DELETE /students/:id`**

### Expected Behavior
- Delete student from the database
- Return HTTP 204 (No Content) on success
- Return HTTP 404 if the student is not found

---

## 🔹 Challenge 3: Input Validation

### Requirement
Add validation rules when creating or updating a student.

### Validation Rules
- id must not be empty
- name must not be empty
- gpa must be between `0.00` and `4.00`

### Expected Behavior
- Return HTTP 400 (Bad Request) if validation fails
- Return clear error messages in JSON format

---

## 🔹 Challenge 4: Error Handling Improvement

### Requirement
Standardize error responses.

### Example Error Response
```json
{
  "error": "Student not found"
}
```
- Use appropriate HTTP status codes
- Do not expose raw database errors directly to the client

---

## 📁 Project Structure (Recommended)
```go
go-api-gin/
├─ main.go
├─ models/
├─ repositories/
├─ services/
├─ handlers/
├─ config/
└─ students.db
```

---
## ✅ Submission Requirements
>`Submit your assignment as a Git repository link (GitHub or GitLab).`

### Requirements:
- The repository must be public
- The project must compile and run successfully
- Include a README.md with:
  - How to run the project
  - Available API endpoints

Do NOT submit ZIP files or screenshots.

#### Example `README.md`:
```md
# Student API with Gin

## How to Run
```bash
go run main.go
```

---
## 📊 Grading Rubric: Mini Assignment (PUT & DELETE API)

**Total Score: 20 Points (+ Bonus)**

### Core Score (5 Points)

| No. | Criteria | Description | Points |
|----|---------|-------------|--------|
| 1 | PUT Student API | `PUT /students/:id` works correctly | 5 |
| 2 | DELETE Student API | `DELETE /students/:id` works correctly | 5 |
| 3 | Input Validation | Validates input data (e.g. GPA range, required fields) | 5 |
| 4 | Project Runs | Project compiles and runs without errors | 2 |
| 5 | Error Handling | Proper HTTP status codes and clear error responses | 2 |
| 6 | README.md | Clear instructions on how to run and test the API | 1 |
| **Total (Core)** | **20 Points** 

---

### ⭐ Bonus: Clean Structure (+2 Point)

| Bonus Item | Description | Bonus |
|-----------|-------------|-------|
| Clean Project Structure | Clear separation of layers (handler, service, repository), readable file structure | +1 |
| Code Quality | Clean code, readable naming, consistent formatting | +1 |

**Maximum Score:** 22 Points
