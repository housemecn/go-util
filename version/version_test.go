package version

import (
	"fmt"
	"runtime"
	"testing"
)

func TestStringify(t *testing.T) {
	version := "1.0.0"
	buildTime := "2020/05/27"

	want := fmt.Sprintf(`Version:      %s
Go version:   %s
Built:        %s
OS/Arch:      %s/%s`,
		version,
		runtime.Version(),
		buildTime,
		runtime.GOOS,
		runtime.GOARCH,
	)

	type args struct {
		version   string
		buildTime string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				version:   version,
				buildTime: buildTime,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Stringify(tt.args.version, tt.args.buildTime); got != tt.want {
				t.Errorf("Stringify() = %v, want %v", got, tt.want)
			}
		})
	}
}
