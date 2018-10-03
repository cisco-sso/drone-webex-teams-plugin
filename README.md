# drone-webex-team

Forked from drone email plugin - https://github.com/Drillster/drone-email

Drone plugin to send build status notifications via [Cisco WebEx Team](https://www.webex.com/products/teams/index.html). For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

## Binary

Build the binary with the following command:

```sh
go get github.com/pkg/errors
go get github.com/aymerick/raymond
go get github.com/Sirupsen/logrus
go get github.com/joho/godotenv
go get github.com/urfave/cli
go get -d github.com/drone/drone-template-li
go build
```

## Docker

Build the docker image with the following commands:

```sh
docker build -t ciscosso/drone-webex-team:latest .
```

### Example
Execute from the working directory:

```sh
docker run --rm \
  -e PLUGIN_ACCESS_TOKEN="<mention WebEx Team bot's access token>" \
  -e PLUGIN_ROOM=WebEx Team Test \
  -e DRONE_REPO_OWNER=octocat \
  -e DRONE_REPO_NAME=hello-world \
  -e DRONE_COMMIT_SHA=7fd1a60b01f91b314f59955a4e4d4e80d8edf11d \
  -e DRONE_COMMIT_BRANCH=master \
  -e DRONE_COMMIT_AUTHOR=octocat \
  -e DRONE_COMMIT_AUTHOR_EMAIL=octocat@test.test \
  -e DRONE_BUILD_NUMBER=1 \
  -e DRONE_BUILD_STATUS=success \
  -e DRONE_BUILD_LINK=http://github.com/octocat/hello-world \
  -e DRONE_COMMIT_MESSAGE="Hello world!" \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  ciscosso/drone-webex-team
```
