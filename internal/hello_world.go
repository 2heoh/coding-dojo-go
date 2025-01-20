package internal

type HelloWorld struct {
}

func (h HelloWorld) HelloWorld() (string, error) {
	return "Hello, World!", nil
}
