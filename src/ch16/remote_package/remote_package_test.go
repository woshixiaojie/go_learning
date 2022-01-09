/* 使用go-get，或者go-mod，来下载包 */
package remote_package

import (
	cm "github.com/easierway/concurrent_map"
	"testing"
)

func TestRemotePackage(t *testing.T) {
/*	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("Key"), 10)
	t.Log(m.Get(cm.StrKey("key")))*/
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}