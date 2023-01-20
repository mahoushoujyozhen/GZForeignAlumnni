package util

import (
	"strconv"
	"sync"
	"time"
)

// Snowflake 雪花算法
type Snowflake struct {
	sync.Mutex       // 同步锁
	timestamp  int64 // 时间戳(ms)
	machineId  int64 // 区域节点Id
	dataId     int64 // 数据中心Id
	serialId   int64 // 序列号Id
}

const (
	startTimestamp = int64(1643644800000) // 起始时间戳(2022-2-1 00:00:00)

	timestampBits = uint(41) // 41位时间戳
	machineIdBits = uint(5)  // 5位区域节点
	dataIdBits    = uint(5)  // 5位数据中心
	serialIdBits  = uint(12) // 12位序列号

	// "timestamp"左移的位数为22位
	// "machineId"左移的位数为17位
	// "dataId"左移的位数位12位

	maxMachineId = int64(1 << machineIdBits) // 最大的区域节点Id
	maxDataId    = int64(1 << dataIdBits)    // 最大的数据中心Id
	maxSerialId  = int64(1 << serialIdBits)  // 最大的序列号
)

// GetSnowflake 生成工厂
func GetSnowflake(timestamp int64, machineId int64, dataId int64, serialId int64) (*Snowflake, error) {
	// 判断参数合法性
	if machineId > maxMachineId || dataId > maxDataId || serialId > maxSerialId {
		return nil, ERR_SNOWFLAKE_PARAM
	} else {
		return &Snowflake{sync.Mutex{}, timestamp, machineId, dataId, serialId}, nil
	}
}

// GenerateUniqueId 生成全局唯一Id
func GenerateUniqueId(Snowflake *Snowflake) (string, error) {
	// 获取锁
	Snowflake.Mutex.Lock()
	// 获取系统当前时间戳(ms)
	var currentTimestamp int64 = time.Now().UnixMilli()
	// 根据不同的时间进行对应的处理
	if currentTimestamp < Snowflake.timestamp {
		// 释放锁
		Snowflake.Mutex.Unlock()
		// 遇到时钟回拨错误,抛出异常
		return "", ERR_CLOCK_BACKWARDS
	}
	if currentTimestamp == Snowflake.timestamp {
		Snowflake.serialId = (Snowflake.serialId + 1) & maxSerialId
		// 判断当前毫秒是否剩余序列号
		if Snowflake.serialId == 0 {
			// 如果序列号使用完毕,等待下一个毫秒
			for currentTimestamp <= Snowflake.timestamp {
				currentTimestamp = time.Now().UnixMilli()
			}
		}
	} else {
		// 如果"时间超前",那么属于新的时间,序列号置为0
		Snowflake.serialId = 0
	}
	// 更新生成Id的时间戳
	Snowflake.timestamp = currentTimestamp
	// 释放锁
	Snowflake.Mutex.Unlock()
	// 转化为"string"类型后返回
	return strconv.FormatInt(Snowflake.timestamp<<22|Snowflake.machineId<<17|Snowflake.dataId<<12|Snowflake.serialId, 10), nil
}
