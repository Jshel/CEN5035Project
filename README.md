# CEN5035Project

Technologies Used:

Backend: GO
    Gorm - https://github.com/go-gorm/gorm
    Gorilla/sessions- https://github.com/gorilla/sessions
    Bcript - golang.org/x/crypto/bcrypt

FrontEnd: Angular

### Application Description:
The purpose of this application is to streamline the contract management for attorneys. Attorneys are faced with the issue of having many contracts to deal with at once. This application allows attorneys to upload files to a database for easy storage and retrieval. The attorneys will be able to view their contracts and manage them all within the portal. The Application allows for attorneys to communicate with clients and other attorneys through a messaging system.

### Demo of Application:

### Cypress Tests:
https://user-images.githubusercontent.com/32618925/164114065-dd51adb7-9c4d-4741-b0f5-13f3cb3756d8.mp4

### Unit Tests:
https://user-images.githubusercontent.com/32618925/unit_test_demo_backend.mp4

### API Documetation:
The API documentation, for both the front end and the back end, can be found on the [wiki](https://github.com/Jshel/CEN5035Project/wiki) page.

### Project Board:
The Project board can be found in the [Projects](https://github.com/Jshel/CEN5035Project/projects/1) tab.

### Team members
Jennifer Sheldon - Jshel - Frontend   
Sunil Ghosal - sunil494 - Frontend   
Akshay Sharma - akshaysharmajs - Backend   
Nicholas Fox - DrBubbles42 - Backend 

### APP Setup

- First, run frontend server (http://localhost:4200/):
    - `cd CEN5035-front-end`
    - `run nmp start`

- Now, run backend server (http://localhost:8080/):

    - `cd CEN5035PROJECT/backend/src` directory
    - `go run main.go`
- Unit tests, 
    - `cd CEN5035PROJECT/backend/src/Test` directory
    - `go test`


