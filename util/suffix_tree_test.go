package util_test

import (
	"testing"

	"github.com/wweir/sower/util"
)

func TestNode_Match(t *testing.T) {
	type test struct {
		arg  string
		want bool
	}
	tests := []struct {
		name  string
		node  *util.Node
		tests []test
	}{{
		"simple",
		util.NewNodeFromRules("a.wweir.cc", "b.wweir.cc"),
		[]test{
			{"a.wweir.cc", true},
			{"b.wweir.cc", true},
		},
	}, {
		"parent",
		util.NewNodeFromRules("wweir.cc", "a.wweir.cc"),
		[]test{
			{"wweir.cc", true},
			{"a.wweir.cc", true},
			{"b.wweir.cc", false},
		},
	}, {
		"fuzz1",
		util.NewNodeFromRules("wweir.cc", "a.wweir.cc", "*.wweir.cc"),
		[]test{
			{"wweir.cc", true},
			{"a.wweir.cc", true},
			{"b.wweir.cc", true},
			{"a.b.wweir.cc", false},
		},
	}, {
		"fuzz2",
		util.NewNodeFromRules("a.*.cc", "c.wweir.*"),
		[]test{
			{"wweir.cc", false},
			{"a.wweir.cc", true},
			{"b.wweir.cc", false},
			{"c.wweir.cc", true},
		},
	}, {
		"fuzz3",
		util.NewNodeFromRules("*.*.cc", "iamp.*.*"),
		[]test{
			{"wweir.cc", false},
			{"a.wweir.cc", true},
			{"b.wweir.cc", true},
			{"iamp.wweir.cc", true},
		},
	}, {
		"fuzz4",
		util.NewNodeFromRules("**.cc", "a.**.com", "**.wweir.*"),
		[]test{
			{"wweir.cc", true},
			{"a.wweir.cc", true},
			{"a.b.wweir.cc", true},
			{"a.fuzz.com", true},
			{"b.fuzz.com", false},
			{"www.wweir.com", true},
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, test := range tt.tests {
				if got := tt.node.Match(test.arg); got != test.want {
					t.Errorf("Node.Match(%s) = %v, want %v", test.arg, got, test.want)
				}
			}
		})
	}
}
