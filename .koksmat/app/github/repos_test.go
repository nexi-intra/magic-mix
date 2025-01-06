package github

import (
	"testing"

	"github.com/google/go-github/v68/github"
)

func TestGetRepos(t *testing.T) {
	type args struct {
		orgName string
	}
	tests := []struct {
		name    string
		args    args
		want    []*github.Repository
		wantErr bool
	}{
		{
			name: "Error ...",
			args: args{
				orgName: "nexi-intra",
			},
			wantErr: false,
		},
	}
	client, err := GetClient()
	if err != nil {
		t.Errorf("Error authenticating: %s", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRepos(client, tt.args.orgName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRepos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == 0 {
				t.Errorf("GetRepos() = %v, want %v", got, tt.want)
			}
		})
	}
}
