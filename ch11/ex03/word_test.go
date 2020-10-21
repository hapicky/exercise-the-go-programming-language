package word

import (
	"gopl.io/ch11/word2"
	"math/rand"
	"testing"
	"time"
)

func randomUnPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n+2)
	runes[0] = rune(0x3042) // あ
	for i := 0; i < n; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i+1] = r
	}
	runes[n+1] = rune(0x3044) // い
	return string(runes)
}

func TestRandomUnPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomUnPalindrome(rng)
		if word.IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true", p)
		}
	}
}
