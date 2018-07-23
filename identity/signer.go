package identity

type Signer interface {
	Sign(data ...[]byte) ([]byte, error)
}
