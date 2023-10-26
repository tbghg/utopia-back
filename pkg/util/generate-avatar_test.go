package utils

import "testing"

func TestGenerateAvatar(t *testing.T) {
	type args struct {
		width       int
		height      int
		blockWidth  int
		blockHeight int
		outputPath  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "generate avatar", args: args{
				width:       120,
				height:      120,
				blockWidth:  30,
				blockHeight: 30,
				outputPath:  "../../output/avatar.png",
			},
			wantErr: false,
		},
		{
			name: "generate avatar", args: args{
				width:       100,
				height:      100,
				blockWidth:  20,
				blockHeight: 20,
				outputPath:  "../../output/avatar2.png",
			},
			wantErr: false,
		},
		{
			name: "generate avatar", args: args{
				width:       420,
				height:      420,
				blockWidth:  140,
				blockHeight: 140,
				outputPath:  "../../output/avatar3.png",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenerateAvatar(tt.args.width, tt.args.height, tt.args.blockWidth, tt.args.blockHeight, tt.args.outputPath); (err != nil) != tt.wantErr {
				t.Errorf("GenerateAvatar() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
