package clusterconfig

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/runtime"

	configv1 "github.com/openshift/api/config/v1"
	operatorfake "github.com/openshift/client-go/operator/clientset/versioned/fake"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corefake "k8s.io/client-go/kubernetes/fake"

	"github.com/openshift/insights-operator/pkg/record"
	"github.com/stretchr/testify/assert"
)

func Test_gatherClusterIngressCertificates(t *testing.T) {
	certData := `
-----BEGIN CERTIFICATE-----
MIIDazCCAlOgAwIBAgIUfTstqHMAhGLL+j3n6pmwLw8vt84wDQYJKoZIhvcNAQEL
BQAwRTELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoM
GEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0yNDAzMDYxMjU5MTFaFw0yNTAz
MDYxMjU5MTFaMEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEw
HwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwggEiMA0GCSqGSIb3DQEB
AQUAA4IBDwAwggEKAoIBAQDoqkPZMeMi5qjkG384ZwpAc3QScOGYBWOEDFAioq5C
YhtDGBSMq2VwS0r8RvEEhbebvXuH5PLcIuEVO/MZRQD9gSacCfLlMLRKZYpv168m
KUYyhx1bXKmUlbQxCnpAPZ7nf14A3Pb0TzLfsKjoUdUOv/1eorA6+oU78StWx/Nt
W94ad9n3o+cjiMPu/RS3g9b+x07bG5mFYuzpWk/Svb5Lb42g8AtonzqFJBbhlStU
A+9UyzmyXMeTlbI9fFmku7mb5Uq0SZ8jhpH+fyCoQOxefTfVrvjrkkdavDn43hjz
5hCwGJ3mV96MU9hh398oBguOHaJ6V3/UHtW1spsFY83RAgMBAAGjUzBRMB0GA1Ud
DgQWBBR/zvuHjFadvifzAGBHegOxmRXnCzAfBgNVHSMEGDAWgBR/zvuHjFadvifz
AGBHegOxmRXnCzAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQBs
E2U+jQzJuEt9e6UEnS1T0cb2NxaGb7CYsSX0TjZK1VgloAKbnxaCLjRruTOOwfm6
s5CFzFjJoIhUASzoA295Np2AR0UEYr5fendIjKCztCMlpj0fp92jFL6/RZWNGM1A
qECHYtZckeqJjg9vUdfHtiBRoyEHJUJ/tsDDlslwzocdJUqKL8V6KerZsh5SIAkS
rJ8EgVyDvwQaLPQMttjk62croI1Wi3FLmkvvtTbNcMgTnVhFfGjyHOiGnQeBfqax
5P0VBuAUCihegskKEUCJB8HFPkC4hqbrEk0+psQ2Gm8kjoll/SpltFLS77Xjhrz9
1qaiDHuWnUSifz6SGpWr
-----END CERTIFICATE-----
`

	tests := []struct {
		name               string
		ingressDefinitions []configv1.Ingress
		secretData         map[string][]byte
		wantRecords        []record.Record
		wantErrCount       int
	}{
		{
			name: "successful retrieval cluster ingress certificates",
			ingressDefinitions: []configv1.Ingress{
				{ObjectMeta: metav1.ObjectMeta{Name: "example-ingress", Namespace: "openshift-ingress-operator"}},
			},
			secretData: map[string][]byte{
				"tls.crt": []byte(certData),
			},
			wantRecords: []record.Record{
				{
					Name: "aggregated/ingress_controllers_certs",
					Item: record.JSONMarshaller{
						Object: []*CertificateInfo{{
							Name:        "example-ingress",
							Namespace:   "openshift-ingress-operator",
							NotBefore:   time.Date(2024, time.March, 6, 12, 59, 11, 0, time.UTC),
							NotAfter:    time.Date(2025, time.March, 6, 12, 59, 11, 0, time.UTC),
							Controllers: []ControllerInfo{},
						},
						},
						// Object: []IngressControllerInfo{
						// 	{
						// 		Name: "example-ingress",
						// 		OperatorGeneratedCertificate: []IngressCertificateInfo{
						// 			{
						// 				Name:      "router-ca",
						// 				NotBefore: time.Date(2024, time.March, 6, 12, 59, 11, 0, time.UTC),
						// 				NotAfter:  time.Date(2025, time.March, 6, 12, 59, 11, 0, time.UTC),
						// 			},
						// 		},
						// 		CustomCertificates: []IngressCertificateInfo{},
						// 	},
						// },
					},
				},
			},
			wantErrCount: 0,
		},
		// {
		// 	name: "failed retrieval cluster ingress certificates",
		// 	ingressDefinitions: []configv1.Ingress{
		// 		{
		// 			ObjectMeta: metav1.ObjectMeta{
		// 				Name: "example-ingress",
		// 			},
		// 			Spec: configv1.IngressSpec{},
		// 		},
		// 	},
		// 	secretData:   map[string][]byte{},
		// 	wantRecords:  nil,
		// 	wantErrCount: 1, // There should be an error due to missing 'tls.crt'
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			objects := make([]runtime.Object, len(tt.ingressDefinitions))
			for i, ingress := range tt.ingressDefinitions {
				objects[i] = &ingress
			}
			operatorClient := operatorfake.NewSimpleClientset(objects...)
			coreClient := corefake.NewSimpleClientset(&corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{Name: "router-ca", Namespace: "openshift-ingress-operator"},
				Data:       tt.secretData,
			})

			records, errs := gatherClusterIngressCertificates(context.TODO(), coreClient.CoreV1(), operatorClient)
			assert.EqualValues(t, tt.wantRecords, records)
			assert.Len(t, errs, tt.wantErrCount)
		})
	}
}

func Test_getCertificateInfoFromSecret(t *testing.T) {
	// Mocks
	mockCertbytes := []byte(`-----BEGIN CERTIFICATE-----
MIIDazCCAlOgAwIBAgIUfTstqHMAhGLL+j3n6pmwLw8vt84wDQYJKoZIhvcNAQEL
BQAwRTELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoM
GEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0yNDAzMDYxMjU5MTFaFw0yNTAz
MDYxMjU5MTFaMEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEw
HwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwggEiMA0GCSqGSIb3DQEB
AQUAA4IBDwAwggEKAoIBAQDoqkPZMeMi5qjkG384ZwpAc3QScOGYBWOEDFAioq5C
YhtDGBSMq2VwS0r8RvEEhbebvXuH5PLcIuEVO/MZRQD9gSacCfLlMLRKZYpv168m
KUYyhx1bXKmUlbQxCnpAPZ7nf14A3Pb0TzLfsKjoUdUOv/1eorA6+oU78StWx/Nt
W94ad9n3o+cjiMPu/RS3g9b+x07bG5mFYuzpWk/Svb5Lb42g8AtonzqFJBbhlStU
A+9UyzmyXMeTlbI9fFmku7mb5Uq0SZ8jhpH+fyCoQOxefTfVrvjrkkdavDn43hjz
5hCwGJ3mV96MU9hh398oBguOHaJ6V3/UHtW1spsFY83RAgMBAAGjUzBRMB0GA1Ud
DgQWBBR/zvuHjFadvifzAGBHegOxmRXnCzAfBgNVHSMEGDAWgBR/zvuHjFadvifz
AGBHegOxmRXnCzAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQBs
E2U+jQzJuEt9e6UEnS1T0cb2NxaGb7CYsSX0TjZK1VgloAKbnxaCLjRruTOOwfm6
s5CFzFjJoIhUASzoA295Np2AR0UEYr5fendIjKCztCMlpj0fp92jFL6/RZWNGM1A
qECHYtZckeqJjg9vUdfHtiBRoyEHJUJ/tsDDlslwzocdJUqKL8V6KerZsh5SIAkS
rJ8EgVyDvwQaLPQMttjk62croI1Wi3FLmkvvtTbNcMgTnVhFfGjyHOiGnQeBfqax
5P0VBuAUCihegskKEUCJB8HFPkC4hqbrEk0+psQ2Gm8kjoll/SpltFLS77Xjhrz9
1qaiDHuWnUSifz6SGpWr
-----END CERTIFICATE-----`)
	b, _ := pem.Decode(mockCertbytes)
	x509cert, _ := x509.ParseCertificate(b.Bytes)

	testCases := []struct {
		name            string
		secret          *corev1.Secret
		expected        *CertificateInfo
		expectErr       bool
		secretName      string
		secretNamespace string
	}{
		{
			name:            "a Secret containing a certificate returns a CertificateInfo struct",
			secretName:      "test",
			secretNamespace: "test-namespace",
			secret: &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: "test-namespace"},
				Data:       map[string][]byte{"tls.crt": mockCertbytes},
			},
			expected: &CertificateInfo{
				Name: "test", Namespace: "test-namespace",
				NotBefore: x509cert.NotBefore, NotAfter: x509cert.NotAfter,
				Controllers: []ControllerInfo{},
			},
		}, {
			name:            "a Secret without a certificate returns an error",
			secretName:      "test",
			secretNamespace: "test-namespace",
			secret:          &corev1.Secret{ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: "test-namespace"}},
			expectErr:       true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			client := corefake.NewSimpleClientset(tc.secret)

			// When
			test, err := getCertificateInfoFromSecret(
				context.Background(), client.CoreV1(), tc.secretNamespace, tc.secretName)

			// Assert
			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.EqualValues(t, tc.expected, test)
		})
	}
}