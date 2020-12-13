package utils

import "strings"

var filters = [31]string{"NodeID",
	"HTMLURL",
	"FollowersURL",
	"ID",
	"NodeID",
	"GravatarID",
	"FollowingURL",
	"GistsURL",
	"StarredURL",
	"SubscriptionURL",
	"SubscriptionsURL",
	"OrganizationsURL",
	"ReposURL",
	"EventsURL",
	"ReceivedEventsURL",
	"SiteAdmin",
	"Type",
	// Repo Filters
	"GitURL",
	"SSHURL",
	"SvnUrl",
	"WatchersCount",
	"HasIssues",
	"HasProjects",
	"HasDownloads",
	"HasWiki",
	"CreatedAt",
	"UpdatedAt",
	"ReleasesURL",
	"DeploymentsURL",
	"NotificationURL",
	"Owner",
}

// Filter ...
func Filter(x string) bool {
	if strings.HasSuffix(x, "URL") {
		return true
	}
	for _, i := range filters {
		if i == x {
			return true
		}
	}
	return false
}
