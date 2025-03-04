//go:build integration_tests
// +build integration_tests

package integration

const (
	// kongSystemServiceCert is a testing TLS certificate with SAN *.kong-system.svc.
	kongSystemServiceCert = `-----BEGIN CERTIFICATE-----
MIIDADCCAeigAwIBAgIUcwR1DfxFoRBBPojxpzXJpGDVKKgwDQYJKoZIhvcNAQEL
BQAwKTELMAkGA1UEBhMCUEwxGjAYBgNVBAMMESoua29uZy1zeXN0ZW0uc3ZjMB4X
DTIxMTIwNjE1MzAzNloXDTIyMDEwNTE1MzAzNlowKTELMAkGA1UEBhMCUEwxGjAY
BgNVBAMMESoua29uZy1zeXN0ZW0uc3ZjMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A
MIIBCgKCAQEA4ICNz0/82BJoKLv7GMskNIDv5zPBzx/Gv3Qy+9FuyPwuuXbCTIok
wEg/t1GUtysHHT/Vm7x313ykVFFkG21dx8BIAY0z4nW3Z1A/ZFpSKVkk3tPDdRgX
yLLQkPzVBbnKQNeUZ3BuJ4pGWl5v8va5UO49seAtOTgUKRlLZMiH7vGWqKwg27M+
5Ap7eNPmw+yWNuIslHFdJr4ufB7ExxKXFNBNS0d15BtwaN9Be7Ox7ux6EPKdrOxS
IZyCzmEpsKMx2CuAeQwlFKpXWb8jhmsh+ieAhqhS5X1bHXL3mCk8pn/d+zcX+dbu
/G1//8LfBZacDR2QgmhZifHXSC8yT/CjsQIDAQABoyAwHjAcBgNVHREEFTATghEq
Lmtvbmctc3lzdGVtLnN2YzANBgkqhkiG9w0BAQsFAAOCAQEAcLwLZD7Co/kQUrLZ
sfuiMmyKi3r1XbanJqKccMOVaqh8L4odz9w0SM0AHne4vk2XMFBjmOdtxT0O9PAc
Jh2vGbyNDaKdL5oYRzJ7j75X0FQXhtpD1iAgWb09EH2WQ90PYeYTWiIFCtM5GdA6
L+Ardq+Cvdy9LigRkyHweDbjelEbe8SPbpbVxp8qOzu94pgclfw+zD5UsudpRYlv
ZcYAmU5vZhpxFwHCEkzW8nUDjT3CIQ9KBaVOX2kXCkVc8Im4Bn0HTVDtd7f97PgM
otVM+9o/Q34HJZS0KdLYHtAUH8S5oUjfIabgcp1nm7/FGwgsHp3tYISC9D4QaFCC
jzqXYQ==
-----END CERTIFICATE-----`
	// kongSystemServiceKey is the private key for the testing certificate kongSystemServiceCert.
	kongSystemServiceKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA4ICNz0/82BJoKLv7GMskNIDv5zPBzx/Gv3Qy+9FuyPwuuXbC
TIokwEg/t1GUtysHHT/Vm7x313ykVFFkG21dx8BIAY0z4nW3Z1A/ZFpSKVkk3tPD
dRgXyLLQkPzVBbnKQNeUZ3BuJ4pGWl5v8va5UO49seAtOTgUKRlLZMiH7vGWqKwg
27M+5Ap7eNPmw+yWNuIslHFdJr4ufB7ExxKXFNBNS0d15BtwaN9Be7Ox7ux6EPKd
rOxSIZyCzmEpsKMx2CuAeQwlFKpXWb8jhmsh+ieAhqhS5X1bHXL3mCk8pn/d+zcX
+dbu/G1//8LfBZacDR2QgmhZifHXSC8yT/CjsQIDAQABAoIBAB7wmJqho28D2mcC
wTBBjtPNkUKD14n8DyADm6Mo0ePRHX9h5pU11KrLSjyxeZVk0K4vRfkYmEuSWfNk
5C8De5Ez5riQBT6IiqYqYRIrgHdCWdp7xMw2bdCzFBtnPNR1LnKRQ1qeHBBG9jsu
GK+bYR7ONqJ1CsZ//AdN/R3+pP1A/CMPwvq43j84KSoqPJrgMHJXullAbUlKlY29
jtrG4GUvt3BiGXC/kKFYI41xAuVI5SdEXu+ycZcU8Ifx2Gpg8Fjp3f4l7GRwvY5K
OE0zP9GSahB6rJF7GZCaKWxOBBix7GjzFPogfcuHCa9Fjwm4IQmC6w3/rYf4+9+v
8M25uKECgYEA+2ki69Hq4jEHUW/B24j+jYKBMlRa3Cd1ci4CP36R1K0080UCz2/4
FmSUyrecIL8C3UZtHQg3EeQrPSJ8h+FrYB48XfLcwDXDWlTIzxPmiWUlG8YK3KLA
LN7kvMuRmOeASBFJIVQuGo9kH+xnFmR7vbYuZ8HUJiMLoWUl8vNuqusCgYEA5Jmr
9JItPPOuWVww6qj+fhbIeBGGjzvUciYdhi5DrIWI2pap9Jn6C08YqsKaKkqPwCjO
VwhwSFwYZbs+pC8YOUxZ9vzkCo623GUcPKe0EjW2GkvufsJohE3ecQkmeve96Qyi
99wmrYmyYfABz+ff1edC7VFk79cuqUiC/3v0TNMCgYAhkAaObsapjZwJfh7mHOLG
p25x9prumwHtzUCVk2MKflj8RPE8GhmHe8P1UA+yu205dwZoAsm/RLOVBL6VMT2x
Zjfu3tYjfsnmjD0GkASNwQf0LjsS+1Mmalck8RQt0nHorQ4TOfaxqwTV0ixs69st
F14YkeKteK47zJIFXgQfIwKBgB8kUCihQUhsafQCeyd8ni7PK8AvowUgQXDLgHon
E1ENX/dnTv/jegzQWavpltbsEWk8Jd/1ZlZ1NV2mhIIZaFNl81uSV/6YMpETtSUO
M5nHd2ddsL/T/CkJ8qOze2qFFXoKHqlldF9vwr1U1OpdzEB3oMZzsCx8Q/8LwczM
NhvBAoGBAPtiu+/HDixPpuFeUsoGqN46rjjyM6sG6AFjjARie+d+jZi5Q7rEfrtY
7Oh25R61i2QlO0fTi01jK6VH6O1G+D2AciGK7cWgLCAJv4ppAA3ysbRbfL1mXS3B
O5PMpLb3dETohTyDk7+r1UgPmKSFDs4OnO5mS1dvQGM1f3OpcgJQ
-----END RSA PRIVATE KEY-----`
)

const (
	// XXX (this hack is tracked in https://github.com/Kong/kubernetes-ingress-controller/issues/1613):
	//
	// The test process (`go test github.com/Kong/kubernetes-ingress-controller/test/integration/...`) serves the webhook
	// endpoints to be consumed by the apiserver (so that the tests can apply a ValidatingWebhookConfiguration and test
	// those validations).
	// In order to make that possible, we needed to allow the apiserver (that gets spun up by the test harness) to access
	// the system under test (which runs as a part of the `go test` process).
	// In the constants below, we're making an audacious assumption that the host running the `go test` process is also
	// the Docker host on the default bridge (therefore it can listen on 172.17.0.1), and that the apiserver
	// is running within a context (such as KIND running on that same docker bridge), from which 172.17.0.1 is routable.
	// This works if the test runs against a KIND cluster, and does not work against cloud providers (like GKE).
	admissionWebhookListenHost = "172.17.0.1"
	admissionWebhookListenPort = 49023
)
