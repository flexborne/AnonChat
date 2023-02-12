package main

type Optional struct {
	Value any
	Valid bool
}

func NewOptional(value any) Optional {
	return Optional{value, true}
}

func (o Optional) HasValue() bool {
	return o.Valid
}

func (o Optional) ValueOr(fallback any) any {
	if o.Valid {
		return o.Value
	}

	return fallback
}

func (o Optional) Get() any {
	if o.Valid {
		return o.Value
	}
	return nil
}
