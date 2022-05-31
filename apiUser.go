package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// User data
type User struct {
	ID              int           `json:"id"`
	Email           string        `json:"email"`
	Login           string        `json:"login"`
	FirstName       string        `json:"first_name"`
	LastName        string        `json:"last_name"`
	UsualFirstName  string        `json:"usual_first_name"`
	URL             string        `json:"url"`
	Phone           interface{}   `json:"phone"`
	Displayname     string        `json:"displayname"`
	UsualFullName   string        `json:"usual_full_name"`
	ImageURL        string        `json:"image_url"`
	Staff           bool          `json:"staff?"`
	CorrectionPoint int           `json:"correction_point"`
	PoolMonth       string        `json:"pool_month"`
	PoolYear        string        `json:"pool_year"`
	Location        interface{}   `json:"location"`
	Wallet          int           `json:"wallet"`
	AnonymizeDate   time.Time     `json:"anonymize_date"`
	Groups          []interface{} `json:"groups"`
	CursusUsers     []struct {
		ID           int           `json:"id"`
		BeginAt      time.Time     `json:"begin_at"`
		EndAt        interface{}   `json:"end_at"`
		BlackholedAt time.Time     `json:"blackholed_at"`
		Grade        interface{}   `json:"grade"`
		Level        float64       `json:"level"`
		Skills       []interface{} `json:"skills"`
		CursusID     int           `json:"cursus_id"`
		HasCoalition bool          `json:"has_coalition"`
		User         struct {
			ID    int    `json:"id"`
			Login string `json:"login"`
			URL   string `json:"url"`
		} `json:"user"`
		Cursus struct {
			ID        int       `json:"id"`
			CreatedAt time.Time `json:"created_at"`
			Name      string    `json:"name"`
			Slug      string    `json:"slug"`
		} `json:"cursus"`
	} `json:"cursus_users"`
	ProjectsUsers  []map[string]interface{} `json:"projects_users"`
	LanguagesUsers []struct {
		ID         int       `json:"id"`
		LanguageID int       `json:"language_id"`
		UserID     int       `json:"user_id"`
		Position   int       `json:"position"`
		CreatedAt  time.Time `json:"created_at"`
	} `json:"languages_users"`
	Achievements []interface{} `json:"achievements"`
	Titles       []interface{} `json:"titles"`
	TitlesUsers  []interface{} `json:"titles_users"`
	Partnerships []interface{} `json:"partnerships"`
	Patroned     []struct {
		ID          int       `json:"id"`
		UserID      int       `json:"user_id"`
		GodfatherID int       `json:"godfather_id"`
		Ongoing     bool      `json:"ongoing"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	} `json:"patroned"`
	Patroning       []interface{} `json:"patroning"`
	ExpertisesUsers []struct {
		ID          int       `json:"id"`
		ExpertiseID int       `json:"expertise_id"`
		Interested  bool      `json:"interested"`
		Value       int       `json:"value"`
		ContactMe   bool      `json:"contact_me"`
		CreatedAt   time.Time `json:"created_at"`
		UserID      int       `json:"user_id"`
	} `json:"expertises_users"`
	Campus []struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		TimeZone string `json:"time_zone"`
		Language struct {
			ID         int       `json:"id"`
			Name       string    `json:"name"`
			Identifier string    `json:"identifier"`
			CreatedAt  time.Time `json:"created_at"`
			UpdatedAt  time.Time `json:"updated_at"`
		} `json:"language"`
		UsersCount  int `json:"users_count"`
		VogsphereID int `json:"vogsphere_id"`
	} `json:"campus"`
	CampusUsers []struct {
		ID        int  `json:"id"`
		UserID    int  `json:"user_id"`
		CampusID  int  `json:"campus_id"`
		IsPrimary bool `json:"is_primary"`
	} `json:"campus_users"`
}

// GetUser fetch user data from 42 api based on login
func (s *APIClient) GetUser(login string) (User, error) {
	endpoint := fmt.Sprintf("%s/v2/users/%s", s.Url, login)
	client := &http.Client{}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return User{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.token.AccessToken))
	resp, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	if resp.StatusCode != 200 {
		return User{}, fmt.Errorf(resp.Status)
	}
	body, err := ReadHTTPResponse(resp)
	if err != nil {
		return User{}, err
	}
	var user User
	_ = json.Unmarshal(body, &user)
	return user, nil
}
