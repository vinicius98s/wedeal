package loader

import (
	"strconv"

	"github.com/graph-gophers/dataloader"
)

func loadBatchError(err error, n int) []*dataloader.Result {
	r := &dataloader.Result{Error: err}
	res := make([]*dataloader.Result, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, r)
	}
	return res
}

func loadKeys(keys dataloader.Keys) ([]int32, error) {
	res := make([]int32, 0, len(keys))
	for _, s := range keys.Keys() {
		v, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return nil, err
		}
		res = append(res, int32(v))
	}
	return res, nil
}
