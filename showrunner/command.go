package showrunner

type Command interface {
	Main(args []string) error
}
