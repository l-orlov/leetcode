package simplify_path

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				path: "/",
			},
			want: "/",
		},
		{
			name: "2",
			args: args{
				path: "//",
			},
			want: "/",
		},
		{
			name: "3",
			args: args{
				path: "///",
			},
			want: "/",
		},
		{
			name: "4",
			args: args{
				path: "/a/",
			},
			want: "/a",
		},
		{
			name: "5",
			args: args{
				path: "/a",
			},
			want: "/a",
		},
		{
			name: "6",
			args: args{
				path: "/.",
			},
			want: "/",
		},
		{
			name: "7",
			args: args{
				path: "/./",
			},
			want: "/",
		},
		{
			name: "8",
			args: args{
				path: "/..",
			},
			want: "/",
		},
		{
			name: "9",
			args: args{
				path: "/../",
			},
			want: "/",
		},
		{
			name: "10",
			args: args{
				path: "/a/..",
			},
			want: "/",
		},
		{
			name: "11",
			args: args{
				path: "/a/../",
			},
			want: "/",
		},
		{
			name: "12",
			args: args{
				path: "/home/",
			},
			want: "/home",
		},
		{
			name: "13",
			args: args{
				path: "/home//foo/",
			},
			want: "/home/foo",
		},
		{
			name: "14",
			args: args{
				path: "/home/user/Documents/../Pictures",
			},
			want: "/home/user/Pictures",
		},
		{
			name: "15",
			args: args{
				path: "/.../a/../b/c/../d/./",
			},
			want: "/.../b/d",
		},
		{
			name: "16",
			args: args{
				path: "/a/./b/../../c/",
			},
			want: "/c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPathV2(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
