package main

import (
	"os"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestGetTags(t *testing.T) {
	f, _ := os.Open("./stub/cnn.html")
	defer f.Close()

	parsed, err := html.Parse(f)

	if err != nil {
		t.Error(err)
	}

	tags := GetTags(parsed, "article")

	if len(tags) == 0 {
		t.Errorf("Expected tags to be more than 0")
	}
}

func TestGetAttribute(t *testing.T) {
	f, _ := os.Open("./stub/cnn.html")
	defer f.Close()

	parsed, err := html.Parse(f)

	if err != nil {
		t.Error(err)
	}

	tag := GetTags(parsed, "a")[0]
	attr, ok := GetAttribute(&tag, "href")

	if !ok {
		t.Errorf("Expected true, but got %v", ok)
	}

	if attr.Val == "" {
		t.Errorf("Expected to be not empty")
	}
}

func TestGetText(t *testing.T) {
	parsed, _ := html.Parse(strings.NewReader(`<a dtr-evt="newsfeed" dtr-sec="" dtr-act="artikel" onclick="_pt(this)" dtr-idx="12" dtr-id="7759540" dtr-ttl="Warga Khawatir Taman Tempat Nongkrong Mengarah Kriminal Jika Buka 24 Jam" href="https://news.detik.com/berita/d-7759540/warga-khawatir-taman-tempat-nongkrong-mengarah-kriminal-jika-buka-24-jam" class="media__link">
                    <div class="replace_title">
                        Warga Khawatir Taman Tempat Nongkrong Mengarah Kriminal Jika Buka 24 Jam                    </div>
                </a>`))
	text, _ := GetText(parsed)

	if text != "Warga Khawatir Taman Tempat Nongkrong Mengarah Kriminal Jika Buka 24 Jam" {
		t.Errorf("Expected: `Warga Khawatir Taman Tempat Nongkrong Mengarah Kriminal Jika Buka 24 Jam`, but got: %s", text)
	}
}

func TestGetTextWithH2(t *testing.T) {
	parsed, _ := html.Parse(strings.NewReader(`<a aria-label="link description" href="https://www.cnnindonesia.com/gaya-hidup/20250131130848-284-1193114/jangan-ditanam-5-tanaman-ini-bisa-mengundang-tikus-masuk-ke-rumah" dtr-evt="box terpopuler" dtr-sec="terpopuler" dtr-act="terpopuler" onclick="_pt(this)" dtr-idx="2" dtr-id="1193114" dtr-ttl="Jangan Ditanam, 5 Tanaman Ini Bisa Mengundang Tikus Masuk ke Rumah" class="group">
                            <span class="text-cnn_grey text-lg absolute left-0"> 02 </span>
                            <h2 class="text-base text-cnn_black_light group-hover:text-cnn_red">Jangan Ditanam, 5 Tanaman Ini Bisa Mengundang Tikus Masuk ke Rumah</h2>
                            <span class="text-xs text-cnn_red">Gaya Hidup</span>
                        </a>`))
	text, _ := GetTextFromChilds(parsed, "h2")

	if text != "Jangan Ditanam, 5 Tanaman Ini Bisa Mengundang Tikus Masuk ke Rumah" {
		t.Errorf("Expected: `Jangan Ditanam, 5 Tanaman Ini Bisa Mengundang Tikus Masuk ke Rumah`, but got: %s", text)
	}
}
