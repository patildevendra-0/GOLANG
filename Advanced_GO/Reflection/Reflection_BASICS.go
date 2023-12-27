

/*
	--------------------------------------------------------------------------------------------------------------------
													REFLECTION 
	--------------------------------------------------------------------------------------------------------------------

		Reflection in Go refers to the ability of a program to examine its own structure and gain information about its types 
		during runtime. It allows you to inspect variables, functions, and even modify their values dynamically. 
		Reflection is a powerful feature, but it should be used judiciously as it may lead to less type-safe and more error-prone code.

	Real-Life Analogy:

		Imagine you have a mysterious box, and you want to know what's inside it. 
		You can't directly see or touch the contents, but you have a set of tools that help you figure out what's in there.

	--------------------------------------------------------------------------------------------------------------------

	What is Reflection in Simple Terms?

		Reflection is like looking at something and figuring out what it is, kind of like looking at a box and knowing what's inside.
	--------------------------------------------------------------------------------------------------------------------

	Laws of Reflection:

			1.Imagine you have a box (a variable). The first law says you can look at the box and find out what's inside.
			2.Now, imagine you know what's inside the box (you have information about a variable), the second law says you can actually use that information to do things.
			3.Let's say you want to change something inside the box using reflection. The third law says you need to make sure you can actually change it.



	example :

			Data := 21

			Data_Type := reflect.TypeOf(Data)      .................. int

			--------------------------------------------------------------------------------------

			Reflection_obj := reflect.Valueof(Data)
			Actual_Data := Reflection_obj.Interface() .............. 21

			-------------------------------------------------------------------------------------

			Reflection_obj := reflect.ValueOf(&Data).Elem()
			Reflection_obj.SetInt(51)               ---------------- set Data 21 to 51

	--------------------------------------------------------------------------------------------------------------------
*/	