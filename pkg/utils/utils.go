package utils

import (
	"fmt"
	"github.com/mchirico/go.etcd/pkg/etcdutils"
	"sync"
	"time"
)

type UT struct {
	count int64
	sync.Mutex
}

func NewUT() *UT {
	u := &UT{}
	return u
}

func (u *UT)GetCount() int64 {
	u.Lock()
	defer u.Unlock()

	return u.count
}

func (u *UT)Email(email, value string,ttl int64) string {
	u.Lock()
	defer u.Unlock()
	u.count += 1

	e,cancel := etcdutils.NewETC("server")
	defer cancel()

	timeStamp := time.Now().String()
	key := fmt.Sprintf("email/%s/%s",email,timeStamp)
	e.PutWithLease(key,value, ttl)
	return key

}

func (u *UT)Status() string {
	u.Lock()
	defer u.Unlock()
	u.count += 1

	e,cancel := etcdutils.NewETC("server")
	defer cancel()

	timeStamp := time.Now().String()
	key := fmt.Sprintf("tasks/status: %v\n",timeStamp)
	e.PutWithLease(key,fmt.Sprint(u.count), 120)
	return key

}


func (u *UT) GetListing(key string) string {
	u.Lock()
	defer u.Unlock()
	s := ""

	e,cancel := etcdutils.NewETC("server")
	defer cancel()

	result, _ := e.GetWithPrefix(key)


	for i, v := range result.Kvs {
		s+=fmt.Sprintf("result.Kvs[%d]: %s %s, ver: %d,  lease: %d\n", i,result.Kvs[i], v.Value, v.Version, v.Lease)
	}
	return s
}