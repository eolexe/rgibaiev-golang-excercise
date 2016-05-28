# rgibaiev-golang-excercise

short coding exercise for Ruslan Gibaiev

Additional question may be provided via issues to this repo. Good luck and have fun! :)

Create HTTP Rest API:

Use echo for web handler
Implement login endpoint with JWT token, and simple middleware that checks header for 'Authorization: Bearer %jwt_token%' in each request. Otherwise return 403 and json struct with error
Implement endpoint that will use OAuth2 authorization for FB, to login and issue access_token
Log each request including status code (logrus)
Implement persistence with MySQL and Gorm (https://github.com/jinzhu/gorm)
Use tool of your choice for DB migrations
Implement save endpoint for Task object
Implement update endpoint for Task object
Implement get endpoint for Task object
Implement delete endpoint for Task object (just update IsDeleted field)
Use CORS (reply with header Access-Control-Allow-Origin: *)
Add support for OPTION HTTP method for each endpoints
Configure daemon over simple YAML config. Specify path as process flag for daemon. Required params: ListenAddress, DatabaseUri.
Implement 3-rd party libs vendoring with tool of your choice. (godeps)
Put in comments below description of taken architectural decisions
Task:

type Task struct {
    Id          int64
    Title       string
    Description string
    Priority    int
    CreatedAt   *time.Time
    UpdatedAt   *time.Time
    CompletedAt *time.Time
    IsDeleted   bool
    IsCompleted bool
}
