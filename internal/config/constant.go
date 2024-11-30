package config

const (
	snowflakeNodeIDKey = "/counters/snowflake_node_id"
)

func GetSnowflakeNodeIDKey() string {
	return snowflakeNodeIDKey
}
