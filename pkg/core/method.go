package core

import "context"

var _ Map[any, any] = NilMap

type (
	Method[ARG, RESULT any] func(context.Context, ARG) (RESULT, error)
	Map[ARG, RESULT any]    func(ARG) RESULT
)

func (m Method[ARG, RESULT]) Call(ctx context.Context, arg ARG) (RESULT, error) {
	if m == nil {
		var blank RESULT

		return blank, nil
	}

	return m(ctx, arg)
}

func (f Map[ARG, RESULT]) Map(arg ARG) RESULT {
	if f == nil {
		var blank RESULT

		return blank
	}

	return f(arg)
}

func MapMethodParameters[AI, AO, RI, RO any](m Method[AI, RI], mapArg Map[AO, AI], mapResult Map[RI, RO]) Method[AO, RO] {
	return func(ctx context.Context, to AO) (RO, error) {
		result, err := m.Call(ctx, mapArg.Map(to))
		if err != nil {
			var blank RO

			return blank, err
		}

		return mapResult.Map(result), nil
	}
}

func NilMap(any) any {
	return nil
}
