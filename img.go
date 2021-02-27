package main

func main() {
	imageName := "nginx:latest"
	err := images.process(imageName)
	if err != nil {
		panic(err)
	}
}
