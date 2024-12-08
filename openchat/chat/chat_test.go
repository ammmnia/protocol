package chat_test

import (
	"encoding/json"
	. "github.com/ammmnia/protocol/openchat/chat"
	common "github.com/ammmnia/protocol/openchat/common"
	"google.golang.org/protobuf/encoding/protojson"
	"testing"
)

func TestFindUserFullInfoReq(t *testing.T) {
	tt := FindUserFullInfoResp{
		Users: []*common.UserFullInfo{
			{
				Birth: 134141,
			},
		},
	}
	s, _ := json.Marshal(tt)
	t.Log(string(s))

	s2, _ := protojson.MarshalOptions{
		AllowPartial:      true,
		UseEnumNumbers:    true,
		EmitDefaultValues: true,
	}.Marshal(&tt)
	t.Log(string(s2))

	//s3, _ := sonic.Config{}.Froze().Marshal(&tt)
	//t.Log(string(s3))
}
