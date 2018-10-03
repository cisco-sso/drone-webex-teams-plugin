package main

const (
	// DefaultApi is the default spark api endpoint
	DefaultApi = "https://api.ciscospark.com/v1"
)

// DefaultTemplate is the default body template to use for the email
const DefaultTemplate = `
[**[{{ build.status }}] {{ repo.owner }}/{{ repo.name }} #{{ build.number }}**]({{ build.link }})
* Commit: [{{ commit.message }}]({{ commit.link }})
* Author: {{ commit.author.name }} {{ commit.author.email }}
* Branch: {{ commit.branch }}
* Event: {{ build.event }}
* Started at: {{ datetime build.created "Mon Jan 2 15:04:05 MST 2006" "Local" }}
`