package main

import (
	"errors"
	"github.com/cbrgm/githubevents/githubevents"
	"github.com/google/go-github/v43/github"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

const (
	// labels
	labelEvent  = "event"
	labelAction = "action"
)

func metricsHandler(r *prometheus.Registry) func(w http.ResponseWriter, r *http.Request) (int, error) {

	githubEventsCounter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "github_webhook_events_total",
		Help: "The total number of webhook events received by type.",
	}, []string{labelEvent, labelAction})

	r.MustRegister(githubEventsCounter)

	handle := githubevents.New(config.GithubWebhookSecret)
	handle.OnBeforeAny(
		func(deliveryID string, eventName string, event interface{}) error {
			switch event := event.(type) {
			case *github.BranchProtectionRuleEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.CheckRunEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.CheckSuiteEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.CommitCommentEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.CreateEvent:
				githubEventsCounter.WithLabelValues(eventName, "").Inc()
				return nil
			case *github.DeleteEvent:
				githubEventsCounter.WithLabelValues(eventName, "").Inc()
				return nil
			case *github.DeployKeyEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.DeploymentEvent:
				githubEventsCounter.WithLabelValues(eventName, "").Inc()
				return nil
			case *github.DeploymentStatusEvent:
				githubEventsCounter.WithLabelValues(eventName, "").Inc()
				return nil
			case *github.DiscussionEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.ForkEvent:
				githubEventsCounter.WithLabelValues(eventName, "").Inc()
				return nil
			case *github.GitHubAppAuthorizationEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.GollumEvent:
				githubEventsCounter.WithLabelValues(eventName, "").Inc()
				return nil
			case *github.InstallationEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.InstallationRepositoriesEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.IssueCommentEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.IssuesEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.LabelEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.MarketplacePurchaseEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.MemberEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.MembershipEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.MetaEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.MilestoneEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.OrganizationEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.OrgBlockEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.PackageEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.PageBuildEvent:
				githubEventsCounter.WithLabelValues(eventName, "").Inc()
				return nil
			case *github.PingEvent:
				githubEventsCounter.WithLabelValues(eventName, "").Inc()
				return nil
			case *github.ProjectEvent:
				githubEventsCounter.WithLabelValues(eventName, "").Inc()
				return nil
			case *github.ProjectCardEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.ProjectColumnEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.PublicEvent:
				githubEventsCounter.WithLabelValues(eventName, "").Inc()
				return nil
			case *github.PullRequestEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.PullRequestReviewEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.PullRequestReviewCommentEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.PushEvent:
				githubEventsCounter.WithLabelValues(eventName, "").Inc()
				return nil
			case *github.ReleaseEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.RepositoryDispatchEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.RepositoryEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.RepositoryVulnerabilityAlertEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.StarEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.StatusEvent:
				githubEventsCounter.WithLabelValues(eventName, "").Inc()
				return nil
			case *github.TeamEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.TeamAddEvent:
				githubEventsCounter.WithLabelValues(eventName, "").Inc()
				return nil
			case *github.WatchEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.WorkflowJobEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			case *github.WorkflowDispatchEvent:
				githubEventsCounter.WithLabelValues(eventName, "").Inc()
				return nil
			case *github.WorkflowRunEvent:
				githubEventsCounter.WithLabelValues(eventName, *event.Action).Inc()
				return nil
			}
			return nil
		},
	)

	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		if r.Method != http.MethodPost {
			return http.StatusMethodNotAllowed, errors.New("method not allowed")
		}

		err := handle.HandleEventRequest(r)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		return http.StatusOK, nil
	}
}
