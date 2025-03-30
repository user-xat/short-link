package configs

import "os"

type envStore struct {
	store map[string]string
}

func newEnvStore(store map[string]string) *envStore {
	return &envStore{
		store: store,
	}
}

func (s envStore) getValue(envName, defaultValue string) string {
	var v string
	var ok bool
	if s.store != nil {
		v, ok = s.store[envName]
	} else {
		v, ok = os.LookupEnv(envName)
	}
	if ok {
		return v
	}
	return defaultValue
}
