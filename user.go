package keybase

// UserLookupParams provides user lookup request parameters.
type UserLookupParams struct {
	Usernames      string `url:"usernames,omitempty"`
	Domain         string `url:"domain,omitempty"`
	Twitter        string `url:"twitter,omitempty"`
	Github         string `url:"github,omitempty"`
	Reddit         string `url:"reddit,omitempty"`
	HackerNews     string `url:"hackernews,omitempty"`
	Coinbase       string `url:"coinbase,omitempty"`
	KeyFingerprint string `url:"key_fingerprint,omitempty"`

	Fields string `url:"fields,omitempty"`
}

// UserLookupResponse provides the user lookup response.
// Per Keybase documentation, if the user making the request is the user being
// fetched, the 'Me' field will be populated. Otherwise, the 'them' field will be.
type UserLookupResponse struct {
	Status status `json:"status"`
	Them   []User `json:"them,omitempty"`
	Me     User   `json:"me,omitempty"`
}

// UserLookup wraps the user/lookup API endpoint.
func UserLookup(params UserLookupParams) (*UserLookupResponse, error) {
	r := new(UserLookupResponse)
	err := get("user/lookup", params, r)

	return r, err
}

// UserAutocompleteResponse contains a response to user/autocomplete requests.
type UserAutocompleteResponse struct {
	Status      status `json:"status"`
	Completions []struct {
		TotalScore float64 `json:"total_score"`
		Components struct {
			Username       acComponent `json:"username"`
			KeyFingerprint struct {
				acComponent
				Algo  int `json:"algo"`
				NBits int `json:"nbits"`
			} `json:"key_fingerprint"`
			FullName   acComponent `json:"full_name"`
			Github     acComponent `json:"github"`
			Reddit     acComponent `json:"reddit"`
			Twitter    acComponent `json:"twitter"`
			Coinbase   acComponent `json:"coinbase"`
			Hackernews acComponent `json:"hackernews"`
			Websites   []struct {
				acComponent
				Protocol string `json:"protocol"`
			} `json:"websites"`
		} `json:"components"`
		UID        string `json:"uid"`
		Thumbnail  string `json:"thumbnail"`
		IsFollowee bool   `json:"is_followee"`
	} `json:"completions"`
}

type acComponent struct {
	Val   string  `json:"val"`
	Score float64 `json:"score"`
}

// UserAutocomplete calls the user/autocomplete API endpoint.
func UserAutocomplete(query string) (*UserAutocompleteResponse, error) {
	r := new(UserAutocompleteResponse)
	err := get("user/autocomplete", struct {
		Query string `url:"q"`
	}{query}, r)

	return r, err
}

// UserDiscoverParams provides user/discover params.
type UserDiscoverParams UserLookupParams

// UserDiscoverResponse contains a user/discover response.
type UserDiscoverResponse struct {
	Status  status `json:"status"`
	Matches struct {
		Twitter    [][]discoverAccount `json:"twitter"`
		Github     [][]discoverAccount `json:"github"`
		Hackernews [][]discoverAccount `json:"hackernews"`
		Web        [][]discoverAccount `json:"web"`
		Coinbase   [][]discoverAccount `json:"coinbase"`
	} `json:"matches"`
}

type discoverAccount struct {
	Thumbnail string `json:"thumbnail"`
	Username  string `json:"username"`
	PublicKey struct {
		KeyFingerprint string `json:"key_fingerprint"`
		Bits           int    `json:"bits"`
		Algo           int    `json:"algo"`
	} `json:"public_key"`
	FullName     string `json:"full_name"`
	CTime        int    `json:"ctime"`
	RemoteProofs struct {
		DNS            []string `json:"dns"`
		GenericWebSite []struct {
			Hostname   string `json:"hostname"`
			Protocol   string `json:"protocol"`
			Searchable string `json:"searchable"`
		} `json:"generic_web_site"`
		Twitter    string `json:"twitter"`
		Github     string `json:"github"`
		Reddit     string `json:"reddit"`
		Hackernews string `json:"hackernews"`
		Coinbase   string `json:"coinbase"`
	} `json:"remote_proofs"`
}

// UserDiscover wraps the user/discover API endpoint.
func UserDiscover(params UserDiscoverParams) (*UserDiscoverResponse, error) {
	r := new(UserDiscoverResponse)
	err := get("user/discover", params, r)

	return r, err
}

// User is what the Keybase API defines as a "User Object".
// Documentation is available at https://keybase.io/docs/api/1.0/user_objects.
type User struct {
	ID     string `json:"id"`
	Basics struct {
		Ctime    int    `json:"ctime"`
		Mtime    int    `json:"mtime"`
		Salt     string `json:"salt"`
		UID      string `json:"uid"`
		Username string `json:"username"`
	} `json:"basics"`
	InvitationStats struct {
		Available int `json:"available"`
		Open      int `json:"open"`
		Power     int `json:"power"`
		Used      int `json:"used"`
	} `json:"invitation_stats"`
	Profile struct {
		Bio      string `json:"bio"`
		FullName string `json:"full_name"`
		Location string `json:"location"`
		Mtime    int    `json:"mtime"`
	} `json:"profile"`
	Emails struct {
		Primary struct {
			Email      string `json:"email"`
			IsVerified int    `json:"is_verified"`
		} `json:"primary"`
	} `json:"emails"`
	PublicKeys struct {
		Primary struct {
			KeyFingerprint string `json:"key_fingerprint"`
			KID            string `json:"kid"`
			KeyType        int    `json:"key_type"`
			Bundle         string `json:"bundle"`
			Ctime          int    `json:"ctime"`
			Mtime          int    `json:"mtime"`
		} `json:"primary"`
	} `json:"public_keys"`
	PrivateKeys struct {
		Bundle  string `json:"bundle"`
		KeyType int    `json:"key_type"`
		KID     string `json:"kid"`
		Ctime   int    `json:"ctime"`
		Mtime   int    `json:"mtime"`
	} `json:"private_keys"`
	CryptoCurrencyAddress struct {
		Bitcoin struct {
			Address string `json:"address"`
			SigID   string `json:"sig_id"`
		} `json:"bitcoin"`
	} `json:"cryptocurrency_addresses"`
}
