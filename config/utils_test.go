package config

import (
	"os"
	"reflect"
	"testing"
)

var TestCases_boolFromEnv = []struct {
	in  string
	def bool
	out bool
}{
	{"true", true, true},
	{"false", true, false},
	{"0", true, false},
	{"", true, true},
	{"true", false, true},
	{"false", false, false},
	{"1", false, true},
	{"", false, false},
}

func Test_boolFromEnv(t *testing.T) {
	for _, tc := range TestCases_boolFromEnv {
		t.Run("", func(t *testing.T) {
			os.Setenv("foo", tc.in)
			defer os.Unsetenv("foo")
			res := boolFromEnv("foo", tc.def)
			if res != tc.out {
				t.Errorf("got %v, want %v", res, tc.out)
			}
		})
	}
}

var TestCases_listFromEnv = []struct {
	in  string
	sep string
	out []string
}{
	{"foo", ":", []string{"foo"}},
	{"foo:bar", ":", []string{"foo", "bar"}},
	{"foo:bar:baz", ":", []string{"foo", "bar", "baz"}},
	{"foo:bar baz", " ", []string{"foo:bar", "baz"}},
	{"", " ", []string{}},
}

func Test_listFromEnv(t *testing.T) {
	for _, tc := range TestCases_listFromEnv {
		t.Run("", func(t *testing.T) {
			os.Setenv("foo", tc.in)
			defer os.Unsetenv("foo")
			res := listFromEnv("foo", tc.sep)
			if !(len(tc.out) == 0 && len(res) == 0) && !reflect.DeepEqual(tc.out, res) {
				t.Errorf("got %v, want %v", res, tc.out)
			}
		})
	}
}

var TestCases_listFromEnvs = []struct {
	pref string
	vars map[string]string
	out  map[string]string
}{
	{
		"foo",
		map[string]string{
		},
		map[string]string{
		},
	},
	{
		"foo",
		map[string]string{
			"foofoo": "bar",
			"barbar": "baz",
		},
		map[string]string{
			"foo": "bar",
		},
	},
	{
		"bar",
		map[string]string{
			"bar1": "",
			"bar2": "foo",
		},
		map[string]string{
			"1": "",
			"2": "foo",
		},
	},
}

func Test_listFromEnvs(t *testing.T) {
	for _, tc := range TestCases_listFromEnvs {
		t.Run("", func(t *testing.T) {
			for key, val := range tc.vars {
				os.Setenv(key, val)
				defer os.Unsetenv(key)
			}

			res := listFromEnvs(tc.pref)
			if !(len(tc.out) == 0 && len(res) == 0) && !reflect.DeepEqual(tc.out, res) {
				t.Errorf("got %v, want %v", res, tc.out)
			}
		})
	}
}
