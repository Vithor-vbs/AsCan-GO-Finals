*** Settings ***
Library       RequestsLibrary
Library       Collections

*** Variables ***
${URL}        http://localhost:8080

&{USER_1}      ID=1
...            Username=Isaac_newton123
...            FirsName=Isaac
...            LastName=Newton
...            Email=gravidade@exemplo.com
...            Password=inercia123
...            Phone=97654-2543
...            UserStatus=0

&{USER_100}    ID=100
...            Username=teste
...            FirsName=teste
...            LastName=teste
...            Email=testando@exemplo.com
...            Password=testedeAPI
...            Phone=77777-7777
...            UserStatus=0

*** Keywords ***

Connect to API
    Create session    api    ${URL}   
    ${HEADERS}     Create Dictionary    content-type=application/json
    Set Suite Variable    ${HEADERS}

#GET All
Get Request: All the users
    ${ANSWER}=    Get on session    api    /user
    log            ${ANSWER.json()}
    Set Test Variable    ${ANSWER}

Check status_code
    [Arguments]    ${STATUS_CODE}
    Should Be Equal As Numbers    ${ANSWER.status_code}    ${STATUS_CODE}

Check reason
    [Arguments]    ${REASON}
    Should Be Equal    ${ANSWER.reason}    ${REASON}
Check content is not empty
    Should Not Be Empty    ${ANSWER.content}

#GET one
Get Request: get user by username "${USERNAME}"
    ${ANSWER}=    Get on session    api    /user/${USERNAME}
    log            ${ANSWER.text}
    Set Test Variable    ${ANSWER}

# Checking user fields consistency
#     Should Not Be Empty    ${ANSWER.json()["username"]}
#     Should Not Be Empty    ${ANSWER.json()["firstName"]}

#POST Create user
Post Request: Create user
    ${ANSWER}=    Post on session    api    /user    
    ...           data={"id": 100,"username": "teste","firsName": "teste","lastName": "teste","email": "testando@exemplo.com","password": "testedeAPI","phone": "77777-7777","userStatus": 0}
    ...           headers=${HEADERS}

    log            ${ANSWER.text}
    Set Test Variable    ${ANSWER}

#PUT Update user
Update user "${USER_ID}"
    ${ANSWER}    PUT On Session    api    /user/${USER_ID}
    ...                              data={"username": "${USER_1.Username}","firstName": "${USER_1.FirsName}","lastName": "${USER_1.LastName}","email": "${USER_1.Email}"}
    ...                              headers=${HEADERS}
    Log                  ${ANSWER.text}
    Set Test Variable    ${ANSWER}
    
#DELETE Delete user
Delete User "${ID_USER}"
    ${ANSWER}    DELETE On Session    api    user/${ID_USER}
    Log            ${ANSWER.text}
    Set Test Variable    ${ANSWER}

Conferir se retorna todos os dados cadastrados do livro "${ID_USER}"
    Conferir livro    ${ID_USER}

Conferir se retorna todos os dados alterados do livro "${ID_USER}"
    Conferir livro    ${ID_USER}

Conferir livro
    [Arguments]     ${ID_USER}

    Dictionary Should Contain Item    ${ANSWER.json()}    Id             ${USER_${ID_USER}.ID}
    Dictionary Should Contain Item    ${ANSWER.json()}    username       ${USER_${ID_USER}.Username}
    Dictionary Should Contain Item    ${ANSWER.json()}    firstName      ${USER_${ID_USER}.FirsName}
    Dictionary Should Contain Item    ${ANSWER.json()}    lastName       ${USER_${ID_USER}.LastName}
    Dictionary Should Contain Item    ${ANSWER.json()}    email          ${USER_${ID_USER}.Email}
    Dictionary Should Contain Item    ${ANSWER.json()}    password       ${USER_${ID_USER}.Password}
    Dictionary Should Contain Item    ${ANSWER.json()}    phone          ${USER_${ID_USER}.Phone}

Check if "${ID_USER}" is empty
    Should Be True     ${ANSWER.json()}