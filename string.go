package util

func P2Str(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func Str2P(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func StrArr2PArr(sa []string) []*string {
	if sa == nil {
		return nil
	}
	result := make([]*string, len(sa))
	for index, s := range sa {
		result[index] = Str2P(s)
	}
	return result
}
