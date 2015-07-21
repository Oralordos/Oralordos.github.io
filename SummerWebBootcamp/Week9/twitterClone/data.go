package twitter

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

func getProfileByEmail(ctx context.Context, email string) (*profile, error) {
	key := datastore.NewKey(ctx, "profile", email, 0, nil)
	var p profile
	err := datastore.Get(ctx, key, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func getProfileByUsername(ctx context.Context, usr string) (*profile, error) {
	query := datastore.NewQuery("profile")
	p := []profile{}
	_, err := query.Filter("Username =", usr).GetAll(ctx, &p)
	if err != nil {
		return nil, err
	}
	if len(p) == 0 {
		return nil, datastore.ErrNoSuchEntity
	} else if len(p) > 1 {
		return nil, datastore.ErrInvalidKey
	}
	return &p[0], nil
}

func getTweets(ctx context.Context, email string) ([]tweet, error) {
	query := datastore.NewQuery("Tweets")
	t := []tweet{}
	if email != "" {
		key := datastore.NewKey(ctx, "profile", email, 0, nil)
		query = query.Ancestor(key)
	}
	query = query.Order("-SubmitTime")
	keys, err := query.GetAll(ctx, &t)
	if err != nil {
		return nil, err
	}
	for i := range t {
		p, err := getProfileByEmail(ctx, keys[i].Parent().StringID())
		if err != nil {
			return nil, err
		}
		t[i].Username = p.Username
	}
	return t, nil
}

func postTweet(ctx context.Context, t *tweet, email string) error {
	profileKey := datastore.NewKey(ctx, "profile", email, 0, nil)
	key := datastore.NewIncompleteKey(ctx, "Tweets", profileKey)
	_, err := datastore.Put(ctx, key, t)
	return err
}
