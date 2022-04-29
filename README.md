# Tugas Besar Strategi Algoritma 3

Author:
- Ignasius Ferry Priguna (13520126)
- Bayu Samudra (13520128)
- Muhammad Gilang Ramadhan (13520137)

# Deskripsi
Fungsi utama yang digunakan untuk mendukung fungsionalitas program ini terletak di backend, tepatnya di /controller/validation.go dan di /lib. File validation.go berisi fungsi FileValidation yang berfungsi untuk memastikan file hanya berisi huruf ACTG tanpa spasi menggunakan regex. Fungsi tersebut mengembalikan string yang kosong jika validasi gagal dan isi file jika berhasil serta boolean yang menyatakan berhasil atau gagalnya validasi.
File yang berada di dalam folder lib berisi algoritma pencarian pola dan fungsi-fungsi pembantunya serta fungsi pengecekan regex untuk input syarat pencarian. Fungsi BoyerMoore digunakan untuk melakukan pencarian suatu pola string dalam teks menggunakan algoritma BoyerMoore. Jika fungsi menemukan kecocokan, akan dikembalikan posisi ditemukannya kecocokan sedangkan jika tidak akan mengembalikan nilai -1. Fungsi BoyerMoore memanfaatkan fungsi pembantu BuildLast yang disimpan di util.go untuk menyimpan posisi kemunculan terakhir setiap alfabet dalam pola. Struktur data yang digunakan untuk fungsi pembantu ini adalah map.
Fungsi KMP digunakan untuk melakukan pencarian pola string dalam teks menggunakan algoritma Knuth-Morris-Pratt.  Sama seperti fungsi BoyerMoore, fungsi KMP mengembalikan posisi ditemukannya pola jika kecocokan ditemukan dan -1 jika tidak ditemukan. Fungsi ini memanfaatkan fungsi pembantu PrefixFunction yang tersimpan di util.go. PrefixFunction mengembalikan array of integer yang menyatakan panjang prefix terbesar yang juga merupakan sufix untuk suatu bagian pola pada posisi tertentu.
Fungsi Similarity digunakan untuk melakukan perhitungan tingkat kecocokan suatu pola dibandingkan dengan teks. Fungsi ini memanfaatkan fungsi bantuan KmpTableGenerator yang akan menghasilkan jump table yang dipakai dalam proses pencarian serta fungsi Lcs. Sebenarnya fungsi similarity ini merupakan pengembangan dari fungsi KMP.

# Tata Cara Penggunaan Program
Program terdiri atas dua komponen, yaitu backend dan frontend. Program backend  menggunakan environment Go lang dan sql, sedangkan program frontend menggunakan Typescript dan framework React. Dipakai juga Docker dan Heroku sebagai untuk deploy dan build websitenya, dengan  dependency yang diperlukan kedua program terdapat pada gambar di bawah. Database Management System yang digunakan adalah MariaDB (development di lokal) dan MySQL  (deployment di server). Skema relasional basis data terdapat pada gambar di bawah.
