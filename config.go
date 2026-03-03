package main

type Project struct {
	Path     string
	Commands []string
}

type Config map[string]Project
