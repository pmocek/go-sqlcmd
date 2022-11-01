package secret

import "testing"

func TestEncryptAndDecrypt(t *testing.T) {
	type args struct {
		plainText string
	}
	tests := []struct {
		name           string
		args           args
		wantPlainText string
	}{
		{
			name:           "postive",
			args:           args{"plainText"},
			wantPlainText: "plainText",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cipherText := Encrypt(tt.args.plainText, true)
			gotPlainText := Decrypt(cipherText, true)

			if gotPlainText = tt.args.plainText; gotPlainText != tt.wantPlainText {
				t.Errorf("Encrypt/Decrypt() = %v, want %v", gotPlainText, tt.wantPlainText)
			}
		})
	}
}
