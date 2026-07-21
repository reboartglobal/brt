# TODO: Perbaikan RGN Website untuk Publish

## ✅ Tahap 1 — Perbaikan Build Critical (SELESAI)
- [x] Hapus setting `theme = "rgn-theme"` (tidak diperlukan, layout sudah di root)
- [x] Hapus folder `themes/rgn-theme/`
- [x] Buat folder `assets/css/` dan pindahkan CSS dari `statics/css/style.css`
- [x] Rapikan duplikasi CSS (gabung jadi satu set yang bersih)
- [x] Hapus folder `statics/` yang salah nama (Hugo hanya kenali `static/`)
- [x] Buat folder `static/` untuk favicon
- [x] Perbaiki Go API files:
  - [x] `_config.go`: tambah import `"os"`
  - [x] `health.go`: rename `Handler` → `HealthHandler`, hapus `getEnv` duplikat
  - [x] `contact.go`: rename `Handler` → `ContactHandler`
  - [x] `info.go`: rename `Handler` → `InfoHandler`
  - [x] `packages.go`: rename `Handler` → `PackagesHandler`

## ✅ Tahap 2 — Perbaikan Konfigurasi & Konten (SELESAI)
- [x] Perbaiki footer: "Netlify" → "Vercel"
- [x] Buat favicon.ico di `static/`
- [x] Perbaiki `_redirects` (sebelumnya salah arah ke `/en/404.html`)
- [x] Perbaiki inkonsistensi harga paket di `_index.md`
- [x] `vercel.json` sudah benar (outputDirectory: "public")

## ✅ Tahap 3 — Finalisasi (SELESAI)
- [x] Validasi struktur project
- [x] Struktur sudah sesuai: `assets/`, `static/`, `layouts/`, `content/`, `api/`
- [x] CSS bisa di-load via Hugo Pipes (`resources.Get "css/style.css"`)
- [x] Go API siap di-compile di Vercel
- [x] Build command sudah sesuai

