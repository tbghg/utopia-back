package implement

import (
	"testing"
	"utopia-back/database"
)

func TestFollowDal_GetFansList(t *testing.T) {
	err := database.TestInit()
	if err != nil {
		t.Errorf("TestInit() error = %v", err)
	}

	type args struct {
		userId uint
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				userId: 9,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FollowDal{}
			gotList, err := f.GetFansList(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFansList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(gotList) > 0 {
				t.Errorf("GetFansList() gotList = %v", gotList)
			}
		})
	}
}
