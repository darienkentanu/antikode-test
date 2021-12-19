# antikode-test
crud-api-test

[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=blue)](https://pkg.go.dev/gorm.io/gorm?tab=doc)

# Table of Content
- [Description](#description)
- [How to use](#how-to-use)
- [Endpoints](#endpoints)
- [Credits](#credits)

# Description
This is submition of coding test at Antikode
```
In this project i use MVC architecture because it's the most common achitecture used by programmer
```

# How to use
- Install Go and MySQL
- Create database on mysql named `antikode`
- Clone this repository in your $PATH:
```
$ git clone https://github.com/darienkentanu/antikode-test
```
- to run this project first you must add your google api key to your system enviroment variable
```
-linux/mac => $export GOOGLE_API_KEY={`YOUR GOOGLE API KEY`}
-windows => >set GOOGLE_API_KEY={`YOUR GOOGLE API KEY`}
type it without '{}' and '``'
```

How to run: 
```
$ ./run.sh
```
-or
```
$ go run app/main.go
```

<br>


# Endpoints

| Method | Endpoint | Description
|:-----|:--------|:----------|
| POST | /brands/add | Add list of brand|
| GET | /brands/getall | Get list of all brands|
| PUT | /brands/edit/{id} | Update brand by id|
| DELETE | /brands/delete{id} | Delete brand by id|
|---|---|---|
| POST | /outlets/add/{brandname} | Add list of outlet by brandname|
| GET | /outlets/getall | Get list of all outlets |
| PUT | /outlets/edit/{id} | Update outlet by id |
| DELETE | /outlets/delete/{id} | Delete outlet by id |
|---|---|---|
| POST | /products/add/{brandname} | Add list of product by brandname|
| GET | /products/getall | Get list of all products |
| PUT | /products/edit/{id} | Update product by id |
| DELETE | /products/delete/{id} | Delete product by id |
|:-----|:--------|:----------|


<br>

## Credits

- [Darien Kentanu](https://github.com/darienkentanu) (Author and maintainer)