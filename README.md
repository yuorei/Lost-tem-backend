# Lost-Item App
The Otoshimono app is a tool that allows you to share your lost items with people around the world simply by taking a picture of them. By using this app, if you find a lost item that does not need to be reported to the police, you can share it immediately.

# Apps
[Otoshimono App](https://otoshimono.gpio.biz/)
Please use from here

## Front
Click here for [Otoshimono App Front](https://github.com/gpioblink/otoshimono-front)
## Preferred language
The backend of this app was developed in Go.

## database
Adopted PostgreSQL
I am using Cloud SQL
## frameworks, libraries
I am using the following frameworks and libraries. The version number is mentioned in go.mod

### gin
Go web framework
It was adopted because of the number of stars on GitHub and the fact that the development members have experience using it.

### gorm
ORM for Go
The reason for adoption was the number of stars on GitHub and the abundance of documents.

## Reproduction of development environment
```
git clone git@github.com:yuorei/lost-item-backend.git
```
Clone from GitHub
Adopted docker for the development environment
```
docker compose up
```
App backend is launched locally

## Deploy
Deploying to Cloud Run
### CI, CD
CI and CD are automatically entered when merged into main on GitHub
I made the settings in Cloud Run
- Edit Continuous Deployment
- Select repository and branch
- Set secrets and environment variables in Cloud Run
