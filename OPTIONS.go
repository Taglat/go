
func Options(args []string) {
	const validOptions = "abcdefghijklmnopqrstuvwxyz"
	var optionBits [4]byte

	for _, arg := range args {
		if len(arg) < 2 || arg[0] != '-' {
			fmt.Println("Invalid Option")
			return
		}

		if Contains(arg, "h") {
			fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
			return
		}

		for _, char := range arg[1:] { // a b c 
			index := IndexRune(validOptions, char) // a => 0 b => 1 c => 2
			if index >= 0 {
				byteIndex := index / 8   // 0 / 8 => 0 | 1 / 8 => 0 | 2 / 8 => 0           // Определяем, к какому байту относится бит
				bitPosition := index % 8 // 0 % 8 => 0 | 1 % 8 => 1 | 2 % 8 => 2       // Определяем позицию внутри байта

				optionBits[byteIndex] |= 1 << bitPosition // optionBits[0] = 00000001, 00000010, 00000100 => 00000111
			} else {
				fmt.Println("Invalid Option")
				return
			}
		}
	}

	for i := len(optionBits) - 1; i >= 0; i-- { // i = 4 - 1
		for j := 7; j >= 0; j-- { // j = 7
			if optionBits[i]&(1<<j) != 0 { // optionBits[3] & 00000001 
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}
		}
		if i > 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
