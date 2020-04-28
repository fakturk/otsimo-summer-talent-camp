# Otsimo Summer Talent Camp Project

# April 27,2020

## Purpose

This projects aims to create a REST API for managing the backend of Otsimo Summer Talent Camp Project.

## Project Description
At Otsimo, we want to start a project called the Summer Talent Camp. At the beginning of the process, the 4th year university students from different departments will apply to the camp. After the selection progress, the interns will be given a task to work on during their internship. Among these interns, some will be hired as full-time team members at the end of the summer. To make the selection process simpler for the team, we would like to develop a platform that provides a way to manage candidates, their applications, and appointments relating to them easily and quickly. On this platform, the assigned team member will be able to quickly assess the candidates and see if there is anything that needs their attention. They will be able to accept or reject the application, find out if they have any appointments with the applicants, etc. As a developer, you will code the functions related to the database. We will use MongoDB as a database and GO as the programming language. All the details, rules are written below and an example DB dump (dumped with mongodump --host localhost:27018 --archive=dump.gz --gzip --db Otsimo) is attached in description.

## Requirements


- [x] Each candidate should have a unique identifier.


- [x] New candidate's Status should be Pending and meeting count should be 0. If meeting count is greater than 0 and smaller than 4, the Status should be In Progress.


- [x] New candidates should have an assignee who is working in the department that the candidate is applying to work.


- [x] Candidates cannot be accepted before the completion of 4 meetings.


- [x] If the next meeting is the last (4th) one, the assignee of the candidate should be changed to Zafer. He is the CEO of Otsimo.



## Bonus Features


- [x] Implementing FindAssigneesCandidates (id string) ([]Candidate, error)


- [x] The email format for the candidate should be example@email.xyz. Otherwise, the candidate should not be inserted to DB because the only way to communicate with the candidate is through email.


- [x] Upps. We forget to add the application_date field for the Candidate object.


- [x] Create a simple HTTP Rest API to manage Candidates and Assignees by using these storage functions. For example we can accept candidate with this.

## Architecture
### Candidate
We have candidates service and controller.

Candidate Controller have following methods;

``` code 
func GetCandidatesFunc(w http.ResponseWriter, r *http.Request)
```
``` code 
func CreateCandidateFunc(w http.ResponseWriter, r *http.Request)
```
``` code 
func ReadCandidateFunc(w http.ResponseWriter, r *http.Request)
```
``` code 
func DeleteCandidateFunc(w http.ResponseWriter, r *http.Request)
```
``` code 
func AcceptCandidateFunc(w http.ResponseWriter, r *http.Request)
```
``` code 
func DenyCandidateFunc(w http.ResponseWriter, r *http.Request)
```
``` code 
func ArrangeMeetingFunc(w http.ResponseWriter, r *http.Request)
```
``` code 
func CompleteMeetingFunc(w http.ResponseWriter, r *http.Request)
```

Candidate Service have following methods;

``` code 
func GetAllCandidates() ([]model.Candidate, error) 
```
``` code 
func CreateCandidate(candidate model.Candidate) (model.Candidate, *mongo.InsertOneResult, error)
```
``` code 
func ReadCandidate(_id string) (model.Candidate, error)
```
``` code 
func DeleteCandidate(_id string) (*mongo.DeleteResult, error) 
```
``` code 
func ArrangeMeeting(_id string, meetingTime *time.Time) (*mongo.UpdateResult, error)
```
``` code 
func CompleteMeeting(_id string) (*mongo.UpdateResult, error) 
```
``` code 
func DenyCandidate(_id string) (*mongo.UpdateResult, error)
```
``` code 
func AcceptCandidate(_id string) (*mongo.UpdateResult, error) 
```

### Assignee
We have assignees service and controller.

Assignee Controller have following methods;

``` code 
func FindAssigneeIDByNameFunc(w http.ResponseWriter, r *http.Request) 
```
``` code 
func FindAssigneesCandidatesFunc(w http.ResponseWriter, r *http.Request) 
```
``` code 
Assignee Controller have following methods;
```
``` code 
func FindAssigneeIDByName(name string) (string, error) 
```
``` code 
func GetAssignee(_id string) (model.Assignee, error) 
```
``` code 
func FindAssigneesCandidates(_id string) ([]model.Candidate, error) 
```

### db

MongoDB connection settings.

We have following function to connect db;

``` code 
func ConnectDB(collectionName string) *mongo.Collection
```

### model

Holds structs of our project

``` code 
type Candidate struct 
```
``` code 
type Assignee struct 
```
``` code 
type MeetingTime struct 
```

### helper

Error model defined in here;

``` code 
type ErrorResponse struct
```
``` code 
func GetError(err error, w http.ResponseWriter)
```

## How to Run

```shell
$ go run main.go
```

## Usage

- Create (Add) Candidate:

    - request type: POST

    - host: localhost:8080/candidate/create
    - example request
```json
{
        "_id": "",
        "first_name": "Fatih",
        "last_name": "Akturk",
        "email": "f.akturk@gmail.com",
        "department": "Development",
        "university": "TOBB",
        "experience": true,
        "status": "Pending",
        "meeting_count": 0,
        "next_meeting": null,
        "application_date":"2020-04-27T13:40:00.000+00:00",
        "assignee": "5bb6368f55c98300013a087d"
    }
```
- Read Candidate:

    - request type: GET

    - host: localhost:8080/candidate/read/{id}

- Delete Candidate:

    - request type: DELETE

    - host: localhost:8080/candidate/delete/{id}

- Accept Candidate:

    - request type: GET

    - host: localhost:8080/candidate/accept/{id}
    
- Deny Candidate:

    - request type: GET

    - host: localhost:8080/candidate/deny/{id}
    
- Get All Candidate:

    - request type: GET

    - host: localhost:8080/candidates
 
 - Arrange Meeting:
 
     - request type: POST
 
     - host: localhost:8080/meeting/arrange/{id}
     - example request
 ```json
{
"meeting_time":"2020-05-03T13:40:00.000+00:00"
}
 ```

- Complete Meeting:

    - request type: GET

    - host: localhost:8080/meeting/complete/{id}
    
- Find Assignee ID By Name:

    - request type: GET

    - host: localhost:8080/assignee/findid/{name}
    
 - Find Assignee Candidates:
 
     - request type: GET
 
     - host: localhost:8080/assignee/candidates/{id}
     
  ## Postman usage
  There is a postman collection added for 10 requests and "otsimo.postman_collection.json" file can be implemented in the Postman and directly called from there.
  
  ![Postman ](/images/otsimo-postman.png)