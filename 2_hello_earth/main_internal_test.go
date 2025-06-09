package main

import "testing"

// func ExampleMain() {
// 	main()
// 	//Output:
// 	// Hello, world!
// }

// func TestGreet_English(t *testing.T) {
// 	lang := language("en")
// 	want := "Hello, world!"
// 	got := greet(lang)

// 	if got != want {
// 		// mark test as failed
// 		t.Errorf("expected: %q, got %q", want, got)
// 	}
// }

// func TestGreet_French(t *testing.T) {
// 	lang := language("fr")
// 	want := "Bonjour, monde!"
// 	got := greet(lang)

// 	if got != want {
// 		// mark test as failed
// 		t.Errorf("expected: %q, got %q", want, got)
// 	}
// }

// func TestGreet_Akkadian(t *testing.T) {
// 	lang := language("ak")
// 	want := ""
// 	got := greet(lang)

// 	if got != want {
// 		// mark test as failed
// 		t.Errorf("expected: %q, got %q", want, got)
// 	}
// }

func TestGreet(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	var tests = map[language]testCase{
		"English": {
			lang: "en",
			want: "Hello, world!",
		},
		"French": {
			lang: "fr",
			want: "Bonjour, monde!",
		},
		"Chinese": {
			lang: "cn",
			want: "你好！",
		},
		"German": {
			lang: "de",
			want: "Hallo, welt!",
		},
		"Undefined": {
			lang: "",
			want: `unsupported language: ""`,
		},
	}

	for name, tc := range tests {
		t.Run(string(name), func(t *testing.T) {
			got := greet(tc.lang)
			if got != tc.want {
				t.Errorf("expected %q, got %q", tc.want, got)
			}
		})
	}

}
