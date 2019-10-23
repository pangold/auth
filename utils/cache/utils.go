package cache

import (
	"fmt"
)

func generateCacheKey(service, key string) string {
	return fmt.Sprintf("%s.%s", service, key)
}
