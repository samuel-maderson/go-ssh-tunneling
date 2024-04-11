package types

type Args struct {
	Rhost string `arg:"-r,--rhost" help:"Remote MySQL host to access"`
	Key   string `arg:"-k,--key" help:"Private SSH key to use for authentication"`
	User  string `arg:"-u,--user" help:"SSH user to use for authentication"`
}
