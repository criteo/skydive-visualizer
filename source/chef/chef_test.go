package chef

import (
	"testing"

	"github.com/go-chef/chef"
	"github.com/stretchr/testify/require"
)

func TestGetChefAttr(t *testing.T) {
	node := chef.Node{
		AutomaticAttributes: map[string]interface{}{
			"a": map[string]interface{}{
				"b": map[string]interface{}{
					"c": "ok1",
				},
			},
		},
		OverrideAttributes: map[string]interface{}{
			"a": map[string]interface{}{
				"b": map[string]interface{}{
					"c": "not_ok",
					"d": "ok2",
				},
			},
		},
		DefaultAttributes: map[string]interface{}{
			"e": "ok3",
		},
	}

	v, ok := getChefAttr(node, "a.b.c")
	require.True(t, ok)
	require.Equal(t, "ok1", v)

	v, ok = getChefAttr(node, "a.b.d")
	require.True(t, ok)
	require.Equal(t, "ok2", v)

	v, ok = getChefAttr(node, "e")
	require.True(t, ok)
	require.Equal(t, "ok3", v)

	v, ok = getChefAttr(node, "not_exists")
	require.False(t, ok)
	require.Equal(t, "", v)

	v, ok = getChefAttr(node, "not_exists.other")
	require.False(t, ok)
	require.Equal(t, "", v)
}
