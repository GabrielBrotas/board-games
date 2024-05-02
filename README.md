# board-games

## Deploy to Heroku

Heroku's container support. Here’s how you can do it:

1. Create a Heroku App

Use the Heroku CLI to create a new app:

```bash
heroku create
```

2. Log in to Container Registry

Make sure you're logged in to Heroku via the CLI and then log in to Heroku Container Registry:

```bash
heroku container:login
```

3. Navigate to the Server Directory

Change to the server directory where your Dockerfile is located:

```bash
cd server
```

4. Build and Push the Docker Image

Build the Docker image and push it to Heroku:

```bash
heroku container:push web -a <your-app-name>
```

Replace your-app-name with the name of your Heroku app.

5. Release the Image

After pushing the image, release it to deploy:

```bash
heroku container:release web -a your-app-name
```

Final Steps
Verify Deployment: Open your application in a web browser using the command `heroku open agile-atoll-94388` or by navigating to `your-app-name.herokuapp.com`.
View Logs: If you encounter any issues, view logs with `heroku logs --tail` to troubleshoot.
