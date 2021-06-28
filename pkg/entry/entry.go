// Copyright (c) 2021, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package entry

import (
	pb "go.etcd.io/etcd/api/v3/etcdserverpb"
	"go.etcd.io/etcd/pkg/v3/pbutil"
)

type CustomInternalRaftRequest struct {
	*pb.InternalRaftRequest
}

func Unmarshal(data []byte) interface{} {
	var raftReq pb.InternalRaftRequest
	if !pbutil.MaybeUnmarshal(&raftReq, data) { // backward compatible
		var r pb.Request
		pbutil.MustUnmarshal(&r, data)
		return &r
	}

	return &raftReq
}
