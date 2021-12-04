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
