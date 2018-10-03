package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/drone/drone-template-lib/template"
	"net/http"
	"encoding/json"
	"bytes"
	"errors"
)

type (
	Repo struct {
		FullName string
		Owner    string
		Name     string
		SCM      string
		Link     string
		Avatar   string
		Branch   string
		Private  bool
		Trusted  bool
	}

	Remote struct {
		URL string
	}

	Author struct {
		Name   string
		Email  string
		Avatar string
	}

	Commit struct {
		Sha     string
		Ref     string
		Branch  string
		Link    string
		Message string
		Author  Author
	}

	Build struct {
		Number   int
		Event    string
		Status   string
		Link     string
		Created  float64
		Started  float64
		Finished float64
	}

	PrevBuild struct {
		Status string
		Number int
	}

	PrevCommit struct {
		Sha string
	}

	Prev struct {
		Build  PrevBuild
		Commit PrevCommit
	}

	Job struct {
		Status   string
		ExitCode int
		Started  int64
		Finished int64
	}

	Yaml struct {
		Signed   bool
		Verified bool
	}

	Config struct {
		Token          string
		Room           string
		RoomId         string
		ApiEndPoint    string
		Attachment     string
		Body           string
	}

	Plugin struct {
		Repo        Repo
		Remote      Remote
		Commit      Commit
		Build       Build
		Prev        Prev
		Job         Job
		Yaml        Yaml
		Tag         string
		PullRequest int
		DeployTo    string
		Config      Config
	}
	
	Message struct {
		RoomId		string   `json:"roomId"`
		Markdown	string   `json:"markdown"`
	}
	
	Rooms struct {
		Items  []struct {
			Id				string   `json:"id"`
			Title			string   `json:"title"`
			Type			string   `json:"type"`
			IsLocked		bool     `json:"isLocked"`
			TeamId			string   `json:"teamId"`
			LastActivity	string   `json:"lastActivity"`
			Created			string   `json:"created"`
		} `json:"items"`
	}
)

// Exec will send emails over SMTP
func (p Plugin) Exec() error {
	if p.Config.Token == "" {
		log.Errorf("Error spark access token is mandatory")
		return nil
	}
	if p.Config.Room == "" {
		if p.Config.RoomId == "" {
			log.Errorf("Error spark room or room_id is mandatory")
			return nil
		}
	}
	if p.Config.ApiEndPoint == "" {
		log.Errorf("Error spark api endpoint is mandatory")
		return nil
	}

	type Context struct {
		Repo        Repo
		Remote      Remote
		Commit      Commit
		Build       Build
		Prev        Prev
		Job         Job
		Yaml        Yaml
		Tag         string
		PullRequest int
		DeployTo    string
	}
	ctx := Context{
		Repo:        p.Repo,
		Remote:      p.Remote,
		Commit:      p.Commit,
		Build:       p.Build,
		Prev:        p.Prev,
		Job:         p.Job,
		Yaml:        p.Yaml,
		Tag:         p.Tag,
		PullRequest: p.PullRequest,
		DeployTo:    p.DeployTo,
	}

	// Render body in HTML and plain text
	renderedBody, err := template.RenderTrim(p.Config.Body, ctx)
	if err != nil {
		log.Errorf("Could not render body template: %v", err)
		return err
	}
	
	// Create http client
	client := &http.Client{}
	
	// Get spark room id
	roomId := ""
	request, err := http.NewRequest("GET", p.Config.ApiEndPoint + "/rooms", nil)
	request.Header.Add("Authorization", "Bearer " + p.Config.Token)
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		log.Info(err)
		return err
	}
	defer response.Body.Close()
	rooms := Rooms{}
	if response.StatusCode == 200 {
		err := json.NewDecoder(response.Body).Decode(&rooms)
		if err != nil {
			return err
		}
		for _, room := range rooms.Items {
			if room.Title == p.Config.Room {
				roomId = room.Id
			}
		}
	} else {
		return errors.New("Error while getting room id from spark, status - " + string(response.StatusCode))
	}
	if roomId == "" {
		roomId = p.Config.RoomId
	}
	
	// Send to spark
	payload := Message{
		RoomId: roomId,
		Markdown: renderedBody,
	}
	payloadByets, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	request, err = http.NewRequest("POST", p.Config.ApiEndPoint + "/messages", bytes.NewBuffer(payloadByets))
	request.Header.Add("Authorization", "Bearer " + p.Config.Token)
	request.Header.Add("Content-Type", "application/json")
	response, err = client.Do(request)
	if err != nil {
		log.Info(err)
		return err
	}
	defer response.Body.Close()
	
	if response.StatusCode != 200 {
		return errors.New("Error while updating spark, status - " + string(response.StatusCode))
	}

	return nil
}
