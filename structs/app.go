package main

import "fmt"

type User struct {
	id        int
	name      string
	age       int
	education []string
	metadata  map[string]string
}

// acts like a method on the struct
func (u User) toString() string {
	return fmt.Sprintf(`
	{
		"id": %v,
		"name": %v,
		"age": %v,
		"education": %v,
		"metadata": %v
	}`, u.id, u.name, u.age, u.education, u.metadata)
}

func stressUser(u User) User {
	u.age = u.age + 1
	return u
}

func main() {
	var newUser User = User{
		id:        1,
		age:       32,
		name:      "Sujal Mandal",
		education: []string{"B.Tech", "PGD cyber security"},
		metadata: map[string]string{
			"email":          "ss6sujal@gmail.com",
			"phone":          "1234567890",
			"currentAddress": "United Arab Emirates",
		},
	}
	fmt.Printf("User fetched ---> %v \n", newUser)

	//individual fields can be updated and accessed using the dot operator
	newUser.id = 2
	fmt.Printf("Updated user object ---> %v \n", newUser)

	//invoke a receiver function on the struct
	jsonString := newUser.toString()
	fmt.Printf("User object as string ---> %v \n", jsonString)

	//pass by value for structs - update the age of the user
	oldUser := stressUser(newUser)

	//log details of the old user
	fmt.Printf("new user object as string (after calling method that updates the User param) ---> %v \n", newUser.toString())
	//log details of the new user
	fmt.Printf("old user object as string ---> %v \n", oldUser.toString())
}
