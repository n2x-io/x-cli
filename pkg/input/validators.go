package input

import (
	"errors"
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/ssh"
)

func ValidUUID(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("missing UUID")
	}

	if _, err := uuid.Parse(id); err != nil {
		return fmt.Errorf("invalid UUID: %v", err)
	}

	return nil
}

func ValidName(val interface{}) error {
	name := val.(string)

	if len(name) == 0 {
		return fmt.Errorf("invalid name: missing identifier")
	}

	if errMsgs := isDNS1035Label(name); len(errMsgs) > 0 {
		err := fmt.Errorf("invalid name")

		for _, errMsg := range errMsgs {
			err = fmt.Errorf("%s | %s", err, errMsg)
		}

		return err
	}

	return nil
}

func ValidTags(val interface{}) error {
	str := val.(string)

	if len(str) == 0 {
		return nil
	}

	tags := strings.Split(strings.ToLower(strings.TrimSpace(str)), ",")

	for _, tag := range tags {
		r, err := regexp.MatchString(`^[a-z]{1}[a-z0-9]+$`, tag)
		if err != nil {
			return err
		}
		if !r {
			return errors.New("invalid tag")
		}
	}

	return nil
}

func ValidFQDN(val interface{}) error {
	fqdn := val.(string)

	if len(fqdn) == 0 {
		return nil
		// return fmt.Errorf("missing fqdn")
	}

	if _, err := net.LookupHost(fqdn); err != nil {
		return fmt.Errorf("invalid fqdn: %v", err)
	}

	return nil
}

func ValidID(val interface{}) error {
	id := val.(string)

	if strings.ToLower(id) == "master" {
		return errors.New("reserved identifier")
	}

	r, err := regexp.MatchString(`^[a-z]{1}[a-z0-9-]+$`, id)
	if err != nil {
		return err
	}
	if !r {
		return errors.New("invalid identifier")
	}

	return nil
}

func ValidEmail(val interface{}) error {
	email := val.(string)

	r, err := regexp.MatchString(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`, email)
	if err != nil {
		return err
	}
	if !r {
		return errors.New("invalid email address")
	}

	return nil
}

func ValidNetwork(val interface{}) error {
	ipv4Addr, ipv4Net, err := net.ParseCIDR(val.(string))
	if err != nil {
		return err
	}

	if !ipv4Addr.Equal(ipv4Net.IP) {
		return fmt.Errorf("invalid network address %v", ipv4Addr)
	}

	cidrMask, _ := ipv4Net.Mask.Size()

	if cidrMask != 16 {
		return errors.New("only /16 networks are supported at the moment")
	}

	return nil
}

func ValidPassword(val interface{}) error {
	// var r bool
	// var err error

	pw := val.(string)

	if len(pw) < 10 {
		return errors.New("invalid password: length must be at least 10 characters")
	}

	/*
		r, err = regexp.MatchString(`[A-Z]`, pw)
		if err != nil {
			return err
		}
		if !r {
			return errors.New("invalid password")
		}

		r, err = regexp.MatchString(`[0-9]`, pw)
		if err != nil {
			return err
		}
		if !r {
			return errors.New("invalid password")
		}

		r, err = regexp.MatchString(`[a-z]`, pw)
		if err != nil {
			return err
		}
		if !r {
			return errors.New("invalid password")
		}
	*/

	return nil
}

func ValidIPNetCIDR(val interface{}) error {
	_, _, err := net.ParseCIDR(val.(string))
	if err != nil {
		return err
	}

	return nil
}

func ValidPort(val interface{}) error {
	port := val.(string)

	p, err := strconv.Atoi(port)
	if err != nil {
		return err
	}

	if p > 0 && p < 65535 {
		return nil
	}

	return errors.New("invalid port")
}

func ValidUint(val interface{}) error {
	n := val.(string)

	i, err := strconv.Atoi(n)
	if err != nil {
		return errors.New("invalid unsigned integer")
	}

	if i > 0 {
		return nil
	}

	return errors.New("invalid unsigned integer")
}

func ValidPrice(val interface{}) error {
	n := val.(string)

	p, err := strconv.Atoi(n)
	if err != nil {
		return errors.New("invalid price")
	}

	if p >= 0 {
		return nil
	}

	return errors.New("invalid price")
}

func ValidSSHPublicKey(val interface{}) error {
	pubKey := val.(string)

	if len(pubKey) == 0 {
		return errors.New("missing SSH public key")
	}

	if _, _, _, _, err := ssh.ParseAuthorizedKey([]byte(pubKey)); err != nil {
		return err
	}

	return nil
}

// regexps

// Ref: https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/util/validation/validation.go

// maxLenError returns a string explanation of a "string too long" validation
// failure.
func maxLenError(length int) string {
	return fmt.Sprintf("must be no more than %d characters", length)
}

// regexError returns a string explanation of a regex validation failure.
func regexError(msg string, fmt string, examples ...string) string {
	if len(examples) == 0 {
		return msg + " (regex used for validation is '" + fmt + "')"
	}

	msg += " (e.g. "
	for i := range examples {
		if i > 0 {
			msg += " or "
		}
		msg += "'" + examples[i] + "', "
	}
	msg += "regex used for validation is '" + fmt + "')"

	return msg
}

// emptyError returns a string explanation of a "must not be empty" validation
// failure.
func emptyError() string {
	return "must be non-empty"
}

const dns1035LabelFmt string = "[a-z]([-a-z0-9]*[a-z0-9])?"

const dns1035LabelErrMsg string = "a DNS-1035 label must consist of lower case alphanumeric characters or '-', start with an alphabetic character, and end with an alphanumeric character"

// DNS1035LabelMaxLength is a label's max length in DNS (RFC 1035)
const DNS1035LabelMaxLength int = 63

var dns1035LabelRegexp = regexp.MustCompile("^" + dns1035LabelFmt + "$")

// IsDNS1035Label tests for a string that conforms to the definition of a label in
// DNS (RFC 1035).
func isDNS1035Label(value string) []string {
	var errs []string

	if len(value) > DNS1035LabelMaxLength {
		errs = append(errs, maxLenError(DNS1035LabelMaxLength))
	}

	if !dns1035LabelRegexp.MatchString(value) {
		errs = append(errs, regexError(dns1035LabelErrMsg, dns1035LabelFmt, "my-name", "abc-123"))
	}
	return errs
}
