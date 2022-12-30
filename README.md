
# Project Book Rent APIs

This is a backend fun-project in the form of APIs for a book lending service that is fully using Golang with the Echo framework, and the other services used are MySQL which is GORM as ORM and Redis.


## Project Recipes

 - [More about Go Echo Minimalist Go Web Framework](https://echo.labstack.com/)
 - [More about GORM Fantastic ORM Library For Golang](https://gorm.io/)
 - [More about Golang Redis](https://redis.uptrace.dev/)

## Architecture
Layered Architecture is applied in this project for some reasons :
- Modularity: Layered architecture allows you to break down a complex system into smaller, more manageable parts. This makes it easier to understand and modify the system, as you can focus on one layer at a time.
- Reusability: Because each layer is designed to perform a specific function, it is easy to reuse the code from one layer in another project. This can save time and effort when developing new software.
- Testability: As mentioned earlier, layered architecture makes it easier to test the different parts of a system separately. This can help ensure that the system is working correctly and can save time and effort when debugging.
- Maintainability: If a problem arises in one layer of the system, it is easier to isolate and fix the issue if the system is organized in layers. This can make it easier to maintain the system over time.
- Scalability: Layered architecture makes it easier to scale a system up or down as needed, as you can add or remove layers as needed without affecting the other parts of the system. This can make it easier to adapt to changing business needs.


### A Bit Notes

This project still has many weaknesses and there is still a lot has to be added and improved. Therefore, there is a very high possible of future updates to this project.


## Installation

To run this project you have to clone my this repostitory to your local :

```git clone https://github.com/husnulnawafil/book-rent-api.git```

Then, you will need to add the following environment variables to your `.env` file inside root folder. Below is mine :

```export APP_PORT="8000"
export MYSQL_DRIVER="mysql"
export MYSQL_NAME="book_rent"
export MYSQL_ADDRESS="127.0.0.1"
export MYSQL_PORT="3306"
export MYSQL_USERNAME="root"
export MYSQL_PASSWORD=""
export REDIS_HOST="127.0.0.1:6379"
export REDIS_PASSWORD=""
```

To make sure `.env` has been read run on your terminal.

```source .env```

And last but not sure it is least.

```go run main.go```

Voilla , it works.

Notes : every time you make a change on your code you have to run ```go run main.go``` again and again. If you do not prefer, you may try [golang air](https://github.com/cosmtrek/air).
## Documentation
Here provides some [Postman](https://www.postman.com/) Collection that may help.

[Documentation](https://github.com/husnulnawafil/book-rent-api/blob/34b8ce4a289a06ee11a904b650c003500d65538f/documentation/book_rent.postman_collection.json)

