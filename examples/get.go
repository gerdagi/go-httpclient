package examples

import "fmt"

type Endpoints struct {
	CurrentUserUrl    string `json:"current_user_url"`
	AuthorizationsUrl string `json:"authorizations_url"`
	RepositoryUrl     string `json:"repository_url"`
}

func GetEndpoints() (*Endpoints, error) {
	response, err := httpClient.Get("https://api.github.com", nil)
	if err != nil {
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Status Code: %d", response.StatusCode()))
	fmt.Println(fmt.Sprintf("Status : %s", response.Status()))
	fmt.Println(fmt.Sprintf("Body : %s", response.String()))

	var endPoints Endpoints
	if err := response.UnmarshalJson(&endPoints); err != nil {
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Repositories URL: %s", endPoints.RepositoryUrl))

	return &endPoints, nil
}
