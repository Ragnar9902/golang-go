package main

func get_builtin(cmd string)func(...string)int{
	funct_name := []string{
		"exit",
		"version",
	}
	f := []func(...string)int{
		exit,
		version,
	}
	for i, buil := range funct_name {
		if (cmd == buil){
			return f[i]
		}
	}
	return nil
}
