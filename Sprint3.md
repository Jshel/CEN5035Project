## What was accomplished (new stuff since the last sprint):

### Frontend:

Improved tests

Added dynamic forms for messages and contracts

### Backend:

Added contract uploading and downloading. Fixed sessions. Redirect from login page and from create account.

Unit tests.

## Demos:

### Frontend:

Run ```npm start``` in CEN5035Project\CEN5035-front-end to run the frontend.

Run ```go run main.go``` in CEN5035Project\backend\src to start up the backend.

Run ```npx cypress run``` in CEN5035Project\CEN5035-front-end to run the Cypress tests.

The video generated in CEN5035Project\CEN5035-front-end\cypress\videos:


https://user-images.githubusercontent.com/32618925/161346092-ebcafedd-0b3b-474e-bfa6-8a96cba92d3d.mp4



### Backend:
DrBubbles42: Nicholas Fox   
akshaysharmajs: Akshay Sharma 

Start the Frontend: ```npm run start```   
Start the Backend: ```go run main.go```   
Start the test server: `live-server`   
Curl to download contract: `curl http://localhost:4200/api/download?attorney_email=a%40a.a%26contract_id=00000001 -o download2.pdf`   

![sprint3BE](https://user-images.githubusercontent.com/25064175/161361184-0072a022-e9bd-4a44-ba50-0f09d3bf06c7.gif)
