package p2p

import "testing"

func TestTCPTransport(t *testing.T) {
	const listenAddr = ":4000"
	tr := NewTCPTransport(listenAddr)

	if tr.listenAddress != listenAddr {
		t.Errorf("listen address does not match, tr: %s, listenAddr: %s", tr.listenAddress, listenAddr)
	}

	err := tr.ListenAndAccept()
	if err != nil {
		t.Errorf("error: tr.istenAndAccept")
	}

	select {}
}
