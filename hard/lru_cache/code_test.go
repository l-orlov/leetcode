package lru_cache

import (
	"strconv"
	"strings"
	"testing"
)

func TestLRUCache_Step(t *testing.T) {
	tests := []struct {
		name         string
		inputActions string
		inputArgs    string
		wantOutput   string
	}{
		{
			name:         "1",
			inputActions: `["LRUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"]`,
			inputArgs:    `[[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]`,
			wantOutput:   `[null,null,null,null,null,null,-1,null,19,17,null,-1,null,null,null,-1,null,-1,5,-1,12,null,null,3,5,5,null,null,1,null,-1,null,30,5,30,null,null,null,-1,null,-1,24,null,null,18,null,null,null,null,-1,null,null,18,null,null,-1,null,null,null,null,null,18,null,null,-1,null,4,29,30,null,12,-1,null,null,null,null,29,null,null,null,null,17,22,18,null,null,null,-1,null,null,null,20,null,null,null,-1,18,18,null,null,null,null,20,null,null,null,null,null,null,null]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actions := strings.Split(tt.inputActions[1:len(tt.inputActions)-1], ",")[1:]
			args := strings.Split(tt.inputArgs[2:len(tt.inputArgs)-2], "],[")
			capacity, err := strconv.Atoi(args[0])
			if err != nil {
				t.Fatalf("strconv.Atoi: %s", err)
			}
			args = args[1:]

			wantOutput := strings.Split(tt.wantOutput[1:len(tt.wantOutput)-1], ",")[1:]

			if len(actions) != len(args) {
				t.Fatal("not valid args len")
			}

			cache := Constructor(capacity)
			for i, action := range actions {
				if i == 48 || i == 25 {
					a := true
					_ = a
				}
				if action == `"put"` {
					tmpArgs := strings.Split(args[i], ",")
					key, ierr := strconv.Atoi(tmpArgs[0])
					if ierr != nil {
						t.Fatalf("strconv.Atoi: %s", ierr)
					}

					val, ierr := strconv.Atoi(tmpArgs[1])
					if ierr != nil {
						t.Fatalf("strconv.Atoi: %s", ierr)
					}

					cache.Put(key, val)
				} else {
					key, ierr := strconv.Atoi(args[i])
					if ierr != nil {
						t.Fatalf("strconv.Atoi: %s", ierr)
					}

					res := cache.Get(key)
					want, ierr := strconv.Atoi(wantOutput[i])
					if ierr != nil {
						t.Fatalf("strconv.Atoi: %s", ierr)
					}

					if want != res {
						t.Fatalf("for action %d %s(%s) got %d, want %d", i, action, args[i], res, want)
					}
				}
			}
		})
	}
}
