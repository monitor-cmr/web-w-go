# web-w-go

## Enviroment Requirement

1. `go version go1.17.3 linux/amd64`
   1. Using go module
   2. `github.com/cosmtrek/air` - Auto reload
2. VS Code
3. Linux OS
4. Go-gin framework
   1. [https://golang.org/doc/tutorial/web-service-gin](https://golang.org/doc/tutorial/web-service-gin)

## Day 01

1. Install Go
   1. wget https://golang.org/dl/go1.17.3.linux-amd64.tar.gz
   2. sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.17.3.linux-amd64.tar.gz
   3. export PATH=$PATH:/usr/local/go/bin
   4. source ~/.bash_profile
   5. go version
2. Install `cosmtrek/air`
   1. curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
   2. sudo mv ./bin/air /usr/local/go/bin/
   3. air -v

## Day 02

1. Set up a hello world REST GIN framework.
2. Using Memory as backend storage.
3. How to code.
   1. Create `mkdir rest` folder && `cd rest`
   2. `go mod init github.com/monitor-cmr/web-w-go`
   3. `go get -u github.com/gin-gonic/gin`
4. Install Auto Reload
   1. `air init`
5. Design API Endpoint.

   | REST HTTP Request | CRUD   | GIN Methods |
   |-------------------|--------|-------------|
   | POST              | Create | GIN.POST    |
   | GET               | Read   | GIN.GET     |
   | PUT               | Update | GIN.PUT     |
   | DELETE            | Delete | GIN.DELETE  |

6. Albums request URL

   | REST HTTP Request | CRUD   | GIN Methods | Description      |
   |-------------------|--------|-------------|------------------|
   | /albums           | Read   | GIN.GET     | Get all albums   |
   | /albums           | Create | GIN.POST    | Create an albums |
   | /albums/:id       | Read   | GIN.GET     | Get an albums    |
   | /albums/:id       | Update | GIN.PUT     | Update an albums |
   | /albums/:id       | Delete | GIN.DELETE  | Delete an albums |

7. Create [`main.go`](rest/main.go) - en entrypoint file

## Day 03

1. Define Table albums in database

   | ID | Title                            | Artist         | Price |
   |----|----------------------------------|----------------|-------|
   | 1  | Blue Train                       | John Coltrane  | 56.99 |
   | 2  | Jeru                             | Gerry Mulligan | 17.99 |
   | 3  | Sarah Vaughan and Clifford Brown | Sarah Vaughan  | 39.99 |

2. Declaration of an album struct. You’ll use this to store album data in memory.

   ```go
   // album represents data about a record album.
   type album struct {
      ID     int
      Title  string
      Artist string
      Price  float64
   }
   ```

3. Declaration a slice of album structs containing data  - Memory Database

   ```go
   // albums slice to seed record album data.
   var albums = []album{
       {ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
       {ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
       {ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
   }
   ```

4. Working on Method POST - Create an albums

   | REST HTTP Request | CRUD   | GIN Methods | Description      |
   |-------------------|--------|-------------|------------------|
   | /albums           | Read   | GIN.GET     | Get all albums   |

5. Implement the `getAlbums` function in `main.go`

## Day 04

1. Working on Method POST - Create an albums

   | REST HTTP Request | CRUD   | GIN Methods | Description      |
   |-------------------|--------|-------------|------------------|
   | /albums           | Create | GIN.POST    | Create an albums |

2. Implement the `postAlbums` function - Method POST

   ```go
   var newAlbum album

    // Call BindJSON to bind the received JSON to newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
    c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
        return
    }

    // Get a length of albums
    lenAlbum := len(albums)

   // Get a last album's ID
    // - index start 0
    // - a last album = length - 1 => lenAlbum - 1
    newAlbum.ID = albums[lenAlbum-1].ID + 1

    // Add the new album to the slice.
    albums = append(albums, newAlbum)

    // response to user request
    c.IndentedJSON(http.StatusCreated, gin.H{"data": newAlbum.ID})
   ```

3. curl -XPOST -d @createAlbum.json http://127.0.0.1:8080/albums

## Day 05

1. Working on Method POST - Create an albums

   | REST HTTP Request | CRUD   | GIN Methods | Description      |
   |-------------------|--------|-------------|------------------|
   | /albums/:id       | Get    | GIN.GET     | Get an albums    |

2. Implement the `getAlbumsByID` function - Method GET

   ```go
   import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
   )

   func getAlbumsByID(c *gin.Context) {
    // /albums/:id => c.Param("id") => string value
    tempID := c.Param("id")

    // Convert string to int
    id, err := strconv.Atoi(tempID)
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
        return
    }

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, gin.H{"data": a})
            return
        }
      }

      // response to user request
      c.IndentedJSON(http.StatusNotFound, gin.H{"data": "album not found"})
   }
   ```

3. curl -XGET http://127.0.0.1:8080/albums/1

## Day 06

1. Working on Method PUT - Create an albums

   | REST HTTP Request | CRUD   | GIN Methods | Description      |
   |-------------------|--------|-------------|------------------|
   | /albums/:id       | Put    | GIN.PUT     | Update a albums  |

2. Implement the `putAlbums` function - Method GET

   ```go
    // /albums/:id => c.Param("id") => string value
    tempID := c.Param("id")

    // Convert string to int
    id, err := strconv.Atoi(tempID)
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
        return
    }

   var updateAlbum album

   // Call BindJSON to bind the received JSON to updateAlbum.
   if err := c.BindJSON(&updateAlbum); err != nil {
      c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
      return
   }

   // Assign updateAlbum to id in Request URL
   updateAlbum.ID = id

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for i, a := range albums {
        if a.ID == id {
         albums[i] = updateAlbum
            c.IndentedJSON(http.StatusOK, gin.H{"data": updateAlbum.ID})
            return
        }
   }

   // response to user request
   c.IndentedJSON(http.StatusNotFound, gin.H{"data": "album not found"})
   ```

3. curl -XPUT -d @updateAlbum.json http://127.0.0.1:8080/albums/3

## Day 07

1. Working on Method DELETE - Create an albums

   | REST HTTP Request | CRUD   | GIN Methods | Description      |
   |-------------------|--------|-------------|------------------|
   | /albums/:id       | Delete | GIN.DELETE  | Delete a albums  |

2. Implement the `deleteAlbums` function - Method DELETE

   ```go
    // /albums/:id => c.Param("id") => string value
    tempID := c.Param("id")

    // Convert string to int
    id, err := strconv.Atoi(tempID)
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
        return
    }

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for i, a := range albums {
        if a.ID == id {
            // i->length(albums)
            // albums = [1,2,3]
            // index = 0,1,2
            // tempAbum = len(albums) - 1 => 2
            tempAlbum := make([]album, len(albums)-1)

            // 0->i
            // found an ID at index = 1
            // cp albums[:index] => tempAlbum[:index]
            copy(tempAlbum[:i], albums[:i])

            // Move to index + 1
            // cp albums[index+1:] => tempAlbum[index:]
            copy(tempAlbum[i:], albums[i+1:])

            // assign albums to tempAlbum
            albums = tempAlbum

            // response to user request
            c.IndentedJSON(http.StatusOK, gin.H{"data": "Delete an albums"})
            return
        }
    }

   // response to user request
   c.IndentedJSON(http.StatusNotFound, gin.H{"data": "album not found"})
   ```

3. curl -XDELETE http://127.0.0.1:8080/albums/3

## Day 08

1. folder structure

   ```bash
   .
   ├── createAlbum.json
   ├── go.mod
   ├── go.sum
   ├── main.go
   ├── router
   │   └── route.go
   └── updateAlbum.json
   ```

2. Create `router/route.go` file
3. Content [`router/route.go`](rest/router/route.go) file
4. Content `main.go` file
   ```go
   package main

   import "github.com/monitor-cmr/web-w-go/router"

   func main() {
    // Get gin.Engine from router/route.go
    route := router.Router()

    // Run App
    route.Run(":8080")
   }
   ```

## Day 09

1. folder structure

   ```bash
   .
   ├── createAlbum.json
   ├── go.mod
   ├── go.sum
   ├── handler
   │   └── album.go
   ├── main.go
   ├── router
   │   └── route.go
   └── updateAlbum.json
   ```

2. Create `handler/album.go` file
3. Content [`handler/album.go`](rest/handler/album.go) file
4. Content [`router/route.go`](rest/router/route.go) file
   ```go
   package router

   import (
    "github.com/gin-gonic/gin"
    "github.com/monitor-cmr/web-w-go/handler"
   )

   // Router ...
   func Router() *gin.Engine {
    router := gin.Default()

    // Albums - Start

    // Create a new AlbumHandler
    al := handler.NewAlbumHandler()

    // curl -XGET http://127.0.0.1:8080/albums
    router.GET("/albums", al.GetAlbums)

    // curl -XPOST -d @createAlbum.json http://127.0.0.1:8080/albums
    router.POST("/albums", al.PostAlbums)

    // curl -XGET http://127.0.0.1:8080/albums/[1-n] - method GET
    router.GET("/albums/:id", al.GetAlbumsByID)

    // curl -XPUT -d @updateAlbum.json http://127.0.0.1:8080/albums/[1-n] - method PUT
    router.PUT("/albums/:id", al.PutAlbums)

    // curl -XDELETE http://127.0.0.1:8080/albums/[1-n] - method DELETE
    router.DELETE("/albums/:id", al.DeleteAlbums)

    // Albums - End

    return router
   }
   ```

## Day 10

1. folder structure

   ```bash
   .
   ├── createAlbum.json
   ├── domain
   │   └── album.go
   ├── go.mod
   ├── go.sum
   ├── handler
   │   └── album.go
   ├── main.go
   ├── router
   │   └── route.go
   └── updateAlbum.json
   ```

2. Create `domain/album.go` file

   ```go
   package domain

   // Album represents data about a record album.
   type Album struct {
    ID     int
    Title  string
    Artist string
    Price  float64
   }
   ```

3. Content [`handler/album.go`](rest/handler/album.go) file

## Day 11

1. Declare AlbumRepository Interface `domain/album.go` file

   ```go
    type AlbumRepository interface {
      SelectAll() ([]Album, error)
      Select(int) (*Album, error)
      Save(Album) (*int, error)
      Update(Album) error
      Delete(int) error
   }
   ```

## Day 12

1. Folder structure

   ```bash
   .
   ├── createAlbum.json
   ├── domain
   │   └── album.go
   ├── go.mod
   ├── go.sum
   ├── handler
   │   └── album.go
   ├── main.go
   ├── router
   │   └── route.go
   ├── storage
   │   └── memory
   │       └── album.go
   └── updateAlbum.json
   ```

2. Create [`storage/memory/album.go`](rest/storage/memory/album.go)
3. Implement all method in `AlbumRepository interface`

## Day 13

1. Update [`handler/album.go`](rest/handler/album.go) to use `AlbumRepository interface`

   ```go
   // AlbumHandler ...
   type AlbumHandler struct {
    repo domain.AlbumRepository
   }

   // NewAlbumHandler ...
   func NewAlbumHandler(re domain.AlbumRepository) *AlbumHandler {
    return &AlbumHandler{
        repo: re,
    }
   }
   ...
   ```

2. Update `router/route.go` to create `memory.NewAlbumMemory` storage

   ```go
   package router

   import (
    "github.com/gin-gonic/gin"
    "github.com/monitor-cmr/web-w-go/handler"
    "github.com/monitor-cmr/web-w-go/storage/memory"
   )

   // Router ...
   func Router() *gin.Engine {
    router := gin.Default()

    // Albums - Start

    // Create Memory storage
    store := memory.NewAlbumMemory()
    // Create a new AlbumHandler
    al := handler.NewAlbumHandler(store)

    // curl -XGET http://127.0.0.1:8080/albums
    router.GET("/albums", al.GetAlbums)

    // curl -XPOST -d @createAlbum.json http://127.0.0.1:8080/albums
    router.POST("/albums", al.PostAlbums)

    // curl -XGET http://127.0.0.1:8080/albums/[1-n] - method GET
    router.GET("/albums/:id", al.GetAlbumsByID)

    // curl -XPUT -d @updateAlbum.json http://127.0.0.1:8080/albums/[1-n] - method PUT
    router.PUT("/albums/:id", al.PutAlbums)

    // curl -XDELETE http://127.0.0.1:8080/albums/[1-n] - method DELETE
    router.DELETE("/albums/:id", al.DeleteAlbums)

    // Albums - End

    return router
   }
   ```

## Day 14

1. Folder structure

   ```bash
   .
   ├── createAlbum.json
   ├── domain
   │   └── album.go
   ├── go.mod
   ├── go.sum
   ├── handler
   │   └── album.go
   ├── main.go
   ├── router
   │   └── route.go
   ├── service
   │   └── album.go
   ├── storage
   │   └── memory
   │       └── album.go
   └── updateAlbum.json
   ```

2. Declare `Service Interface`

   ```go
   // AlbumServiceRepository ....
   type AlbumServiceRepository interface {
    FetchAll() ([]domain.Album, error)
    Fetch(int) (*domain.Album, error)
    Save(domain.Album) (*int, error)
    Update(domain.Album) error
    Delete(int) error
   }
   ```

3. Implement all method in `Service Interface` in [service/album.go](rest/service/album.go)

## Day 15

1. MySQL with Docker
   1. docker run --name mysql8 -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:8.0
   2. docker container inspect mysql8 |grep IPAddress
   3. docker exec -it mysql8 bash
   4. mysql -u root -p
   5. CREATE DATABASE MYSQLTEST;
   6. use MYSQLTEST
   7. Create albums table

      ```sql
      CREATE TABLE albums (
         id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
         title VARCHAR(30) NOT NULL,
         artist VARCHAR(30) NOT NULL,
         price FLOAT(50,2)
      );
      ```

   8. Insert data into table

      ```sql
      insert into albums(title, artist, price) values("Blue Train", "John Coltrane", 56.99);
        insert into albums(title, artist, price) values("Jeru", "Gerry Mulligan", 17.99);
        insert into albums(title, artist, price) values("Sarah Vaughan and Clifford", "Sarah Vaughan", 39.99);
      ```

   9. Create MySQL User
      ```sql
      CREATE USER 'alochym'@'%' IDENTIFIED BY 'alochym@123';
      GRANT ALL ON MYSQLTEST.* TO 'alochym'@'%';
      ```

2. Storage layer using MySQL
   1. Create [`storage/mysql/album.go`](rest/storage/mysql/album.go)
3. Create MySQL Connection function in `main.go` file

   ```go
   // MysqlConn ...
   func MysqlConn(host string, port int64, username, password, dbname string) *sql.DB {
    dbSource := fmt.Sprintf(
        "%s:%s@tcp(%s:%d)/%s?charset=utf8",
        username,
        password,
        host,
        port,
        dbname,
    )

    dbConn, err := sql.Open("mysql", dbSource)

    if err != nil {
        fmt.Println(err)
        panic(err)
    }

    return dbConn
   }
   ```

4. Implement all methods in [`AlbumRepository interface`](rest/storage/mysql/album.go)
5. Update router/router using Storage with MySQL
   ```go
    // Create Memory storage
   store := memory.NewAlbumMySQL()
   ```

