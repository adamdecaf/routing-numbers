package routing

import (
	"fmt"
	"errors"
	"strconv"
	"strings"
)

const (
	RTNLength = 9
)

func Check(rtn string) error {
	rtn = strings.TrimSpace(rtn)

	// Pad with leading zero(s)
	if (len(rtn) - RTNLength) <= 2 {
		rtn = strings.Repeat("0", RTNLength - len(rtn)) + rtn
	}
	if len(rtn) != RTNLength {
		return fmt.Errorf("routing number %s is not %d digits long", rtn, RTNLength)
	}

	split := strings.Split(rtn, "")

	r := make([]int, RTNLength)
	for i := range r {
		v, err := strconv.Atoi(split[i])
		if err != nil {
			return err
		}
		r[i] = v
	}

	sum := 3*(r[0] + r[3] + r[6]) + 7*(r[1] + r[4] + r[7]) + (r[2] + r[5] + r[8])
	if sum % 10 == 0 {
		return nil
	}
	return errors.New("routing number does not pass check digit")
}

// Valid is a handy method for running Check but also performing the
// comparison that it returns no error.
func Valid(rtn string) bool {
	return Check(rtn) == nil
}
