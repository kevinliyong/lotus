// 创建一个选择器，使除密封的第一个流程外，其他流程都使用第一个流程用的worker
// 但是怎么保证第一个worker依然可用，如果不可用怎么办？
package sectorstorage

import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
)

type mySelector struct {
	sid2worker map[abi.SectorID]Worker
}

func myNewSelector() *mySelector {
	return &mySelector{
		sid2worker: make(map[abi.SectorID]Worker, 32), //一个扇区32G好像，初始化就32个扇区吧，其实在密封任务结束后是可以删除的
	}
}

func (s *mySelector) Ok(ctx context.Context, sector abi.SectorID, task sealtasks.TaskType, spt abi.RegisteredSealProof, whnd *workerHandle) (bool, error) {
	if s.sid2worker[sector] == whnd.w {
		return false, nil
	}

	return true, nil
}

func (s *mySelector) Cmp(ctx context.Context, sector abi.SectorID, task sealtasks.TaskType, a, b *workerHandle) (bool, error) {
	// 实际上，只有一个worker，不用比较
	return true, nil
}

var _ WorkerSelector = &mySelector{}
