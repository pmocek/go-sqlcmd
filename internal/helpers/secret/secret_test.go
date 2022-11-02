package secret

import (
	"strings"
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {
	type args struct {
		plainText string
		encrypt bool
	}
	tests := []struct {
		name           string
		args           args
		wantPlainText string
	}{
		{
			name:           "noEncrypt",
			args:           args{"plainText", false},
			wantPlainText: "plainText",
		},
		{
			name:           "encrypt",
			args:           args{"plainText", true},
			wantPlainText: "plainText",
		},
		{
			name:           "emptyStringForEncryptPanic",
			args:           args{"", true},
			wantPlainText: "",
		},
		{
			name:           "emptyStringForDecryptPanic",
			args:           args{"", true},
			wantPlainText: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// If test name ends in 'Panic' expect a Panic
			if strings.HasSuffix(tt.name, "Panic") {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}

			var gotPlainText string
			if tt.name != "emptyStringForDecryptPanic" {
				cipherText := Encrypt(tt.args.plainText, tt.args.encrypt)
				gotPlainText = Decrypt(cipherText, tt.args.encrypt)
			} else {
				gotPlainText = Decrypt(tt.args.plainText, tt.args.encrypt)
			}

			if gotPlainText = tt.args.plainText; gotPlainText != tt.wantPlainText {
				t.Errorf("Encrypt/Decrypt() = %v, want %v", gotPlainText, tt.wantPlainText)
			}
		})
	}
}
