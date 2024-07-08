package bootdev

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type BootDevUser struct {
	StripeCustomerID        string    `json:"StripeCustomerID"`
	IsAdmin                 bool      `json:"IsAdmin"`
	ReferringUserUUID       any       `json:"ReferringUserUUID"`
	DiscordUserID           any       `json:"DiscordUserID"`
	DiscordUserHandle       any       `json:"DiscordUserHandle"`
	GithubAuthToken         any       `json:"GithubAuthToken"`
	SyncedGoogleID          any       `json:"SyncedGoogleID"`
	SyncedGithubID          any       `json:"SyncedGithubID"`
	MembershipExpiresAt     any       `json:"MembershipExpiresAt"`
	MembershipPaymentType   any       `json:"MembershipPaymentType"`
	DiscordBotPrivate       bool      `json:"DiscordBotPrivate"`
	NotificationPreferences any       `json:"NotificationPreferences"`
	Email                   string    `json:"Email"`
	Currency                any       `json:"Currency"`
	Xp                      int       `json:"XP"`
	Level                   int       `json:"Level"`
	XPForLevel              int       `json:"XPForLevel"`
	XPTotalForLevel         int       `json:"XPTotalForLevel"`
	Role                    string    `json:"Role"`
	Gems                    int       `json:"Gems"`
	Armor                   int       `json:"Armor"`
	IsMember                bool      `json:"IsMember"`
	UUID                    string    `json:"UUID"`
	RecruitersCanContact    bool      `json:"RecruitersCanContact"`
	CreatedAt               time.Time `json:"CreatedAt"`
	UpdatedAt               time.Time `json:"UpdatedAt"`
	FirstName               string    `json:"FirstName"`
	LastName                string    `json:"LastName"`
	Handle                  string    `json:"Handle"`
	Bio                     any       `json:"Bio"`
	JobTitle                any       `json:"JobTitle"`
	Location                any       `json:"Location"`
	City                    any       `json:"City"`
	Country                 any       `json:"Country"`
	TwitterHandle           any       `json:"TwitterHandle"`
	LinkedinURL             any       `json:"LinkedinURL"`
	GithubHandle            string    `json:"GithubHandle"`
	WebsiteURL              string    `json:"WebsiteURL"`
	ProfileImageURL         string    `json:"ProfileImageURL"`
	ResumeURL               any       `json:"ResumeURL"`
	IsRecruiter             any       `json:"IsRecruiter"`
}

func GetUser(username string) (BootDevUser, error) {
	user := BootDevUser{}
	if len(username) == 0 {
		return user, fmt.Errorf("invalid username")
	}
	encodedUsername := url.PathEscape(strings.ToLower(username))
	url := fmt.Sprintf("https://api.boot.dev/v1/users/public/%s", encodedUsername)
	resp, err := http.Get(url)
	if err != nil {
		return user, err
	}
	if resp.StatusCode != 200 {
		return user, fmt.Errorf("user %s not found", strings.ToLower(username))
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}
