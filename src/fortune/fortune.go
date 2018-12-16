package fortune

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"

	"github.com/ukko/cowsay/src/redis"
)

// New return new fortune message (or from cache)
func New() (string, error) {
	k := fmt.Sprintf("f:%d", time.Now().Second())
	f, err := redis.Get(k)
	if err != nil {
		return "", err
	}

	if f != nil {
		return string(f), nil
	}

	cmd := exec.Command("/usr/games/fortune")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}

	if err := redis.Set(k, []byte(out.String())); err != nil {
		return "", err
	}

	return out.String(), nil
}
