package gobacklog

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/url"
)

// PrintResponseJSON is bool
var PrintResponseJSON = false

// Space returns
func (c *Client) Space() (*Space, error) {
	bytes, er := c.Get("/api/v2/space", url.Values{})

	if er != nil {
		return nil, er
	}

	var space *Space
	json.Unmarshal(bytes, &space)
	return space, nil
}

// SpaceActivities `/api/v2/space/activities`
func (c *Client) SpaceActivities(option *ActivitiesOption) (ActivitySlice, error) {
	var params url.Values
	if option == nil {
		params = url.Values{}
	}

	params, er := option.Values()
	if er != nil {
		return nil, er
	}

	bytes, er := c.Get("/api/v2/space/activities", params)

	if er != nil {
		return nil, er
	}

	if PrintResponseJSON {
		fmt.Println(string(bytes))
	}

	var model ActivitySlice
	json.Unmarshal(bytes, &model)

	return model, nil
}

// SpaceNotification /api/v2/space/notification
func (c *Client) SpaceNotification() (*SpaceNotification, error) {
	bytes, er := c.Get("/api/v2/space/notification", url.Values{})

	if er != nil {
		return nil, er
	}

	if PrintResponseJSON {
		fmt.Println(string(bytes))
	}

	var spaceNotification *SpaceNotification
	json.Unmarshal(bytes, &spaceNotification)
	return spaceNotification, nil
}

// DiskUsage /api/v2/space/diskUsage
func (c *Client) DiskUsage() (*DiskUsage, error) {
	bytes, er := c.Get("/api/v2/space/diskUsage", url.Values{})

	if er != nil {
		return nil, er
	}

	if PrintResponseJSON {
		fmt.Println(string(bytes))
	}

	var diskUsage *DiskUsage
	json.Unmarshal(bytes, &diskUsage)
	return diskUsage, nil
}

// Users /api/v2/users
func (c *Client) Users() (UserSlice, error) {
	bytes, er := c.Get("/api/v2/users", url.Values{})

	if er != nil {
		return nil, er
	}

	if PrintResponseJSON {
		fmt.Println(string(bytes))
	}

	var users UserSlice
	json.Unmarshal(bytes, &users)

	return users, nil
}

// User /api/v2/users/:userId
func (c *Client) User(userID int) (*User, error) {
	endpoint := fmt.Sprintf("/api/v2/users/%d", userID)
	bytes, er := c.Get(endpoint, url.Values{})

	if er != nil {
		return nil, er
	}

	if PrintResponseJSON {
		fmt.Println(string(bytes))
	}

	var user *User
	json.Unmarshal(bytes, &user)

	return user, nil
}

// Myself returns
func (c *Client) Myself() (*User, error) {
	bytes, er := c.Get("/api/v2/users/myself", url.Values{})

	if er != nil {
		return nil, er
	}

	if PrintResponseJSON {
		fmt.Println(string(bytes))
	}

	var user *User
	json.Unmarshal(bytes, &user)

	return user, nil
}

// ProjectsWithOption returns project information.
// /api/v2/projects
func (c *Client) ProjectsWithOption(option *ProjectsOption) (ProjectSlice, error) {
	params, er := option.Values()
	if er != nil {
		return nil, er
	}

	bytes, er := c.Get("/api/v2/projects", params)

	if er != nil {
		return nil, er
	}

	if PrintResponseJSON {
		fmt.Println(string(bytes))
	}

	var projs ProjectSlice
	json.Unmarshal(bytes, &projs)

	return projs, nil
}

// ProjectWithID returns project information.
// /api/v2/projects/:projectIdOrKey
func (c *Client) ProjectWithID(projectID int) (*Project, error) {
	endpoint := fmt.Sprintf("/api/v2/projects/%d", projectID)
	bytes, er := c.Get(endpoint, url.Values{})

	if er != nil {
		return nil, er
	}

	if PrintResponseJSON {
		fmt.Println(string(bytes))
	}

	var proj *Project
	json.Unmarshal(bytes, &proj)

	return proj, nil
}

// ProjectWithKey returns project information.
// /api/v2/projects/:projectIdOrKey
func (c *Client) ProjectWithKey(projectKey string) (*Project, error) {
	if len(projectKey) == 0 {
		return nil, errors.New("invalid arguments")
	}

	endpoint := fmt.Sprintf("/api/v2/projects/%s", projectKey)
	bytes, er := c.Get(endpoint, url.Values{})

	if er != nil {
		return nil, er
	}

	if PrintResponseJSON {
		fmt.Println(string(bytes))
	}

	var proj *Project
	json.Unmarshal(bytes, &proj)

	return proj, nil
}

// Issues is
func (c *Client) Issues() (IssueSlice, error) {
	bytes, er := c.Get("/api/v2/issues", url.Values{})

	if er != nil {
		return nil, er
	}

	if PrintResponseJSON {
		fmt.Println(string(bytes))
	}

	var issues IssueSlice
	json.Unmarshal(bytes, &issues)

	return issues, nil
}

// IssuesWithOption is
func (c *Client) IssuesWithOption(opt *IssuesOption) (IssueSlice, error) {
	params, er := opt.Values()
	if er != nil {
		return nil, er
	}
	bytes, er := c.Get("/api/v2/issues", params)

	if er != nil {
		return nil, er
	}

	if PrintResponseJSON {
		fmt.Println(string(bytes))
	}

	var issues IssueSlice
	json.Unmarshal(bytes, &issues)

	return issues, nil
}

// IssueWithKey is
func (c *Client) IssueWithKey(issueIDOrKey string) (*Issue, error) {
	bytes, er := c.Get("/api/v2/issues/"+issueIDOrKey, url.Values{})

	if er != nil {
		return nil, er
	}

	if PrintResponseJSON {
		fmt.Println(string(bytes))
	}

	var issue *Issue
	json.Unmarshal(bytes, &issue)

	return issue, nil

}

// DownloadAttachment returns
// /api/v2/issues/:issueIdOrKey/attachments/:attachmentId
func (c *Client) DownloadAttachment(issueIDOrKey string, attachmentID int) (io.ReadCloser, string, error) {
	if len(issueIDOrKey) == 0 {
		return nil, "", errors.New("issueIDOrKey is too short")
	}
	endpoint := fmt.Sprintf("/api/v2/issues/%s/attachments/%d", issueIDOrKey, attachmentID)
	resp, er := c.executeReturnsResponse("GET", endpoint, url.Values{})

	if er != nil {
		return nil, "", er
	}

	a, _ := resp.Header["Content-Disposition"]
	_, params, err := mime.ParseMediaType(a[0])

	if err != nil {
		return nil, "", er
	}

	return resp.Body, params["filename"], nil
}

// StatusesWithProject returns project statuses.
// /api/v2/projects/:projectIdOrKey/statuses
func (c *Client) StatusesWithProject(projectKey string) (*StatusSlice, error) {
	if len(projectKey) == 0 {
		return nil, errors.New("invalid arguments")
	}

	endpoint := fmt.Sprintf("/api/v2/projects/%s/statuses", projectKey)
	bytes, er := c.Get(endpoint, url.Values{})

	if er != nil {
		return nil, er
	}

	if PrintResponseJSON {
		fmt.Println(string(bytes))
	}

	var statuses *StatusSlice
	json.Unmarshal(bytes, &statuses)

	return statuses, nil
}
