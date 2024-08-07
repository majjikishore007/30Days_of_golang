package main

const (
    ErrNotFound   = DictionaryErr("could not find the word you were looking for")
    ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
    return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string,error) {
	res, ok := d[key]
	if !ok{
		return "",DictionaryErr(ErrNotFound)
	}
	return res,nil
}

func (d Dictionary) Add(key,value string) error {
	_,err:= d.Search(key)
	switch err{
	case DictionaryErr(ErrNotFound):
		d[key]=value
	case nil:
		return DictionaryErr(ErrWordExists)
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key,value string)  {
	d[key]=value
}

func main()  {
	
}