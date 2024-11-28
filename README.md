# **Building a Go API for Current Toronto Time with MySQL Database Logging**

This Go application provides an API that returns the current time in Toronto in JSON format and logs each request to a MySQL database.

---

## **Getting Started**

### **Prerequisites**
Ensure you have the following installed on your system:
- **Go 1.16+**
- **MySQL server**
- **curl/Postman** (for testing the API)

---

### **Database Setup**
1. **Install MySQL**: 
   Install MySQL server if not already installed. You can follow the instructions from the official MySQL website or use a package manager for installation.

2. **Create Database and Table**:
   Open your MySQL client and create a new database and table:
   ```sql
   CREATE DATABASE toronto_time_db;

   USE toronto_time_db;

   CREATE TABLE time_log (
       id INT AUTO_INCREMENT PRIMARY KEY,
       timestamp DATETIME NOT NULL
   );

3. **Test the application**

- run main.go
- I have passed the password argument directly in terminal or we can use environment variables

   
![image](https://github.com/user-attachments/assets/ed4ba2af-dece-4aaf-b15d-66a0a22c268d)

4. **Open brower and enter url or curl from the terminal**


![image](https://github.com/user-attachments/assets/1f357c3e-a02d-48ce-9185-c22af6736290)

5. **Enter query in mysql to view the rows**


![image](https://github.com/user-attachments/assets/7e47c419-af7d-4055-9ab3-37b097480dd8)




   
