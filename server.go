package main /


import (
	"fmt"       
	"io/ioutil" 
	"log"       
	"math"      
	"net/http"
	"os"     
	"sort"    
	"strconv" 
	"strings" 
	"time"
)


func main() {
	fmt.Println("Go Eğtimine Hoş Geldiniz") 
	const pi float64 = 3.14159265           
	const egitimTarihi int = 2021
	var yas int = 14
	var boy float32 = 1.74
	var guzelMi bool = true
	var isim string = "Jennifer Ariston"

	dogumYili := egitimTarihi - yas 

	fmt.Println("PI Sayısı:   ", pi)
	
	fmt.Printf("Öğrenci adı: %s\nÖğrenci yaşı: %d\nDoğum Yılı: %d\nÖğrenci Boyu: %.2f\nGüzel mi?%t\n", isim, yas, dogumYili, boy, guzelMi)



	fmt.Printf("%d \n", yas)       
	fmt.Printf("%b \n", dogumYili) 
	fmt.Printf("%c \n", yas)       
	fmt.Printf("%x \n", yas)       
	fmt.Printf("%o \n", yas)      
	fmt.Printf("%e \n", pi)        

	
	ornekString := "Merhaba, Dunya!"
	virgulluString := "Metin,Ali,Feyyaz"

	fmt.Println(strings.Contains(ornekString, "Me"))      
	fmt.Println(strings.Index(ornekString, "D"))         
	fmt.Println(strings.Count(ornekString, "a"))           
	fmt.Println(strings.Replace(ornekString, "u", "ü", 1))
	fmt.Println(strings.Split(virgulluString, ","))        
	
	karakterListesi := []string{"z", "y", "x", "w", "v", "u", "t", "s", "r", "q", "p", "o", "m", "l", "k", "j", "i", "h", "g", "f", "e", "d", "c", "b", "a"}
	sort.Strings(karakterListesi)
	fmt.Println("karakterListesi:", karakterListesi)

	
	rastgeleBirTamsayi := 3
	rastgeleBirFloat := 3.14
	birBaskaString := "!000000"
	birBaskaStringDaha := "1.37"

	tamSayiydiFloatOldu := float64(rastgeleBirTamsayi)
	floatIkenTamsayiOldu := int(rastgeleBirFloat)

	
	yeniBirTamsayi, _ := strconv.ParseInt(birBaskaString, 0, 64)
	yeniBirFloat, _ := strconv.ParseFloat(birBaskaStringDaha, 64)

		_, _, _, _ = tamSayiydiFloatOldu, floatIkenTamsayiOldu, yeniBirTamsayi, yeniBirFloat 
	


	sayac := 0

	fmt.Println("While alternatifi olarak for döngüsü: ")

	for sayac <= 10 {
		fmt.Println(sayac)
		sayac++ 
	}

	fmt.Println("Tipik for döngüsü: ")

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	var yasBilgisi int 

	fmt.Print("Yaşınızı Giriniz: ")

	fmt.Scanf("%d", &yasBilgisi) 
	fmt.Printf("Yaşınız %d olduğuna göre:\n\n", yasBilgisi)

	// If ile akış kontrolu
	if yasBilgisi > 16 && yasBilgisi < 18 {
		fmt.Println("Geçici ehliyet alabilirsiniz ve alkol satılan mekanlara giremezsiniz")
	} else if yasBilgisi >= 18 && yasBilgisi <= 65 {
		fmt.Println("Hem ehliyet alabilir hem alkollü mekanlarda bulanabilirsiniz ama lütfen alkollüyken araç kullanmayın!")
	} else if yasBilgisi > 65 {
		fmt.Println("Ehliyet için her yıl sağlık kontorlünden geçmelisiniz.")
	} else {
		fmt.Println("Hayatın en eğlenceli ve tasasız zamanları :)")
	}

	// swtich/case ile akış kontrolu
	var notBilgisi int

	fmt.Print("Lütfen sınav notunuzu giriniz (0-5): ")
	fmt.Scanf("%d", &notBilgisi)

	fmt.Println("Notunuz:", notBilgisi)

	switch notBilgisi {
	case 0:
		fmt.Println("Kaldın! :(")
	case 1:
		fmt.Println("Pek iyi görünmüyor, çok çalışman gerekecek!")
	case 2:
		fmt.Println("Daha çok çalışmalısın!")
	case 3:
		fmt.Println("Gelecek sefere daha iyisini yaparsın")
	case 4:
		fmt.Println("Başarılı")
	case 5:
		fmt.Println("Çok Başarılı")
	default:
		fmt.Println("Farklı bir not sistemine göre giriş yaptın gibi görünüyor!")
	}

	// Diziler aynı türdeki verilerden oluşan değerlerin gruplardır
	rakamlar := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// for iterasyon için de kullanılabilir range ile birlikte
	fmt.Println("for ile iterasyon")

	for _, v := range rakamlar {
		fmt.Println(v)
	}

	diziler := [4]string{"Doctor Who", "American Gods", "Game Of Thrones", "Orange is new black"}

	for i, v := range diziler {
		fmt.Printf("%d: %s\n", i, v)
	}


	rakamDilimi := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0} // [] dizlerden farklı olarak boş bırakılıyor
	rakamDilimi = append(rakamDilimi, -1, -2, -3)      // dilimimize itemler ekledik

	fmt.Println("rakamDilimi:", rakamDilimi)

	
	rakamDilimi = sil(rakamDilimi, len(rakamDilimi)-1) // sondan bir element silen fonksiyonumuzdan dönen değeri atıyoruz
	rakamDilimi = sil(rakamDilimi, len(rakamDilimi)-1) // böylece dilim bir element kısalıyor
	rakamDilimi = sil(rakamDilimi, len(rakamDilimi)-1) // her seferinde

	fmt.Println("rakamDilimi:", rakamDilimi)

	yarimDilim := rakamDilimi[:5] // Bir endeks aralığını alırken ilk ya da son endeksi belirtmek opsiyoneldir.

	for endeks, deger := range yarimDilim { 
		fmt.Printf("Endeks: %d = Değer: %d\n", endeks, deger)
	}

	// make ile dilim oluşturmak
	birDilimDaha := make([]int, 3)

	fmt.Println("birDilimDaha make ile oluşturuldu:", birDilimDaha)

	// Python dilinde dictionary, Ruby dilinde hash olarak adlandırılan anahtar değer eşleşmeleri:
	// yani map veri tipi
	karne := make(map[string]int)
	karne["Matematik"] = 3
	karne["Fizik"] = 4
	karne["Edebiyat"] = 5

	for anahtar, deger := range karne {
		fmt.Printf("%s: %d\n", anahtar, deger)
	}

	karne["Biyoloji"] = 4 // Biyoloji anahtarı 4 tamsayısı ile eşleştirilerek eklendi(pair de diyebiliriz.)

	// Silme işleminden önce
	fmt.Println("Karne:", karne)

	delete(karne, "Biyoloji") // Biyoloji => 4 eşleşmesi map içinden silindi

	// Silme işleminden sonra
	fmt.Println("Karne", karne)

	// main içerisinde fonksiyon çağrısı
	rakamlarToplami := diziDegerleriniTopla(rakamlar) // Görüldüğü gibi bir fonksiyon çağrısın değişkene atayabiliriz

	fmt.Println("rakamlar isimli dizideki değerlerin toplamı:", rakamlarToplami)

	// Go Dilinde fonksiyonlar birden fazla değer dönebilir
	karesi, ikiKati := birdenFazlaDegerDon(4)
	fmt.Printf("Karesi: %d\nİki Katı: %d\n", karesi, ikiKati)

	// Go Dilinde bir fonksia geçeceğimiz argüman saysı belirsizse
	fmt.Println("Birden fazla sayıda argümanın birbiri ile çarpımı", birCokArgumanBirbiriyleCarp(1, 2, 3, 4, 5))

	// Go dilinde closure oluşturma
	birSayiDaha := 2

	katmerle := func() int {
		birSayiDaha *= 2
		return birSayiDaha
	}

	katmerle()
	katmerle()
	katmerle()

	fmt.Println("birSayiDaha:", birSayiDaha)

	// İç içe fonksiyonlar oluşturulabilir
	tamSayi := 7

	fmt.Printf("%d sayısının faktöriyeli %d sayısıdır.\n", tamSayi, faktoriyel(tamSayi))

	// defer ile bir kod grubunun en son çalışması sağlanabilir, hatalar yakalanabilir
	// Son giren ilk çıkar mantığı ile çalıştırılır birden fazla defer
	defer ikiYazdir() // en son çalıştırılacak
	birYazdir()

	// Bir sıfıra bölme hatası yakalayalım
	fmt.Println("Bölme işlemi: 3 / 0 =")
	guvenliBolme(3, 0)
	fmt.Println("Bölme işlemi: 8 / 2 =", guvenliBolme(8, 2))

	// PANIC ile hata yakalama
	panicYakala()

	// değişkenin değerini pointer kullanarak fonksiyon içinde değiştir
	degeriDegisecek := 9

	degeriniKaresiYap(&degeriDegisecek) //& ile bellekteki adresi geçtik

	fmt.Println("degeriDegisecek değişkeninin halihazırdaki değeri:", degeriDegisecek)

	// Kendi veri tiplerimizi struct tanımları içerisinde tasarlayabiliriz
	dGen1 := Dikdortgen{yukseklik: 2.0, en: 3.5}
	dGen1.setYukseklik(4.0)
	dGen1.setEn(7.0)
	cmbr := Cember{yaricap: 2.0}
	cmbr.setYaricap(2.4)
	fmt.Printf("Çember Yarıçap: %.4f\n", cmbr.yaricap)
	fmt.Printf("Dikdörtgen Yükseklik: %.4f\nDikdörtgen En: %.4f\n", dGen1.getYukseklik(), dGen1.getEn())
	fmt.Println("Çemberin Alanı:", alanHesapla(cmbr))
	fmt.Println("Dikdörtgenin Alanı:", alanHesapla(dGen1))

	// dosya işlemleri
	dosya, err := os.Create("ornek.txt")

	// Dosya oluşturulurken bir hata oluştuysa yakalayıp logladık
	if err != nil {
		log.Fatal(err)
	}

	dosya.WriteString("Dosyamıza string bi veri ekledik.")

	dosya.Close()

	// io/ioutil kütüphanesi ile dosya akışı okuma
	dosyaAkisi, err := ioutil.ReadFile("ornek.txt")

	if err != nil {
		fmt.Printf("%s\n", err)
		log.Fatal(err)
	}

	dosyayiOku := string(dosyaAkisi)
	fmt.Println("Okundu:", dosyayiOku)

	// Go Routine
	for k := 0; k < 10; k++ {
		go say(k)
	}

	time.Sleep(time.Millisecond * 10000)

	//CHANNEL
	stringCHAN := make(chan string)

	for i := 0; i < 5; i++ {
		go sandvicYap(stringCHAN)
		go mayonezEkle(stringCHAN)
		go hesabaEkle(stringCHAN)
		time.Sleep(time.Second * 2)
	}

	// Web Server
    
    fmt.Println("Web Sunucuyu kapatmak için: CTRL C")

	http.HandleFunc("/yonetici", isleyiciAdmin)
	http.HandleFunc("/", isleyici)

	http.ListenAndServe(":3333", nil) // localhost:3333
}

// INTERFACE, STRUCT. GETTER, SETTER

type Sekıl interface {
	alan() float64
}

type Dikdortgen struct {
	yukseklik float64
	en        float64
}

func (dgen Dikdortgen) getYukseklik() float64 {
	return dgen.yukseklik
}

func (dgen *Dikdortgen) setYukseklik(deger float64) {
	dgen.yukseklik = deger
}

func (dgen Dikdortgen) getEn() float64 {
	return dgen.en
}

func (dgen *Dikdortgen) setEn(deger float64) {
	dgen.en = deger
}

func (dgen Dikdortgen) alan() float64 {
	return dgen.yukseklik * dgen.en
}

type Cember struct {
	yaricap float64
}

func (c Cember) getYaricap() float64 {
	return c.yaricap
}

func (c *Cember) setYaricap(deger float64) {
	c.yaricap = deger
}

func (c Cember) alan() float64 {
	return math.Pi * math.Pow(c.yaricap, 2)
}

func alanHesapla(s Sekil) float64 {
	return s.alan()
}

// INTERFACE, STRUCT, GETTER, SETTER TANIMLARI BİTTİ

// FONKSIYON TANIMLARI
func diziDegerleriniTopla(rakamlar [10]int) int {
	toplam := 0

	for _, deger := range rakamlar {
		toplam += deger // toplam = toplam + deger ifadesinin eşleniği
	}
	return toplam
}

func birdenFazlaDegerDon(x int) (int, int) {
	return x * x, x * 2
}

func birCokArgumanBirbiriyleCarp(args ...int) int {
	carpim := 1

	for _, v := range args {
		carpim *= v
	}
	return carpim
}

// Rekürsif (iç içe) fonksiyon
func faktoriyel(n int) int {
	if n == 0 {
		return 1
	}
	return n * faktoriyel(n-1)
}

func birYazdir() {
	fmt.Println("Bir")
}

func ikiYazdir() {
	fmt.Println("İki")
}

func guvenliBolme(sayi1, sayi2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	cozum := sayi1 / sayi2
	return cozum
}

func panicYakala() {
	defer func() {
		fmt.Println(recover())
	}()

func degeriniKaresiYap(x *int) {
	*x *= *x // aadresteki değerleri alıp üzerinde işleme yaptık
}

func sil(s []int, i int) []int {
	s = append(s[:i], s[i+1:]...)
	return s
}

func say(i int) {
	for j := 0; j < 10; j++ {
		fmt.Println(i, ":", j)
		time.Sleep(time.Millisecond * 1000)
	}
}

// Channel
var sandvicSayisi = 0
var sandvicIsmi = ""

func sandvicYap(stringCHAN chan string) {
	sandvicSayisi++
	sandvicIsmi = "Sandviç No:" + strconv.Itoa(sandvicSayisi)
	fmt.Println(sandvicIsmi, "hazır, mayonez eklemeye gönder!")

	stringCHAN <- sandvicIsmi

	time.Sleep(time.Second * 1)
}

func mayonezEkle(stringCHAN chan string) {
	sandvic := stringCHAN

	fmt.Println("Mayonez eklendi, buyrun:", sandvic)

	stringCHAN <- sandvicIsmi
	time.Sleep(time.Second * 1)
}

func hesabaEkle(stringCHAN chan string) {
	sandvic := <-stringCHAN

	fmt.Println(sandvic, "hesaba eklendi! Müşteriye gönder!")

	time.Sleep(time.Second * 1)
}

//WEB SERVER 
func isleyici(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sunucuma Hoş Geldiniz!\n")
}

func isleyiciAdmin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Yönetici Sayfasına Hoş Geldiniz(ŞİFRE GEREKLİ)")
}