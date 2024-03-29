// Autor: Tiago Rodrigues da Cunha Cabral
// Matrícula: 15/0150296
// Disciplina: Segurança Computacional - 117927 - Turma A
// Algoritmo: Troca de Chaves Diffie-Hellman

package main

var (
	// commonPrime      int64   = 1091
	// commonPrime      int64   = 105929
	// commonPrime      int64   = 1301077
	// commonPrime      int64   = 15487457
	// commonPrime      int64   = 86033551
	// commonPrime      int64   = 122955661
	// commonPrime      int64   = 160487039
	// commonPrime      int64   = 236893021
	commonPrime      int64   = 548609707
	commonSquareRoot float64 = 5
	aliceSecretKey   float64 = 424124212412
	bobSecretKey     float64 = 351232133213
	// commonSquareRoot float64 = 9
)

func modularExponent(x int64, y int64, modulos int64) int64 {
	if y == 0 {
		return 1
	}
	if y%2 == 1 {
		return (x * modularExponent(x, y-1, modulos)) % modulos
	}
	t := modularExponent(x, y/2, modulos)
	return (t * t) % modulos
}

// mixKeys Returns the "Mix" of the Secret Key operated with the common prime and square root base.
// (( commonSquare ** privateKey ) mod commonPrime )
func mixKeys(privateKey float64) int64 {
	mixedKey := modularExponent(int64(commonSquareRoot), int64(privateKey), commonPrime)
	return int64(int64(mixedKey) % commonPrime)
}

// mixSecretKeys Returns a "Mix" of the received mixed key from the other part
// in the communication, and operates it with own's secretKey
// (( receivedMixKey ** ownSecretKey ) mod commonPrime )
func mixSecretKeys(receivedMixKey int64, ownSecretKey float64) int64 {
	mixedKey := modularExponent(int64(receivedMixKey), int64(ownSecretKey), commonPrime)
	return int64(int64(mixedKey) % commonPrime)
}

func main() {

	PrintSecretKeys(aliceSecretKey, bobSecretKey)

	aliceMixedKey := mixKeys(aliceSecretKey)
	bobMixedKey := mixKeys(bobSecretKey)

	PrintMixedKeys(aliceMixedKey, bobMixedKey)

	PrintKeyExchange(aliceMixedKey, bobMixedKey)

	aliceMixedSecret := mixSecretKeys(bobMixedKey, aliceSecretKey)
	bobMixedSecret := mixSecretKeys(aliceMixedKey, bobSecretKey)

	PrintCommonSecretKey(aliceMixedSecret, bobMixedSecret)
}
