package gen

import (
	"testing"
	"github.com/gogf/gf/test/gtest"
	"github.com/gogf/gf/os/glog"
)

func TestSnakeCase(t *testing.T) {
	ret := SnakeCase("testMd5")
	glog.Debug(ret)
	gtest.Assert(ret, "test_md5")
}
