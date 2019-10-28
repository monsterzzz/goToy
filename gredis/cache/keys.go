package cache

import "fmt"

const (
	VIEW = "view:"
	PLAY = "play:"
)

func GetStudentViewKey(id uint) string {
	return fmt.Sprintf("%s%s%d", VIEW, "student:", int(id))
}

func GetStudentPlayKey(id uint) string {
	return fmt.Sprintf("%s%s%d", PLAY, "student:", int(id))
}
