# ToDoList

Project Summary	"One of the most common tools used by professionals all over the world to capture tasks is a To Do list. They come in all shapes, forms, and features. In this project you will build Backend of a SAAS application that allows its users to build and maintain a To Do list.

A task consist of following properties: creation date time, title, description, file attachments (e.g. an image), due date time, completion status (true/false), completion date time."


Datastore	Use flat file as your database. Store list of tasks in a text or json file.


REST	Unless absolutely required, all the communication between a client (user) and the server (website) must be based on REST and JSON.
Core operations	"Allow a user to perform following operations
- Create a new task
- Edit a task
- Delete a task
- View list of tasks
- Attach file(s) with an existing task
NOTE: Data returned from all of the APIS must be paginated"


Reports	"Allow a user to generate following reports (each report should be served from a separate endpoint)
- Count of total tasks, completed tasks, and remaining tasks (aggregate all 3 in parallel)
- Average number of tasks completed per day (aggregate average in parallel for each day)
- On what date, maximum number of tasks were completed in a single day
- Count maximum number of tasks added on a particular day. (It should return date and number of tasks, if multiple, return multiple dates)

Note: If report for particular use case exists on disk it should be returned without generating a new one otherwise new report will be generated and persisted to the disk."
Similar Task Detection	Return user a list of similar tasks. Two tasks A and B are considered similar if all the words in the task A exist in task B or vice versa.
Caching of Reports	A generated report should remain valid for 15 minutes on disk. After 15 minutes expire the report from the disk
Logging	All the REST API endpoint calls should be logged in a log file
API Documentation	Generate documentation of API that can be consumed by front end developers


REST API collection (Postman)	REST API collection using any API client of all the endpoints.


Language and Framework	Go language with standard http package for rest. Assignment in any other language will not be accepted