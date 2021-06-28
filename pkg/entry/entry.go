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
