/*
	--------------------------------------------------------------------------------------------------------------------
														INTERFACE
	--------------------------------------------------------------------------------------------------------------------													
	an interface is a type that specifies a set of method signatures. 
	It defines a contract for what methods a type must have, but it doesn't provide 
	the implementation of those methods.
	Instead, it leaves the implementation to the types that satisfy the interface.

	--------------------------------------------------------------------------------------------------------------------
													    Embedding

	Struct Embedding	  : In Go, a struct can include another struct, and the fields and methods of the embedded struct become 
						  part of the outer struct.

	Interface Embedding : Similarly, interfaces can be embedded within other interfaces, and a type that 
						  implements the outer interface also implicitly implements the embedded interfaces.

	--------------------------------------------------------------------------------------------------------------------
*/