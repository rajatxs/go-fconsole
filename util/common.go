package util

func Attempt(err error) {
	if err != nil {
		Log.Fatal(err.Error())
	}
}
