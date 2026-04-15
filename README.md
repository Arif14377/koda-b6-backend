# Ecommerce Coffee Shop
Project mandiri bootcamp koda batch 6. Aplikasi web pemesanan kopi dan makanan

## Fitur User
1. Authentication
    - Register
    - Login
    - Forgot Password
2. Recommended Product
    Menampilkan 4 produk rekomendasi berdasarkan.....
3. Reviews
    Review dari customer
4. List Product
    Menampilkan list semua produk dengan pagination.
5. Filter Product
    - keyword (search)
    - kategori
6. Detail Produk
    Menampilkan detail sebuah produk.
7. Keranjang dan Checkout
    - Menambahkan dan menghapus produk dari keranjang, dengan variasi dan size produk.
    - Menghapus produk dari keranjang jika checkout
8. List Histori Pemesanan
    Menampilkan list history belanja.
9. Detail Histori Pemesanan
    Menampilkan detail pemesanan berdasarkan order id.

# Dokumentasi API
| Group | Method | Endpoint | Auth | Request Body |
| :--- | :--- | :--- | :--- | :--- |
| Auth | POST | /auth/register | - | {"fullName":"string", "email":"string", "password":"string", "confirmPassword:"string"}
| Auth | POST | /auth/login | - | {"email":"string", "password":"string"}
| Auth | POST | /users/forgot-password | - | {"email":"string"}
| Auth | POST | /users/forgot-password/verifikasi-otp | - | {"email":"string", "code":integer}
| Auth | POST | /users/forgot-password/change | - | {"email":"string", "newPassword":"string", "confirmPassword":"string"}
| Products | GET | /users/products | - | - |
| Products | GET | /users/recommended-products | - | - |
| Products | GET | /users/products/:id | - | {"id": integer}
| Reviews | GET | /users/reviews | - | - |
|  |  |  |  |  |

