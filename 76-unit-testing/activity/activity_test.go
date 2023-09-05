package activity_test

import (
	"demo/activity"
	mockactivity "demo/activity/mock"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestGetActivity(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockActivity := mockactivity.NewMockIActivity(ctrl)

	expectedOutput := &activity.ActivityType{Activity: "Have a jam session with your friends", Type: "music", Participants: 5, Price: 0.1, Key: "2715253", Acessibility: 0.3}

	mockActivity.EXPECT().GetActivity("https://www.boredapi.com/api/activity").Return(expectedOutput, nil).MaxTimes(1)

}
