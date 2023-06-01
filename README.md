 
# Twitch-general-follows 
API that returns list of general follows by nickname  

## Configuration  

API needs config.toml and twitch-cfg.toml files

* config.toml
```toml
port=":8080"
```
* twitch-cgf.toml
```
user_info="https://api.twitch.tv/helix/users"
user_get_follow_list="https://api.twitch.tv/helix/users/follows"
get_token="https://id.twitch.tv/oauth2/token"
validate_token="https://id.twitch.tv/oauth2/validate"


client_id="id"
client_secret="secret"
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
3. Install dependencies  

~~~bash  
  go get
~~~
___
4. Start the server:

~~~bash  
  make run
~~~
or
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
* /ping - test server endpoint