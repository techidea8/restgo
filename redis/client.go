package redis

import "github.com/techidea8/restgo"

func InitRedis(conf Conf) {
	redisPool := NewRedisPool(conf)
	RedisEngin = NewRedisClient(redisPool)
}

type RedisClient struct {
	redisPool *redis.Pool
}

func (this *RedisClient) Exec(cmd string, key interface{}, args ...interface{}) (interface{}, error) {
	con := this.redisPool.Get()
	if err := con.Err(); err != nil {
		return nil, err
	}
	defer con.Close()
	parmas := make([]interface{}, 0)
	parmas = append(parmas, key)

	if len(args) > 0 {
		for _, v := range args {
			parmas = append(parmas, v)
		}
	}
	return con.Do(cmd, parmas...)
}

func NewRedisClient(p *redis.Pool) *RedisClient {
	return &RedisClient{
		redisPool: p,
	}
}
func (this *RedisClient) Set(k, v string) {
	this.Exec("set", k, v)
}

func (this *RedisClient) HSet(k, f,v string) {
	//HSET KEY_NAME FIELD VALUE 
	this.Exec("hset", k,f, v)
}
func (this *RedisClient) HGet(k, f string)(r string, err error) {
	//HSET KEY_NAME FIELD VALUE 
	result, e :=this.Exec("hget", k,f)
	if e != nil {
		return "", e
	}else{
		return fmt.Sprintf("%s", result), e
	}
}


func (this *RedisClient) Get(k string) (r string, err error) {
	result, e := this.Exec("get", k)
	if e != nil {
		return "", e
	}
	return fmt.Sprintf("%s", result), nil
}

func (this *RedisClient) SetKeyExpire(k string, ex int) {

	this.Exec("EXPIRE", k, ex)

}

//
func (this *RedisClient) SetKeyExpireHourLater(k string, ex int) {
	this.SetKeyExpire(k, ex*3600)
}

//
func (this *RedisClient) SetKeyExpireMinitusLater(k string, ex int) {
	this.SetKeyExpire(k, ex*60)
}
func (this *RedisClient) SetKeyExpireSecondLater(k string, ex int) {
	this.SetKeyExpire(k, ex)
}
func (this *RedisClient) Exists(k string) bool {
	c := this.redisPool.Get()
	defer c.Close()
	exist, err := redis.Bool(c.Do("EXISTS", k))

	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return exist
	}
}

//获得键值时间
func (this *RedisClient) Ttl(k string) int64 {
	r, err := this.Exec("TTL", k)
	if err != nil {
		return -1
	} else {
		return r.(int64)
	}
}

func (this *RedisClient) DelKey(k string) error {
	_, err := this.Exec("DEL", k)
	return err
}
