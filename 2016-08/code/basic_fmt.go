func main() {
	trace.Start(os.Stderr)
	... // your concurrent code
	trace.Stop()
}
