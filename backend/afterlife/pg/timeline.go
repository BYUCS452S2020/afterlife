package pg

import (
	"context"
	"errors"
	"time"

	"github.com/byuoitav/afterlife"
)

func (*DataService) Timeline(ctx context.Context, token string) (afterlife.Timeline, error) {
	if token != "12345" {
		return afterlife.Timeline{}, errors.New("invalid token")
	}

	return afterlife.Timeline{
		afterlife.Event{
			Name: "Wife Graduation",
			At:   time.Now().Add(30 * 24 * time.Hour),
			Type: afterlife.EventTypeEmail,
			Email: &afterlife.EventEmail{
				To:      []string{"dgrandall@me.com"},
				Subject: "Congrats!",
				Body:    "Yay! You graduated!",
			},
		},
		afterlife.Event{
			Name: "Starting Kindergarten",
			At:   time.Now().Add(120 * 24 * time.Hour),
			Type: afterlife.EventTypeEmail,
			Email: &afterlife.EventEmail{
				To:      []string{"dgrandall@me.com"},
				Subject: "Good luck!",
				Body:    "Good luck at school today!\nLove,\nDanny\n\n\n\n\n\n\nlol i'm danny!\n\n\n test again!\n\nhi",
			},
		},
		afterlife.Event{
			Name: "Starting Kindergarten",
			At:   time.Now().Add(120 * 24 * time.Hour),
			Type: afterlife.EventTypeEmail,
			Email: &afterlife.EventEmail{
				To:      []string{"dgrandall@me.com"},
				Subject: "Good luck!",
				Body:    "Good luck at school today!\nLove,\nDanny\n\n\n\n\n\n\nlol i'm danny!\n\n\n test again!\n\nhi",
			},
		},
		afterlife.Event{
			Name: "Starting Kindergarten",
			At:   time.Now().Add(120 * 24 * time.Hour),
			Type: afterlife.EventTypeEmail,
			Email: &afterlife.EventEmail{
				To:      []string{"dgrandall@me.com"},
				Subject: "Good luck!",
				Body:    "Good luck at school today!\nLove,\nDanny\n\n\n\n\n\n\nlol i'm danny!\n\n\n test again!\n\nhi",
			},
		},
		afterlife.Event{
			Name: "Starting Kindergarten",
			At:   time.Now().Add(120 * 24 * time.Hour),
			Type: afterlife.EventTypeEmail,
			Email: &afterlife.EventEmail{
				To:      []string{"dgrandall@me.com"},
				Subject: "Good luck!",
				Body:    "Good luck at school today!\nLove,\nDanny\n\n\n\n\n\n\nlol i'm danny!\n\n\n test again!\n\nhjkfldsjkklklkalsklakslskldasklsaklsaklsaklsakslakaslksalkaslkaslksalkaslksalkaslkaslkask asklsakaslkasl saklsakalsksal saklsaksal akslsaklask lsaklsak alsks lasklsaklaskslak lsak lskasl ksalkas l ksalkaslksalksa lsalkas ksak aslksa lkaslkasksal ksa",
			},
		},
	}, nil
}
