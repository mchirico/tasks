package utils

import (
	"fmt"
	"testing"
)

func TestNewUT(t *testing.T) {
	u := NewUT()
	r := u.Email("test@task","test",25)
	t.Logf("%v\n",r)

    r2 := u.GetListing("email/test@task")
    t.Logf("end: %s\n",r2)

}

func TestUT_Status(t *testing.T) {
	u := NewUT()
	fmt.Println(u.Status())
}