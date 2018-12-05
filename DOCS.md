Use the Cisco Spark plugin for sending build status notifications via Cisco WebEx Team

## Config
You can configure the plugin using the following parameters:

* **access_token** 	- Access token of WebEx Team Bot
* **room** 		- WebEx Team Room Name
* **room_id** 		- WebEx Team Room ID - Optional if room name not chossen
* **api_endpoint** 	- WebEx Team API End point - Default: https://api.ciscospark.com/v1
* **body** 		- WebEx Team body template

## Secrets or environment variable
You can use the following secrets to secure the sensitive configuration of this plugin:

* **access_token** 	- Access token of WebEx Team Bot or
* **SPARK_ACCESS_TOKEN** - Variable in drone agent

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
pipeline:
  notify spark:
    image: ciscosso/drone-webex-team
    pull: true
    room: "Cisco WebEx Bot Test"
    access_token: "<mention WebEx Team bot's access token>"
    when:
      status: [ changed, failure, success ]
```

### Custom Templates

In some cases you may want to customize the look and feel of the email message
so you can use custom templates. For the use case we expose the following
additional parameters, all of the accept a custom handlebars template, directly
provided as a string or as a remote URL which gets fetched and parsed:

* **body** - A handlebars template to create a custom template. For more
  details take a look at the [docs](http://handlebarsjs.com/). You can see the
  default template [here](defaults.go)
