package worth

import(
	"context"
	"testing"
	"encoding/json"
)

func TestGetToken(t *testing.T) {
	cfg := NewConfiguration()
	client := NewAPIClient(cfg)

	accessKey := "30ec62934072bd5f3313cd90958318fa"
	secret_key := "f815ab5f87091bdb09cd7cb869cce8bc"

	r, _, err := client.AuthenticatorApi.GenToken(context.Background(), accessKey, secret_key)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(r.Data["token"])
	token, ok := r.Data["token"].(string)
	if !ok {
		t.Fatal(err)
	}

	{
		actionObj := map[string]interface{}{
			"ACTION": "SET_FEED_DEVICE_CONTROL",
			"VALUE": map[string]interface{}{
				"weight": 1,
			},
		}

		data, err := json.Marshal(actionObj)
		if err != nil {
			t.Fatal(err)
		}

		action := ActionReq{
			DeviceId: "3105000000000123",
			Data: string(data),
		}

		ctx := context.WithValue(context.Background(), ContextAccessToken, token)

		r, _, err := client.DeviceApi.SendAction(ctx, action)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(r)
	}
}
