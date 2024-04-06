package main

type Object struct {
	obid string
	group string
}

type start struct {

}

type __Frame__ struct {
	Object
	start    string `fid:"StartFrame"`
	main     string `fid:"MainFrame"`
	detail   string `fid:"DetailFrame"`
	wallet   string `fid:"WalletFrame"`
	settings string `fid:"SettingsFrame"`
	info     string `fid:"InfoFrame"`
	profile  string `fid:"ProfileFrame"`
	licenses string `fid:"LicensesFrame"`
}

var objects = make(map[string]Object) //for test only

func fglob(id string) *Object {
	if obj, ok := objects[id]; ok {
		return &obj
	}
	return nil
}

func main() {}
