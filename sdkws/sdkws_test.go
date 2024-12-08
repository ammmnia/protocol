package sdkws_test

import (
	"encoding/json"
	. "github.com/ammmnia/protocol/sdkws"
	"google.golang.org/protobuf/encoding/protojson"
	"testing"
)

func TestNewGroupInfo(t *testing.T) {
	tt := GroupInfo{
		Status: 0,
	}
	s, _ := json.Marshal(tt)
	t.Log(string(s))

	s2, _ := protojson.MarshalOptions{
		EmitDefaultValues: true,
	}.Marshal(&tt)
	t.Log(string(s2))
}
