# CEN5035Project

Technologies Used:

Backend: GO
    Gorm - https://github.com/go-gorm/gorm   
    Gorilla/sessions- https://github.com/gorilla/sessions   
    Bcript - https://golang.org/x/crypto/bcrypt   

FrontEnd: Angular

### Application Description:
The purpose of this application is to streamline the contract management for attorneys. Attorneys are faced with the issue of having many contracts to deal with at once. This application allows attorneys to upload files to a database for easy storage and retrieval. The attorneys will be able to view their contracts and manage them all within the portal. The Application allows for attorneys to communicate with clients and other attorneys through a messaging system.

### Demo of Application:


https://user-images.githubusercontent.com/32618925/164346781-64a2a599-360f-422c-9a66-b51d8f0756db.mp4


### Cypress Tests:

To run the tests and generate the video, run "npm start" from `\CEN5035Project\CEN5035-front-end`, `go run main.go` from `CEN5035Project\backend\src`, and `npx cypress ru`" from the `\CEN5035Project\CEN5035-front-end` folder. The video can be found at `CEN5035Project\CEN5035-front-end\cypress\videos\spec.ts` and the tests themselves can be found at `CEN5035Project\CEN5035-front-end\cypress\integration\spec.ts`.


https://user-images.githubusercontent.com/32618925/164347743-0d908848-03de-4996-a089-6cb498df2382.mp4


### Unit Tests:

To run the unit tests for the backend, navigate the the backend source folder `cd CEN5035Project\backend\src` and run `go run main.go`. Then, navigate to the Test folder `cd CEN5035Project\backend\src\Test` and run `go test`.



https://user-images.githubusercontent.com/32618925/164350805-dc844ce1-ac5f-484d-bb54-2acfd88b24cf.mp4



### API Documetation:
The API documentation, for both the front end and the back end, can be found on the [wiki](https://github.com/Jshel/CEN5035Project/wiki) page.

### Project Board:
The Project board can be found in the [Projects](https://github.com/Jshel/CEN5035Project/projects/1) tab.

### Sprint 4:
sprint 4 can be found [here](https://github.com/Jshel/CEN5035Project/blob/main/Sprint4.md).

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


