# URL Shortener

# Table of Contents
- [Usage](#usage)
  - [Start the server](#start-the-server-on-localhost8080)
  - [Specify a port](#specify-a-port)
  - [Packages used](#packages-used)
  - [Prepopulate database](#prepopulate-database)
- [Endpoints](#endpoints)
  - [GET /{endpoint}](#get-endpoint)
  - [GET /shortcuts](#get-shortcuts)
  - [POST /shortcut](#post-shortcut)
  - [PUT /{endpoint}](#put-endpoint)
  - [DELETE /{endpoint}](#delete-endpoint)

## Usage
This is a simple URL shortening app that allows users to create URL shortcuts using an API. 

Simply copy these files into a directory and run ```go run main.go``` to start the server on localhost:8080

### Start the server on localhost:8080
```
go run main.go
```

Going to an endpoint that is stored in the SQLite .db file generated will redirect the user to the URL stored

![redirect example](/redirect.png)

The above would redirect a user to google if GET /shortcuts contained:
```json
{
    "shortcuts":[
      {"endpoint":"example1","url":"https://www.google.com"},
      {"endpoint":"example2","url":"https://example.com/"}
    ]
  }
```

 You can view all the endpoints and URLs by calling [GET /shortcuts](#get-shortcuts). Full CRUD API documentation can be found [here](#endpoints).

### Specify a port:
- Default port: 8080
```
go run main.go -port_number 9090
```

### Prepopulate database:
Specify a yaml file to prepopulate the database with

Example YAML file
```yaml
- endpoint: "urlshort"
  url: "https://github.com/joshuabl97/urlShort#prepopulate-the-database"
- endpoint: "urlshort-inspo"
  url: "https://github.com/gophercises/urlshort"
```


Usage 
```
// starts the server
 go run main.go -yaml_filepath example.yaml 
```

### Packages used:
- [chi](https://github.com/go-chi/chi)
- [SQLite](https://www.sqlite.org/index.html)
- [zerolog](https://github.com/rs/zerolog)


## Endpoints

### GET /{endpoint}

Redirects you to the url where the endpoint is linked in the database (view all endpoints by calling GET /shortcuts)

- **Endpoint:** `/{endpoint}`
- **Method:** `GET`

Example Usage:
In the browser go to an endpoint listed in shortcuts and it will redirect you to the URL
i.e 
localhost:8080/example1 would bring you to google.com

### GET /shortcuts

Get all the shortcuts available to be used by GET /{endpoint}

- **Endpoint:** `/shortcuts`
- **Method:** `GET`

- **Response:**
  ```json
  {
    "shortcuts":[
      {"endpoint":"example1","url":"https://www.google.com"},
      {"endpoint":"example2","url":"https://example.com/"}
    ]
  }
  ```

Example Usage:
```
curl localhost:8080/shortcuts | jq
```

### POST /shortcut

Creates a new shortcut available to be used by [GET /{endpoint}](#get-endpoint)

You can view all the endpoints and URLs by calling [GET /shortcuts](#get-shortcuts)

- **Endpoint:** `/shortcut`
- **Method:** `POST`
- **Parameters:**
    - **required**
    - **Content-Type: application/json**

| Parameter   | Type      | Description                                                 |
|-------------|-----------|-------------------------------------------------------------|
| endpoint    | string    | endpoint to be used in [GET /{endpoint}](#get-endpoint)     |
| url         | string    | url that [GET /{endpoint}](#get-endpoint) will route you to |

Example Usage:
```
curl -X POST 
-H "Content-Type: application/json" 
-d '{"endpoint":"test","url":"https://godaddy.com"}' 
http://localhost:8080/shortcut
```

### PUT /{endpoint}

Updates an existing shortcut available to be used by [GET /{endpoint}](#get-endpoint)

You can view all the endpoints and URLs by calling [GET /shortcuts](#get-shortcuts)

- **Endpoint:** `/{endpoint}`
- **Method:** `PUT`
- **Parameters:**
    - **required**
    - **Content-Type: application/json**

| Parameter   | Type      | Description                                                 |
|-------------|-----------|-------------------------------------------------------------|
| endpoint    | string    | endpoint to be used in [GET /{endpoint}](#get-endpoint)     |
| url         | string    | the url you would like to replace                           |

Example Usage:
```
curl -X PUT 
-H "Content-Type: application/json" 
-d '{"endpoint":"example1","url":"https://anotherSite.com"}' 
http://localhost:8080/shortcut
```

### DELETE /{endpoint}
Deletes an endpoint stored in the database

You can view all the endpoints and URLs by calling [GET /shortcuts](#get-shortcuts)

- **Endpoint:** `/{endpoint}`
- **Method:** `DELETE`

Example Usage:
```
curl -X DELETE localhost:8080/example1
```