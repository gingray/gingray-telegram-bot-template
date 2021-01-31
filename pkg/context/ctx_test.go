package context

import (
	"reflect"
	_ "devJoyTelegramBot/testing"
	"testing"
)

func TestGetCtx(t *testing.T) {
	tests := []struct {
		name string
		want *BotCtx
	}{
		{
			name: "Init case",
			want: GetCtx(),
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCtx(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
