 
# Twitch-general-follows 
API that returns list of general follows by nickname  

## Configuration  

API needs config.toml and twitch-cfg.toml files

* config.yaml
```yaml
port: 8080
client_id: id
client_secret: secret
```
___

## Run Locally  

1. Clone the project  

~~~bash  
  git clone https://github.com/modaniru/twitch-general-follows
~~~
___
2. Go to the project directory  

~~~bash  
  cd twitch-general-follows
~~~
___
*hint: you can start this API and download all dependencies 
with one conmmand (if you can run Makefile commands)*
* run
~~~bash
make
~~~ 
* build
~~~bash
make build
~~~
below contains information how run API without Make.
___
3. Install dependencies  

~~~bash  
  go mod install
~~~
___
4. Start the server:
~~~bash  
  go run src/main.go
~~~
___
5. Build project
~~~bash  
  go build src/main.go
~~~
___
## Endpoints
* /get?login=modaniru&login=... - get general follows
```javascript
{
  "status_code": int
  "message": string
  "data":
    [
      {
        "id": string,
        "login": string,
        "display_name": string,
        "type": string,
        "broadcaster_type": string,
        "description": string,
        "profile_image_url": string,
        "offline_image_url": string,
        "view_count": int,
        "email": string,
        "created_at": string
      },
      ...
    ]
}
```
* /ping - test server endpoint
```javascript
{
  "status_code": int
  "message": string
  "data": null
}
```