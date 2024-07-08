package ranks

import (
	"bootdev-clan/internal/bootdev"
	"cmp"
	"log"
	"slices"
	"sync"
	"time"
)

type Ranks struct {
	LastUpdated time.Time
	Usernames   []string
	Users       []bootdev.BootDevUser
	mu          sync.Mutex
}

func (r *Ranks) UpdateRanks() {
	r.mu.Lock()
	defer r.mu.Unlock()
	// don't fetch new data if it's under 2 minutes old
	if time.Since(r.LastUpdated) < 2*time.Minute {
		return
	}
	// get the latest ranks for each user
	users := []bootdev.BootDevUser{}
	for _, username := range r.Usernames {
		if len(username) > 0 {
			user, err := bootdev.GetUser(username)
			if err != nil {
				log.Printf("ranks: %s", err)
			} else {
				users = append(users, user)
			}
		}
	}
	// sort by xp, descending
	slices.SortFunc(users,
		func(a, b bootdev.BootDevUser) int {
			return cmp.Compare(b.Xp, a.Xp)
		})
	r.Users = users
	r.LastUpdated = time.Now()
	log.Println("ranks: update succeeded")
}
