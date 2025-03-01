package configs

type envStore map[string]string

func (s envStore) getValue(envName, defaultValue string) string {
	v, ok := s[envName]
	if ok {
		return v
	}
	return defaultValue
}
