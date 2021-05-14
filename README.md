# BlogsAPI

API/Backend written in Golang for practice

## Features:

_API/Backend to fetch and display blogs_

- Markdown parsing for blogs
- Dockerized the whole setup
- A simple frontend to display blogs
- Admin Auth so only they can add new blogs

### Setup locally

- Clone the repo using `git clone https://github.com/ShauryaAg/BlogsAPI.git`
- Move into the project folder `cd BlogsAPI/`
- Create a `.env` file in the project folder

```
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=postgres
POSTGRES_HOST=postgres
SECRET=<Secret>
```

##### Using Docker <a href="https://www.docker.com/"> <img alt="VaccineNotifier" src="https://www.docker.com/sites/default/files/d8/styles/role_icon/public/2019-07/vertical-logo-monochromatic.png" width="50" /> </a>

- Run using `sudo docker-compose up`

### **OR**

##### Using Golang <a href="https://golang.org/"> <img alt="VaccineNotifier" src="https://golang.org/lib/godoc/images/go-logo-blue.svg" width="50" /> </a>

- Install the dependecies using `go mod download`
- Run using `go run server.go`

## Screenshots

Home Page:
![image](https://user-images.githubusercontent.com/31778302/115430234-cfde3980-a221-11eb-9f4a-ea2e2b7789ed.png)

Single blog:
![image](https://user-images.githubusercontent.com/31778302/115430701-382d1b00-a222-11eb-81ed-93e402cb38d2.png)

## TODO

- [ ] Add better markdown parsing
