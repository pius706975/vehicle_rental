package main

// func (r *reserve_repo) CreateReservation(data *models.Reservation) (*models.Reservation, error) {

// 	transaction := r.db.Begin()
// 	defer func() {
// 		recover := recover()
// 		if recover != nil {
// 			transaction.Rollback()
// 			return
// 		}
// 	}()

// 	errT := transaction.Error
// 	if errT != nil {
// 		return nil, errT
// 	}

// 	// user check
// 	var user models.User

// 	err := r.db.Where("user_id = ?", data.UsersID).First(&user).Error
// 	if err != nil {
// 		transaction.Rollback()
// 		return nil, errors.New("user not found")
// 	}

// 	// vehicle check 
// 	var vehicle models.Vehicle

// 	errV := r.db.Where("vehicle_id = ?", data.VehicleID).First(&vehicle).Error
// 	if errV != nil {
// 		transaction.Rollback()
// 		return nil, errors.New("vehicle not found")
// 	}

// 	// stock check
// 	vehicleStock := vehicle.Stock - data.Qt 
// 	if vehicleStock < data.Qt {
// 		transaction.Rollback()
// 		return nil, errors.New("not enough stock")
// 	}

// 	errC := transaction.Create(data).Error
// 	if errC != nil {
// 		transaction.Rollback()
// 		return nil, errC
// 	}

// 	// create the reservation history
// 	history := &models.History{
// 		ReservationID: data.ReservationID,
// 		CreatedAt:     data.CreatedAt,
// 		UpdatedAt:     data.UpdatedAt,
// 	}

// 	errCH := transaction.Create(history).Error
// 	if errCH != nil {
// 		transaction.Rollback()
// 		return nil, errors.New("history cannot be created")
// 	}


// 	if data.PaymentMethod == "cash" {
// 		err := r.db.Model(&vehicle).Where("vehicle_id = ?", data.VehicleID).Update("stock", vehicleStock).Error
// 		if err != nil {
// 			transaction.Rollback()
// 			return nil, errors.New("update failed")
// 		}
// 	}

// 	if err := transaction.Preload("Users", func (db *gorm.DB) *gorm.DB {
// 		return db.Select("user_id, first_name, email, mobile_number, created_at, updated_at")
// 	}).Preload("Vehicles", func(db *gorm.DB) *gorm.DB {
// 		return db.Select("vehicle_id, model, location, price, category_id, rating, created_at, updated_at")
// 	}).Preload("Vehicle.Category").First(&data).Error; err != nil {
// 		return nil, err
// 	}

// 	return data, transaction.Commit().Error
// }

// func (s *reserve_service) CreateReservation(data *models.Reservation) *helper.Response {

// 	paymentCode, err := GeneratePaymentCode(10)
// 	if err != nil {
// 		return helper.New(err.Error(), 400, true)
// 	}
// 	if data.PaymentMethod == "cash" {
// 		data.PaymentStatus = "The rent has been paid"
// 	}

// 	data.PaymentCode = fmt.Sprintf("%s%s", "PC-", paymentCode)

// 	result, err := s.repo.CreateReservation(data)
// 	if err != nil {
// 		return helper.New(err.Error(), 400, true)
// 	}

// 	return helper.New(result, 200, false)
// }

// func GeneratePaymentCode(length int) (string, error) {

// 	charset := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 	max := big.NewInt(int64(len(charset)))
// 	b := make([]byte, length)
// 	for i := range b {
// 		n, err := rand.Int(rand.Reader, max)
// 		if err != nil {
// 			return "", err
// 		}
// 		b[i] = charset[n.Int64()]
// 	}
// 	return string(b), nil
// }