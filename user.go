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

// User is what the Keybase API defines as a "User Object".
// Documentation is available at https://keybase.io/docs/api/1.0/user_objects.
type User struct {
	ID                    string                  `json:"id"`
	Basics                basics                  `json:"basics"`
	InvitationStats       invitationStats         `json:"invitation_stats"`
	Profile               profile                 `json:"profile"`
	Emails                emails                  `json:"emails"`
	PublicKeys            publicKeys              `json:"public_keys"`
	PrivateKeys           privateKeys             `json:"private_keys"`
	CryptoCurrencyAddress cryptoCurrencyAddresses `json:"cryptocurrency_addresses"`
}

type basics struct {
	Ctime    int    `json:"ctime"`
	Mtime    int    `json:"mtime"`
	Salt     string `json:"salt"`
	UID      string `json:"uid"`
	Username string `json:"username"`
}

type invitationStats struct {
	Available int `json:"available"`
	Open      int `json:"open"`
	Power     int `json:"power"`
	Used      int `json:"used"`
}

type profile struct {
	Bio      string `json:"bio"`
	FullName string `json:"full_name"`
	Location string `json:"location"`
	Mtime    int    `json:"mtime"`
}

type emails struct {
	Primary email `json:"primary"`
}

type email struct {
	Email      string `json:"email"`
	IsVerified int    `json:"is_verified"`
}

type publicKeys struct {
	Primary publicKey `json:"primary"`
}

type publicKey struct {
	KeyFingerprint string `json:"key_fingerprint"`
	KID            string `json:"kid"`
	KeyType        int    `json:"key_type"`
	Bundle         string `json:"bundle"`
	Ctime          int    `json:"ctime"`
	Mtime          int    `json:"mtime"`
}

type privateKeys struct {
	Bundle  string `json:"bundle"`
	KeyType int    `json:"key_type"`
	KID     string `json:"kid"`
	Ctime   int    `json:"ctime"`
	Mtime   int    `json:"mtime"`
}

type cryptoCurrencyAddresses struct {
	Bitcoin ccAddress `json:"bitcoin"`
}

type ccAddress struct {
	Address string `json:"address"`
	SigID   string `json:"sig_id"`
}

// UserLookup wraps the user/lookup API endpoint.
func UserLookup(params UserLookupParams) (*UserLookupResponse, error) {
	r := new(UserLookupResponse)
	err := get("user/lookup", params, r)

	return r, err
}