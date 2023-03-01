*** Settings ***
Documentation     This is a test suite for testing the API requests functionality 
Resource           ascanApi_resource.robot
Suite Setup       Connect to API

***Variables***
${DELETE_USER_ID}        35
${UPDATE_USER}           10


*** Test Cases ***
GET all the users
    Get Request: All the users
    Check status_code    200
    Check reason    OK
    Check content is not empty
GET specific user
    get Request: get user by username "vito_santos"
    Check status_code    200
    Check reason    OK
    # Checking user fields consistency

POST create user
    Post Request: Create user
    Check status_code    200
    Check reason    OK
    Check content is not empty

PUT update user
    Update user "${UPDATE_USER}"
    Check status_code    200
    Check reason   OK

DELETE user
    Delete user "${DELETE_USER_ID}"
    Check status_code    200
    Check reason   OK
    Check if "${DELETE_USER_ID}" is empty