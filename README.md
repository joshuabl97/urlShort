# [URL Shortener](https://github.com/gophercises/urlshort)

# Table of Contents
1. [Usage](#usage)
2. [Endpoints](#endpoints)
    - [GET /{endpoint}](#get-endpoint)
    - [GET /shortcuts](#get-shortcuts)
    - [POST /shortcut](#post-shortcut)
3. [Exercise Modifications](#exercise-modifications)
4. [Exercise Details](#exercise-details)

## Usage
This is a simple URL shortening app that allows users to create URL shortcuts using an API. 

Simply copy these files into a directory and run ```go run main.go``` to start the server on localhost:8080

Going to an endpoint that is stored in the SQLite .db file generated will redirect the user to the URL stored. You can view all the endpoints and URLs by calling [GET /shortcuts](#get-shortcuts). Full CRUD documentation can be found [here](#endpoints).

You can specify a flag to choose your own port
Please note: this is not required - default = 8080
```
go run main.go -p 9090
```

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

## Exercise modifications

I've modified the example for this gophercises [URL Shortener](https://github.com/gophercises/urlshort) exercise (details below) to use the Chi router instead of Gorilla Mux. I've turned this into a basic CRUD api which demonstrates my ability to work with JSON, http requests, and databases rather than inputting the data using flags for YAML/JSON.

[chi router](https://github.com/go-chi/chi)

[gorilla mux](https://github.com/gorilla/mux)

[link to original exercise (gophercises)](https://github.com/gophercises/urlshort)

[SQLite](https://www.sqlite.org/index.html)


## Exercise details

The goal of this exercise is to create an [http.Handler](https://golang.org/pkg/net/http/#Handler) that will look at the path of any incoming web request and determine if it should redirect the user to a new page, much like URL shortener would.

For instance, if we have a redirect setup for `/dogs` to `https://www.somesite.com/a-story-about-dogs` we would look for any incoming web requests with the path `/dogs` and redirect them.

To complete this exercises you will need to implement the stubbed out methods in [handler.go](https://github.com/gophercises/urlshort/blob/master/handler.go). There are a good bit of comments explaining what each method should do, and there is also a [main/main.go](https://github.com/gophercises/urlshort/blob/master/main/main.go) source file that uses the package to help you test your code and get an idea of what your program should be doing.

I suggest first commenting out all of the code in main.go related to the `YAMLHandler` function and focusing on implementing the `MapHandler` function first.

Once you have that working, focus on parsing the YAML using the [gopkg.in/yaml.v2](https://godoc.org/gopkg.in/yaml.v2) package. *Note: You will need to `go get` this package if you don't have it already.*

After you get the YAML parsing down, try to convert the data into a map and then use the MapHandler to finish the YAMLHandler implementation. Eg you might end up with some code like this:

```go
func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
  parsedYaml, err := parseYAML(yaml)
  if err != nil {
    return nil, err
  }
  pathMap := buildMap(parsedYaml)
  return MapHandler(pathMap, fallback), nil
}
```

But in order for this to work you will need to create functions like `parseYAML` and `buildMap` on your own. This should give you ample experience working with YAML data.


## Bonus

As a bonus exercises you can also...

1. Update the [main/main.go](https://github.com/gophercises/urlshort/blob/master/main/main.go) source file to accept a YAML file as a flag and then load the YAML from a file rather than from a string.
2. Build a JSONHandler that serves the same purpose, but reads from JSON data.
3. Build a Handler that doesn't read from a map but instead reads from a database. Whether you use BoltDB, SQL, or something else is entirely up to you.