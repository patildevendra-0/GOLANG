/*
	--------------------------------------------------------------------------------------------------------------------
														INTERFACE
	--------------------------------------------------------------------------------------------------------------------													
	an interface is a type that specifies a set of method signatures. 
	It defines a contract for what methods a type must have, but it doesn't provide 
	the implementation of those methods.
	Instead, it leaves the implementation to the types that satisfy the interface.


	#   In Go, we don't explicitly state that a type implements an interface............imp ///////////////#########
		If a type provides implementations for all the methods declared by an interface, it is considered to implement that interface.

	--------------------------------------------------------------------------------------------------------------------
													    Embedding

	Struct Embedding	  : In Go, a struct can include another struct, and the fields and methods of the embedded struct become 
						  part of the outer struct.

	Interface Embedding : Similarly, interfaces can be embedded within other interfaces, and a type that 
						  implements the outer interface also implicitly implements the embedded interfaces.


	--------------------------------------------------------------------------------------------------------------------
													Empty Interface
	
		An empty interface in Go is denoted by interface{}.
		It represents a type that can hold values of any type.

		Use Cases:

				It's commonly used when you want to write functions or structures that can work with values of any type.
				It is often used in situations where you need to handle different types dynamically.

		Real-Life Analogy:

				Think of an empty box that can hold anything - books, toys, or even food. Similarly, an empty interface can hold values of any type.

	--------------------------------------------------------------------------------------------------------------------
													Type Assertion
	
		Type assertion is a way to find out and use the actual type of a value inside an interface.
		It helps you deal with different types stored in an empty interface.
		
	Real-Life Analogy:

		Imagine you have a mystery box. Type assertion is like opening the box, checking what's inside, 
		and then using it accordingly.



	Summary:

			Empty interfaces let you work with values of any type.
			Type assertions help you figure out and use the specific type inside an interface.		
	--------------------------------------------------------------------------------------------------------------------
*/