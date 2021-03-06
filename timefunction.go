//
// TITLE           : Translate str date format to str date froamt2
// AUTHOR          : GitArt
// DATE & TIME     : 10.05.2016 14:39
// DESCRIPTION     : Возврат двух дат с разницей в часах между текущим временем и вычисляемым
//                   Дата возвращается в строковом виде
//                   На вход подается параметр в часах (строковой тип)
//                   на выходе два параметра текущая дата и дата на несколько часов назад (входной параметр)
//

// *****************************************************************************************************************************
// TITLE           : Period In Hours
// AUTHOR          : Савченко Артур
// DATE & TIME     : 10.05.2016 14:39
// DESCRIPTION     : Возврат двух дат с разницей в часах между текущим временем и вычисляемым
//                   Дата возвращается в строковом виде
//                   На вход подается параметр в часах (строковой тип)
//                   на выходе два параметра текущая дата и дата на несколько часов назад (входной параметр)
// *****************************************************************************************************************************
func PHrs(Hrs string) (string, string) {

	F := "2006-01-02T15:04:05"
	R, _ := strconv.Atoi(Hrs)           // Преобразование  стринга в инт
	V := time.Duration(R * -1)          // Обратный отсчет в часах
	T := time.Now()                     // Текущее время
	S := T.Add(time.Hour * V).Format(F) // Старт
	E := T.Format(F)                    // Финиш

	if Hrs == "" || Hrs == "0" {
		S = E
	}
	return S, E
}

func FormHrs(Hrs string) string {
	F := "02-01-2006 15:04"
	S, _ := time.Parse(F, Hrs)
	return S.String()
}
