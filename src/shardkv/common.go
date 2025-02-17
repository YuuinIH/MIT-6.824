package shardkv

//
// Sharded key/value server.
// Lots of replica groups, each running Raft.
// Shardctrler decides which group serves each shard.
// Shardctrler may change shard assignment from time to time.
//
// You will have to modify these definitions.
//

const (
	OK                 = "OK"
	ErrNoKey           = "ErrNoKey"
	ErrWrongGroup      = "ErrWrongGroup"
	ErrWrongLeader     = "ErrWrongLeader"
	ErrShardNoValid    = "ErrShardNoValid"
	ErrConfigOutOfDate = "ErrConfigOutOfDate"
)

type Err string

// Put or Append
type PutAppendArgs struct {
	// You'll have to add definitions here.
	Key   string
	Value string
	Op    string // "Put" or "Append"
	// You'll have to add definitions here.
	// Field names must start with capital letters,
	// otherwise RPC will break.
	ClientID  int64
	Seq       int64
	ConfigNum int
}

type PutAppendReply struct {
	Err Err
}

type GetArgs struct {
	Key string
	// You'll have to add definitions here.
	ClientID  int64
	Seq       int64
	ConfigNum int
}

type GetReply struct {
	Err   Err
	Value string
}

type PushShardArgs struct {
	Shard     map[int]ShardComponent // shard id -> shard component
	ConfigNum int
}

type PushShardReply struct {
	Err       Err
	ConfigNum int
}
