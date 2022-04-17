# VIDEO CHAT BACKEND
This is a backend application for a video chatting application I am currently creating written in golang. It utilizes the following:
1. golang
2. postgresql
3. tokens
4. webrtc
5. gorm

----------

## RUNNING THE APPLICATION
Please make sure environment variables are set up before running the application or else errors might occur during run time

### WITH DOCKER
1. If you have docker installed you can just run `docker build -t <name_of_image> .` to create an image for this project
2. After creating an image I would advice to create a **.env** file for your environment variables (which would be easier) or create variables manually on your operating system of choice
3. Then run `docker run -p <host_port>:<application_port> --env-file <path_of_env_file> --name <name_of_container> <name_of_local_created_image>`
4. You can remove the `--env-file` tag if you aren't creating a **.env** file

### DEVELOPMENT
1. In the directory of the project you can simply run `go run server.go` and the application will work fine 
2. Alternatively you can dowload gin live server and run `gin -p <host_port> -a <application_port>` so live reload functionalities


----------
## EVIRONMENT VARIABLES NEEDED
These environment variables are need in order for the application to run 
### SERVER_PORT
- Port where the application will run 

### SECRET_KEY
- Access token secret key

### TOKEN_EXPIRE_TIME 
- Expiration time of access token (must be a number)

### POSTGRES VARIABLES
Because I used Postgresql to create this application the following variables and their values will be neeeded.
1. POSTGRES_USER 
2. POSTGRES_PASSWORD 
3. POSTGRES_DBNAME 
4. POSTGRES_PORT 

