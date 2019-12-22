package gobacklog

import (
	"time"
)

// Issue represents
// +gen * slice:"Where,Count,SortBy,GroupBy[string],GroupBy[int],GroupBy[bool],first,MaxBy,MinBy,Distinct,DistinctBy,Shuffle"
type Issue struct {
	ID             *int             `json:"id,omitempty"`
	ProjectID      *int             `json:"projectId,omitempty"`
	IssueKey       *string          `json:"issueKey,omitempty"`
	KeyID          *int             `json:"keyId,omitempty"`
	IssueType      IssueType        `json:"issueType,omitempty"`
	Summary        *string          `json:"summary,omitempty"`
	Description    *string          `json:"description,omitempty"`
	Resolution     Resolution       `json:"resolution,omitempty"`
	Priority       Priority         `json:"priority,omitempty"`
	Status         Status           `json:"status,omitempty"`
	Assignee       Assignee         `json:"assignee,omitempty"`
	Category       CategorySlice    `json:"category,omitempty"`
	Versions       VersionSlice     `json:"versions,omitempty"`
	Milestone      VersionSlice     `json:"milestone,omitempty"`
	StartDate      *time.Time       `json:"startDate,omitempty"`
	DueDate        *time.Time       `json:"dueDate,omitempty"`
	EstimatedHours *int             `json:"estimatedHours,omitempty"`
	ActualHours    *int             `json:"actualHours,omitempty"`
	ParentIssueID  *int             `json:"parentIssueId,omitempty"`
	CreatedUser    *User            `json:"createdUser,omitempty"`
	Created        *time.Time       `json:"created,omitempty"`
	UpdatedUser    *User            `json:"updatedUser,omitempty"`
	Updated        *time.Time       `json:"updated,omitempty"`
	CustomFields   CustomFieldSlice `json:"customFields,omitempty"`
	Attachments    AttachmentSlice  `json:"attachments,omitempty"`
	SharedFiles    SharedFileSlice  `json:"sharedFiles,omitempty"`
	Stars          StarSlice        `json:"stars,omitempty"`
}

// Issues represents
type Issues struct {
	Issues IssueSlice
}

// IssueType represents
type IssueType struct {
	ID           *int    `json:"id,omitempty"`
	ProjectID    *int    `json:"projectId,omitempty"`
	Name         *string `json:"name,omitempty"`
	Color        *string `json:"color,omitempty"`
	DisplayOrder *int    `json:"displayOrder,omitempty"`
}

// Resolution represents
type Resolution struct {
	ID   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Priority represents
type Priority struct {
	ID   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Status represents
// +gen * slice:"Where,Count,SortBy,GroupBy[string],GroupBy[int],GroupBy[bool],first,MaxBy,MinBy,Distinct,DistinctBy,Shuffle"
type Status struct {
	ID           *int    `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	ProjectID    *int    `json:"projectId,omitempty"`
	Color        *string `json:"color,omitempty"`
	DisplayOrder *int    `json:"displayOrder,omitempty"`
}

// Assignee represents
type Assignee struct {
	ID          *int    `json:"id,omitempty"`
	UserID      *string `json:"userId,omitempty"`
	Name        *string `json:"name,omitempty"`
	RoleType    *int    `json:"roleType,omitempty"`
	Lang        *string `json:"lang,omitempty"`
	MailAddress *string `json:"mailAddress,omitempty"`
}

// Category represents
// +gen * slice:"Where,Count,SortBy,GroupBy[string],GroupBy[int],GroupBy[bool],first,MaxBy,MinBy,Distinct,DistinctBy,Shuffle"
type Category struct {
	ID           *int    `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	DisplayOrder *int    `json:"displayOrder,omitempty"`
}

// Version represents
// +gen * slice:"Where,Count,SortBy,GroupBy[string],GroupBy[int],GroupBy[bool],first,MaxBy,MinBy,Distinct,DistinctBy,Shuffle"
type Version struct {
	ID             *int       `json:"id,omitempty"`
	ProjectID      *int       `json:"projectId,omitempty"`
	Name           *string    `json:"name,omitempty"`
	Description    *string    `json:"description,omitempty"`
	StartDate      *time.Time `json:"startDate,omitempty"`
	ReleaseDueDate *time.Time `json:"releaseDueDate,omitempty"`
	Archived       *bool      `json:"archived,omitempty"`
	DisplayOrder   *int       `json:"displayOrder,omitempty"`
}

// User represents
// +gen * slice:"Where,Count,SortBy,GroupBy[string],GroupBy[int],GroupBy[bool],first,MaxBy,MinBy,Distinct,DistinctBy,Shuffle"
type User struct {
	ID           *int    `json:"id,omitempty"`
	UserID       *string `json:"userId,omitempty"`
	Name         *string `json:"name,omitempty"`
	RoleType     *int    `json:"roleType,omitempty"`
	Lang         *string `json:"lang,omitempty"`
	MailAddress  *string `json:"mailAddress,omitempty"`
	NulabAccount *string `json:"nulabAccount,omitempty"`
}

// CustomField represents
// +gen * slice:"Where,Count,SortBy,GroupBy[string],GroupBy[int],GroupBy[bool],first,MaxBy,MinBy,Distinct,DistinctBy,Shuffle"
type CustomField struct {
	ID          *int                   `json:"id,omitempty"`
	Name        *string                `json:"name,omitempty"`
	FieldTypeID *int                   `json:"fieldTypeId,omitempty"`
	Value       *CustomFieldValueSlice `json:"value,omitempty"`
}

// CustomFieldValue represents
// +gen * slice:"Where,Count,SortBy,GroupBy[string],GroupBy[int],GroupBy[bool],first,MaxBy,MinBy,Distinct,DistinctBy,Shuffle"
type CustomFieldValue struct {
	ID           *int    `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	DisplayOrder *int    `json:"displayOrder,omitempty"`
}

// Attachment represents
// +gen * slice:"Where,Count,SortBy,GroupBy[string],GroupBy[int],GroupBy[bool],first,MaxBy,MinBy,Distinct,DistinctBy,Shuffle"
type Attachment struct {
	ID          *int       `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Size        *int64     `json:"size,omitempty"`
	CreatedUser *User      `json:"createdUser,omitempty"`
	Created     *time.Time `json:"created,omitempty"`
}

// SharedFile represents
// +gen * slice:"Where,Count,SortBy,GroupBy[string],GroupBy[int],GroupBy[bool],first,MaxBy,MinBy,Distinct,DistinctBy,Shuffle"
type SharedFile struct {
	ID          *int       `json:"id,omitempty"`
	Type        *string    `json:"type,omitempty"`
	Directory   *string    `json:"dir,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Size        *int64     `json:"size,omitempty"`
	CreatedUser *User      `json:"createdUser,omitempty"`
	Created     *time.Time `json:"created,omitempty"`
	UpdatedUser *User      `json:"updatedUser,omitempty"`
	Updated     *time.Time `json:"updated,omitempty"`
}

// Star represents
// +gen * slice:"Where,Count,SortBy,GroupBy[string],GroupBy[int],GroupBy[bool],first,MaxBy,MinBy,Distinct,DistinctBy,Shuffle"
type Star struct {
	ID        *int       `json:"id,omitempty"`
	Comment   *string    `json:"comment,omitempty"`
	URL       *string    `json:"url,omitempty"`
	Title     *string    `json:"name,omitempty"`
	Created   *time.Time `json:"created,omitempty"`
	Presenter *User      `json:"presenter,omitempty"`
}

// Space represents
type Space struct {
	SpaceKey           *string    `json:"spaceKey,omitempty"`
	Name               *string    `json:"name,omitempty"`
	OwnerID            *int       `json:"ownerId,omitempty"`
	Lang               *string    `json:"lang,omitempty"`
	Timezone           *string    `json:"timezone,omitempty"`
	ReportSendTime     *string    `json:"reportSendTime,omitempty"`
	TextFormattingRule *string    `json:"textFormattingRule,omitempty"`
	Created            *time.Time `json:"created,omitempty"`
	Updated            *time.Time `json:"updated,omitempty"`
}

// SpaceNotification represents
type SpaceNotification struct {
	Content *string    `json:"content,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
}

// DiskUsage represents
type DiskUsage struct {
	Capacity    *int                 `json:"capacity,omitempty"`
	Issue       *int                 `json:"issue,omitempty"`
	Wiki        *int                 `json:"wiki,omitempty"`
	File        *int                 `json:"file,omitempty"`
	Subversion  *int                 `json:"subversion,omitempty"`
	Git         *int                 `json:"git,omitempty"`
	PullRequest *int                 `json:"pullRequest,omitempty"`
	Details     DiskUsageDetailSlice `json:"details,omitempty"`
}

// DiskUsageDetail represents
// +gen * slice:"Where,Count,SortBy,GroupBy[string],GroupBy[int],GroupBy[bool],first,MaxBy,MinBy,Distinct,DistinctBy,Shuffle"
type DiskUsageDetail struct {
	ProjectID   *int `json:"projectId,omitempty"`
	Issue       *int `json:"issue,omitempty"`
	Wiki        *int `json:"wiki,omitempty"`
	File        *int `json:"file,omitempty"`
	Subversion  *int `json:"subversion,omitempty"`
	Git         *int `json:"git,omitempty"`
	PullRequest *int `json:"pullRequest,omitempty"`
}

// Total returns total disk usage, byte unit.
func (d *DiskUsageDetail) Total() int {
	total := *d.Issue +
		*d.Wiki +
		*d.File +
		*d.Subversion +
		*d.Git +
		*d.PullRequest
	return total
}

// Project represents
// +gen * slice:"Where,Count,SortBy,GroupBy[string],GroupBy[int],GroupBy[bool],first,MaxBy,MinBy,Distinct,DistinctBy,Shuffle"
type Project struct {
	ID                                *int    `json:"id,omitempty"`
	ProjectKey                        *string `json:"projectKey,omitempty"`
	Name                              *string `json:"name,omitempty"`
	ChartEnabled                      *bool   `json:"chartEnabled,omitempty"`
	SubtaskingEnabled                 *bool   `json:"subtaskingEnabled,omitempty"`
	ProjectLeaderCanEditProjectLeader *bool   `json:"projectLeaderCanEditProjectLeader,omitempty"`
	UseWikiTreeView                   *bool   `json:"useWikiTreeView,omitempty"`
	TextFormattingRule                *string `json:"textFormattingRule,omitempty"`
	Archived                          *bool   `json:"archived,omitempty"`
	DisplayOrder                      *int    `json:"displayOrder,omitempty"`
}

// Activity represents
// +gen * slice:"Where,Count,SortBy,GroupBy[string],GroupBy[int],GroupBy[bool],first,MaxBy,MinBy,Distinct,DistinctBy,Shuffle"
type Activity struct {
	ID      *int     `json:"id,omitempty"`
	Project *Project `json:"project,omitempty"`
	Type    *int     `json:"type,omitempty"`
	Content *Content `json:"content,omitempty"`
	// Notifications *Notification `json:"notifications,omitempty"` // TODO
	CreatedUser *User      `json:"createdUser,omitempty"`
	Created     *time.Time `json:"created,omitempty"`
}

// Content represents
// +gen * slice:"Where,Count,SortBy,GroupBy[string],GroupBy[int],GroupBy[bool],first,MaxBy,MinBy,Distinct,DistinctBy,Shuffle"
type Content struct {
	ID          *int            `json:"id,omitempty"`
	KeyID       *int            `json:"key_id,omitempty"`
	Summary     *string         `json:"summary,omitempty"`
	Description *string         `json:"description,omitempty"`
	Comment     *Comment        `json:"comment,omitempty"`
	Changes     *ChangeSlice    `json:"changes,omitempty"`
	Attachments AttachmentSlice `json:"attachments,omitempty"`
	SharedFile  *SharedFile     `json:"shared_files,omitempty"`
}

// Change represents
// +gen * slice:"Where,Count,SortBy,GroupBy[string],GroupBy[int],GroupBy[bool],first,MaxBy,MinBy,Distinct,DistinctBy,Shuffle"
type Change struct {
	Field    *string `json:"field,omitempty"`
	NewValue *string `json:"new_value,omitempty"`
	OldValue *string `json:"old_value,omitempty"`
	Type     *string `json:"type,omitempty"`
}

// Comment represents
// +gen * slice:"Where,Count,SortBy,GroupBy[string],GroupBy[int],GroupBy[bool],first,MaxBy,MinBy,Distinct,DistinctBy,Shuffle"
type Comment struct {
	ID      *int    `json:"id,omitempty"`
	Content *string `json:"content,omitempty"`
}
