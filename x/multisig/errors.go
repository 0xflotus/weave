package multisig

import (
	"fmt"

	"github.com/iov-one/weave/errors"
)

// ABCI Response Codes
// multisig takes 1030-1040
const (
	CodeInvalidMsg             = 1030
	CodeMultisigAuthentication = 1031
)

var (
	errMissingSigs       = fmt.Errorf("Missing sigs")
	errInvalidThreshold  = fmt.Errorf("Activation threshold must be lower than or equal to the number of sigs")
	errContractDuplicate = fmt.Errorf("Contract already exists")

	errUnauthorizedMultiSig = fmt.Errorf("Multisig authentication failed")
)

func ErrMissingSigs() error {
	return errors.WithCode(errMissingSigs, CodeInvalidMsg)
}
func ErrInvalidActivationThreshold() error {
	return errors.WithCode(errInvalidThreshold, CodeInvalidMsg)
}
func ErrInvalidChangeThreshold() error {
	return errors.WithCode(errInvalidThreshold, CodeInvalidMsg)
}
func ErrContractDuplicate(contract []byte) error {
	msg := fmt.Sprintf("author=%X", contract)
	return errors.WithLog(msg, errContractDuplicate, CodeInvalidMsg)
}
func IsInvalidMsgErr(err error) bool {
	return errors.HasErrorCode(err, CodeInvalidMsg)
}

func ErrUnauthorizedMultiSig(contract []byte) error {
	msg := fmt.Sprintf("contrac=%X", contract)
	return errors.WithLog(msg, errUnauthorizedMultiSig, CodeMultisigAuthentication)
}
func IsMultiSigAuthenticationErr(err error) bool {
	return errors.HasErrorCode(err, CodeMultisigAuthentication)
}
