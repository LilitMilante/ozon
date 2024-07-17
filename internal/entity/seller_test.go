package entity

import (
	"strings"
	"testing"
)

func TestSeller_Normalize(t *testing.T) {
	t.Parallel()

	got := Seller{
		FullName: "   name     ",
		Login:    "   login    ",
	}

	got.Normalize()

	want := Seller{
		FullName: "name",
		Login:    "login",
	}

	if want != got {
		t.Fatalf("want: %s\ngot:%s", want, got)
	}
}

func TestSeller_Validate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		seller Seller
		want   func(t *testing.T, got error)
	}{
		{
			name: "All fields are filled in correctly",
			seller: Seller{
				FullName: strings.Repeat("П", 50),
				Login:    strings.Repeat("Л", 50),
				Password: "12345678",
			},
			want: func(t *testing.T, got error) {
				if got != nil {
					t.Fatalf("want nil\ngot %v", got)
				}
			},
		},
		{
			name: "The full name field is empty",
			seller: Seller{
				FullName: "",
				Login:    "test@.com",
				Password: "12345678",
			},
			want: func(t *testing.T, got error) {
				wantText := "the full name must not be empty and must be between 1 and 50 characters long"
				if got == nil || got.Error() != wantText {
					t.Fatalf("want %v\ngot %v", wantText, got)
				}
			},
		},
		{
			name: "The full name more than 50 characters",
			seller: Seller{
				FullName: "Napu Amo Hala Ona Ona Aneka Wehi Milestones Ona Hiwea Nena Wawa Keho Onka Kahe Hea Leke",
				Login:    "test@.com",
				Password: "12345678",
			},
			want: func(t *testing.T, got error) {
				wantText := "the full name must not be empty and must be between 1 and 50 characters long"
				if got == nil || got.Error() != wantText {
					t.Fatalf("want %v\ngot %v", wantText, got)
				}
			},
		},
		{
			name: "The login field is empty",
			seller: Seller{
				FullName: "Testov Test Testovich",
				Login:    "",
				Password: "12345678",
			},
			want: func(t *testing.T, got error) {
				wantText := "the login must not be empty and must be between 3 and 50 characters long"
				if got == nil || got.Error() != wantText {
					t.Fatalf("want %v\ngot %v", wantText, got)
				}
			},
		},
		{
			name: "The login less than 3 characters",
			seller: Seller{
				FullName: "Testov Test Testovich",
				Login:    "t@",
				Password: "12345678",
			},
			want: func(t *testing.T, got error) {
				wantText := "the login must not be empty and must be between 3 and 50 characters long"
				if got == nil || got.Error() != wantText {
					t.Fatalf("want %v\ngot %v", wantText, got)
				}
			},
		},
		{
			name: "The login more than 50 characters",
			seller: Seller{
				FullName: "Testov Test Testovich",
				Login:    "SuperLongUsernameThatExceedsFiftyCharactersInLengthAndIsUnique@.com",
				Password: "12345678",
			},
			want: func(t *testing.T, got error) {
				wantText := "the login must not be empty and must be between 3 and 50 characters long"
				if got == nil || got.Error() != wantText {
					t.Fatalf("want %v\ngot %v", wantText, got)
				}
			},
		},
		{
			name: "The password field is empty",
			seller: Seller{
				FullName: "Testov Test Testovich",
				Login:    "test@.com",
				Password: "",
			},
			want: func(t *testing.T, got error) {
				wantText := "empty password"
				if got == nil || got.Error() != wantText {
					t.Fatalf("want %v\ngot %v", wantText, got)
				}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			got := testCase.seller.Validate()
			testCase.want(t, got)
		})
	}
}
