package main

//Bingo containes information about a single serial killer trait
type Bingo struct {
	ID int `json:"id" binding:"required"`
	Marked bool `json:"marked"`
	Trait string `json:"trait" binding:"required"`
}

//List of traits
var Bingos = []Bingo{
	Bingo{1, false, "Drives a van"},
	Bingo{1, false, "Arson"},
	Bingo{1, false, "Childhood Head Injury"},
	Bingo{1, false, "Abandonment"},
	Bingo{1, false, "Mother a prostitute"},
	Bingo{1, false, "Bunker"},
	Bingo{1, false, "Someone is in foster care"},
	Bingo{1, false, "Cannibalism"},
	Bingo{1, false, "Caucasian"},
	Bingo{1, false, "Tortured or killed pets",
	Bingo{1, false, "Weak or absent father"},
	Bingo{1, false, "FREE or Prison Break"},
	Bingo{1, false, "Someone has a missing limb"},
	Bingo{1, false, "Bedwetting"},
	Bingo{1, false, "Sexual Deviance"},
	Bingo{1, false, "Three first names"},
	Bingo{1, false, "Dancing"},
	Bingo{1, false, "Religious fanatics"},
	Bingo{1, false, "Oddly specific preferences"},
	Bingo{1, false, "Broom as a weapon"},
	Bingo{1, false, "Two in one night"},
	Bingo{1, false, "Mommy issues"},
	Bingo{1, false, "High IQ"},
	Bingo{1, false, "Catchy nickname"},
	}