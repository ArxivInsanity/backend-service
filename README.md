# Backend Service
This is the backend service for the Arxiv Insanity application. Application is built using Golang and uses the Gin framework.

To deploy the application:
1. First provision the GKE cluster using terraform by running the github actions [here](https://github.com/ArxivInsanity/terraform-infra/actions/workflows/terraform.yml).
2. Next run the github actions to deploy the application to GKE using terraform [here](https://github.com/ArxivInsanity/backend-service/actions/workflows/GKE-deploy.yaml).

For local setup, use dev containers:
1. Install Dev containers extension in vscode
2. Open the root folder using vscode, and select "Reopen code in container"
3. Create a launch.json file with the below configuration (.vscode/launch.json)
```
{
    "version": "0.2.0",
    "configurations": [
      {
        "name": "Launch",
        "type": "go",
        "request": "launch", 
        "mode": "debug",
        "host": "127.0.0.1",
        "program": "${workspaceFolder}/src",
        "env": {
            "JWT_SECRET":"REPLACE_ME",
            "OAUTH2_CLIENT_ID" : "REPLACE_ME",
            "OAUTH2_SECRET" : "REPLACE_ME",
            "MONGO_URL" : "REPLACE_ME"
        },
        "args": [], 
        "showLog": true
      }
    ]
  }
```
4. Press F5 (or click on Run and Debug and click start)
5. Navigate to [localhost](http://localhost:8080/docs/index.html) to see the swagger UI.
6. To authenticate with google OAUTH2, navigate [here](http://localhost:8080/auth/google), which will redirect you back to the swagger ui after authenticating
