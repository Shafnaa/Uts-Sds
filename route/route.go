package route

import (
	"uts/database"
	"uts/models"

	"github.com/gofiber/fiber/v2"
)

// buatlah endpoint Insert Data sesuai dengan Colection Postman
func InsertData(c *fiber.Ctx) error {

	//Tulis Jawaban Code di Sini :))

	return c.JSON(fiber.Map{
		"Pesan": "testing fork and clone",
	})
}

// Lengkapi Code Berikut untuk untuk Mengambil data untuk semua user - user
func GetAllData(c *fiber.Ctx) error {

	

	return c.JSON(fiber.Map{
		"data": user,
	})

}

//Lengkapi Code berikut Untuk Mengambil data dari id_user berdasarkan Parameter

func GetUserByid(c *fiber.Ctx) error {



	return c.JSON(fiber.Map{
		"data": user,
	})
}

func Delete(c *fiber.Ctx) error {

	var user models.User

	id_user := c.Params("id_user")

	database.DB.Where("id_user = ?", id_user).Delete(user)

	return c.JSON(fiber.Map{
		"Pesan": "Data telah di hapus",
	})
}

// mengupdate data user berdasarkan id_user yang di tempatkan di parameter
func Update(c *fiber.Ctx) error {

	id_user := c.Params("id_user")

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var users models.User
	database.DB.Find(&users)
	//data yang di ubah 
	//membuat variable user berdasarkan model user
	var user models.User

	update := models.User{
		Nama:     data["nama"],
		Email:    data["email"],
		Password: data["password"],
	}
	//mengambil database untuk di update

	database.DB.Model(&user).Where("id_user = ?", id_user).Updates(update)

	return c.JSON(fiber.Map{
		"Pesan": "Data User telah di Update",
	
	})
}
