package c4

var initialBoard = [][]Colour{
	{Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}},
	{Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}},
	{Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}},
	{Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}},
	{Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}},
	{Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}},
}

var colWinBoard = [][]Colour{
	{Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{"yellow"}},
	{Colour{""}, Colour{""}, Colour{""}, Colour{"yellow"}, Colour{""}, Colour{""}, Colour{"red"}},
	{Colour{""}, Colour{""}, Colour{"yellow"}, Colour{"red"}, Colour{""}, Colour{""}, Colour{"red"}},
	{Colour{""}, Colour{""}, Colour{"yellow"}, Colour{"yellow"}, Colour{""}, Colour{""}, Colour{"red"}},
	{Colour{""}, Colour{"yellow"}, Colour{"yellow"}, Colour{"red"}, Colour{""}, Colour{""}, Colour{"yellow"}},
	{Colour{"red"}, Colour{"red"}, Colour{"yellow"}, Colour{"yellow"}, Colour{""}, Colour{"yellow"}, Colour{"yellow"}},
}

var rowWinBoard = [][]Colour{
	{Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{"yellow"}},
	{Colour{""}, Colour{""}, Colour{""}, Colour{"yellow"}, Colour{""}, Colour{""}, Colour{"yellow"}},
	{Colour{""}, Colour{""}, Colour{"yellow"}, Colour{"red"}, Colour{""}, Colour{""}, Colour{"red"}},
	{Colour{""}, Colour{""}, Colour{"red"}, Colour{"yellow"}, Colour{""}, Colour{""}, Colour{"red"}},
	{Colour{""}, Colour{"yellow"}, Colour{"yellow"}, Colour{"red"}, Colour{"red"}, Colour{"red"}, Colour{"red"}},
	{Colour{"red"}, Colour{"red"}, Colour{"yellow"}, Colour{"yellow"}, Colour{"red"}, Colour{"yellow"}, Colour{"yellow"}},
}

var diagwinBLTRBoard = [][]Colour{
	{Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{"yellow"}},
	{Colour{""}, Colour{""}, Colour{""}, Colour{"yellow"}, Colour{""}, Colour{""}, Colour{"yellow"}},
	{Colour{""}, Colour{""}, Colour{"yellow"}, Colour{"red"}, Colour{""}, Colour{""}, Colour{"red"}},
	{Colour{""}, Colour{""}, Colour{"red"}, Colour{"yellow"}, Colour{""}, Colour{"red"}, Colour{"red"}},
	{Colour{""}, Colour{"yellow"}, Colour{"yellow"}, Colour{"red"}, Colour{"red"}, Colour{"yellow"}, Colour{"red"}},
	{Colour{"red"}, Colour{"red"}, Colour{"yellow"}, Colour{"red"}, Colour{"red"}, Colour{"yellow"}, Colour{"yellow"}},
}

var diagwinTLBRBoard = [][]Colour{
	{Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{""}, Colour{"yellow"}},
	{Colour{""}, Colour{""}, Colour{""}, Colour{"yellow"}, Colour{""}, Colour{""}, Colour{"yellow"}},
	{Colour{""}, Colour{""}, Colour{"yellow"}, Colour{"yellow"}, Colour{""}, Colour{""}, Colour{"red"}},
	{Colour{""}, Colour{""}, Colour{"red"}, Colour{"yellow"}, Colour{"yellow"}, Colour{"yellow"}, Colour{"red"}},
	{Colour{""}, Colour{"yellow"}, Colour{"yellow"}, Colour{"red"}, Colour{"red"}, Colour{"yellow"}, Colour{"red"}},
	{Colour{"red"}, Colour{"red"}, Colour{"yellow"}, Colour{"red"}, Colour{"red"}, Colour{"yellow"}, Colour{"yellow"}},
}
