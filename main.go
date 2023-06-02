package main

import (
	basic "./basic"
	basic_flow "./basic_flow_control"
	basic_structers "./basic_structures"
)

func main() {
	basic.Packages()
	basic.Functions()
	basic.Variables_with_initializers()
	basic.Zero_Values()
	basic.Type_Conversions()
	basic.Type_Inference()
	basic.Constants()
	basic.Type_Conversions()
	basic_flow.For()
	basic_flow.If()
	basic_flow.Switch()
	basic_flow.Switch_True()
	basic_flow.Defer()
	basic_flow.Stacking_Defers()
	basic_structers.Struct()
	basic_structers.Slices()
	basic_structers.Pointers()
	basic_structers.Arrays()
	basic_structers.Slice_len_cap()
	basic_structers.Slice_literals()
	basic_structers.Range()
	basic_structers.Maps()
	basic_structers.Mutating_Maps()
	basic_structers.Function_Values()
	basic_structers.Function_Closures()
}
