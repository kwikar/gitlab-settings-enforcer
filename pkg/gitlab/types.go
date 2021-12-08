package gitlab

import (
	"github.com/xanzy/go-gitlab"
)

// currentState stores current Project state for each project interacted with
type ProjectSettings struct {
	Approval gitlab.ProjectApprovals `json:"approval_settings,omitempty"`
	General  gitlab.Project          `json:"project_settings,omitempty"`
}

type groupsClient interface {
	ListGroupProjects(gid interface{}, opt *gitlab.ListGroupProjectsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Project, *gitlab.Response, error)
	ListSubgroups(gid interface{}, opt *gitlab.ListSubgroupsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Group, *gitlab.Response, error)
}

type projectsClient interface {
	ChangeApprovalConfiguration(pid interface{}, opt *gitlab.ChangeApprovalConfigurationOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectApprovals,
		*gitlab.Response, error)
	GetApprovalConfiguration(pid interface{}, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectApprovals, *gitlab.Response, error)
	GetProject(pid interface{}, opt *gitlab.GetProjectOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Project, *gitlab.Response, error)
	EditProject(pid interface{}, opt *gitlab.EditProjectOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Project, *gitlab.Response, error)
	EditProjectPushRule(pid interface{}, opt *gitlab.EditProjectPushRuleOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectPushRules,
		*gitlab.Response, error)
	GetProjectPushRules(pid interface{}, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectPushRules, *gitlab.Response, error)
}

type protectedBranchesClient interface {
	ProtectRepositoryBranches(pid interface{}, opt *gitlab.ProtectRepositoryBranchesOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProtectedBranch,
		*gitlab.Response, error)
	UnprotectRepositoryBranches(pid interface{}, branch string, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error)
	GetProtectedBranch(pid interface{}, branch string, options ...gitlab.RequestOptionFunc) (*gitlab.ProtectedBranch, *gitlab.Response, error)
}

type protectedTagsClient interface {
	ProtectRepositoryTags(pid interface{}, opt *gitlab.ProtectRepositoryTagsOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProtectedTag,
		*gitlab.Response, error)
	UnprotectRepositoryTags(pid interface{}, tag string, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error)
	GetProtectedTag(pid interface{}, tag string, options ...gitlab.RequestOptionFunc) (*gitlab.ProtectedTag, *gitlab.Response, error)
}

type branchesClient interface {
	CreateBranch(pid interface{}, opt *gitlab.CreateBranchOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Branch, *gitlab.Response, error)
	GetBranch(pid interface{}, branch string, options ...gitlab.RequestOptionFunc) (*gitlab.Branch, *gitlab.Response, error)
}

var (
	listGroupProjectOps = &gitlab.ListGroupProjectsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 100,
		},
		Archived: gitlab.Bool(false),
	}

	listSubgroupOps = &gitlab.ListSubgroupsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 100,
		},
	}
)
