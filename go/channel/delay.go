package channel

type (
	// Decrypter attempts to decrypt a key
	Decrypter interface {
		Decrypt(cipherText string) (string, error)
	}
)

// https://github.com/DataDog/datadog-lambda-go/blob/1cdcde5b6f3e4c2d3d237e380fab7f6d9f1c1294/internal/wrapper/wrap_handler.go#L44-L70
func decryptAPIKey(decrypter Decrypter, kmsAPIKey string) <-chan string {

	ch := make(chan string)

	// Lambdaのフレームワークなので、ユーザー定義関数が実行される前に
	// 余計なレスポンスラグがかからないようにしている。
	// gorutineで処理を逃して、実行の遅さが立ち上げに影響しないようにしている。
	go func() {
		result, err := decrypter.Decrypt(kmsAPIKey)
		if err != nil {
			// logger.Error(fmt.Errorf("Couldn't decrypt api kms key %s", err))
		}
		ch <- result
		close(ch)
	}()
	return ch
}
