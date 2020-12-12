package utils

var filters = [16]string{"NodeID",
	"HTMLURL",
	"FollowersURL",
	"ID",
	"NodeID",
	"GravatarID",
	"FollowingURL",
	"GistsURL",
	"StarredURL",
	"SubscriptionURL",
	"OrganizationsURL",
	"ReposURL",
	"EventsURL",
	"ReceivedEventsURL",
	"SiteAdmin",
	"Type"}

func Filter(x string) bool {
	for _, i := range filters {
		if i == x {
			return true
		}
	}
	return false
}
