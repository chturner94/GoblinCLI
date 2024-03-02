package goblincli

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/term"
)

// App serves as the main struct of our application, you can populate the App with
// MenuOptions to provide new options within your command line application.
type App struct {
	MenuOptions []MenuOptions
}

// MenuOptions represent each application component you want reperesented within the application.
type MenuOptions struct {
	EventFunction EventFunction
	MenuEntry     string
}

// Event functions have few constraints and represents the task that selecting the menu option
// represents.
type EventFunction func()

// The run method is the primary method associated with the App struct, and is responsible for
// setting up the application and providing the event loop and exit condition for the application.
// This function should extend and be a jumping off point and act as a new main() func for managing
// your applet. We recommend extending the widget struct here.
// Here is how you would set up this application:
//
//	func main() {
//	  myApp := App{
//	    MenuOptions: []MenuOptions{
//	      {EventFunction: option1, MenuEntry: "1. Your description here."},
//	      {EventFunction: option2, MenueEntry: "2. Your description here."}
//	    }
//	  }
//	}
//
// The only thing left for you to provide are functions accessible to the main function which can serve
// as the eventFunctions.
func (app *App) Run() {
	oldState, err := term.GetState(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatal(err)
	}
	print("\033[H\033[2J")
	defer func() {
		fmt.Println("Program is quitting...")
		term.Restore(int(os.Stdin.Fd()), oldState)
		os.Exit(0)
	}()
	terminal := NewTerminal()
	for {
		fmt.Println("Select an option:")
		for _, option := range app.MenuOptions {
			terminal.WriteLine(option.MenuEntry)
		}
		a := []any{fmt.Sprintf("%d. Exit", len(app.MenuOptions)+1)}
		err := terminal.WriteLine(fmt.Sprint(a...))
		if err != nil {
			terminal.WriteLine("Error writing to terminal" + err.Error())
			return
		}

		terminal.WriteLine("Enter your choice: ")
		choiceStr, err := terminal.ReadLine()
		if err != nil {
			terminal.WriteLine("Error reading input: " + err.Error())
			return
		}

		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			terminal.WriteLine("Invalid input. Please enter a number.")
			continue
		}

		if choice > 0 && choice <= len(app.MenuOptions) {
			app.MenuOptions[choice-1].EventFunction()
		} else if choice == len(app.MenuOptions)+1 {
			terminal.WriteLine("Exiting...")
			return
		} else {
			terminal.WriteLine("Invalid choice. Please try again.")
		}
		terminal.WriteLine("")
	}
}
