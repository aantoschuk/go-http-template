# go-template

## Overview

### Router

I use Chi as a router to access features such as route grouping and middleware, 
since it is possible to implement these, but I don't want to have to implement such methods all the time.

### Cors

The `cors.go` file contains default values for *cors.Options*. 
Add env variables and provide a fallback if env is not set.

### .env

I don't configure the `.env` filesa, as this could limit flexibility.


