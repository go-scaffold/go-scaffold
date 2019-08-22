package promptcli

type runnablePrompt interface {
	Run() (string, error)
}
