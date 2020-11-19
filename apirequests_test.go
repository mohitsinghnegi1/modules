package modules

import (
	"reflect"
	"testing"
)

func TestSendServerHeartBeat(t *testing.T) {
	type args struct {
		serverAddress     string
		activeConnections int
		gameMode          string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Test for SendServerHeartBeat Function",
			args: args{
				serverAddress:     "https://localhost:4500",
				activeConnections: 23,
				gameMode:          "cardParty",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasSendedHeartBeat := SendServerHeartBeat(tt.args.serverAddress, tt.args.activeConnections, tt.args.gameMode)
			if got := hasSendedHeartBeat; got != tt.want {
				t.Errorf("hasSendedHeartBeat = %v, want %v", got, tt.want)
			}
			if got := hasSendedHeartBeat; reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("hasSendedHeartBeat return type = %v, expected  return type %v", reflect.TypeOf(got), reflect.TypeOf(tt.want))
			}
		})
	}
}
