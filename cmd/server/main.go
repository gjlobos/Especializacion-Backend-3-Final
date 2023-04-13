package main

import (
	"log"

	"database/sql"

	"github.com/gin-gonic/gin"

	"final/cmd/server/handler"
	"final/internal/appointment"
	"final/internal/dentist"
	"final/internal/patient"
	"final/pkg/middleware"
	"final/pkg/store"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/finalEspBack3DB")
	if err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	storageDentist := store.NewSqlStoreDentist(db)
	repoDentist := dentist.NewRepository(storageDentist)
	serviceDentist := dentist.NewService(repoDentist)
	dentistHandler := handler.NewDentistHandler(serviceDentist)

	storagePatient := store.NewSqlStorePatient(db)
	repoPatient := patient.NewRepository(storagePatient)
	servicePatient := patient.NewService(repoPatient)
	patientHandler := handler.NewPatientHandler(servicePatient)

	storageAppointment := store.NewSqlStoreAppointment(db)
	repoAppointment := appointment.NewRepository(storageAppointment)
	serviceAppointment := appointment.NewService(repoAppointment)
	AppointmentHandler := handler.NewAppointmentHandler(serviceAppointment)

	r := gin.Default()

	dentists := r.Group("/dentists")
	{
		dentists.GET(":id", dentistHandler.GetDentistByID())
		dentists.POST("", middleware.Authentication(), dentistHandler.PostDentist())
		dentists.PUT(":id", middleware.Authentication(), dentistHandler.PutDentist())
		dentists.PATCH(":id", middleware.Authentication(), dentistHandler.PatchDentist())
		dentists.DELETE(":id", middleware.Authentication(), dentistHandler.DeleteDentist())
	}

	patients := r.Group("/patients")
	{
		patients.GET(":id", patientHandler.GetPatientByID())
		patients.POST("", middleware.Authentication(), patientHandler.PostPatient())
		patients.PUT(":id", middleware.Authentication(), patientHandler.PutPatient())
		patients.PATCH(":id", middleware.Authentication(), patientHandler.PatchPatient())
		patients.DELETE(":id", middleware.Authentication(), patientHandler.DeletePatient())
	}

	appointments := r.Group("/appointments")
	{
		appointments.GET(":id", AppointmentHandler.GetAppointmentByID())
		appointments.GET("personal_id_number/:personal_id_number", AppointmentHandler.GetAppointmentByPersonalIdNumber())
		appointments.POST("", middleware.Authentication(), AppointmentHandler.PostAppointment())
		appointments.PUT(":id", middleware.Authentication(), AppointmentHandler.PutAppointment())
		appointments.PATCH(":id", middleware.Authentication(), AppointmentHandler.PatchAppointment())
		appointments.DELETE(":id", middleware.Authentication(), AppointmentHandler.DeleteAppointment())
	}

	r.Run(":8080")
}
