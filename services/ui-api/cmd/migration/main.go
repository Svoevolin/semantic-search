package migration

import (
	"fmt"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/config"
)

func main() {
	cfg, err := config.New[config.App]()
	if err != nil {
		panic(fmt.Sprintf("Error loading config: %v", err))
	}
	logger := logger.
}
func parseFlag() (string, error) {
	allowed := []string{"up", "down"}

}
