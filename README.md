# ARGUS GO Client 

This is the official go library for the [ARGUS Engine](https://github.com/Khelechy/argus), this library helps go developers and applications seamlessly integrate to the ARGUS Engine, authentication and event listening.

### Install the package 

Ensure you have golang +v1.19 installed. 

```sh
    go get github.com/khelechy/argus
```

### Import the package in your code

```go
    import (
        ...
        "github.com/khelechy/argus"
        ...
    )
```

### Using the package

```go
    argus, err := argus.Connect(&argus.Argus{
		Username: "testuser",
		Password: "testpassword",
	})

	if err != nil {
		fmt.Println()
		return
	}

	for {
		select {
		case event := <-argus.Events:
			fmt.Println(event.ActionDescription)
		case message := <-argus.Messages:
			log.Println(message)
		case err := <-argus.Errors:
			log.Println("Error:", err)
		}
	}
```