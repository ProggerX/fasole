package main


type task struct {
	name string
	is_complete bool
}

type creationDialog struct {
	isShown bool
	name string
}

type model struct {
	creationd creationDialog
	tasks []task
	showingHelp bool
	cursor int
}

func initialModel(tasks []task) model {
	return model{
		creationd: creationDialog{
			isShown: false,
			name: "",
		},
		showingHelp: false,
		tasks: tasks,
		cursor: 0,
	}
}
