package main

import "github.com/01-edu/z01"

type Door struct {
    state bool
}

func CloseDoor(ptrDoor *Door) bool {
    for _, r := range "Door Closing..." {
        z01.PrintRune(r)
    }
    z01.PrintRune('\n')
    ptrDoor.state = false
    return true
}

func OpenDoor(ptrDoor *Door) bool {
    for _, r := range "Door Opening..." {
        z01.PrintRune(r)
    }
    z01.PrintRune('\n')
    ptrDoor.state = true
    return true
}

func IsDoorOpen(ptrDoor *Door) bool {
    for _, r := range "is the Door opened ?" {
        z01.PrintRune(r)
    }
    z01.PrintRune('\n')
    return ptrDoor.state
}

func IsDoorClose(ptrDoor *Door) bool {
    for _, r := range "is the Door closed ?" {
        z01.PrintRune(r)
    }
    z01.PrintRune('\n')
    return !ptrDoor.state
}

func main() {
    door := &Door{}

    OpenDoor(door)
    if IsDoorClose(door) {
        OpenDoor(door)
    }
    if IsDoorOpen(door) {
        CloseDoor(door)
    }
    if door.state {
        CloseDoor(door)
    }
}
