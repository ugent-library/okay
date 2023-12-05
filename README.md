[![Go Reference](https://pkg.go.dev/badge/github.com/ugent-library/okay.svg)](https://pkg.go.dev/github.com/ugent-library/okay)

# ugent-library/okay

Simple validation for Go.

## Install

```sh
go get -u github.com/ugent-library/okay
```

## Examples

Sugared validation:

```go
func Validate(u *User) error {
	return okay.Validate(
		okay.NotEmpty("username", f.Username),
		okay.LengthBetween("username", f.Username, 1, 100),
        // with custom message
		okay.LengthBetween("username", f.Username, 1, 100).WithMessage("Username is too short or too long"),
	)
}
```

Building validation errors manually:

```go
func Validate(r *Rec) error {
    errs := okay.NewErrors()
    for i, link := range r.Links
        if isExpired(link) {
            errs.Add(okay.NewError(fmt.Sprintf("/rec/links/%d", i), "link.expired"))
        }
    }
    return errs.ErrorOrNil()
}
```

Rules and parameters:

```go
err := okay.LengthBetween("username", f.Username, 1, 100)
msg := fmt.Sprintf("%s should have %s %d and %d", err.Key, err.Rule, err.Params...)
// msg is "username should have length_between 1 and 100"
```

I18n:

```go
import "github.com/leonelquinteros/gotext"

func ErrorMessages(loc *gotext.Locale, errs *okay.Errors) msgs []string {
    for _, err := range errs.Errors {
        msgs = append(msgs, loc.Get(err.Rule, err.Params...))
    }
    return
}
```