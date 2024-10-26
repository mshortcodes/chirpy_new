package auth

import "testing"

func TestCheckPasswordHash(t *testing.T) {
	password1 := "firstpw123"
	password2 := "secondpw456"
	hash1, _ := HashPassword(password1)
	hash2, _ := HashPassword(password2)

	tests := []struct {
		name     string
		password string
		hash     string
		wantErr  bool
	}{
		{
			name:     "correct password",
			password: password1,
			hash:     hash1,
			wantErr:  false,
		},
		{
			name:     "incorrect password",
			password: "wrongpw",
			hash:     hash1,
			wantErr:  true,
		},
		{
			name:     "password doesn't match hash",
			password: password1,
			hash:     hash2,
			wantErr:  true,
		},
		{
			name:     "empty password",
			password: "",
			hash:     hash1,
			wantErr:  true,
		},
		{
			name:     "invalid hash",
			password: password1,
			hash:     "invalidhash",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckPasswordHash(tt.password, tt.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPasswordHash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
