package util

func Ptr2Str(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func Str2Ptr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func Strs2Ptrs(strs []string) []*string {
	if strs == nil {
		return nil
	}
	result := make([]*string, len(strs))
	for index, str := range strs {
		result[index] = Str2Ptr(str)
	}
	return result
}

func Ptrs2Strs(ptrs []*string) []string {
	if ptrs == nil {
		return nil
	}
	result := make([]string, len(ptrs))
	for index, ptr := range ptrs {
		result[index] = Ptr2Str(ptr)
	}
	return result
}

func StringFallback(str, fallback string) string {
	if str == "" {
		return fallback
	}
	return str
}
