package activity

//go:generate mockgen -source=activity.go -destination=mock/activity_mock.go -package=mockactivity
import (
	"encoding/json"
	"net/http"
)

type IActivity interface {
	GetActivity(url string) (*ActivityType, error)
}

type ActivityType struct {
	Activity     string  `json:"activity"`
	Type         string  `json:"type"`
	Participants int     `json:"participants"`
	Price        float32 `json:"price"`
	Link         string  `json:"link"`
	Key          string  `json:"key"`
	Acessibility float32 `json:"accessibility"`
}

func (a *ActivityType) GetActivity(url string) (*ActivityType, error) {

	//url := "https://www.boredapi.com/api/activity"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := new(http.Client)

	//res, err := client.Get(url)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(a)
	if err != nil {
		return nil, err
	}

	return a, nil
}
