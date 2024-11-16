package main



/*

Step-by-step Plan:

Define an abstract factory interface with methods for creating abstract products.

Create concrete factories that implement the abstract factory interface.


Define abstract product interfaces.


Create concrete products that implement the abstract product interfaces.

Use the abstract factory to create families of related objects.

*/

import "fmt"

// Abstract Product A
type Button interface {
    Render() string
}

// Abstract Product B
type Checkbox interface {
    Render() string
}

// Concrete Product A1
type WindowsButton struct{}

func (b *WindowsButton) Render() string {
    return "Rendering Windows Button"
}

// Concrete Product B1
type WindowsCheckbox struct{}

func (c *WindowsCheckbox) Render() string {
    return "Rendering Windows Checkbox"
}

// Concrete Product A2
type MacButton struct{}

func (b *MacButton) Render() string {
    return "Rendering Mac Button"
}

// Concrete Product B2
type MacCheckbox struct{}

func (c *MacCheckbox) Render() string {
    return "Rendering Mac Checkbox"
}

// Abstract Factory
type GUIFactory interface {
    CreateButton() Button
    CreateCheckbox() Checkbox
}

// Concrete Factory 1
type WindowsFactory struct{}

func (f *WindowsFactory) CreateButton() Button {
    return &WindowsButton{}
}

func (f *WindowsFactory) CreateCheckbox() Checkbox {
    return &WindowsCheckbox{}
}

// Concrete Factory 2
type MacFactory struct{}

func (f *MacFactory) CreateButton() Button {
    return &MacButton{}
}

func (f *MacFactory) CreateCheckbox() Checkbox {
    return &MacCheckbox{}
}

// Client code
func main() {
    var factory GUIFactory

    // Change the factory type to create different families of products
    factory = &WindowsFactory{}
    button := factory.CreateButton()
    checkbox := factory.CreateCheckbox()

    fmt.Println(button.Render())
    fmt.Println(checkbox.Render())

    factory = &MacFactory{}
    button = factory.CreateButton()
    checkbox = factory.CreateCheckbox()

    fmt.Println(button.Render())
    fmt.Println(checkbox.Render())
}


/*


Explanation
Abstract Products: Button and Checkbox interfaces define the methods that all concrete products must implement.

Concrete Products: WindowsButton, WindowsCheckbox, MacButton, and MacCheckbox implement the abstract product interfaces.

Abstract Factory: GUIFactory interface declares methods for creating abstract products.

Concrete Factories: WindowsFactory and MacFactory implement the GUIFactory interface to create concrete products.

Client Code: Uses the abstract factory to create families of related objects without knowing their concrete classes.
main





*/