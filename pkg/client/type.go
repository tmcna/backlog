package client

// Error structure.
type Error struct {
	Message  string `json:"message"`
	Code     int    `json:"code"`
	MoreInfo string `json:"moreInfo"`
}

// ErrorResponse structure.
// see https://developer.nulab.com/docs/backlog/error-response/#error-response
type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

// IssueResponseCount
const (
	IssueResponseCount        int = 10
	ActivityResponseCount     int = 10
	NotificationResponseCount int = 10
)

// DisplayOrderDesc
const (
	DisplayOrderDesc int = 0
	DisplayOrderAsc  int = 1
)

// ActivityResponse structure is the response body of /api/v2/space/activities.
type ActivityResponse struct {
	ID      int             `json:"id"`
	Project ProjectResponse `json:"project"`
	Type    int             `json:"type"`
	Reason  int             `json:"reason"`
	User    UserResponse    `json:"createdUser"`
}

// CategoryResponse structure is the response body of /api/v2/projects/:projectIdOrKey/categories.
type CategoryResponse struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	DisplayOrder int    `json:"displayOrder"`
}

// CommentResponse structure is the response body of /api/v2/issues/:issueIdOrKey/comments
type CommentResponse struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

// IssueResponse structure is the response body of /api/v2/issues.
type IssueResponse struct {
	ID             int               `json:"id"`
	ProjectID      int               `json:"projectId"`
	IssueKey       string            `json:"issueKey"`
	KeyID          int               `json:"keyId"`
	IssueType      IssueTypeResponse `json:"issueType"`
	Summary        string            `json:"summary"`
	Description    string            `json:"description"`
	Resolutions    string            `json:"resolutions"`
	Priority       PriorityResponse  `json:"priority"`
	Status         StatusResponse    `json:"status"`
	Assignee       UserResponse      `json:"assignee"`
	StartDate      string            `json:"startDate"`
	DueDate        string            `json:"dueDate"`
	EstimatedHours string            `json:"estimatedHours"`
}

// IssueTypeResponse structure is the response body of /api/v2/projects/:projectIdOrKey/issueTypes.
type IssueTypeResponse struct {
	ID           int    `json:"id"`
	ProjectID    int    `json:"projectId"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	DisplayOrder int    `json:"displayOrder"`
}

// MilestoneResponse structure is the response body of /api/v2/projects/:projectIdOrKey/versions.
type MilestoneResponse struct {
	ID             int    `json:"id"`
	ProjectID      int    `json:"projectId"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	StartDate      string `json:"startDate"`
	ReleaseDueDate string `json:"releaseDueDate"`
	Archived       bool   `json:"archived"`
	DisplayOrder   int    `json:"displayOrder"`
}

// NotificationResponse structure is the response body of /api/v2/notifications.
type NotificationResponse struct {
	ID    int           `json:"id"`
	Issue IssueResponse `json:"issue"`
}

// PriorityResponse structure is the response body of api/v2/priorities.
type PriorityResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ProjectResponse structure is the response body of /api/v2/projects.
type ProjectResponse struct {
	ID         int    `json:"id"`
	ProjectKey string `json:"projectKey"`
	Name       string `json:"name"`
}

// SpaceUsageResponse structure is the response body of api/v2/space/diskUsage.
type SpaceUsageResponse struct {
	Capacity   int `json:"capacity"`
	Issue      int `json:"issue"`
	Wiki       int `json:"wiki"`
	File       int `json:"file"`
	Subversion int `json:"subversion"`
	Git        int `json:"git"`
	GitLFS     int `json:"gitLFS"`
	Details    []struct {
		ProjectID  int `json:"projectId"`
		Issue      int `json:"issue"`
		Wiki       int `json:"wiki"`
		File       int `json:"file"`
		Subversion int `json:"subversion"`
		Git        int `json:"git"`
		GitLFS     int `json:"gitLFS"`
	} `json:"details"`
}

// StatusResponse structure is the response body of api/v2/projects/:projectIdOrKey/statuses
type StatusResponse struct {
	ID           int    `json:"id"`
	ProjectID    int    `json:"projectId"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	DisplayOrder int    `json:"displayOrder"`
}

// UserResponse structure is the response body of /api/v2/users.
type UserResponse struct {
	ID          int    `json:"id"`
	UserID      string `json:"userId"`
	Name        string `json:"name"`
	RoleType    int    `json:"roleType"`
	Lang        string `json:"lang"`
	MailAddress string `json:"mailAddress"`
}
