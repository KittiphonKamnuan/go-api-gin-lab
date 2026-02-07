ทดสอบ API ครบทุก endpoint ผลลัพธ์ทั้งหมดผ่าน                                                                              
                           
  │         Test         │         Endpoint          │                      Result                       │                         
                           
  │ GET all students     │ GET /students             │ 200 OK                                            │                         
                           
  │ GET by ID            │ GET /students/:id         │ 200 OK                                            │                         
                           
  │ POST create          │ POST /students            │ 201 Created                                       │                         
                           
  │ PUT update           │ PUT /students/66090002    │ 200 OK + return updated data                      │                         
                           
  │ PUT not found        │ PUT /students/99999999    │ 404 {"error":"Student not found"}                 │
  
  │ DELETE success       │ DELETE /students/66090002 │ 204 No Content                                    │                         
  
  │ DELETE not found     │ DELETE /students/99999999 │ 404 {"error":"Student not found"}                 │
  
  │ POST empty id        │                           │ 400 {"error":"id must not be empty"}              │
  
  │ POST empty name      │                           │ 400 {"error":"name must not be empty"}            │
  
  │ POST gpa > 4.0       │                           │ 400 {"error":"gpa must be between 0.00 and 4.00"} │
  
  │ POST gpa < 0         │                           │ 400 {"error":"gpa must be between 0.00 and 4.00"} │
  
  │ PUT empty name       │                           │ 400 {"error":"name must not be empty"}            │
  
  │ PUT gpa out of range │                           │ 400 {"error":"gpa must be between 0.00 and 4.00"} │
