package conf

import (
	"fmt"
	queue "service-queue"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

type customBridge struct {
	connDial *redis.Pool
}

func (c *customBridge) Push(key interface{}) bool {
	con := c.connDial.Get()
	redis.Strings(con.Do("RPUSH", "list", key))
	defer con.Close()
	return true
}

func (c *customBridge) Pop() interface{} {
	con := c.connDial.Get()
	var count, _ = con.Do("LLEN", "list")
	var index, err = redis.String(con.Do("LINDEX", "list", "0"))
	if err != nil {
		fmt.Println("Data is empty")
	}
	fmt.Println(count)
	if count.(int64) <= 0 {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Data is empty")
			}
		}()
	}
	if _, err2 := con.Do("LPOP", "list"); err2 != nil {
		panic(err2)
	}

	defer con.Close()
	return index
}

func (c *customBridge) Keys() []interface{} {
	con := c.connDial.Get()
	var result, err = redis.Strings(con.Do("LRANGE", "list", "0", "-1"))
	if err != nil {
		panic(err)
	}
	if len(result) <= 0 {
		fmt.Println("Data is empty")
	}
	resultParse := make([]interface{}, 0)
	for _, each := range result {
		resultParse = append(resultParse, each)
	}
	defer con.Close()
	return resultParse

}

func (c *customBridge) Len() int {
	con := c.connDial.Get()
	var count, _ = redis.Int(con.Do("LLEN", "list"))

	if count <= 0 {
		fmt.Println("Data is empty")
	}
	defer con.Close()
	return count
}

func (c *customBridge) Contains(key interface{}) bool {
	con := c.connDial.Get()
	var exist bool
	var result, err = redis.Strings(con.Do("LRANGE", "list", "0", "-1"))
	if err != nil {
		panic(err)
	}
	if len(result) == 1 {
		exist = true
	} else {
		for _, g := range result[:len(result)-1] {
			gConvert, _ := strconv.Atoi(g)
			if key == gConvert {
				if _, errD := con.Do("RPOP", "list"); errD != nil {
					panic(err)
				}

				exist = false
			} else {
				exist = true
			}
		}
	}
	defer con.Close()
	return exist
}

// New is take main func
func New(conn *redis.Pool) queue.Service {
	return &customBridge{conn}
}
