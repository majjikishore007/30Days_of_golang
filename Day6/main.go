package main

type Dictionary map[string]string

const (
    ErrNotFound   = DictionaryErr("could not find the word you were looking for")
    ErrWordExists = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
    return string(e)
}

func (d Dictionary) Search(key string) (string, error)   {
	value,null:=d[key]
	if !null{
		return "",ErrNotFound
	}
	return value,nil
}

func (d Dictionary) Add(key,value string) error {
	d[key]=value
	return nil
}

func (d Dictionary) Update(key,newValue string) error  {
	  _, err := d.Search(key)

    switch err {
    case ErrNotFound:
        return ErrWordDoesNotExist
    case nil:
        d[key] = newValue
    default:
        return err
    }

    return nil
}

func (d Dictionary) Delete(key string)  {
	delete(d,key)
}