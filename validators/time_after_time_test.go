package validators

import (
	"testing"
	"time"

	"github.com/drlogout/validate/v3"
	"github.com/stretchr/testify/require"
)

func Test_TimeAfterTime(t *testing.T) {
	r := require.New(t)
	now := time.Now()
	v := TimeAfterTime{
		FirstName: "Opens At", FirstTime: now.Add(100000),
		SecondName: "Now", SecondTime: now,
	}

	errors := validate.NewErrors()
	v.IsValid(errors)
	r.Equal(0, errors.Count())

	v.SecondTime = now.Add(200000)
	v.IsValid(errors)

	r.Equal(1, errors.Count())
	r.Equal(errors.Get("opens_at"), []string{"Opens At must be after Now."})

	errors = validate.NewErrors()
	v.Message = "OpensAt must be later than Now."

	v.IsValid(errors)

	r.Equal(1, errors.Count())
	r.Equal(errors.Get("opens_at"), []string{"OpensAt must be later than Now."})

	firstTime := now.AddDate(260, 0, 0)
	v = TimeAfterTime{
		FirstName: "Opens At", FirstTime: firstTime,
		SecondName: "Now", SecondTime: now,
	}

	errors = validate.NewErrors()
	v.IsValid(errors)
	r.Equal(0, errors.Count())
}
