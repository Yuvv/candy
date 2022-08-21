package function

// IFunction represents a function that accepts one argument and produces a result.
type IFunction[T, R any] interface {
	// Apply applies this function to the given argument.
	// Params:
	//   t – the function argument
	// Returns:
	//   the function result
	Apply(t T) R
}

// Function is func type of IFunction
type Function[T, R any] func(t T) R

// IBiFunction represents a function that accepts two arguments and produces a result.
// This is the two-arity specialization of IFunction.
type IBiFunction[T, U, R any] interface {
	// Apply applies this function to the given arguments.
	Apply(t T, u U) R
}

// BiFunction is func type of IBiFunction
type BiFunction[T, U, R any] func(t T, u U) R

// IBinaryOperator represents an operation upon two operands of the same type,
// producing a result of the same type as the operands.
// This is a specialization of IBiFunction for the case where the operands
// and the result are all the same type.
type IBinaryOperator[T any] interface {
	IBiFunction[T, T, T]
}

// BinaryOperator is func type of IBinaryOperator
type BinaryOperator[T any] func(t, u T) T

// ISupplier represents a supplier of results.
// There is no requirement that a new or distinct result be returned each
// time the supplier is invoked.
type ISupplier[T any] interface {
	// Get is used to get a result
	Get() T
}

// Supplier is func type of ISupplier
type Supplier[T any] func() T

// IPredicate represents a predicate (boolean-valued function) of one argument.
type IPredicate[T any] interface {
	// Test evaluates this predicate on the given argument.
	// Returns true if the input argument matches the predicate, otherwise false
	Test(t T) bool
}

// Predicate is func type of IPredicate
type Predicate[T any] func(t T) bool

// IConsumer represents an operation that accepts a single input argument and
// returns no result. Unlike most other functional interfaces, IConsumer is expected
// to operate via side-effects.
type IConsumer[T any] interface {
	// Accept performs this operation on the given argument.
	Accept(t T)
}

// Consumer is func type of IConsumer
type Consumer[T any] func(t T)

// IBiConsumer Represents an operation that accepts two input arguments and returns
// no result. This is the two-arity specialization of IConsumer. Unlike most other
// functional interfaces, IBiConsumer is expected to operate via side-effects.
type IBiConsumer[T, U any] interface {
	// Accept performs this operation on the given arguments.
	Accept(t T, u U)
}

// BiConsumer is func type of IBiConsumer
type BiConsumer[T, U any] func(t T, u U)
