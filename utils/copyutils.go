package utils

import "github.com/jinzhu/copier"

func DeepCopy[T any](to *T, from interface{}) (*T, error) {
	res := new(T)
	if err := copier.CopyWithOption(res, to, copier.Option{
		DeepCopy:    true,
		IgnoreEmpty: true,
	}); err != nil {
		return nil, err
	}

	if err := copier.CopyWithOption(res, from, copier.Option{
		DeepCopy:    true,
		IgnoreEmpty: true,
	}); err != nil {
		return nil, err
	}

	return res, nil
}
