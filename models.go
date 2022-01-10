package raptorproto

// Config 用户提交的任务配置
type Config struct {
	ID               string        // 唯一标识，后台自动生成
	Name             string        // 配置名
	TargetService    string        // 被调服务
	Task                           // 任务
	Trigger                        // 触发条件
	Dependencies     []*Dependency // 任务依赖
	RetryStrategy                  // 失败处理策略
	ShardingStrategy               // 分片策略
}

type Task struct {
	Type    string // 任务类型
	Content string // 任务内容
	Args    string // 任务参数
}

type Trigger struct {
}

type Dependency struct {
}

type RetryStrategy struct {
}

type ShardingStrategy struct {
}
