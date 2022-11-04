package webgit

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kluctl/template-controller/api/v1alpha1"
	"github.com/kluctl/template-controller/controllers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"time"
)

type Note interface {
	GetId() string
	GetBody() string

	UpdateBody(body string) error
	GetCreatedAt() time.Time
}

type WebgitInterface interface {
	GetProject(projectId string) (ProjectInterface, error)
}

type ProjectInterface interface {
	ListMergeRequests(state v1alpha1.MergeRequestState) ([]MergeRequestInterface, error)
	GetMergeRequest(mrId string) (MergeRequestInterface, error)
}

type MergeRequestInfo struct {
	ID           int                        `json:"id"`
	TargetBranch string                     `json:"targetBranch"`
	SourceBranch string                     `json:"sourceBranch"`
	Title        string                     `json:"title"`
	State        v1alpha1.MergeRequestState `json:"state"`
	CreatedAt    metav1.Time                `json:"createdAt"`
	UpdatedAt    metav1.Time                `json:"updatedAt"`
	Author       string                     `json:"author"`
	Labels       []string                   `json:"labels"`
	Draft        bool                       `json:"draft"`
}

type MergeRequestInterface interface {
	Info() (*MergeRequestInfo, error)

	HasApproved() (bool, error)
	Approve() error
	Unapprove() error

	CreateMergeRequestNote(body string) (Note, error)
	GetMergeRequestNote(noteId string) (Note, error)
	ListMergeRequestNotes() ([]Note, error)
	ListMergeRequestNotesAfter(t time.Time) ([]Note, error)
}

func BuildWebgit(ctx context.Context, client client.Client, namespace string, spec any, defaults any) (ProjectInterface, error) {
	merged, err := mergeSpec(spec, defaults)
	if err != nil {
		return nil, err
	}

	if m, ok := merged["gitlab"]; ok {
		var gitlab v1alpha1.GitlabProject
		err = mapToObject(m, &gitlab)
		if err != nil {
			return nil, err
		}

		project, err := BuildWebgitGitlab(ctx, client, namespace, gitlab)
		if err != nil {
			return nil, err
		}
		return project, nil
	} else if m, ok := merged["github"]; ok {
		var github v1alpha1.GithubProject
		err = mapToObject(m, &github)
		if err != nil {
			return nil, err
		}

		project, err := BuildWebgitGithub(ctx, client, namespace, github)
		if err != nil {
			return nil, err
		}
		return project, nil
	} else {
		return nil, fmt.Errorf("no git project spec provided")
	}
}

func BuildWebgitMergeRequest(ctx context.Context, client client.Client, namespace string, spec any, defaults any) (MergeRequestInterface, error) {
	project, err := BuildWebgit(ctx, client, namespace, spec, defaults)
	if err != nil {
		return nil, err
	}

	merged, err := mergeSpec(spec, defaults)
	if err != nil {
		return nil, err
	}

	if m, ok := merged["gitlab"]; ok {
		var gitlab v1alpha1.GitlabMergeRequestRef
		err = mapToObject(m, &gitlab)
		if err != nil {
			return nil, err
		}
		if gitlab.MergeRequestId == nil {
			return nil, fmt.Errorf("missing mergeRequestId")
		}
		return project.GetMergeRequest(fmt.Sprintf("%d", *gitlab.MergeRequestId))
	} else if m, ok := merged["github"]; ok {
		var github v1alpha1.GithubPullRequestRef
		err = mapToObject(m, &github)
		if err != nil {
			return nil, err
		}
		if github.PullRequestId == nil {
			return nil, fmt.Errorf("missing pullRequestId")
		}
		return project.GetMergeRequest(fmt.Sprintf("%d", *github.PullRequestId))
	} else {
		return nil, fmt.Errorf("no pullRequest spec provided")
	}
}

func mergeSpec(spec any, defaults any) (map[string]any, error) {
	var err error
	var mergedMap map[string]any

	if defaults == nil {
		mergedMap, err = objectToMap(spec)
		if err != nil {
			return nil, err
		}
	} else {
		mergedMap, err = objectToMap(defaults)
		if err != nil {
			return nil, err
		}
		m2, err := objectToMap(spec)
		if err != nil {
			return nil, err
		}
		controllers.MergeMap2(mergedMap, m2, true)
	}
	return mergedMap, nil
}

func objectToMap(obj any) (map[string]any, error) {
	b, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	var ret map[string]any
	err = json.Unmarshal(b, &ret)
	if err != nil {
		return nil, err
	}
	if ret == nil {
		ret = map[string]any{}
	}
	return ret, nil
}

func mapToObject(m any, out any) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, out)
	if err != nil {
		return err
	}
	return nil
}
