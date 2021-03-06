## What was accomplished (new stuff since the last sprint):

### Frontend:

Added functionality to upload, search for, and view contracts and messages. Added functionality to display counts of messages and contracts. Added support for cookies and completed lazy loading for module. Fixed cookie related bugs and linked front-end routes to backend API.

### Backend:

Fixed some bugs and errors. Endpoint for file upload and download fixed. Added messaging endpoints. Added endpoints to Get sessions, emails, and count messages and contracts. Added wildcard support in endpoints that query contracts and messages. Added support for sending lists of contracts and messages. Unit tests for all handlers.

## Demos:

### Frontend:

Run ```npm start``` in CEN5035Project\CEN5035-front-end to run the frontend.

Run ```go run main.go``` in CEN5035Project\backend\src to start up the backend.

Run ```npx cypress run``` in CEN5035Project\CEN5035-front-end to run the Cypress tests.

The video generated in CEN5035Project\CEN5035-front-end\cypress\videos:

https://user-images.githubusercontent.com/32618925/164355948-dac65bf5-05cc-4c50-b4b6-8d3b2a256d4c.mp4


### Backend:
DrBubbles42: Nicholas Fox   
akshaysharmajs: Akshay Sharma 

Start the Frontend: ```npm run start```   
Start the Backend: ```go run main.go```   
Start the test server: `live-server` 
Run Unit tests: `go test`
![sprint4](https://user-images.githubusercontent.com/25064175/164350687-4ff3d806-fae6-48c8-8897-16fbf2d20bfb.gif)
