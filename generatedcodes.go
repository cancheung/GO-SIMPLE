/*************************************************************************************************************************************************
 * G E N E R A T E D  K E Y S  A N D  U I D
 * Гереация ключа
 * randStr(16, "alphanum")
 * randStr(16, "alpha")
 * randStr(12, "number")
 *************************************************************************************************************************************************/
func Keygen(strSize int, randType string) string {
	var dictionary string
	if randType == "alphanumsmall" { dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"	}
	if randType == "alphanum"      { dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"}
	if randType == "alpha"         { dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"}
	if randType == "number"        { dictionary = "0123456789"}

	var bytes = make([]byte, strSize)
	rand.Read(bytes)

	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

/******************************************************************************************************************************************************
 * Генерация уникального ключа
 * Format : MM-DD-YYYY-UID
 * Gen_corppoint("C3","C03","RTYY")
 ******************************************************************************************************************************************************/
func Gen_corppoint(Corp, Net, Point string) string {
	Tm := Corp + ":" + Net + ":" + Point + ":" + Sys_keyid()
	return Tm
}

/******************************************************************************************************************************************************
 * Генерация уникального ключа
 * Format : 1234567
 ******************************************************************************************************************************************************/
func Gen_unix() int64 {
	Tm := time.Now().Unix() //UnixNano()
	return Tm
}

/******************************************************************************************************************************************************
 * Генерация уникального ключа
 * Format : MM-DD-YYYY-UID
 ******************************************************************************************************************************************************/
func KEY() string {
	Tm := time.Now().Format("01.02.2006.150405.999999") // Tm := time.Now().Format("01-02-2006-15-04-05.999999")
	GuidStr := Tm + "-" + Sys_Round()                   // GENID()  Gd := os.Getuid()
	return GuidStr
}

/******************************************************************************************************************************************************
 * Генерация уникального ключа
 * Format: MM-DD-YYYY-UID
 ******************************************************************************************************************************************************/
func SKEY() string {
	format := "01-02-2006-15-04-05.999999"
	format =  "01-022006150405.999999"
	Tm := time.Now().Format(format)
	Tm = strings.Replace(Tm, ".", "-", 1)
	return Tm
}

/******************************************************************************************************************************************************
 * Генерация уникального кода вида : pGCsYA==
 ******************************************************************************************************************************************************/
func Sys_Round() string {
	size := 5
	rb := make([]byte, size)
	_, err := rand.Read(rb)

	// Errorapi/dat
	if err != nil {
		return ""
	}

	rs := base64.URLEncoding.EncodeToString(rb)
	return rs
}

/******************************************************************************************************************************************************
 * Формирование UID (6 знаков)
 *
 * newUUID generates a random UUID according to RFC 4122
 * uuid, err := newUUID()
 * 	if err != nil {
 * 		fmt.Printf("error: %v\n", err)
 * 	}
 * 	fmt.Printf("%s\n", uuid)
 ******************************************************************************************************************************************************/
func GENID() string {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)

	if n != len(uuid) || err != nil {
		return " "
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80 // variant bits; see section 4.1.1
	uuid[6] = uuid[6]&^0xf0 | 0x40 // version 4 (pseudo-random); see section 4.1.3

	// return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
	return fmt.Sprintf("-%x", uuid[0:10]) // return fmt.Sprintf("%x", uuid[0:6]), nil
}

/*********************************************************************************************
 *	 Код сформирован случайным образом
 *	 на основе GUID и приведен к верхнему регистру
 *********************************************************************************************/
func Sys_keyid() string {
	 s := strings.ToUpper(strings.Replace(GENID(), "-", "", 1))
	 return s
}

/************************************************************************************************************************************************
 * Гереация ключа
 * randStr(16, "alphanum")
 * randStr(16, "alpha")
 * randStr(12, "number")
 ************************************************************************************************************************************************/
func randStr(strSize int, randType string) string {
	var dictionary string
	if randType == "alphanum" {	dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"}
	if randType == "alpha"    {	dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"	}
	if randType == "number"   {	dictionary = "0123456789"}
	
	var bytes = make([]byte, strSize)
	rand.Read(bytes)

	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

/******************************************************************************************************
 * TITLE  :  Расчет кодирование SHA1 Hash
 * REMARK :
 ******************************************************************************************************/
func Sys_sha(Salt, Body string) string {

	if Salt == "" || Body == "" {
	   return "SHA Error."
	}

	b := Salt + Body
	h := sha1.New()
	h.Write([]byte(b))
	s := hex.EncodeToString(h.Sum(nil))
	return s
}

/******************************************************************************************************
 * TITLE  :  Расчет кодирование SHA1 Hash
 * REMARK :
 ******************************************************************************************************/
func Sys_sha256(Body, Salt string) string {
	b := Salt + Body
	h := sha256.New()
	h.Write([]byte(b))
	s := hex.EncodeToString(h.Sum(nil))
	return s
}

/***********************************************************************************************************************************
 * TITLE  :  Расчет кодирование SHA1 Hash
 * REMARK :
 ***********************************************************************************************************************************/
func Sys_Encode(Body string) string {
	s := hex.EncodeToString([]byte(Body))
	return s
}

/**************************************************************************************************************************************
 * Конвертация Int64 to Str
 ******************************************************************************************************************************************************/
func InttoStr(Ints int) string {
	//str := strconv.FormatInt(Intt64, 10)      // Выдает конвертацию 2000-wqut
	//str := strconv.Itoa64(Int64)              // use base 10 for sanity purpose
	str := fmt.Sprintf("%d", Ints)
	return str
}

/******************************************************************************************************************************************************
 * Конвертация Int64 to Str
 ******************************************************************************************************************************************************/
func Int64toStr(Int64 int64) string {
	//str := strconv.FormatInt(Intt64, 10)      // Выдает конвертацию 2000-wqut
	//str := strconv.Itoa64(Int64)              // use base 10 for sanity purpose
	str := fmt.Sprintf("%d", Int64)
	return str
}

/******************************************************************************************************************************************************
 * Конвертация Float64 to Str (с 6 десятичными )
 ******************************************************************************************************************************************************/
func FloatToString(input_num float64) string {
	 return strconv.FormatFloat(input_num, 'f', 6, 64)
}

/******************************************************************************************************************************************************
 * Конвертация Float64 to Str без десятичных
 ******************************************************************************************************************************************************/
func FloatToStr(input_num float64) string {
	 return strconv.FormatFloat(input_num, 'f', 0, 64)
}
