# Planner-Service

This app is for those who is learning to build golang microservice. It helps to understand golang packages and code structuring.  It is consumed by [microfrontend](https://github.com/BishnuPandey/Planner-Frontend) for data handling. 

### Run the project

Clone the repo. Then in the project directory, run:

```
> git clone git@github.com:BishnuPandey/Planner-Service.git
> cd /your_folder_path
> go run main.go
```

Runs the app in the development mode.
Open [http://localhost:9000/api/tasks](http://localhost:9000/api/tasks) to view task list in your browser.

### JSON Response
```
[
    {
        "_id": "65d2f86c9b0b24a5b82d640b",
        "status": false,
        "task": "Hello Go"
    },
    {
        "_id": "65d319bd541392e8899c973d",
        "status": false,
        "task": "SAMRIDDHI PANDEY"
    },
    {
        "_id": "65d32e0ac971a75c569bf7d4",
        "status": true,
        "task": "One"
    },
    {
        "_id": "65d3301f121105cf4cba469e",
        "status": false,
        "task": "Two"
    }
]
```