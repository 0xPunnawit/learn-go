package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type Cover struct {
	Id       int
	Fullname string
	Address  string
	Phone    string
}

var db *sqlx.DB

func main() {
	// โหลดไฟล์ .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// อ่านค่าจาก .env
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	// สร้าง connection string
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbname,
	)

	// เปิดการเชื่อมต่อ DB ผ่าน pgx driver
	var err error
	db, err = sqlx.Open("pgx", dsn)
	if err != nil {
		log.Fatal("❌ Failed to open connection:", err)
	}

	// เพิ่มข้อมูล cover
	// cover := Cover{
	// 	Fullname: "สมชาย ใจดี",
	// 	Address:  "123 หมู่ 4 ตำบลสุขใจ อำเภอเมือง จังหวัดกรุงเทพฯ 10100",
	// 	Phone:    "0812345600",
	// }
	// err = AddCover(cover)
	// if err != nil {
	// 	panic(err)
	// }

	// แก้ไขข้อมูล cover
	// cover := Cover{
	// 	Id:       16,
	// 	Fullname: "สมชาย ใจดี Update",
	// 	Address:  "123 หมู่ 4 ตำบลสุขใจ อำเภอเมือง จังหวัดกรุงเทพฯ 10100 Update",
	// 	Phone:    "08Update",
	// }
	// err = UpdateCover(cover)
	// if err != nil {
	// 	panic(err)
	// }

	// ลบข้อมูล cover
	// err = DeleteCover(Cover{Id: 17})
	// if err != nil {
	// 	panic(err)
	// }

	// ดึงข้อมูล cover ทั้งหมด
	// covers, err := GetCovers()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// for _, cover := range covers {
	// 	fmt.Println(cover)
	// }

	// ดึงข้อมูล cover ตาม id
	cover, err := GetCoverX(19)
	if err != nil {
		panic(err)
	}
	fmt.Println(cover)

}

func GetCoversX() ([]Cover, error) {
	query := "SELECT id, fullname, address, phone FROM cover"
	covers := []Cover{}
	err := db.Select(&covers, query)
	if err != nil {
		return nil, err
	}

	return covers, nil
}

func GetCoverX(id int) (*Cover, error) {
	query := "SELECT id, fullname, address, phone FROM cover WHERE id=$1"
	cover := Cover{}
	err := db.Get(&cover, query, id)
	if err != nil {
		return nil, err
	}

	return &cover, nil
}

// ดึงข้อมูล cover ทั้งหมด
func GetCovers() ([]Cover, error) {
	// ทดสอบ ping
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("✅ Connected to PostgreSQL successfully!")

	query := "SELECT id, fullname, address, phone FROM cover"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	covers := []Cover{}
	for rows.Next() {
		cover := Cover{}
		err := rows.Scan(&cover.Id, &cover.Fullname, &cover.Address, &cover.Phone)
		if err != nil {
			return nil, err
		}

		covers = append(covers, cover)
	}

	return covers, nil
}

// ดึงข้อมูล cover ตาม id
func GetCover(id int) (*Cover, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	query := "SELECT id, fullname, address, phone FROM cover WHERE id=$1"
	row := db.QueryRow(query, id)

	cover := Cover{}
	err = row.Scan(&cover.Id, &cover.Fullname, &cover.Address, &cover.Phone)
	if err != nil {
		return nil, err
	}

	return &cover, nil
}

func AddCover(cover Cover) error {

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	query := "INSERT INTO cover (fullname, address, phone) VALUES ($1, $2, $3)"
	result, err := tx.Exec(query, cover.Fullname, cover.Address, cover.Phone)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	if affected <= 0 {
		return errors.New("cannot insert")
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil

}

func UpdateCover(cover Cover) error {

	query := "UPDATE cover SET fullname=$1, address=$2, phone=$3 WHERE id=$4"
	result, err := db.Exec(query, cover.Fullname, cover.Address, cover.Phone, cover.Id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected <= 0 {
		return errors.New("cannot update")
	}

	return nil

}

func DeleteCover(cover Cover) error {
	query := "DELETE FROM cover WHERE id=$1"
	result, err := db.Exec(query, cover.Id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("no record found with id %d", cover.Id)
	}

	return nil
}
