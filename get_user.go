package go_dummy

func GetUser() map[string]interface{} {
	user := map[string]interface{}{
		"id": "1",
		"name": "John Doe",
		"age": 27,
		"gender": "Male",
		"country": "United Kingdom",
		"city": "London",
		"address": "Baker Street",
		"phone": "1234567890",
		"email": "johndoe@gmail.com",
	}
	return user;
}