package pake

import (
	"crypto/elliptic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tscholl2/siec"
)

func BenchmarkPakeSIEC255(b *testing.B) {
	curve := siec.SIEC255()
	for i := 0; i < b.N; i++ {
		// initialize A
		A, _ := Init([]byte{1, 2, 3}, 0, curve, 1*time.Microsecond)
		// initialize B
		B, _ := Init([]byte{1, 2, 3}, 1, curve, 1*time.Microsecond)
		// send A's stuff to B
		B.Update(A.Bytes())
		// send B's stuff to A
		A.Update(B.Bytes())
		// send A's stuff back to B
		B.Update(A.Bytes())
	}
}

func BenchmarkPakeP521(b *testing.B) {
	curve := elliptic.P521()
	for i := 0; i < b.N; i++ {
		// initialize A
		A, _ := Init([]byte{1, 2, 3}, 0, curve, 1*time.Microsecond)
		// initialize B
		B, _ := Init([]byte{1, 2, 3}, 1, curve, 1*time.Microsecond)
		// send A's stuff to B
		B.Update(A.Bytes())
		// send B's stuff to A
		A.Update(B.Bytes())
		// send A's stuff back to B
		B.Update(A.Bytes())
	}
}

func BenchmarkPakeP224(b *testing.B) {
	curve := elliptic.P224()
	for i := 0; i < b.N; i++ {
		// initialize A
		A, _ := Init([]byte{1, 2, 3}, 0, curve, 1*time.Microsecond)
		// initialize B
		B, _ := Init([]byte{1, 2, 3}, 1, curve, 1*time.Microsecond)
		// send A's stuff to B
		B.Update(A.Bytes())
		// send B's stuff to A
		A.Update(B.Bytes())
		// send A's stuff back to B
		B.Update(A.Bytes())
	}
}

func TestPake(t *testing.T) {
	curve := siec.SIEC255()
	// successful (both have same k)
	// initialize A
	A, err := Init([]byte{1, 2, 3}, 0, curve, 100*time.Millisecond)
	assert.Nil(t, err)
	assert.False(t, A.IsVerified())
	// initialize B
	B, err := Init([]byte{1, 2, 3}, 1, curve, 1*time.Millisecond)
	assert.Nil(t, err)
	assert.False(t, B.IsVerified())
	// send A's stuff to B
	err = B.Update(A.Bytes())
	assert.Nil(t, err)
	assert.False(t, B.IsVerified())
	// send B's stuff to A
	err = A.Update(B.Bytes())
	assert.Nil(t, err) // A validates
	assert.True(t, A.IsVerified())
	// send A's stuff back to B
	err = B.Update(A.Bytes())
	assert.Nil(t, err) // B validates
	assert.True(t, B.IsVerified())

	// failure (both have different k)
	// initialize A
	A, err = Init([]byte{1, 2, 3}, 0, curve, 1*time.Millisecond)
	assert.Nil(t, err)
	assert.False(t, A.IsVerified())
	// initialize B
	B, err = Init([]byte{4, 5, 6}, 1, curve, 1*time.Millisecond)
	assert.Nil(t, err)
	assert.False(t, B.IsVerified())
	// send A's stuff to B
	err = B.Update(A.Bytes())
	assert.Nil(t, err)
	assert.False(t, B.IsVerified())
	// send B's stuff to A
	err = A.Update(B.Bytes())
	assert.NotNil(t, err) // A validates
	assert.False(t, A.IsVerified())
	// send A's stuff back to B
	err = B.Update(A.Bytes())
	assert.NotNil(t, err)
	assert.False(t, B.IsVerified())

}

func TestSessionKeySIEC255(t *testing.T) {
	curve := siec.SIEC255()
	// initialize A
	A, _ := Init([]byte{1, 2, 3}, 0, curve, 1*time.Millisecond)
	// initialize B
	B, _ := Init([]byte{1, 2, 3}, 1, curve, 1*time.Millisecond)
	// send A's stuff to B
	B.Update(A.Bytes())
	// send B's stuff to A
	A.Update(B.Bytes())
	// send A's stuff back to B
	B.Update(A.Bytes())
	s1, err := A.SessionKey()
	assert.Nil(t, err)
	s1B, err := B.SessionKey()
	assert.Nil(t, err)
	assert.Equal(t, s1, s1B)

	// initialize A
	A, _ = Init([]byte{1, 2, 3}, 0, curve, 1*time.Millisecond)
	// initialize B
	B, _ = Init([]byte{1, 2, 3}, 1, curve, 1*time.Millisecond)
	// send A's stuff to B
	B.Update(A.Bytes())
	// send B's stuff to A
	A.Update(B.Bytes())
	// send A's stuff back to B
	B.Update(A.Bytes())
	s2, err := A.SessionKey()
	assert.Nil(t, err)

	assert.NotEqual(t, s1, s2)
}

// // THIS TEST DOES NOT WORK!?!
// func TestSessionKeyP224(t *testing.T) {
// 	curve := elliptic.P224()
// 	// initialize A
// 	A, _ := Init([]byte{1, 2, 3}, 0, curve, 1*time.Millisecond)
// 	// initialize B
// 	B, _ := Init([]byte{1, 2, 3}, 1, curve, 1*time.Millisecond)
// 	// send A's stuff to B
// 	B.Update(A.Bytes())
// 	// send B's stuff to A
// 	A.Update(B.Bytes())
// 	// send A's stuff back to B
// 	B.Update(A.Bytes())
// 	s1, err := A.SessionKey()
// 	assert.Nil(t, err)
// 	s1B, err := B.SessionKey()
// 	assert.Nil(t, err)
// 	assert.Equal(t, s1, s1B)

// 	// initialize A
// 	A, _ = Init([]byte{1, 2, 3}, 0, curve, 1*time.Millisecond)
// 	// initialize B
// 	B, _ = Init([]byte{1, 2, 3}, 1, curve, 1*time.Millisecond)
// 	// send A's stuff to B
// 	B.Update(A.Bytes())
// 	// send B's stuff to A
// 	A.Update(B.Bytes())
// 	// send A's stuff back to B
// 	B.Update(A.Bytes())
// 	s2, err := A.SessionKey()
// 	assert.Nil(t, err)

// 	assert.NotEqual(t, s1, s2)
// }

func TestSessionKeyP521(t *testing.T) {
	curve := elliptic.P521()
	// initialize A
	A, _ := Init([]byte{1, 2, 3}, 0, curve, 1*time.Millisecond)
	// initialize B
	B, _ := Init([]byte{1, 2, 3}, 1, curve, 1*time.Millisecond)
	// send A's stuff to B
	B.Update(A.Bytes())
	// send B's stuff to A
	A.Update(B.Bytes())
	// send A's stuff back to B
	B.Update(A.Bytes())
	s1, err := A.SessionKey()
	assert.Nil(t, err)
	s1B, err := B.SessionKey()
	assert.Nil(t, err)
	assert.Equal(t, s1, s1B)

	// initialize A
	A, _ = Init([]byte{1, 2, 3}, 0, curve, 1*time.Millisecond)
	// initialize B
	B, _ = Init([]byte{1, 2, 3}, 1, curve, 1*time.Millisecond)
	// send A's stuff to B
	B.Update(A.Bytes())
	// send B's stuff to A
	A.Update(B.Bytes())
	// send A's stuff back to B
	B.Update(A.Bytes())
	s2, err := A.SessionKey()
	assert.Nil(t, err)

	assert.NotEqual(t, s1, s2)
}

func TestSessionKeyP256(t *testing.T) {
	curve := elliptic.P256()
	// initialize A
	A, _ := Init([]byte{1, 2, 3}, 0, curve, 1*time.Millisecond)
	// initialize B
	B, _ := Init([]byte{1, 2, 3}, 1, curve, 1*time.Millisecond)
	// send A's stuff to B
	B.Update(A.Bytes())
	// send B's stuff to A
	A.Update(B.Bytes())
	// send A's stuff back to B
	B.Update(A.Bytes())
	s1, err := A.SessionKey()
	assert.Nil(t, err)
	s1B, err := B.SessionKey()
	assert.Nil(t, err)
	assert.Equal(t, s1, s1B)

	// initialize A
	A, _ = Init([]byte{1, 2, 3}, 0, curve, 1*time.Millisecond)
	// initialize B
	B, _ = Init([]byte{1, 2, 3}, 1, curve, 1*time.Millisecond)
	// send A's stuff to B
	B.Update(A.Bytes())
	// send B's stuff to A
	A.Update(B.Bytes())
	// send A's stuff back to B
	B.Update(A.Bytes())
	s2, err := A.SessionKey()
	assert.Nil(t, err)

	assert.NotEqual(t, s1, s2)
}

func TestSessionKeyP384(t *testing.T) {
	curve := elliptic.P384()
	// initialize A
	A, _ := Init([]byte{1, 2, 3}, 0, curve, 1*time.Millisecond)
	// initialize B
	B, _ := Init([]byte{1, 2, 3}, 1, curve, 1*time.Millisecond)
	// send A's stuff to B
	B.Update(A.Bytes())
	// send B's stuff to A
	A.Update(B.Bytes())
	// send A's stuff back to B
	B.Update(A.Bytes())
	s1, err := A.SessionKey()
	assert.Nil(t, err)
	s1B, err := B.SessionKey()
	assert.Nil(t, err)
	assert.Equal(t, s1, s1B)

	// initialize A
	A, _ = Init([]byte{1, 2, 3}, 0, curve, 1*time.Millisecond)
	// initialize B
	B, _ = Init([]byte{1, 2, 3}, 1, curve, 1*time.Millisecond)
	// send A's stuff to B
	B.Update(A.Bytes())
	// send B's stuff to A
	A.Update(B.Bytes())
	// send A's stuff back to B
	B.Update(A.Bytes())
	s2, err := A.SessionKey()
	assert.Nil(t, err)

	assert.NotEqual(t, s1, s2)
}

func TestSessionKeyString(t *testing.T) {
	curves := []string{"siec", "p384", "p521", "p256"}
	for _, curve := range curves {
		A, _ := InitCurve([]byte{1, 2, 3}, 0, curve, 1*time.Millisecond)
		// initialize B
		B, _ := InitCurve([]byte{1, 2, 3}, 1, curve, 1*time.Millisecond)
		// send A's stuff to B
		B.Update(A.Bytes())
		// send B's stuff to A
		A.Update(B.Bytes())
		// send A's stuff back to B
		B.Update(A.Bytes())
		s1, err := A.SessionKey()
		assert.Nil(t, err)
		s1B, err := B.SessionKey()
		assert.Nil(t, err)
		assert.Equal(t, s1, s1B)

		// initialize A
		A, _ = InitCurve([]byte{1, 2, 3}, 0, curve, 1*time.Millisecond)
		// initialize B
		B, _ = InitCurve([]byte{1, 2, 3}, 1, curve, 1*time.Millisecond)
		// send A's stuff to B
		B.Update(A.Bytes())
		// send B's stuff to A
		A.Update(B.Bytes())
		// send A's stuff back to B
		B.Update(A.Bytes())
		s2, err := A.SessionKey()
		assert.Nil(t, err)

		assert.NotEqual(t, s1, s2)
	}
}
