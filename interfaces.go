package main

import "fmt"

type User interface {
	Permisos() int
	Nombre() string
}

type Admin struct {
	nombre string
}

func (this Admin) Permisos() int {
	return 5
}

func (this Admin) Nombre() string {
	return this.nombre
}

type Editor struct {
	nombre string
}

func (this Editor) Permisos() int {
	return 4
}

func (this Editor) Nombre() string {
	return this.nombre
}

func auth(user User) string {
	permisos := user.Permisos()
	if permisos > 4 {
		return user.Nombre() + " tiene permisos de Administrador"
	} else if permisos == 4 {
		return user.Nombre() + " tiene permisos de Editor"
	}

	return ""
}

func main() {
	/*
		admin := Admin{nombre: "Alfredo"}
		editor := Editor{nombre:"Editor"}
		fmt.Println(auth(admin))
		fmt.Println(auth(editor))
	*/
	//Otra manera de realizar el ejemplo
	usuarios := []User{Admin{"Admin"}, Editor{"Edito"}}
	for _, usuario := range usuarios {
		fmt.Println(auth(usuario))
	}
}
