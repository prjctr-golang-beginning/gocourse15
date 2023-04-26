package main

import (
	"cohcop/pkg/imageprocessing"
	order2 "cohcop/pkg/order"
	"image"
	"log"
)

func main() {
	// high_coupling
	// Створення нового користувача
	user := order2.User{
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "password",
	}

	// Створення нового замовлення
	order := order2.Order{
		User:       user,
		Product:    "Laptop",
		TotalPrice: 1000,
	}

	// Обробка замовлення та відправлення повідомлення користувачу
	order2.ProcessOrder(order)

	// low_coupling
	op := order2.NewOrderProcessor(nil, nil, nil)
	if err := op.ProcessOrder(order2.OrderIndependent{}); err != nil {
		log.Println(`Some errors`)
	}

	//high_cohesion
	i := imageprocessing.Image{}
	i.AdjustBrightness(0.65)
	i.Resize(65, 200)
	i.ApplyGlareEffect(0.65)
	if err := i.Do(); err != nil {
		log.Println(`Some errors`)
	}

	//low_cohesion
	var iGray image.Image
	iGray = &image.Gray{}
	iGray = imageprocessing.ConvertToGray(iGray)
	iGray = imageprocessing.Resize(iGray, 100, 200)

	log.Println(iGray)
}
